package runner

import (
	"context"
	"net/http"
	"time"

	_ "go.uber.org/automaxprocs"

	"github.com/rudderlabs/rudder-go-kit/logger"
	_ "github.com/rudderlabs/rudder-go-kit/maxprocs"

	"github.com/rudderlabs/rudder-server/app"
	"github.com/rudderlabs/rudder-server/app/apphandlers"
	"github.com/rudderlabs/rudder-server/warehouse"
)

// ReleaseInfo holds the release information
type ReleaseInfo struct {
	Version         string
	Commit          string
	BuildDate       string
	BuiltBy         string
	EnterpriseToken string
}

// Runner is responsible for running the application
type Runner struct {
	appType                   string
	application               app.App
	releaseInfo               ReleaseInfo
	warehouseMode             string
	warehouseApp              *warehouse.App
	enableSuppressUserFeature bool
	logger                    logger.Logger
	appHandler                apphandlers.AppHandler
	gracefulShutdownTimeout   time.Duration
}

// New creates and initializes a new Runner
func New(releaseInfo ReleaseInfo) *Runner { _ = "STUB: not implemented"; return nil }

// Run runs the application and returns the exit code
func (r *Runner) Run(ctx context.Context, shutdownFn func(), args []string) int {
	_ = "STUB: not implemented"
	// Start stats
	return 0
}

// TODO: remove as soon as we update the configuration with statsExcludedTags where necessary

// application & backend setup should be done before starting any new goroutines.

// Prepare databases in sequential order, so that failure in one doesn't affect others (leaving dirty schema migration state)

// Start admin server

// Start rudder core

// we don't want to exit if we can't send server features

// Start warehouse
// initialize warehouse service after core to handle non-normal recovery modes

// clearing zap Log buffer to std output

// Assume graceful shutdown failed, log remain goroutines and force kill

func runAllInit() { _ = "STUB: not implemented"; return }

func (r *Runner) versionInfo() map[string]any { _ = "STUB: not implemented"; return nil }

func (r *Runner) versionHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (r *Runner) printVersion() { _ = "STUB: not implemented"; return }

func (r *Runner) canStartServer() bool { _ = "STUB: not implemented"; return false }

func (r *Runner) canStartWarehouse() bool { _ = "STUB: not implemented"; return false }
