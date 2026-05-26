package router

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/controlplane"
	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/utils/types"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/repo"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
)

const defaultUploadPriority = 100

type (
	workerIdentifierMapKey = string
	jobID                  = int64
)

type Router struct {
	destType string

	db          *sqlquerywrapper.DB
	stagingRepo *repo.StagingFiles
	uploadRepo  *repo.Uploads

	triggerStore       *sync.Map
	createUploadAlways createUploadAlwaysLoader

	logger       logger.Logger
	conf         *config.Config
	statsFactory stats.Stats

	warehouses           []model.Warehouse
	workspaceBySourceIDs map[string]string
	configSubscriberLock sync.RWMutex

	workerChannelMap     map[string]chan *UploadJob
	workerChannelMapLock sync.RWMutex

	createJobMarkerMap     map[string]time.Time
	createJobMarkerMapLock sync.RWMutex

	inProgressMap     map[workerIdentifierMapKey][]jobID
	inProgressMapLock sync.RWMutex

	processingMu sync.Mutex

	scheduledTimesCache     map[string][]int
	scheduledTimesCacheLock sync.RWMutex

	activeWorkerCount atomic.Int32
	now               func() time.Time
	nowSQL            string

	backgroundGroup *errgroup.Group

	tenantManager    *multitenant.Manager
	bcManager        *bcm.BackendConfigManager
	uploadJobFactory UploadJobFactory
	notifier         *notifier.Notifier

	config struct {
		maxConcurrentUploadJobs           int
		allowMultipleSourcesForJobsPickup bool
		waitForWorkerSleep                time.Duration
		uploadAllocatorSleep              time.Duration
		uploadStatusTrackFrequency        time.Duration
		shouldPopulateHistoricIdentities  bool
		uploadFreqInS                     config.ValueLoader[int64]
		noOfWorkers                       config.ValueLoader[int]
		enableJitterForSyncs              config.ValueLoader[bool]
		maxParallelJobCreation            config.ValueLoader[int]
		mainLoopSleep                     config.ValueLoader[time.Duration]
		stagingFilesBatchSize             config.ValueLoader[int]
		warehouseSyncFreqIgnore           config.ValueLoader[bool]
		cronTrackerRetries                config.ValueLoader[int64]
		uploadBufferTimeInMin             config.ValueLoader[time.Duration]
	}

	stats struct {
		processingPendingJobsStat        stats.Gauge
		processingAvailableWorkersStat   stats.Gauge
		processingPickupLagStat          stats.Timer
		processingPickupWaitTimeStat     stats.Timer
		schedulerWarehouseLengthStat     stats.Gauge
		schedulerTotalSchedulingTimeStat stats.Timer
		cronTrackerExecTimestamp         stats.Gauge
	}
}

func New(
	reporting types.Reporting,
	destType string,
	conf *config.Config,
	log logger.Logger,
	statsFactory stats.Stats,
	db *sqlquerywrapper.DB,
	notifier *notifier.Notifier,
	tenantManager *multitenant.Manager,
	controlPlaneClient *controlplane.Client,
	bcManager *bcm.BackendConfigManager,
	encodingFactory *encoding.Factory,
	triggerStore *sync.Map,
	createUploadAlways createUploadAlwaysLoader,
) *Router {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Backend Config subscriber subscribes to backend-config and gets all the configurations that includes all sources, destinations and their latest values.
func (r *Router) backendConfigSubscriber(ctx context.Context) { _ = "STUB: not implemented"; return }

// non-blocking populate historic identities

// spawn one worker for each unique destID_namespace
// check this commit to https://github.com/rudderlabs/rudder-server/pull/476/commits/fbfddf167aa9fc63485fe006d34e6881f5019667
// to avoid creating goroutine for disabled sources/destinations

// workerIdentifier get name of the worker (`destID_namespace`) to be stored in map wh.workerChannelMap
func (r *Router) workerIdentifier(warehouse model.Warehouse) (identifier string) {
	_ = "STUB: not implemented"
	return ""
}

func (r *Router) initWorker() chan *UploadJob { _ = "STUB: not implemented"; return nil }

func (r *Router) incrementActiveWorkers() { _ = "STUB: not implemented"; return }

func (r *Router) decrementActiveWorkers() { _ = "STUB: not implemented"; return }

func (r *Router) getActiveWorkerCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Router) setDestInProgress(warehouse model.Warehouse, jobID int64) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) removeDestInProgress(warehouse model.Warehouse, jobID int64) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) isUploadJobInProgress(warehouse model.Warehouse, jobID int64) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (r *Router) getInProgressNamespaces() []string { _ = "STUB: not implemented"; return nil }

func (r *Router) checkInProgressMap(jobID int64, identifier string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (r *Router) runUploadJobAllocator(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) uploadsToProcess(ctx context.Context, availableWorkers int, skipIdentifiers []string) ([]*UploadJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Router) processingStats(availableWorkers int, jobStats model.UploadJobsStats) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) mainLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Router) createJobs(ctx context.Context, warehouse model.Warehouse) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) handlePriorityForWaitingUploads(ctx context.Context, warehouse model.Warehouse) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If it is present do nothing else delete it

func (r *Router) uploadStartAfterTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (r *Router) createUploadJobsFromStagingFiles(ctx context.Context, warehouse model.Warehouse, stagingFiles []*model.StagingFile, priority int, uploadStartAfter time.Time) error {
	_ = "STUB: not implemented"
	// count := 0
	// Process staging files in batches of stagingFilesBatchSize
	// E.g. If there are 1000 pending staging files and stagingFilesBatchSize is 100,
	// Then we create 10 new entries in wh_uploads table each with 100 staging files
	return nil
}

// The following will be populated by staging files:
// FirstEventAt:     0,
// LastEventAt:      0,
// UseRudderStorage: false,
// SourceTaskRunID:  "",
// SourceJobID:      "",
// SourceJobRunID:   "",

// reset upload trigger if the upload was triggered

func (r *Router) uploadFrequencyExceeded(warehouse model.Warehouse, syncFrequency string) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Router) uploadFreqInS(syncFrequency string) int64 { _ = "STUB: not implemented"; return 0 }

func (r *Router) updateCreateJobMarker(warehouse model.Warehouse, lastProcessedTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) loadReloadableConfig(whName string) { _ = "STUB: not implemented"; return }

func (r *Router) loadStats() { _ = "STUB: not implemented"; return }

func (r *Router) copyWarehouses() []model.Warehouse { _ = "STUB: not implemented"; return nil }

func (r *Router) getNowSQL() string { _ = "STUB: not implemented"; return "" }
