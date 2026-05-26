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
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	customDestinationManager "github.com/rudderlabs/rudder-server/router/customdestinationmanager"
	"github.com/rudderlabs/rudder-server/router/internal/eventorder"
	"github.com/rudderlabs/rudder-server/router/internal/partition"
	"github.com/rudderlabs/rudder-server/router/isolation"
	"github.com/rudderlabs/rudder-server/router/throttler"
	"github.com/rudderlabs/rudder-server/router/transformer"
	"github.com/rudderlabs/rudder-server/router/types"
	routerutils "github.com/rudderlabs/rudder-server/router/utils"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	"github.com/rudderlabs/rudder-server/services/rsources"
	transformerFeaturesService "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/services/transientsource"
)

const module = "router"

// Handle is the handle to this module.
type Handle struct {
	// external dependencies
	jobsDB                     jobsdb.JobsDB
	throttlerFactory           throttler.Factory
	backendConfig              backendconfig.BackendConfig
	Reporting                  reporter
	transientSources           transientsource.Service
	rsourcesService            rsources.JobService
	transformerFeaturesService transformerFeaturesService.FeaturesService
	debugger                   destinationdebugger.DestinationDebugger
	adaptiveLimit              func(int64) int64

	// configuration
	reloadableConfig                   *reloadableConfig
	destType                           string
	guaranteeUserEventOrder            bool
	netClientTimeout                   time.Duration
	transformerTimeout                 time.Duration
	enableBatching                     bool
	noOfWorkers                        int
	eventOrderKeyThreshold             config.ValueLoader[int]
	eventOrderDisabledStateDuration    config.ValueLoader[time.Duration]
	eventOrderHalfEnabledStateDuration config.ValueLoader[time.Duration]
	deliveryThrottlerTimeout           config.ValueLoader[time.Duration]
	drainConcurrencyLimit              config.ValueLoader[int]
	maxNoOfJobsPerChannel              int // maximum capacity of each worker channel (hard capacity limit of the underlying go channel)
	noOfJobsPerChannel                 int // requested capacity of each worker channel (important when job buffering is being calculated using the standard method)
	saveDestinationResponse            bool
	saveDestinationResponseOverride    config.ValueLoader[bool]
	reportJobsdbPayload                config.ValueLoader[bool]

	diagnosisTickerTime time.Duration

	// state

	logger                         logger.Logger
	tracer                         stats.Tracer
	telemetry                      *Diagnostic
	netHandle                      NetHandle
	customDestinationManager       customDestinationManager.DestinationManager
	transformer                    transformer.Transformer
	destinationsMapMu              sync.RWMutex
	destinationsMap                map[string]*routerutils.DestinationWithSources // destinationID -> destination
	connectionsMap                 map[types.SourceDest]types.ConnectionWithID
	isBackendConfigInitialized     bool
	backendConfigInitialized       chan bool
	responseQ                      chan workerJobStatus
	throttlingCosts                atomic.Pointer[types.EventTypeThrottlingCost]
	batchSizeHistogramStat         stats.Measurement
	batchInputCountStat            stats.Measurement
	batchOutputCountStat           stats.Measurement
	routerTransformInputCountStat  stats.Measurement
	routerTransformOutputCountStat stats.Measurement
	batchInputOutputDiffCountStat  stats.Measurement
	routerResponseTransformStat    stats.Measurement
	processRequestsHistogramStat   stats.Measurement
	processRequestsCountStat       stats.Measurement
	processJobsHistogramStat       stats.Measurement
	processJobsCountStat           stats.Measurement
	throttlingErrorStat            stats.Measurement
	throttledStat                  stats.Measurement
	isolationStrategy              isolation.Strategy
	backgroundGroup                *errgroup.Group
	backgroundCtx                  context.Context
	backgroundCancel               context.CancelFunc
	backgroundWait                 func() error
	startEnded                     chan struct{}
	barrier                        *eventorder.Barrier

	eventOrderingDisabledForWorkspace   func(workspaceID string) bool
	eventOrderingDisabledForDestination func(destinationID string) bool

	limiter struct {
		pickup    kitsync.Limiter
		transform kitsync.Limiter
		batch     kitsync.Limiter
		process   kitsync.Limiter
		stats     struct {
			pickup    *partition.Stats
			transform *partition.Stats
			batch     *partition.Stats
			process   *partition.Stats
		}
	}

	drainer              routerutils.Drainer
	drainingPartitionsMu sync.RWMutex
	drainingPartitions   map[string]struct{} // keeps track of router partitions which are currently draining
}

