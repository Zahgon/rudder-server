package router

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/processor/integrations"
	"github.com/rudderlabs/rudder-server/router/internal/eventorder"
	"github.com/rudderlabs/rudder-server/router/transformer"
	"github.com/rudderlabs/rudder-server/router/types"
	routerutils "github.com/rudderlabs/rudder-server/router/utils"
	"github.com/rudderlabs/rudder-server/utils/cache"
)

// worker a structure to define a worker for sending events to sinks
type worker struct {
	id        int // identifies the worker
	partition string

	rt     *Handle // handle to router
	logger logger.Logger

	workLoopThroughput stats.Histogram // stat to record throughput of the worker's processing loop

	ctx          context.Context    // context for the worker
	cancelFunc   context.CancelFunc // cancel function for the worker context
	workerBuffer *workerBuffer      // the worker's input buffer

	barrier *eventorder.Barrier // barrier to ensure ordering of events

	deliveryTimeStat          stats.Measurement
	routerDeliveryLatencyStat stats.Measurement
	routerProxyStat           stats.Measurement

	// Cache for transformer outgoing request metrics using StatsCache
	deliveryLatencyStatsCache *cache.StatsCache[deliveryMetricLabels]
	deliveryCountStatsCache   *cache.StatsCache[deliveryMetricLabels]
}

type workerJob struct {
	job         *jobsdb.JobT
	parameters  *routerutils.JobParameters
	assignedAt  time.Time
	drainReason string
}

// acceptWorkerJob accepts a worker job and returns a router job if batching/router transformation is enabled.
//
//   - If the job is aborted, it sends an aborted job status to responseQ.
//   - If the job needs to wait due to event ordering, it sends a waiting job status to responseQ.
//   - If no batching or router transformation is enabled, it processes the job immediately, otherwise it returns a router job for batching or transformation.
func (w *worker) acceptWorkerJob(workerJob workerJob) *types.RouterJobT {
	_ = "STUB: not implemented"
	return nil
}

// send aborted job status to responseQ and continue

// Enhancing job parameter with the drain reason.

// send waiting job status to responseQ and continue

// mark job as waiting if prev job from same user has not succeeded yet

// check

// destination or connection not found, skip processing

// add the job to the batch

// then process the current single job

func (w *worker) workLoop() { _ = "STUB: not implemented"; return }

// start the worker loop

func (w *worker) transformRouterJobs(routerJobs []types.RouterJobT) []types.DestinationJobT {
	_ = "STUB: not implemented"
	return nil
}

// OAuth destinations need to be transformed separately per destination to avoid mixing OAuth credentials/tokens across destinations.
// non-OAuth destinations can be transformed together in a single batch as they don't have the issue of mixing credentials/tokens.
// So we group router jobs by OAuth vs non-OAuth destinations and transform them separately.

// assume oauth on error

func (w *worker) transform(routerJobs []types.RouterJobT) []types.DestinationJobT {
	_ = "STUB: not implemented"
	// transform limiter with dynamic priority
	return nil
}

// the following stats (in combination with the limiter's timer stats) are used to capture the transform stage
// average latency, batching efficiency and max processing capacity

func (w *worker) batchTransform(routerJobs []types.RouterJobT) []types.DestinationJobT {
	_ = "STUB: not implemented"
	// batch limiter with dynamic priority
	return nil
}

// the following stats (in combination with the limiter's timer stats) are used to capture the batch stage
// average latency, batching efficiency and max processing capacity

func (w *worker) process(destinationJobs []types.DestinationJobT) {
	_ = "STUB: not implemented"
	// process limiter with dynamic priority
	return
}

// TODO: use w.ctx and handle graceful shutdown scenario

