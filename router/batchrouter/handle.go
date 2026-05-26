package batchrouter

import (
	"context"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	asynccommon "github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
	"github.com/rudderlabs/rudder-server/router/batchrouter/isolation"
	routerutils "github.com/rudderlabs/rudder-server/router/utils"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	"github.com/rudderlabs/rudder-server/services/diagnostics"
	"github.com/rudderlabs/rudder-server/services/rsources"
	"github.com/rudderlabs/rudder-server/services/transientsource"
	"github.com/rudderlabs/rudder-server/utils/types"
	"github.com/rudderlabs/rudder-server/warehouse/client"
)

const module = "batch_router"

type Handle struct {
	destType string
	// dependencies

	conf               *config.Config
	logger             logger.Logger
	netHandle          *http.Client
	jobsDB             jobsdb.JobsDB
	reporting          types.Reporting
	backendConfig      backendconfig.BackendConfig
	fileManagerFactory filemanager.Factory
	transientSources   transientsource.Service
	rsourcesService    rsources.JobService
	warehouseClient    *client.Warehouse
	debugger           destinationdebugger.DestinationDebugger
	Diagnostics        diagnostics.DiagnosticsI
	adaptiveLimit      func(int64) int64
	isolationStrategy  isolation.Strategy
	now                func() time.Time

	// configuration

	maxEventsInABatch            int
	maxPayloadSizeInBytes        int
	maxFailedCountForJob         config.ValueLoader[int]
	maxFailedCountForSourcesJob  config.ValueLoader[int]
	asyncUploadTimeout           config.ValueLoader[time.Duration]
	asyncUploadWorkerTimeout     config.ValueLoader[time.Duration]
	retryTimeWindow              config.ValueLoader[time.Duration]
	sourcesRetryTimeWindow       config.ValueLoader[time.Duration]
	schemaGenerationWorkers      config.ValueLoader[int]
	reportingEnabled             bool
	jobQueryBatchSize            config.ValueLoader[int]
	pollStatusLoopSleep          config.ValueLoader[time.Duration]
	payloadLimit                 config.ValueLoader[int64]
	jobsDBCommandTimeout         config.ValueLoader[time.Duration]
	jobdDBQueryRequestTimeout    config.ValueLoader[time.Duration]
	jobdDBMaxRetries             config.ValueLoader[int]
	minIdleSleep                 config.ValueLoader[time.Duration]
	uploadFreq                   config.ValueLoader[time.Duration]
	pingFrequency                config.ValueLoader[time.Duration]
	disableEgress                bool
	warehouseServiceMaxRetryTime config.ValueLoader[time.Duration]
	transformerURL               string
	datePrefixOverride           config.ValueLoader[string]
	customDatePrefix             config.ValueLoader[string]
	maxImportingQueryIterations  config.ValueLoader[int]

	drainer routerutils.Drainer

	// state

	backgroundGroup  *errgroup.Group
	backgroundCtx    context.Context
	backgroundCancel context.CancelFunc
	backgroundWait   func() error

	backendConfigInitializedOnce sync.Once
	backendConfigInitialized     chan bool

	configSubscriberMu       sync.RWMutex                                   // protects the following fields
	destinationsMap          map[string]*routerutils.DestinationWithSources // destinationID -> destination
	connectionWHNamespaceMap map[string]string                              // connectionIdentifier -> warehouseConnectionIdentifier(+namepsace)
	uploadIntervalMap        map[string]time.Duration

	encounteredMergeRuleMapMu sync.Mutex
	encounteredMergeRuleMap   map[string]map[string]bool

	limiter struct {
		read   kitsync.Limiter
		upload kitsync.Limiter
	}

	batchRequestsMetricMu sync.RWMutex
	batchRequestsMetric   []batchRequestMetric

	warehouseServiceFailedTimeMu sync.RWMutex
	warehouseServiceFailedTime   time.Time

	dateFormatProvider *storageDateFormatProvider

	diagnosisTicker          *time.Ticker
	uploadedRawDataJobsCache map[string]map[string]bool
	asyncDestinationStruct   map[string]*asynccommon.AsyncDestinationStruct

	asyncPollTimeStat           stats.Measurement
	asyncFailedJobsTimeStat     stats.Measurement
	asyncSuccessfulJobCount     stats.Measurement
	asyncFailedJobCount         stats.Measurement
	asyncAbortedJobCount        stats.Measurement
	asyncGetImportingIterations stats.Measurement
}

