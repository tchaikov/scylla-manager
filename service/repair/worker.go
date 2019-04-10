// Copyright (C) 2017 ScyllaDB

package repair

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/scylladb/go-log"
	"github.com/scylladb/mermaid/internal/dht"
	"github.com/scylladb/mermaid/internal/timeutc"
	"github.com/scylladb/mermaid/scyllaclient"
	"go.uber.org/atomic"
	"go.uber.org/multierr"
)

// unitWorker holds hostWorkers.
type unitWorker []*hostWorker

// hostWorker manages shardWorkers.
type hostWorker struct {
	Config   *Config
	Run      *Run
	Unit     int
	Host     string
	Segments segments

	Service *Service
	Client  *scyllaclient.Client
	Logger  log.Logger

	shards []*shardWorker
	ffabrt atomic.Bool
}

func (w *hostWorker) init(ctx context.Context) error {
	w.Logger.Info(ctx, "Initialising repair")

	// continue from a savepoint
	prog, err := w.Service.getHostProgress(w.Run, w.Unit, w.Host)
	if err != nil {
		return errors.Wrap(err, "failed to get host progress")
	}

	// split segments to shards
	p, err := w.partitioner(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get partitioner")
	}
	shards := w.splitSegmentsToShards(ctx, p)

	// check if savepoint can be used
	if err := validateShardProgress(shards, prog); err != nil {
		if len(prog) > 1 {
			w.Logger.Info(ctx, "Starting from scratch: invalid progress info", "error", err.Error(), "progress", prog)
		}
		prog = nil
	}

	w.shards = make([]*shardWorker, len(shards))

	for i, segments := range shards {
		// prepare progress
		var p *RunProgress
		if prog != nil {
			p = prog[i]
		} else {
			p = &RunProgress{
				ClusterID:    w.Run.ClusterID,
				TaskID:       w.Run.TaskID,
				RunID:        w.Run.ID,
				Unit:         w.Unit,
				Host:         w.Host,
				Shard:        i,
				SegmentCount: len(segments),
			}
		}

		// prepare labels
		labels := prometheus.Labels{
			"cluster":  w.Run.clusterName,
			"task":     w.Run.TaskID.String(),
			"keyspace": w.Run.Units[w.Unit].Keyspace,
			"host":     w.Host,
			"shard":    fmt.Sprint(i),
		}

		w.shards[i] = &shardWorker{
			parent:   w,
			segments: segments,
			progress: p,
			logger:   w.Logger.With("shard", i),

			repairSegmentsTotal:   repairSegmentsTotal.With(labels),
			repairSegmentsSuccess: repairSegmentsSuccess.With(labels),
			repairSegmentsError:   repairSegmentsError.With(labels),
			repairDurationSeconds: repairDurationSeconds.With(labels),
		}
		w.shards[i].init(ctx)
	}

	return nil
}

func (w *hostWorker) partitioner(ctx context.Context) (*dht.Murmur3Partitioner, error) {
	shardCount, err := w.Client.ShardCount(ctx, w.Host)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shard count")
	}
	return dht.NewMurmur3Partitioner(shardCount, uint(w.Config.ShardingIgnoreMsbBits)), nil
}

func (w *hostWorker) splitSegmentsToShards(ctx context.Context, p *dht.Murmur3Partitioner) []segments {
	shards := w.Segments.splitToShards(p)
	if err := w.Segments.validateShards(shards, p); err != nil {
		w.Logger.Error(ctx, "Suboptimal sharding", "error", err)
	}

	for i := range shards {
		shards[i] = shards[i].merge()
		shards[i] = shards[i].split(int64(w.Config.SegmentTokensMax))
	}

	return shards
}