/*
	Batch
	[u1e1, u2e1, u1e2, u2e2, u1e3, u2e3]
	[b1, b2, b3]
	b1 will send if success
	b2 will send if b2 failed then will drop b3

	Router transform
	[u1e1, u2e1, u1e2, u2e2, u1e3, u2e3]
	200, 200, 500, 200, 200, 200

	Case 1:
	u1e1 will send - success
	u2e1 will send - success
	u1e2 will drop because transformer gave 500
	u2e2 will send - success
	u1e3 should be dropped because u1e2 should be retried
	u2e3 will send

	Case 2:
	u1e1 will send - success
	u2e1 will send - failed 5xx
	u1e2 will send
	u2e2 will drop - because request to destination failed with 5xx
	u1e3 will send
	u2e3 will drop - because request to destination failed with 5xx

	Case 3:
	u1e1 will send - success
	u2e1 will send - failed 4xx
	u1e2 will send
	u2e2 will send - because previous job is aborted
	u1e3 will send
	u2e3 will send
*/

// START: request to destination endpoint

// TODO: remove trackStuckDelivery once we verify it is not needed,
//			router_delivery_exceeded_timeout -> goes to zero

// limiting the log to print 10KB of transformed payload

// Record the new transformer_outgoing_request metrics

// If this is the last iteration, use respStatusCodes & respBodyTemps as is
// If this is not the last iteration, mark all the jobs as failed.

// Record the new transformer_outgoing_request metrics

// respStatusCodes are already populated. Prepare respBodys from respBodyArrs
// Never the case

// Are these useful?

// END: request to destination endpoint

// Struct to hold unique users in the batch (worker.destinationJobs)

// This means more than two jobs of the same user are in the batch & the batch job is failed
// Only one job is marked failed and the rest are marked waiting
// Job order logic requires that at any point of time, we should have only one failed job per user
// This is introduced to ensure the above statement

// Used to send this as a directive for transformer to not let this job batch with other jobs

// NOTE: Sending live events to config backend after the status objects are built completely.

// Sending only one destination live event for every destinationJob

// the following stat (in combination with the limiter's timer stats) are used to capture the process stage
// average latency and max processing capacity

func consolidateRespBodys(respBodyArrs []map[int64]string) map[int64]string {
	_ = "STUB: not implemented"
	return nil
}

func anyNonTerminalCode(respStatusCodes map[int64]int) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *worker) proxyRequest(ctx context.Context, destinationJob types.DestinationJobT, val integrations.PostParametersT) transformer.ProxyRequestResponse {
	_ = "STUB: not implemented"
	return *new(transformer.ProxyRequestResponse)
}

// setting metadata

func (w *worker) hydrateRespStatusCodes(destinationJob types.DestinationJobT, respStatusCodes map[int64]int, respBodys map[int64]string) {
	_ = "STUB: not implemented"
	return
}

func (w *worker) updateFailedJobOrderKeys(failedJobOrderKeys map[eventorder.BarrierKey]struct{}, destinationJob *types.DestinationJobT, respStatusCodes map[int64]int) {
	_ = "STUB: not implemented"
	return
}

// if barrier is disabled, we shouldn't need to track the failed job

func (w *worker) prepareRouterJobResponses(destinationJob types.DestinationJobT, respStatusCodes map[int64]int, respBodys map[int64]string, errorAt string) []*JobResponse {
	_ = "STUB: not implemented"
	return nil
}

// Failure - Save response body
// Success - Skip saving response body
// By default we get some config from dest def
// We can override via env saveDestinationResponseOverride

// TODO: remove this once we enforce the necessary validations in the transformer's response

// assigning the destinationJobMetadata to a local variable (_destinationJobMetadata), so that
// elements in routerJobResponses have pointer to the right destinationJobMetadata.

func (w *worker) prepareResponsesForJobs(destinationJob *types.DestinationJobT, respStatusCode int, respBody string) (map[int64]int, map[int64]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *worker) canSendJobToDestination(failedJobOrderKeys map[eventorder.BarrierKey]struct{}, destinationJob *types.DestinationJobT) bool {
	_ = "STUB: not implemented"
	return false
}

// if guaranteeUserEventOrder is false, letting the next jobs pass

// If the destinationJob has come through router transform / batch transform,
// drop the request if it is of a failed user, else send

