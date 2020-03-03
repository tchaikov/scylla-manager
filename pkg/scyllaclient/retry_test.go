// Copyright (C) 2017 ScyllaDB

package scyllaclient_test

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/scylladb/mermaid/pkg/scyllaclient"
	"github.com/scylladb/mermaid/pkg/scyllaclient/scyllaclienttest"
)

func fastRetry(config *scyllaclient.Config) {
	config.Backoff.MaxRetries = 3
	config.Backoff.WaitMin = 50 * time.Millisecond
}

func TestRetrySingleHost(t *testing.T) {
	t.Parallel()

	t.Run("error", func(t *testing.T) {
		host, port, closeServer := scyllaclienttest.MakeServer(t, scyllaclienttest.RespondStatus(t, 999, 999, 999, 999, 200))
		defer closeServer()
		client := scyllaclienttest.MakeClient(t, host, port, fastRetry)

		_, err := client.NodeInfo(context.Background(), host)
		if err == nil {
			t.Fatalf("NodeInfo() expected error")
		}
		if !strings.Contains(err.Error(), "giving up after 4 attempts") {
			t.Fatalf("NodeInfo() error = %s, expected giving up after 4 attempts", err.Error())
		}
	})

	t.Run("success", func(t *testing.T) {
		host, port, closeServer := scyllaclienttest.MakeServer(t, scyllaclienttest.RespondStatus(t, 999, 999, 999, 200))
		defer closeServer()
		client := scyllaclienttest.MakeClient(t, host, port, fastRetry)

		_, err := client.NodeInfo(context.Background(), host)
		if err != nil {
			t.Fatal("NodeInfo() error", err)
		}
	})
}

func TestRetryHostPool(t *testing.T) {
	t.Parallel()

	t.Run("error", func(t *testing.T) {
		statusCode := map[string]int{
			"127.0.0.1": 999,
			"127.0.0.2": 999,
			"127.0.0.3": 999,
			"127.0.0.4": 999,
			"127.0.0.5": 999,
		}

		_, port, closeServer := scyllaclienttest.MakeServer(t,
			scyllaclienttest.RespondHostStatus(t, statusCode),
			scyllaclienttest.ServerListenOnAddr(t, ":0"),
		)
		defer closeServer()

		multiHost := func(config *scyllaclient.Config) {
			config.Hosts = []string{
				"127.0.0.1",
				"127.0.0.2",
				"127.0.0.3",
				"127.0.0.4",
				"127.0.0.5",
			}
		}
		client := scyllaclienttest.MakeClient(t, "", port, multiHost)

		_, err := client.ClusterName(context.Background())
		if err == nil {
			t.Fatalf("ClusterName() expected error")
		}
		if !strings.Contains(err.Error(), "giving up after 5 attempts") {
			t.Fatalf("ClusterName() error = %s, expected giving up after 5 attempts", err.Error())
		}
	})

	t.Run("success", func(t *testing.T) {
		statusCode := map[string]int{
			"127.0.0.1": 999,
			"127.0.0.2": 200,
		}

		_, port, closeServer := scyllaclienttest.MakeServer(t,
			scyllaclienttest.RespondHostStatus(t, statusCode),
			scyllaclienttest.ServerListenOnAddr(t, ":0"),
		)
		defer closeServer()

		multiHost := func(config *scyllaclient.Config) {
			config.Hosts = []string{
				"127.0.0.1",
				"127.0.0.2",
			}
		}
		client := scyllaclienttest.MakeClient(t, "", port, multiHost)

		_, err := client.ClusterName(context.Background())
		if err != nil {
			t.Fatal("ClusterName() error", err)
		}
	})
}

func TestNoRetry(t *testing.T) {
	t.Parallel()

	host, port, closeServer := scyllaclienttest.MakeServer(t, scyllaclienttest.RespondStatus(t, 999, 200))
	defer closeServer()
	client := scyllaclienttest.MakeClient(t, host, port, fastRetry)

	ctx := scyllaclient.NoRetry(context.Background())
	_, err := client.NodeInfo(ctx, host)
	if err == nil {
		t.Fatalf("NodeInfo() expected error")
	}
}

func TestRetryCancelContext(t *testing.T) {
	t.Parallel()

	table := []struct {
		Name    string
		Handler http.Handler
	}{
		{
			Name:    "Repeat",
			Handler: scyllaclienttest.RespondStatus(t, 999, 999, 999, 200),
		},
		{
			Name:    "Wait",
			Handler: http.HandlerFunc(func(http.ResponseWriter, *http.Request) { time.Sleep(time.Second) }),
		},
	}

	for i := range table {
		test := table[i]

		t.Run(test.Name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			time.AfterFunc(50*time.Millisecond, cancel)

			host, port, closeServer := scyllaclienttest.MakeServer(t, test.Handler)
			defer closeServer()
			client := scyllaclienttest.MakeClient(t, host, port, fastRetry)

			_, err := client.NodeInfo(ctx, host)
			t.Log("NodeInfo() error", err)

			if err == nil {
				t.Fatalf("NodeInfo() expected error")
			}
			if errors.Cause(err) != context.Canceled {
				t.Fatalf("NodeInfo() error=%s, expected context.Canceled", err)
			}
		})
	}
}