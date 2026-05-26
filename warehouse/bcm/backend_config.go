package bcm

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	cpclient "github.com/rudderlabs/rudder-server/warehouse/client/controlplane"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/repo"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
)

func New(
	c *config.Config,
	db *sqlquerywrapper.DB,
	tenantManager *multitenant.Manager,
	log logger.Logger,
	stats stats.Stats,
) *BackendConfigManager {
	_ = "STUB: not implemented"
	return nil
}

// BackendConfigManager is used to handle the backend configuration in the Warehouse
type BackendConfigManager struct {
	conf                       *config.Config
	schema                     *repo.WHSchema
	tenantManager              *multitenant.Manager
	internalControlPlaneClient cpclient.InternalControlPlane
	logger                     logger.Logger
	stats                      stats.Stats

	InitialConfigFetched          chan struct{}
	closeInitialConfigFetchedOnce sync.Once

	subscriptions   []chan []model.Warehouse
	subscriptionsMu sync.Mutex

	// variables to store the backend configuration
	warehouses   []model.Warehouse
	warehousesMu sync.RWMutex

	connectionsMap   map[string]map[string]model.Warehouse // destID -> sourceID -> warehouse map
	connectionsMapMu sync.RWMutex

	sourceIDsByWorkspace   map[string][]string // workspaceID -> []sourceIDs
	sourceIDsByWorkspaceMu sync.RWMutex
}

func (bcm *BackendConfigManager) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (bcm *BackendConfigManager) Subscribe(ctx context.Context) <-chan []model.Warehouse {
	_ = "STUB: not implemented"
	return nil
}

func (bcm *BackendConfigManager) processData(ctx context.Context, data map[string]backendconfig.ConfigT) {
	_ = "STUB: not implemented"
	return
}

// PseudoWarehouseDestinationMap is being used instead of WarehouseDestinations because
// SnowpipeStreaming validation requires the destination to be in this workspace config

// TODO how is this used? because we are duplicating data

// namespace gives the namespace for the warehouse in the following order
//  1. user set name from destinationConfig
//  2. from existing record in wh_schemas with same source + dest combo
//  3. convert source name
func (bcm *BackendConfigManager) namespace(ctx context.Context, source backendconfig.SourceT, destination backendconfig.DestinationT) string {
	_ = "STUB: not implemented"
	return ""
}

func (bcm *BackendConfigManager) IsInitialized() bool { _ = "STUB: not implemented"; return false }

func (bcm *BackendConfigManager) Connections() map[string]map[string]model.Warehouse {
	_ = "STUB: not implemented"
	return nil
}

func (bcm *BackendConfigManager) ConnectionSourcesMap(destID string) (map[string]model.Warehouse, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (bcm *BackendConfigManager) SourceIDsByWorkspace() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// WarehousesBySourceID gets all WHs for the given source ID
func (bcm *BackendConfigManager) WarehousesBySourceID(sourceID string) []model.Warehouse {
	_ = "STUB: not implemented"
	return nil
}

// WarehousesByDestID gets all WHs for the given destination ID
func (bcm *BackendConfigManager) WarehousesByDestID(destID string) []model.Warehouse {
	_ = "STUB: not implemented"
	return nil
}

func (bcm *BackendConfigManager) attachSSHTunnellingInfo(
	ctx context.Context,
	upstream backendconfig.DestinationT,
) backendconfig.DestinationT {
	_ = "STUB: not implemented"
	// at destination level, do we have tunnelling enabled.
	return *new(backendconfig.DestinationT)
}

func deepCopy(src, dest any) error { _ = "STUB: not implemented"; return nil }

func (bcm *BackendConfigManager) persistSSLFileErrorStat(
	workspaceID, destType, destName,
	destID, sourceName, sourceID,
	errTag string,
) {
	_ = "STUB: not implemented"
	return
}
