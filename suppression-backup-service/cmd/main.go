package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rudderlabs/rudder-go-kit/logger"
	obskit "github.com/rudderlabs/rudder-observability-kit/go/labels"

	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
)

var pkgLogger = logger.NewLogger().Child("suppression-backup-service")

func main() {
	pkgLogger.Infon("Starting suppression backup service")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	err := Run(ctx)
	if ctx.Err() == nil {
		cancel()
	}
	if err != nil {
		pkgLogger.Errorn("Could not run suppression backup service", obskit.Error(err))
		os.Exit(1)
	}
}

func Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func getIdentity(ctx context.Context) (identity.Identifier, error) {
	_ = "STUB: not implemented"
	return *new(identity.Identifier), nil
}

// exportPath creates a tmp dir and returns the path to it
func exportPath() (baseDir string, err error) { _ = "STUB: not implemented"; return "", nil }