func (w *worker) updateReqMetrics(respStatusCodes map[int64]int, diagnosisStartTime *time.Time) {
	_ = "STUB: not implemented"
	return
}

func (w *worker) allowRouterAbortedAlert(errorAt string) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *worker) updateAbortedMetrics(destinationID, workspaceId, statusCode, errorAt string) {
	_ = "STUB: not implemented"
	return
}

// To indicate if the failure should be alerted for router-aborted-count

// To specify at which point failure happened

func (w *worker) postStatusOnResponseQ(respStatusCode int, destinationJob *types.DestinationJobT,
	respContentType string, destinationJobMetadata *types.JobMetadataT, status *jobsdb.JobStatusT,
	errorAt string,
) {
	_ = "STUB: not implemented"
	// Enhancing status.ErrorResponse with firstAttemptedAt
	return
}

// destinationJob.Message is the actual payload we tried to send to destination
// destinationJobMetadata.JobT.EventPayload is the router input payload
// capture router output payload in workerJobStatus if reportJobsdbPayload is false
// by default reportJobsdbPayload is true so we capture router input payload in workerJobStatus

// TODO: update default/remove this flag after monitoring the payload sizes

// we always capture router input payload if we see error from destination transformer

// includes ERROR_AT_DEL, ERROR_AT_CUST

// TODO: update after observing the sizes of the payloads

// the job failed

// NOTE: Old key used was "error_response"

// don't delay retry time if retry limit is reached, so that the job can be aborted immediately on the next loop

func (w *worker) sendRouterResponseCountStat(status *jobsdb.JobStatusT, destination *backendconfig.DestinationT, errorAt string) {
	_ = "STUB: not implemented"
	return
}

// To indicate if the failure should be alerted for router-aborted-count

// To specify at which point failure happened

func (w *worker) sendEventDeliveryStat(destinationJobMetadata *types.JobMetadataT, status *jobsdb.JobStatusT, destination *backendconfig.DestinationT) {
	_ = "STUB: not implemented"
	return
}

func (w *worker) sendDestinationResponseToConfigBackend(payload json.RawMessage, destinationJobMetadata *types.JobMetadataT, status *jobsdb.JobStatusT, sourceIDs []string) {
	_ = "STUB: not implemented"
	// Sending destination response to config backend
	return
}

// AvailableSlots returns the number of available slots in the worker's input channel
func (w *worker) AvailableSlots() int { _ = "STUB: not implemented"; return 0 }

// Reserve tries to reserve a slot in the worker's input channel, if available
func (w *worker) ReserveSlot() *reservedSlot { _ = "STUB: not implemented"; return nil }

func (w *worker) trackStuckDelivery() chan struct{} { _ = "STUB: not implemented"; return nil }

// do nothing

func (w *worker) countTransformedJobStatuses(transformType string, transformedJobs []types.DestinationJobT) {
	_ = "STUB: not implemented"
	return
}

// Input Stats for batch/router transformation

// recordTransformerOutgoingRequestMetrics records both transformer_outgoing_request_latency and transformer_outgoing_request_count metrics
// for router deliveries to match transformer's metric structure
func (w *worker) recordTransformerOutgoingRequestMetrics(
	postParams integrations.PostParametersT,
	destinationJob types.DestinationJobT,
	respStatus int,
	duration time.Duration,
) {
	_ = "STUB: not implemented"
	// if EndpointPath is missing, set it to "default" to avoid an empty label value
	return
}

// Get or create cached stats objects using StatsCache

// Record metrics using cached stats

// deliveryMetricLabels represents a unique key for caching stats based on labels
type deliveryMetricLabels struct {
	DestType         string
	TransformerProxy bool
	EndpointPath     string
	StatusCode       int
	RequestMethod    string
	Module           string
	WorkspaceID      string
	DestinationID    string
}

// ToStatTags converts deliveryMetricLabels to stats.Tags for StatsCacheKey interface
func (l deliveryMetricLabels) ToStatTags() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