func (w *hostWorker) exec(ctx context.Context) error {
	w.Logger.Info(ctx, "Repairing")

	// check if host is available
	if _, err := w.Client.Ping(ctx, w.Host); err != nil {
		return errors.Wrap(err, "ping failed")
	}

	// check if no other repairs are running
	if active, err := w.Client.ActiveRepairs(ctx, w.Run.Units[w.Unit].hosts); err != nil {
		w.Logger.Error(ctx, "Active repair check failed", "error", err)
	} else if len(active) > 0 {
		return errors.Errorf("active repair on hosts: %s", strings.Join(active, ", "))
	}

	// limit nr of shards running in parallel
	limit := len(w.shards)
	if l := w.Config.ShardParallelMax; l > 0 && l < limit {
		limit = l
	}

	idx := atomic.NewInt32(0)

	// run shard workers
	wch := make(chan error)
	for j := 0; j < limit; j++ {
		go func() {
			for {
				i := int(idx.Inc()) - 1
				if i >= len(w.shards) {
					return
				}
				shardCtx := log.WithFields(ctx, "host", w.Host, "shard", i)
				wch <- errors.Wrapf(w.shards[i].exec(shardCtx), "shard %d", i)
			}
		}()
	}

	// run metrics updater
	u := newProgressMetricsUpdater(w.Run, w.Service.getProgress, w.Logger)
	go u.Run(ctx, 5*time.Second)

	// join shard workers
	var werr error

	for range w.shards {
		if err := <-wch; err != nil {
			werr = multierr.Append(werr, err)
			if w.Run.failFast {
				w.ffabrt.Store(true)
			}
		}
	}
	// join metrics updater
	u.Stop()

	// try killing any remaining repairs
	killCtx := log.CopyTraceID(context.Background(), ctx)
	if err := w.Client.KillAllRepairs(killCtx, w.Host); err != nil {
		w.Logger.Error(killCtx, "Failed to terminate repairs", "error", err)
	}

	// if repair was canceled return ctx.Err()
	if ctx.Err() != nil {
		w.Logger.Info(ctx, "Repair stopped")
		return ctx.Err()
	}

	return werr
}

func (w *hostWorker) segmentErrors() int {
	v := 0
	for _, s := range w.shards {
		v += s.progress.SegmentError
	}
	return v
}

// shardWorker repairs a single shard.
type shardWorker struct {
	parent   *hostWorker
	segments segments
	progress *RunProgress
	logger   log.Logger

	repairSegmentsTotal   prometheus.Gauge
	repairSegmentsSuccess prometheus.Gauge
	repairSegmentsError   prometheus.Gauge
	repairDurationSeconds prometheus.Summary
}

func (w *shardWorker) init(ctx context.Context) {
	w.updateProgress(ctx)
}

func (w *shardWorker) exec(ctx context.Context) error {
	if w.progress.complete() {
		w.logger.Info(ctx, "Already repaired, skipping")
		return nil
	}

	if w.progress.SegmentError > w.parent.Config.ShardFailedSegmentsMax {
		w.logger.Info(ctx, "Starting from scratch: too many errors")
		w.resetProgress()
		w.updateProgress(ctx)
	}

	w.logger.Info(ctx, "Repairing", "percent_complete", w.progress.PercentComplete())

	var err error
	if w.progress.completeWithErrors() {
		err = w.repair(ctx, w.newRetryIterator())
	} else {
		err = w.repair(ctx, w.newForwardIterator())
	}

	if err == nil {
		w.logger.Info(ctx, "Repair ended", "percent_complete", w.progress.PercentComplete())
	} else {
		w.logger.Info(ctx, "Repair ended", "percent_complete", w.progress.PercentComplete(), "error", err)
	}

	return err
}

func (w *shardWorker) newRetryIterator() *retryIterator {
	return &retryIterator{
		segments:          w.segments,
		progress:          w.progress,
		segmentsPerRepair: w.parent.Config.SegmentsPerRepair,
	}
}

func (w *shardWorker) newForwardIterator() *forwardIterator {
	return &forwardIterator{
		segments:          w.segments,
		progress:          w.progress,
		segmentsPerRepair: w.parent.Config.SegmentsPerRepair,
	}
}

