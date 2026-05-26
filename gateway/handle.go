package gateway

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-schemas/go/stream"

	"github.com/rudderlabs/rudder-server/app"
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	gwstats "github.com/rudderlabs/rudder-server/gateway/internal/stats"
	"github.com/rudderlabs/rudder-server/gateway/throttler"
	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
	"github.com/rudderlabs/rudder-server/gateway/webhook"
	"github.com/rudderlabs/rudder-server/gateway/webhook/auth"
	"github.com/rudderlabs/rudder-server/jobsdb"
	sourcedebugger "github.com/rudderlabs/rudder-server/services/debugger/source"
	"github.com/rudderlabs/rudder-server/services/rsources"
	"github.com/rudderlabs/rudder-server/utils/types"
)

type messageValidator interface {
	Validate(payload []byte, message *stream.MessageProperties) (bool, error)
}

type Handle struct {
	// dependencies

	config          *config.Config
	logger          logger.Logger
	stats           stats.Stats
	tracer          stats.Tracer
	application     app.App
	backendConfig   backendconfig.BackendConfig
	jobsDB          jobsdb.JobsDB
	rateLimiter     throttler.Throttler
	versionHandler  func(w http.ResponseWriter, r *http.Request)
	rsourcesService rsources.JobService
	sourcehandle    sourcedebugger.SourceDebugger

	// statistic measurements initialised during Setup

	batchSizeStat                                 stats.Measurement
	requestSizeStat                               stats.Measurement
	dbWritesStat                                  stats.Measurement
	dbWorkersBufferFullStat, dbWorkersTimeOutStat stats.Measurement
	bodyReadTimeStat                              stats.Measurement
	addToWebRequestQWaitTime                      stats.Measurement
	addToBatchRequestQWaitTime                    stats.Measurement
	processRequestTime                            stats.Measurement
	emptyAnonIdHeaderStat                         stats.Measurement

	// state initialised during Setup

	diagnosisTicker                *time.Ticker
	userWorkerBatchRequestQ        chan *userWorkerBatchRequestT
	batchUserWorkerBatchRequestQ   chan *batchUserWorkerBatchRequestT
	irh                            RequestHandler
	rrh                            RequestHandler
	webhook                        webhook.WebhookRequestHandler
	whProxy                        http.Handler
	suppressUserHandler            types.UserSuppression
	backgroundCancel               context.CancelFunc
	backgroundWait                 func() error
	userWebRequestWorkers          []*userWebRequestWorkerT
	backendConfigInitialisedChan   chan struct{}
	transformerFeaturesInitialised chan struct{}
	now                            func() time.Time

	// other state

	backendConfigInitialised bool
	inFlightRequests         *sync.WaitGroup

	trackCounterMu    sync.Mutex // protects trackSuccessCount and trackFailureCount
	trackSuccessCount int
	trackFailureCount int

	// backendconfig state
	configSubscriberLock              sync.RWMutex
	writeKeysSourceMap                map[string]backendconfig.SourceT
	sourceIDSourceMap                 map[string]backendconfig.SourceT
	nonEventStreamSources             map[string]bool
	blockedEventsWorkspaceTypeNameMap map[string]map[string]map[string]bool

	conf struct { // configuration parameters
		webPort, maxUserWebRequestWorkerProcess, maxDBWriterProcess                       int
		maxUserWebRequestBatchSize, maxDBBatchSize, maxHeaderBytes, maxConcurrentRequests int
		userWebRequestBatchTimeout, dbBatchWriteTimeout                                   config.ValueLoader[time.Duration]

		maxReqSize                           config.ValueLoader[int]
		enableRateLimit                      config.ValueLoader[bool]
		enableSuppressUserFeature            bool
		diagnosisTickerTime                  time.Duration
		ReadTimeout                          time.Duration
		ReadHeaderTimeout                    time.Duration
		WriteTimeout                         time.Duration
		IdleTimeout                          time.Duration
		allowReqsWithoutUserIDAndAnonymousID config.ValueLoader[bool]
		gwAllowPartialWriteWithErrors        config.ValueLoader[bool]
		webhookV2HandlerEnabled              bool
	}

	// additional internal http handlers
	internalHttpHandlers map[string]http.Handler

	streamMsgValidator func(message *stream.Message) error

	// internal batch validator
	msgValidator messageValidator

	webhookAuthMiddleware *auth.WebhookAuth

	// leakyUploader is an optional function that can be set to handle uploading of invalid payloads
	leakyUploader func(upload msgToUpload)
}

// findUserWebRequestWorker finds and returns the worker that works on a particular `userID`.
// This is done so that requests with a userID keep going to the same worker, which would maintain the consistency in event ordering.
func (gw *Handle) findUserWebRequestWorker(userID string) *userWebRequestWorkerT {
	_ = "STUB: not implemented"
	return nil
}

//	userWebRequestBatcher listens on the `webRequestQ` channel of a worker.
//	Based on `userWebRequestBatchTimeout` and `maxUserWebRequestBatchSize` parameters,
//	batches them together and queues the batch of webreqs in the `batchRequestQ` channel of the worker
//
// Every webRequestWorker keeps doing this concurrently.
func (gw *Handle) userWebRequestBatcher(userWebRequestWorker *userWebRequestWorkerT) {
	_ = "STUB: not implemented"
	return
}