// activePartitions returns the list of active partitions, depending on the active isolation strategy
func (rt *Handle) activePartitions(ctx context.Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// pickup picks up jobs from the jobsDB for the provided partition and returns the number of jobs picked up and whether the limits were reached or not
// picked up jobs are distributed to the workers
func (rt *Handle) pickup(ctx context.Context, partition string, workers []*worker, pickupBatchSizeGauge Gauge[int]) (pickupCount int, limitsReached bool) {
	_ = "STUB: not implemented"
	// pickup limiter with dynamic priority
	return 0, false
}

//#JobOrder (See comment marked #JobOrder

// keep track of which partitions are draining

// exit the limiter before sleeping

// Total wall-clock time spent in findWorkerSlot across this pickup loop.
// Published once at the end so dashboards can spot pathological slot-resolution latency
// (e.g. when buffers saturate or the calculator gets expensive).

// Mark the jobs as executing

// Identify jobs which can be processed

// check

// stop the iterator and count all additional jobs discarded by operator by using the same reason as the last job that was discarded

// the following stat (in combination with the limiter's timer stats) is used to track the pickup loop's average latency  and max processing capacity

// sleep for a while if we are discarding too many jobs, so that we don't have a loop running continuously without producing events

// If the discarded ratio is greater than the penalty threshold,
// sleep for a while to avoid having a loop running continuously without producing events

// exit the limiter before sleeping

// getAdaptedJobQueryBatchSize returns the adapted job query batch size based on the throttling limits
func (*Handle) getAdaptedJobQueryBatchSize(input int, pickupThrottlers func() []throttler.PickupThrottler, readSleep time.Duration, maxLimit int, draining bool) int {
	_ = "STUB: not implemented"
	return 0
}

// rounding up to the nearest second
// Calculate the total limit of all active throttlers:
//
//   - if there is a global throttler, use its limit
//   - if there are event type specific throttlers, use the sum of their limits

// global throttler, total limit is the limit of the first throttler

// throttler per event type, total limit is the sum of all recently used throttler per event type limits

// If throttling is enabled then we need to adapt job query batch size:
//
//  - we assume that we will read for more than readSleep seconds (min 1 second)
//  - we will be setting the batch size to be min(totalLimit * readSleepSeconds, maxLimit)

// rounding up to the nearest second

// cap to maxLimit

// if we are draining and the throttling batch size is less than standard one, we should not reduce it

func (rt *Handle) stopIteration(err error, destinationID string) bool {
	_ = "STUB: not implemented"
	// if the context is cancelled, we can stop iteration
	return false
}

// if we are not guaranteeing user event order, we can stop iteration if there are no more slots available

// delegate to the isolation strategy for the final decision

// commitStatusList commits the status of the jobs to the jobsDB
func (rt *Handle) commitStatusList(workerJobStatuses *[]workerJobStatus) {
	_ = "STUB: not implemented"
	return
}

// send response code to throttler
// Update metrics maps
// REPORTING - ROUTER - START

// REPORTING - ROUTER - END

// tracking router errors

// REPORTING - ROUTER - START

// REPORTING - ROUTER - END

// Update the status

// rsources stats

//#JobOrder (see other #JobOrder comment)

// End #JobOrder

func (rt *Handle) getJobsFn(parentContext context.Context) func(context.Context, jobsdb.GetQueryParams, jobsdb.MoreToken) (*jobsdb.MoreJobsResult, error) {
	_ = "STUB: not implemented"
	return nil
}

// parentContext.Err() != nil means we are shutting down
//nolint:nilerr

func (rt *Handle) getQueryParams(partition string, pickUpCount int) jobsdb.GetQueryParams {
	_ = "STUB: not implemented"
	return *new(jobsdb.GetQueryParams)
}

type workerJobSlot struct {
	slot        *reservedSlot
	drainReason string
}

// reserveAnyWorkerSlot picks a random starting offset and walks the workers slice
// with wrap-around, reserving the first worker that has a free slot. Returns nil if
// no worker has capacity. ReserveSlot is the atomic decision point — there is no
// pre-filter step (which would be redundant work and would race with the reservation).
func reserveAnyWorkerSlot(workers []*worker) *reservedSlot { _ = "STUB: not implemented"; return nil }

func (rt *Handle) findWorkerSlot(ctx context.Context, workers []*worker, job *jobsdb.JobT, parameters routerutils.JobParameters, blockedOrderKeys map[eventorder.BarrierKey]struct{}) (*workerJobSlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if job's aborted, then send it to its worker right away

// Pick a random starting offset and walk the workers slice with wrap-around,
// reserving the first worker that has a free slot. ReserveSlot is the atomic
// decision point — checking AvailableSlots() first would be redundant work
// and would also leave a race window between the check and the reservation.

// checking if the orderKey is in blockedOrderKeys. If yes, returning nil.
// this check is done to maintain order.

//#JobOrder (see other #JobOrder comment)
// backoff

//#EndJobOrder

// checks if job is configured to drain or if it's retry limit is reached
func (rt *Handle) drainOrRetryLimitReached(createdAt time.Time, destID, sourceJobRunID string, jobStatus *jobsdb.JobStatusT) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (rt *Handle) retryLimitReached(status *jobsdb.JobStatusT) bool {
	_ = "STUB: not implemented"
	return false
}

// 5xx errors

// retry time window exceeded

func (*Handle) shouldBackoff(job *jobsdb.JobT) bool { _ = "STUB: not implemented"; return false }

func (rt *Handle) shouldThrottle(ctx context.Context, job *jobsdb.JobT, destinationID, eventType string) (limited bool) {
	_ = "STUB: not implemented"
	return false
}

// throttlerFactory could be nil when throttling is disabled or misconfigured.
// in case of misconfiguration, logging errors are emitted.

// we can't throttle, let's hit the destination, worst case we get a 429

func (rt *Handle) getThrottlingCost(job *jobsdb.JobT, eventType string) (cost int64) {
	_ = "STUB: not implemented"
	return 0
}

func (*Handle) crashRecover() {
	_ = "STUB: not implemented"
	// NO-OP
	return
}
