package warehouse

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/app"
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/controlplane"
	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/utils/types"
	whadmin "github.com/rudderlabs/rudder-server/warehouse/admin"
	"github.com/rudderlabs/rudder-server/warehouse/api"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	"github.com/rudderlabs/rudder-server/warehouse/constraints"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
	"github.com/rudderlabs/rudder-server/warehouse/router"
	"github.com/rudderlabs/rudder-server/warehouse/source"
)

type App struct {
	app                app.App
	reporting          types.Reporting
	conf               *config.Config
	logger             logger.Logger
	statsFactory       stats.Stats
	bcConfig           backendconfig.BackendConfig
	db                 *sqlquerywrapper.DB
	notifier           *notifier.Notifier
	tenantManager      *multitenant.Manager
	controlPlaneClient *controlplane.Client
	bcManager          *bcm.BackendConfigManager
	api                *api.Api
	grpcServer         *api.GRPC
	constraintsManager *constraints.Manager
	encodingFactory    *encoding.Factory
	fileManagerFactory filemanager.Factory
	sourcesManager     *source.Manager
	admin              *whadmin.Admin
	triggerStore       *sync.Map
	createUploadAlways *atomic.Bool

	appName string

	config struct {
		host     string
		user     string
		password string
		database string
		sslMode  string
		port     int

		mode                       string
		runningMode                string
		shouldForceSetLowerVersion bool
		dbQueryTimeout             time.Duration
		maxOpenConnections         int

		configBackendURL string
		region           string
	}
}

func New(
	app app.App,
	conf *config.Config,
	log logger.Logger,
	statsFactory stats.Stats,
	bcConfig backendconfig.BackendConfig,
	fileManagerFactory filemanager.Factory,
) *App {
	_ = "STUB: not implemented"
	return nil
}

func (a *App) Setup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *App) setupDatabase(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *App) connectionString(componentName string) string { _ = "STUB: not implemented"; return "" }

// checkForWarehouseEnvVars checks if the required database environment variables are set
func (a *App) checkForWarehouseEnvVars() bool { _ = "STUB: not implemented"; return false }

func (a *App) migrate() error { _ = "STUB: not implemented"; return nil }

func (a *App) migrateAlways() error { _ = "STUB: not implemented"; return nil }

// Run runs the warehouse service
func (a *App) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Report warehouse features

// We don't want to exit if we fail to send features

// Gets the config from config backend and extracts enabled write keys
func (a *App) monitorDestRouters(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *App) onConfigDataEvent(
	configMap map[string]backendconfig.ConfigT,
	dstToWhRouter map[string]*router.Router,
) map[string]*router.Router {
	_ = "STUB: not implemented"
	return nil
}