// mainLoop is responsible for pinging the workers periodically for every active partition
func (brt *Handle) mainLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// activePartitions returns the list of active partitions, depending on the active isolation strategy
func (brt *Handle) activePartitions(ctx context.Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// getWorkerJobs returns the list of jobs for a given partition. Jobs are grouped by destination
func (brt *Handle) getWorkerJobs(partition string) (workerJobs []*DestinationJobs) {
	_ = "STUB: not implemented"
	return nil
}

// keep trying to get jobs while no jobs are returned because ds limits are being reached

// upload the given batch of jobs to the given object storage provider
func (brt *Handle) upload(provider string, batchJobs *BatchedJobs, isWarehouse bool) UploadResult {
	_ = "STUB: not implemented"
	return *new(UploadResult)
}

// do not add to staging file if the event is a rudder_identity_merge_rules record
// and has been previously added to it

// assumes events from warehouse have receivedAt in metadata

// received_at set in rudder-server has timezone component
// whereas first_event_at column in wh_staging_files is of type 'timestamp without time zone'
// convert it to UTC before saving to wh_staging_files

// used to be earlier default

// pingWarehouse notifies the warehouse about a new data upload (staging files)
func (brt *Handle) pingWarehouse(batchJobs *BatchedJobs, output UploadResult) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (brt *Handle) generateSchemaMap(batchJobs *BatchedJobs) map[string]map[string]string {
	_ = "STUB: not implemented"
	// First group jobs by table name
	return nil
}

// Process each table's jobs in parallel

// Process all jobs for this table

// updateJobStatus updates the statuses for the provided batch of jobs in jobsDB
func (brt *Handle) updateJobStatus(batchJobs *BatchedJobs, isWarehouse bool, errOccurred error, notifyWarehouseErr bool) {
	_ = "STUB: not implemented"
	return
}

// skipcq: GO-R4002

// skipcq: GO-R4002

// We keep track of number of failed attempts in case of failure and number of events uploaded in case of success in stats

// change job state to abort state after warehouse service is continuously failing more than warehouseServiceMaxRetryTimeinHr time

// REPORTING - START

// Update metrics maps

// REPORTING - END

// tracking batch router errors

// REPORTING - START

// REPORTING - END

// Mark the status of the jobs

// rsources stats

// uploadInterval calculates the upload interval for the destination
func (brt *Handle) uploadInterval(destinationConfig map[string]any) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// skipFetchingJobs returns true if the destination type is async and the there are still jobs in [importing] state for this destination type
func (brt *Handle) skipFetchingJobs(partition string) bool { _ = "STUB: not implemented"; return false }

// splitBatchJobsOnTimeWindow splits the batchJobs based on a timeWindow if the destination requires so, otherwise a single entry is returned using the zero value as key
func (brt *Handle) splitBatchJobsOnTimeWindow(batchJobs BatchedJobs) map[time.Time]*BatchedJobs {
	_ = "STUB: not implemented"
	return nil
}

// return only one batchJob if the destination type is not time window destinations

// split batchJobs based on timeWindow

// ignore error as receivedAt will always be in the expected format

// create batchJob for timeWindow if it does not exist

func (brt *Handle) retryLimitReached(status *jobsdb.JobStatusT) bool {
	_ = "STUB: not implemented"
	return false
}

// retry time window exceeded
