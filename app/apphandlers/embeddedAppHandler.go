package apphandlers

import (
	"context"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/app"
)

// embeddedApp is the type for embedded type implementation
type embeddedApp struct {
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
	}
}

func (a *embeddedApp) Setup() error { _ = "STUB: not implemented"; return nil }

func (a *embeddedApp) StartRudderCore(ctx context.Context, shutdownFn func(), options *app.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// This separate gateway db is created just to be used with gateway because in case of degraded mode,
// the earlier created gwDb (which was created to be used mainly with processor) will not be running, and it
// will cause issues for gateway because gateway is supposed to receive jobs even in degraded mode.

// setup partition migrator

// always run finally to clean up resources regardless of error

// using a cache so that multiple routers can share the same cache and not hit the DB every time

// using a cache so that multiple batch routers can share the same cache and not hit the DB every time

// This should happen only after setupDatabaseTables() is called and journal table migrations are done
// because if this start before that then there might be a case when ReadDB will try to read the owner table
// which gets created after either Write or ReadWrite DB is created.
