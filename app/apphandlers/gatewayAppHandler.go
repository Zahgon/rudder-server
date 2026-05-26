package apphandlers

import (
	"context"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/app"
)

// gatewayApp is the type for Gateway type implementation
type gatewayApp struct {
	setupDone      bool
	app            app.App
	versionHandler func(w http.ResponseWriter, r *http.Request)
	log            logger.Logger
}

func (a *gatewayApp) Setup() error { _ = "STUB: not implemented"; return nil }

func (a *gatewayApp) StartRudderCore(ctx context.Context, _ func(), options *app.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// wrapping Stop call in an anonymous function
// so that we can decorate gwWODB later for partition migrations
// and call Stop on the decorated instance
