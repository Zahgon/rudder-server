package manager

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/router"
	"github.com/rudderlabs/rudder-server/router/batchrouter"
)

type LifecycleManager struct {
	logger        logger.Logger
	rt            *router.Factory
	brt           *batchrouter.Factory
	backendConfig backendconfig.BackendConfig
	currentCancel context.CancelFunc
	waitGroup     *errgroup.Group
}

// Start starts a Router, this is not a blocking call.
// If the router is not completely started and the data started coming then also it will not be problematic as we
// are assuming that the DBs will be up.
func (r *LifecycleManager) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stops the Router, this is a blocking call.
func (r *LifecycleManager) Stop() { _ = "STUB: not implemented"; return }

// New creates a new Router instance
func New(rtFactory *router.Factory, brtFactory *batchrouter.Factory,
	backendConfig backendconfig.BackendConfig, logger logger.Logger,
) *LifecycleManager {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpAsyncDestinationsLogsDir() { _ = "STUB: not implemented"; return }

// Gets the config from config backend and extracts enabled write-keys
func (r *LifecycleManager) monitorDestRouters(
	ctx context.Context, routerFactory *router.Factory, batchrouterFactory *batchrouter.Factory,
) {
	_ = "STUB: not implemented"
	return
}

// Crash recover routerDB, batchRouterDB
// Note: The following cleanups can take time if there are too many
// rt / batch_rt tables and there would be a delay reading from the 'ch' channel
// However, this shouldn't be the problem since backend config pushes config
// to its subscribers in separate goroutines to prevent blocking.

// Remove all contents of aysnc destinations logs directory

// For batch router destinations
