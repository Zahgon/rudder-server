package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	obskit "github.com/rudderlabs/rudder-observability-kit/go/labels"

	"github.com/rudderlabs/rudder-server/regulation-worker/internal/service"
)

var pkgLogger = logger.NewLogger().Child("regulation-worker")

func main() {
	pkgLogger.Infon("Starting regulation-worker")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	err := Run(ctx)
	if ctx.Err() == nil {
		cancel()
	}
	if err != nil {
		pkgLogger.Errorn("Running regulation worker", obskit.Error(err))
		os.Exit(1)
	}
}

func Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func withLoop(svc service.JobSvc) *service.Looper { _ = "STUB: not implemented"; return nil }

func createHTTPClient(conf *config.Config, httpTimeout time.Duration) *http.Client {
	_ = "STUB: not implemented"
	return nil
}
