// Copyright (C) 2017 ScyllaDB

package restapi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/scylladb/golog"
	"github.com/scylladb/mermaid"
)

func init() {
	render.Respond = httpErrorRender
}

// Services contains REST API services.
type Services struct {
	Cluster   ClusterService
	Repair    RepairService
	Scheduler SchedService
}

// New returns an http.Handler implementing mermaid v1 REST API.
func New(svc *Services, logger log.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(traceIDMiddleware)
	r.Use(recoverPanicsMiddleware)

	r.Use(heartbeat("/ping"))
	r.Use(prometheusMiddleware("/metrics"))

	r.Use(middleware.RequestLogger(httpLogger{logger}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	if svc.Cluster != nil {
		r.Mount("/api/v1/", newClusterHandler(svc.Cluster))
	}
	if svc.Repair != nil {
		r.With(clusterFilter{svc: svc.Cluster}.clusterCtx).
			Mount("/api/v1/cluster/{cluster_id}/repair/", newRepairHandler(svc.Repair))
	}
	if svc.Scheduler != nil {
		r.With(clusterFilter{svc: svc.Cluster}.clusterCtx).
			Mount("/api/v1/cluster/{cluster_id}/", newSchedHandler(svc.Scheduler))
	}
	r.Get("/api/v1/version", newVersionHandler())

	// NotFound registered last due to https://github.com/go-chi/chi/issues/297
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		respondError(w, r, mermaid.ErrNotFound, "")
	})

	return r
}

func httpErrorRender(w http.ResponseWriter, r *http.Request, v interface{}) {
	if err, ok := v.(error); ok {
		httpErr, _ := v.(*httpError)
		if httpErr == nil {
			httpErr = &httpError{
				Err:        err,
				StatusCode: http.StatusInternalServerError,
				Message:    "unexpected error, consult logs",
				TraceID:    log.TraceID(r.Context()),
			}
		}

		if le, _ := middleware.GetLogEntry(r).(*httpLogEntry); le != nil {
			le.AddFields("Error", httpErr.Error())
		}

		render.Status(r, httpErr.StatusCode)
		render.DefaultResponder(w, r, httpErr)
		return
	}

	render.DefaultResponder(w, r, v)
}