//	userWebRequestWorkerProcess listens on the `batchRequestQ` channel of the webRequestWorker for new batches of webRequests
//	Goes over the webRequests in the batch and filters them out(`rateLimit`, `maxReqSize`).
//	And creates a `jobList` which is then sent to `userWorkerBatchRequestQ` of the gateway and waits for a response
//	from the `dbwriterWorker`s that batch them and write to the db.
//
// Finally sends responses(error) if any back to the webRequests over their `done` channels
func (gw *Handle) userWebRequestWorkerProcess(userWebRequestWorker *userWebRequestWorkerT) {
	_ = "STUB: not implemented"
	return
}

// Saving the event data read from req.request.Body to the splice.
// Using this to send event schema to the config backend.

// no error

// Sending events to config backend

// getJobDataFromRequest parses the request body and returns the jobData or an error if
// - the request payload is invalid JSON
// - the request payload does not correspond to a rudder event
// - the write key is invalid
// - the write key is not allowed to send events
// - the request is rate limited
// - the payload doesn't contain a valid identifier (userId, anonymousId)
// - the payload is too large
// - user in the payload is not allowed to send events (suppressed)
func (gw *Handle) getJobDataFromRequest(req *webRequestT) (jobData *jobFromReq, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Should be function of body

// values retrieved from first event in batch

// tracing

// map to hold modified/filtered events of the batch

// facts about the batch populated as we iterate over events

// calculate version

// skipcq: CRT-A0007

// hashing combination of userIDFromReq + anonIDFromReq, using colon as a delimiter

// In case of "batch" requests, if rate-limiter returns true for LimitReached, just drop the event batch and continue.

func (gw *Handle) isNonIdentifiable(anonIDFromReq, userIDFromReq, eventType string) bool {
	_ = "STUB: not implemented"
	return false
}

// extract or rETL event is allowed without user id and anonymous id

func buildUserID(userIDHeader, anonIDFromReq, userIDFromReq string) string {
	_ = "STUB: not implemented"
	return ""
}

// memoizedIsUserSuppressed is a memoized version of isUserSuppressed
func (gw *Handle) memoizedIsUserSuppressed() func(workspaceID, userID, sourceID string) bool {
	_ = "STUB: not implemented"
	return nil
}

// isUserSuppressed checks if the user is suppressed or not
func (gw *Handle) isUserSuppressed(workspaceID, userID, sourceID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (gw *Handle) memoizedIsEventBlocked() func(workspaceID, sourceID, eventType, eventName string) bool {
	_ = "STUB: not implemented"
	return nil
}

// isEventBlocked checks if an event should be blocked based on workspace settings
func (gw *Handle) isEventBlocked(workspaceID, sourceID, eventType, eventName string) bool {
	_ = "STUB: not implemented"
	// Event blocking is only supported for track events
	return false
}

// getPayload reads the request body and returns the payload's bytes or an error if the payload cannot be read
func (gw *Handle) getPayload(arctx *gwtypes.AuthRequestContext, r *http.Request, reqType string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gw *Handle) getPayloadFromRequest(r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
addToWebRequestQ finds the worker for a particular userID and queues the webrequest with the worker(pushes the req into the webRequestQ channel of the worker).
They are further batched together in userWebRequestBatcher
*/
func (gw *Handle) addToWebRequestQ(_ *http.ResponseWriter, req *http.Request, done chan string, reqType string, requestPayload []byte, arctx *gwtypes.AuthRequestContext) {
	_ = "STUB: not implemented"
	return
}

// If the request comes through proxy, proxy would already send this. So this shouldn't be happening in that case

func (gw *Handle) internalBatchHandlerFunc() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// TODO: add tracing

// Sending events to config backend

type jobWithMetadata struct {
	job                    *jobsdb.JobT
	stat                   gwstats.SourceStat
	skipLiveEventRecording bool
}

type jobParams struct {
	MessageID           string `json:"message_id"`
	SourceID            string `json:"source_id"`
	SourceJobRunID      string `json:"source_job_run_id"`
	SourceTaskRunID     string `json:"source_task_run_id"`
	UserID              string `json:"user_id"`
	TraceParent         string `json:"traceparent"`
	DestinationID       string `json:"destination_id,omitempty"`
	SourceCategory      string `json:"source_category"`
	IsBot               bool   `json:"is_bot,omitempty"`
	BotName             string `json:"bot_name,omitempty"`
	BotURL              string `json:"bot_url,omitempty"`
	BotIsInvalidBrowser bool   `json:"bot_is_invalid_browser,omitempty"`
	BotAction           string `json:"bot_action,omitempty"`
	IsEventBlocked      bool   `json:"is_event_blocked,omitempty"`
}

func (gw *Handle) extractJobsFromInternalBatchPayload(reqType string, body []byte) (
	[]jobWithMetadata, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only needed for live-events

// Upload the raw payload via leaky uploader if available

// only live-events will not work if writeKey is not found

// events suppressed - but return success

func getRudderId(userIDFromReq, anonIDFromReq string) (uuid.UUID, error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

func (gw *Handle) getSourceConfigFromSourceID(sourceID string) (src backendconfig.SourceT, ok bool) {
	_ = "STUB: not implemented"
	return *new(backendconfig.SourceT), false
}

func startStoreJobsWatchdog(writeTimeout, gracePeriod time.Duration, jobsCount int, panicFn func(any)) func() {
	_ = "STUB: not implemented"
	return nil
}

func (gw *Handle) storeJobs(ctx context.Context, jobs []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// rsources stats
