package api

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/controlplane"
	proto "github.com/rudderlabs/rudder-server/proto/warehouse"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	cpclient "github.com/rudderlabs/rudder-server/warehouse/client/controlplane"
	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/repo"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
)

const (
	triggeredSuccessfully         = "Triggered successfully"
	noPendingEvents               = "No pending events to sync for this destination"
	downloadFileNamePattern       = "downloadfile.*.tmp"
	noSuchSync                    = "No such sync exist"
	syncFrequencyThresholdMinutes = 30
)

type GRPC struct {
	proto.UnimplementedWarehouseServer

	conf               *config.Config
	logger             logger.Logger
	isMultiWorkspace   bool
	cpClient           cpclient.InternalControlPlane
	connectionManager  *controlplane.ConnectionManager
	tenantManager      *multitenant.Manager
	bcManager          *bcm.BackendConfigManager
	tableUploadsRepo   *repo.TableUploads
	stagingRepo        *repo.StagingFiles
	schemaRepo         *repo.WHSchema
	uploadRepo         *repo.Uploads
	triggerStore       *sync.Map
	fileManagerFactory filemanager.Factory
	now                func() time.Time

	config struct {
		region         string
		cpRouterUseTLS bool
		instanceID     string
		controlPlane   struct {
			url      string
			userName string
			password string
		}
		enableTunnelling              bool
		defaultLatencyAggregationType model.LatencyAggregationType
		maxLatencyQueryLookbackDays   int
	}
}

func NewGRPCServer(
	conf *config.Config,
	logger logger.Logger,
	statsFactory stats.Stats,
	db *sqlmw.DB,
	tenantManager *multitenant.Manager,
	bcManager *bcm.BackendConfigManager,
	triggerStore *sync.Map,
) (*GRPC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (g *GRPC) processData(configData map[string]backendconfig.ConfigT) {
	_ = "STUB: not implemented"
	// Only 1 connection flag is enough, since they are all the same in multi-workspace environments
	return
}

func (*GRPC) GetHealth(context.Context, *emptypb.Empty) (*wrapperspb.BoolValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) GetWHUploads(ctx context.Context, request *proto.WHUploadsRequest) (*proto.WHUploadsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) GetWHUpload(ctx context.Context, request *proto.WHUploadRequest) (*proto.WHUploadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) TriggerWHUploads(ctx context.Context, request *proto.WHUploadsRequest) (*proto.TriggerWhUploadsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.

func (g *GRPC) TriggerWHUpload(ctx context.Context, request *proto.WHUploadRequest) (*proto.TriggerWhUploadsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.

func (g *GRPC) RetryWHUploads(ctx context.Context, req *proto.RetryWHUploadsRequest) (response *proto.RetryWHUploadsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retry request should trigger on these cases.
// 1. Either provide the retry interval.
// 2. Or provide the List of Upload id's that needs to be re-triggered.

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.
// TODO: Also, get rid of the message in here as well.

func (g *GRPC) CountWHUploadsToRetry(ctx context.Context, req *proto.RetryWHUploadsRequest) (response *proto.RetryWHUploadsResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Remove http status code and use grpc status code. Since it requires compatibility on the cp router side, leaving it as it is for now.
// TODO: Also, get rid of the message in here as well.

func (g *GRPC) Validate(ctx context.Context, req *proto.WHValidationRequest) (*proto.WHValidationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// adding ssh tunnelling info, given we have
// useSSH enabled from upstream

// TODO: We can get rid of the Error field in the response. Since it requires compatibility on the cp router side, leaving it as it is for now.

func (g *GRPC) manageTunnellingSecrets(ctx context.Context, config map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

type validateObjectStorageRequest struct {
	Type   string         `json:"type"`
	Config map[string]any `json:"config"`
}

type invalidDestinationCredErr struct {
	Base      error
	Operation string
}

func (err invalidDestinationCredErr) Error() string { _ = "STUB: not implemented"; return "" }

func (g *GRPC) ValidateObjectStorageDestination(ctx context.Context, request *proto.ValidateObjectStorageRequest) (response *proto.ValidateObjectStorageResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkMapForValidKey checks the presence of key in map
// and if yes verifies that the key is string and non-empty.
func checkMapForValidKey(configMap map[string]any, key string) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *GRPC) validateObjectStorage(ctx context.Context, request validateObjectStorageRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// overrideWithEnv overrides the config keys in the fileManager settings
// with fallback values pulled from env. Only supported for S3 for now.
func overrideWithEnv(ctx context.Context, settings *filemanager.Settings) {
	_ = "STUB: not implemented"
	return
}

func ifNotExistThenSet(keyToReplace string, replaceWith any, configMap map[string]any) {
	_ = "STUB: not implemented"
	return
}

// In case we don't have the key, simply replace it with replaceWith

func (g *GRPC) RetrieveFailedBatches(
	ctx context.Context,
	req *proto.RetrieveFailedBatchesRequest,
) (*proto.RetrieveFailedBatchesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) RetryFailedBatches(
	ctx context.Context,
	req *proto.RetryFailedBatchesRequest,
) (*proto.RetryFailedBatchesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) SyncWHSchema(ctx context.Context, req *proto.SyncWHSchemaRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func statsInterceptor(statsFactory stats.Stats) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (g *GRPC) GetFirstAbortedUploadInContinuousAbortsByDestination(
	ctx context.Context,
	request *proto.FirstAbortedUploadInContinuousAbortsByDestinationRequest,
) (*proto.FirstAbortedUploadInContinuousAbortsByDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSyncLatency returns the sync latency for the given workspace, destination, start time and aggregation minutes
// If sourceID is provided, it will return the sync latency for the given sourceID
// If sourceID is not provided, it will return the sync latency for all sources of the given destination
func (g *GRPC) GetSyncLatency(ctx context.Context, request *proto.SyncLatencyRequest) (*proto.SyncLatencyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GRPC) getLatencyAggregationType(
	srcMap map[string]model.Warehouse, sourceID string,
) (model.LatencyAggregationType, error) {
	_ = "STUB: not implemented"
	return *new(model.LatencyAggregationType), nil
}

// GetDestinationNamespaces returns the most recent namespace for each source for a given destination ID.
func (g *GRPC) GetDestinationNamespaces(ctx context.Context, request *proto.GetDestinationNamespacesRequest) (*proto.GetDestinationNamespacesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate input

// Get namespace mappings from repository

// Convert to proto response