func (w *shardWorker) repair(ctx context.Context, ri repairIterator) error {
	w.updateProgress(ctx)

	var (
		start int
		end   int
		id    int32
		ok    bool

		lastPercentComplete = w.progress.PercentComplete()
	)

	if w.progress.LastCommandID != 0 {
		id = w.progress.LastCommandID
	}

	next := func() {
		start, end, ok = ri.Next()
	}

	savepoint := func() {
		if ok {
			w.progress.LastStartToken = w.segments[start].StartToken
		} else {
			w.progress.LastStartToken = 0
		}

		if id != 0 {
			w.progress.LastCommandID = id
			w.progress.LastStartTime = timeutc.Now()
		} else {
			w.progress.LastCommandID = 0
			w.progress.LastStartTime = time.Time{}
		}

		w.updateProgress(ctx)
	}

	next()

	for {
		// log progress information every 10%
		if p := w.progress.PercentComplete(); p > lastPercentComplete && p%10 == 0 {
			w.logger.Info(ctx, "Repair in progress...", "percent_complete", p)
			lastPercentComplete = p
		}

		// no more segments
		if !ok {
			break
		}

		// fail fast abort triggered, return immediately
		if w.parent.ffabrt.Load() {
			return nil
		}

		var err error

		// request segment repair
		if id == 0 {
			id, err = w.runRepair(ctx, start, end)
			if err != nil {
				w.logger.Error(ctx, "Failed to request repair", "error", err)
				err = errors.Wrap(err, "failed to request repair")
			} else {
				savepoint()
			}
		}

		// if requester a repair wait for the command to finish
		if err == nil {
			err = w.waitCommand(ctx, id)
			if err != nil {
				w.logger.Error(ctx, "Command error", "command", id, "error", err)
				err = errors.Wrapf(err, "command %d error", id)
			} else {
				ri.OnSuccess()
			}
		}

		// handle error
		if err != nil {
			// if repair was stopped return immediately
			if ctx.Err() != nil {
				return err
			}

			ri.OnError()

			if w.parent.Run.failFast {
				next()
				savepoint()
				return err
			}

			w.logger.Info(ctx, "Pausing repair due to an error", "wait", w.parent.Config.ErrorBackoff)
			w.updateProgress(ctx)
			t := time.NewTimer(w.parent.Config.ErrorBackoff)
			select {
			case <-ctx.Done():
				t.Stop()
				return ctx.Err()
			case <-t.C:
				// continue
			}
		}

		// reset command id and move on
		id = 0
		next()
		savepoint()

		// check if there is not too many errors
		if w.progress.SegmentError > w.parent.Config.ShardFailedSegmentsMax {
			return errors.New("too many failed segments")
		}
	}

	return nil
}

func (w *shardWorker) runRepair(ctx context.Context, start, end int) (int32, error) {
	u := w.parent.Run.Units[w.parent.Unit]

	cfg := &scyllaclient.RepairConfig{
		Keyspace: u.Keyspace,
		Ranges:   w.segments[start:end].dump(),
		Hosts:    w.parent.Run.WithHosts,
	}
	if !u.allDCs {
		cfg.DC = w.parent.Run.DC
	}
	if !u.AllTables {
		cfg.Tables = u.Tables
	}
	id, err := w.parent.Client.Repair(ctx, w.parent.Host, cfg)

	return id, err
}

func (w *shardWorker) waitCommand(ctx context.Context, id int32) error {
	start := timeutc.Now()
	defer func() {
		w.repairDurationSeconds.Observe(timeutc.Since(start).Seconds())
	}()

	t := time.NewTicker(w.parent.Config.PollInterval)
	defer t.Stop()

	u := w.parent.Run.Units[w.parent.Unit]

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			s, err := w.parent.Client.RepairStatus(ctx, w.parent.Host, u.Keyspace, id)
			if err != nil {
				return err
			}
			switch s {
			case scyllaclient.CommandRunning:
				// continue
			case scyllaclient.CommandSuccessful:
				return nil
			case scyllaclient.CommandFailed:
				return errors.New("repair failed consult Scylla logs")
			default:
				return errors.Errorf("unknown command status %q", s)
			}
		}
	}
}

func (w *shardWorker) resetProgress() {
	w.progress.SegmentSuccess = 0
	w.progress.SegmentError = 0
	w.progress.SegmentErrorStartTokens = nil
	w.progress.LastStartToken = 0
	w.progress.LastStartTime = time.Time{}
	w.progress.LastCommandID = 0
}

func (w *shardWorker) updateProgress(ctx context.Context) {
	if err := w.parent.Service.putRunProgress(ctx, w.progress); err != nil {
		w.logger.Error(ctx, "Cannot update the run progress", "error", err)
	}

	w.repairSegmentsTotal.Set(float64(w.progress.SegmentCount))
	w.repairSegmentsSuccess.Set(float64(w.progress.SegmentSuccess))
	w.repairSegmentsError.Set(float64(w.progress.SegmentError))
}