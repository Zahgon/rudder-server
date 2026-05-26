package backendconfig

//go:generate mockgen -destination=../mocks/backend-config/mock_backendconfig.go -package=mock_backendconfig github.com/rudderlabs/rudder-server/backend-config BackendConfig
//go:generate mockgen -destination=./mock_workspaceconfig.go -package=backendconfig -source=./backend-config.go workspaceConfig

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/backend-config/internal/cache"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/services/diagnostics"
	"github.com/rudderlabs/rudder-server/utils/pubsub"
	"github.com/rudderlabs/rudder-server/utils/sysUtils"
	"github.com/rudderlabs/rudder-server/utils/types"
	"github.com/rudderlabs/rudder-server/utils/types/deployment"
)

var (
	// environment variables
	configBackendURL            string
	cpRouterURL                 string
	pollInterval                config.ValueLoader[time.Duration]
	configJSONPath              string
	configFromFile              bool
	configEnvReplacementEnabled bool
	dbCacheEnabled              bool

	LastSync           string
	LastRegulationSync string

	// DefaultBackendConfig will be initialized be Setup to either a WorkspaceConfig or MultiWorkspaceConfig.
	DefaultBackendConfig     BackendConfig
	pkgLogger                = logger.NewLogger().Child("backend-config")
	IoUtil                   = sysUtils.NewIoUtil()
	Diagnostics              diagnostics.DiagnosticsI
	cacheOverride            cache.Cache
	incrementalConfigUpdates bool
)

func disableCache() { _ = "STUB: not implemented"; return }

type noCache struct{}

func (*noCache) Get(context.Context) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type workspaceConfig interface {
	SetUp() error
	// Deprecated: use Identity() instead.
	AccessToken() string
	Get(context.Context) (map[string]ConfigT, error)
	Identity() identity.Identifier
}

type BackendConfig interface {
	workspaceConfig
	WaitForConfig(ctx context.Context)
	Subscribe(ctx context.Context, topic Topic) pubsub.DataChannel
	Stop()
	StartWithIDs(ctx context.Context, workspaces string)
}

type backendConfigImpl struct {
	workspaceConfig
	eb                *pubsub.PublishSubscriber
	ctx               context.Context
	cancel            context.CancelFunc
	blockChan         chan struct{}
	initializedLock   sync.RWMutex
	initialized       bool
	curSourceJSON     map[string]ConfigT
	curSourceJSONLock sync.RWMutex
	usingCache        bool
	cache             cache.Cache
}

func loadConfig() { _ = "STUB: not implemented"; return }

func Init() { _ = "STUB: not implemented"; return }

func trackConfig(preConfig, curConfig ConfigT) { _ = "STUB: not implemented"; return }

func filterProcessorEnabledWorkspaceConfig(config map[string]ConfigT) map[string]ConfigT {
	_ = "STUB: not implemented"
	return nil
}

func filterProcessorEnabledDestinations(config ConfigT) ConfigT {
	_ = "STUB: not implemented"
	return *new(ConfigT)
}

// TODO skipcq: CRT-P0006

func (bc *backendConfigImpl) configUpdate(ctx context.Context) { _ = "STUB: not implemented"; return }

// try to get config from cache

// sorting the sourceJSON.
// json unmarshal does not guarantee order. For DeepEqual to work as expected, sorting is necessary

// only use diagnostics if there is one workspace

// TODO fix concurrent access

func (bc *backendConfigImpl) pollConfigUpdate(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

/*
Subscribe subscribes a channel to a specific topic of backend config updates.

Channel will receive a new pubsub.DataEvent each time the backend configuration is updated.

Data of the DataEvent should be a backendconfig.ConfigT struct.

Available topics are:

- TopicBackendConfig: Will receive complete backend configuration

- TopicProcessConfig: Will receive only backend configuration of processor enabled destinations

- TopicRegulations: Will receive all regulations
*/
func (bc *backendConfigImpl) Subscribe(ctx context.Context, topic Topic) pubsub.DataChannel {
	_ = "STUB: not implemented"
	return *new(pubsub.DataChannel)
}

func newForDeployment(deploymentType deployment.Type, region string, configEnvHandler types.ConfigEnvI) (BackendConfig, error) {
	_ = "STUB: not implemented"
	return *new(BackendConfig), nil
}

// Setup backend config
func Setup(configEnvHandler types.ConfigEnvI) (err error) { _ = "STUB: not implemented"; return nil }

func (bc *backendConfigImpl) StartWithIDs(ctx context.Context, _ string) {
	_ = "STUB: not implemented"
	return
}

// the only reason why we should resume by using no cache,
// would be if no database configuration has been set

func (bc *backendConfigImpl) Stop() { _ = "STUB: not implemented"; return }

// WaitForConfig waits until backend config has been initialized
func (bc *backendConfigImpl) WaitForConfig(ctx context.Context) { _ = "STUB: not implemented"; return }

func GetConfigBackendURL() string { _ = "STUB: not implemented"; return "" }

func getNotOKError(respBody []byte, statusCode int) error { _ = "STUB: not implemented"; return nil }

func (bc *backendConfigImpl) Identity() identity.Identifier {
	_ = "STUB: not implemented"
	return *new(identity.Identifier)
}

// in case of a cached config the ID is not set when operating in single workspace mode
