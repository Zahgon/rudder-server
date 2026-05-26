package apphandlers

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/app"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// processorApp is the type for Processor type implementation
type processorApp struct {
	setupDone      bool
	app            app.App
	versionHandler func(w http.ResponseWriter, r *http.Request)
	log            logger.Logger
	config         struct {
		eschDSLimit    config.ValueLoader[int]
		arcDSLimit     config.ValueLoader[int]
		rtDSLimit      config.ValueLoader[int]
		batchrtDSLimit config.ValueLoader[int]
		gwDSLimit      config.ValueLoader[int]
		http           struct {
			ReadTimeout       time.Duration
			ReadHeaderTimeout time.Duration
			WriteTimeout      time.Duration
			IdleTimeout       time.Duration
			webPort           int
			MaxHeaderBytes    int
		}
	}
}

func (a *processorApp) Setup() error { _ = "STUB: not implemented"; return nil }

func (a *processorApp) StartRudderCore(ctx context.Context, shutdownFn func(), options *app.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// setup partition migrator

// always run finally to clean up resources regardless of error

// using a cache so that multiple routers can share the same cache and not hit the DB every time

// using a cache so that multiple batch routers can share the same cache and not hit the DB every time

// This should happen only after setupDatabaseTables() is called and journal table migrations are done
// because if this start before that then there might be a case when ReadDB will try to read the owner table
// which gets created after either Write or ReadWrite DB is created.

func (a *processorApp) startHealthWebHandler(ctx context.Context, db jobsdb.JobsDB) error {
	_ = "STUB: not implemented"
	// Port where Processor health handler is running
	return nil
}
