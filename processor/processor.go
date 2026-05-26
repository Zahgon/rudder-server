package processor

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/cachettl"
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/enterprise/trackedusers"
	"github.com/rudderlabs/rudder-server/internal/enricher"
	"github.com/rudderlabs/rudder-server/jobsdb"
	transformerutils "github.com/rudderlabs/rudder-server/processor/internal/transformer"
	"github.com/rudderlabs/rudder-server/processor/isolation"
	"github.com/rudderlabs/rudder-server/processor/transformer"
	"github.com/rudderlabs/rudder-server/processor/types"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	transformationdebugger "github.com/rudderlabs/rudder-server/services/debugger/transformation"
	deduptypes "github.com/rudderlabs/rudder-server/services/dedup/types"
	"github.com/rudderlabs/rudder-server/services/fileuploader"
	"github.com/rudderlabs/rudder-server/services/rmetrics"
	"github.com/rudderlabs/rudder-server/services/rsources"
	transformerFeaturesService "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/services/transientsource"
	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/utils/tracing"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	reportingtypes "github.com/rudderlabs/rudder-server/utils/types"
)

// Custom type definitions for deeply nested map
type (
	DestinationID      string
	SourceID           string
	ConsentProviderKey string
)

type (
	ConsentProviderMap map[ConsentProviderKey]GenericConsentManagementProviderData
	DestConsentMap     map[DestinationID]ConsentProviderMap
	SourceConsentMap   map[SourceID]DestConsentMap
)

const (
	MetricKeyDelimiter    = "!<<#>>!"
	UserTransformation    = "USER_TRANSFORMATION"
	DestTransformation    = "DEST_TRANSFORMATION"
	EventFilter           = "EVENT_FILTER"
	sourceCategoryWebhook = "webhook"
)

func NewHandle(c *config.Config, transformerClients transformer.TransformerClients) *Handle {
	_ = "STUB: not implemented"
	return nil
}

type sourceObserver interface {
	ObserveSourceEvents(source *backendconfig.SourceT, events []types.TransformerEvent)
}

type trackedUsersReporter interface {
	ReportUsers(ctx context.Context, reports []*trackedusers.UsersReport, tx *Tx) error
	GenerateReportsFromJobs(jobs []*jobsdb.JobT, sourceIdFilter map[string]bool) []*trackedusers.UsersReport
}

// Handle is a handle to the processor module
type Handle struct {
	conf               *config.Config
	tracer             *tracing.Tracer
	backendConfig      backendconfig.BackendConfig
	transformerClients transformer.TransformerClients

	gatewayDB                  jobsdb.JobsDB
	routerDB                   jobsdb.JobsDB
	batchRouterDB              jobsdb.JobsDB
	eventSchemaDB              jobsdb.JobsDB
	archivalDB                 jobsdb.JobsDB
	pendingEventsRegistry      rmetrics.PendingEventsRegistry
	logger                     logger.Logger
	enrichers                  []enricher.PipelineEnricher
	dedup                      deduptypes.Dedup
	reporting                  reportingtypes.Reporting
	reportingEnabled           bool
	backgroundCtx              context.Context
	backgroundWait             func() error
	backgroundCancel           context.CancelFunc
	statsFactory               stats.Stats
	stats                      processorStats
	payloadLimit               config.ValueLoader[int64]
	jobsDBCommandTimeout       config.ValueLoader[time.Duration]
	jobdDBQueryRequestTimeout  config.ValueLoader[time.Duration]
	jobdDBMaxRetries           config.ValueLoader[int]
	transientSources           transientsource.Service
	fileuploader               fileuploader.Provider
	utSamplingFileManager      filemanager.FileManager
	storeSamplingFileManager   filemanager.FileManager
	mirrorFilteredCache        *cachettl.Cache[string, bool]
	rsourcesService            rsources.JobService
	transformerFeaturesService transformerFeaturesService.FeaturesService
	destDebugger               destinationdebugger.DestinationDebugger
	transDebugger              transformationdebugger.TransformationDebugger
	isolationStrategy          isolation.Strategy
	limiter                    struct {
		read         kitsync.Limiter
		preprocess   kitsync.Limiter
		srcHydration kitsync.Limiter
		pretransform kitsync.Limiter
		utransform   kitsync.Limiter
		dtransform   kitsync.Limiter
		store        kitsync.Limiter
	}
	config struct {
		isolationMode                             isolation.Mode
		mainLoopTimeout                           time.Duration
		enablePipelining                          bool
		pipelineBufferedItems                     int
		subJobSize                                int
		pipelinesPerPartition                     int
		pingerSleep                               config.ValueLoader[time.Duration]
		readLoopSleep                             config.ValueLoader[time.Duration]
		maxLoopSleep                              config.ValueLoader[time.Duration]
		storeTimeout                              config.ValueLoader[time.Duration]
		maxEventsToProcess                        config.ValueLoader[int]
		sourceIdDestinationMap                    map[string][]backendconfig.DestinationT
		sourceIdSourceMap                         map[string]backendconfig.SourceT
		workspaceLibrariesMap                     map[string]backendconfig.LibrariesT
		oneTrustConsentCategoriesMap              map[string][]string
		connectionConfigMap                       map[connection]backendconfig.Connection
		ketchConsentCategoriesMap                 map[string][]string
		genericConsentManagementMap               SourceConsentMap
		batchDestinations                         []string
		configSubscriberLock                      sync.RWMutex
		enableDedup                               bool
		transformTimesPQLength                    int
		captureEventNameStats                     config.ValueLoader[bool]
		transformerURL                            string
		GWCustomVal                               string
		asyncInit                                 *misc.AsyncInit
		eventSchemaV2Enabled                      bool
		archivalEnabled                           config.ValueLoader[bool]
		eventAuditEnabled                         map[string]bool
		credentialsMap                            map[string][]types.Credential
		nonEventStreamSources                     map[string]bool
		enableConcurrentStore                     config.ValueLoader[bool]
		userTransformationMirroringSanitySampling config.ValueLoader[float64]
		userTransformationMirroringFireAndForget  config.ValueLoader[bool]
		userTransformMirrorURL                    string
		pythonTransformMirrorURL                  string
		mirrorFilterCacheTTL                      time.Duration
		pythonTransformConfig                     transformerutils.PythonTransformConfig
		userTransformationMirroringBlockedIDs     config.ValueLoader[[]string]
		storeSamplerEnabled                       config.ValueLoader[bool]
		archiveInPreProcess                       bool
	}

	drainConfig struct {
		jobRunIDs config.ValueLoader[[]string]
	}

	namespace  string
	instanceID string

	adaptiveLimit func(int64) int64
	storePlocker  kitsync.PartitionLocker

	sourceObservers      []sourceObserver
	trackedUsersReporter trackedUsersReporter
}
type processorStats struct {
	statGatewayDBR                func(partition string) stats.Measurement
	statGatewayDBW                func(partition string) stats.Measurement
	statDBR                       func(partition string) stats.Measurement
	statDBW                       func(partition string) stats.Measurement
	validateEventsTime            func(partition string) stats.Measurement // TODO: stop using it in dashboards and delete
	statNumRequests               func(partition string) stats.Measurement
	statNumEvents                 func(partition string) stats.Measurement
	statDBWriteRouterPayloadBytes func(partition string) stats.Measurement // TODO: stop using it in dashboards and delete
	statDBWriteBatchPayloadBytes  func(partition string) stats.Measurement // TODO: stop using it in dashboards and delete
	statDestNumOutputEvents       func(partition string) stats.Measurement
	statBatchDestNumOutputEvents  func(partition string) stats.Measurement
	trackedUsersReportGeneration  func(partition string) stats.Measurement // TODO: stop using it in dashboards and delete

	statReadStageCount         func(partition string) stats.Measurement
	statPretransformStageCount func(partition string) stats.Measurement
	statPreprocessStageCount   func(partition string) stats.Measurement
	statSrcHydrationStageCount func(partition string) stats.Measurement
	statUtransformStageCount   func(partition string) stats.Measurement
	statDtransformStageCount   func(partition string) stats.Measurement
	statStoreStageCount        func(partition string) stats.Measurement

	utMirroringEqualResponses            func(partition, transformationID string) stats.Measurement
	utMirroringDifferentResponses        func(partition, transformationID string) stats.Measurement
	utMirroringFilteredResponses         func(partition, transformationID string) stats.Measurement
	utMirroringBlockedByTransformationID func(partition, transformationID string) stats.Measurement
	utMirroringDatetimeForgivenResponses func(partition, transformationID string) stats.Measurement
}
type DestStatT struct {
	numEvents               stats.Measurement
	numOutputSuccessEvents  stats.Measurement
	numOutputFailedEvents   stats.Measurement
	numOutputFilteredEvents stats.Measurement
	transformTime           stats.Measurement
}

type ParametersT struct {
	SourceID                string `json:"source_id"`
	SourceName              string `json:"source_name"`
	DestinationID           string `json:"destination_id"`
	ReceivedAt              string `json:"received_at"`
	TransformAt             string `json:"transform_at"`
	MessageID               string `json:"message_id"`
	GatewayJobID            int64  `json:"gateway_job_id"`
	SourceTaskRunID         string `json:"source_task_run_id"`
	SourceJobID             string `json:"source_job_id"`
	SourceJobRunID          string `json:"source_job_run_id"`
	EventName               string `json:"event_name"`
	EventType               string `json:"event_type"`
	SourceDefinitionID      string `json:"source_definition_id"`
	DestinationDefinitionID string `json:"destination_definition_id"`
	SourceCategory          string `json:"source_category"`
	RecordID                any    `json:"record_id"`
	WorkspaceId             string `json:"workspaceId"`
	TraceParent             string `json:"traceparent"`
	ConnectionID            string `json:"connection_id"`
}

type MetricMetadata struct {
	sourceID                string
	destinationID           string
	sourceTaskRunID         string
	sourceJobID             string
	sourceJobRunID          string
	sourceDefinitionID      string
	destinationDefinitionID string
	sourceCategory          string
	transformationID        string
	transformationVersionID string
	trackingPlanID          string
	trackingPlanVersion     int
}

// procErrorJob wraps a jobsdb.JobT with the original parsed events,
// avoiding an expensive marshal/unmarshal round-trip through EventPayload.
type procErrorJob struct {
	*jobsdb.JobT
	events []types.SingularEventT
}

func procErrorJobs(jobs []procErrorJob) []*jobsdb.JobT { _ = "STUB: not implemented"; return nil }

type NonSuccessfulTransformationMetrics struct {
	failedJobs       []procErrorJob
	failedMetrics    []*reportingtypes.PUReportedMetric
	failedCountMap   map[string]int64
	filteredJobs     []procErrorJob
	filteredMetrics  []*reportingtypes.PUReportedMetric
	filteredCountMap map[string]int64
}

type (
	SourceIDT string
)

func buildStatTags(sourceID, workspaceID string, destination *backendconfig.DestinationT, transformationType string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) newUserTransformationStat(
	sourceID, workspaceID string, destination *backendconfig.DestinationT, mirroring bool,
) *DestStatT {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) newDestinationTransformationStat(sourceID, workspaceID, transformAt string, destination *backendconfig.DestinationT) *DestStatT {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) newEventFilterStat(sourceID, workspaceID string, destination *backendconfig.DestinationT) *DestStatT {
	_ = "STUB: not implemented"
	return nil
}

// Setup initializes the module
func (proc *Handle) Setup(
	ctx context.Context,
	backendConfig backendconfig.BackendConfig,
	gatewayDB, routerDB, batchRouterDB,
	eventSchemaDB, archivalDB jobsdb.JobsDB,
	reporting reportingtypes.Reporting,
	transientSources transientsource.Service,
	fileuploader fileuploader.Provider,
	rsourcesService rsources.JobService,
	transformerFeaturesService transformerFeaturesService.FeaturesService,
	destDebugger destinationdebugger.DestinationDebugger,
	transDebugger transformationdebugger.TransformationDebugger,
	enrichers []enricher.PipelineEnricher,
	trackedUsersReporter trackedusers.UsersReporter,
	pendingEventsRegistry rmetrics.PendingEventsRegistry,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Stats

// periodically publish a zero counter for ensuring that stuck processing pipeline alert
// can always detect a stuck processor

func (proc *Handle) setupReloadableVars() { _ = "STUB: not implemented"; return }

// Start starts this processor's main loops.
func (proc *Handle) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// limiters

// pinger loop

// waiting for init group

// proceed

// waiting for transformer features

// proceed

func (proc *Handle) activePartitions(ctx context.Context) []string {
	_ = "STUB: not implemented"
	return nil
}

// TODO: retry?

func (proc *Handle) Shutdown() { _ = "STUB: not implemented"; return }

func (proc *Handle) loadConfig() { _ = "STUB: not implemented"; return }

// If isolation mode is not none, we need to reduce the values for some of the config variables to more sensible defaults

// Enable dedup of incoming events by default

// GWCustomVal is used as a key in the jobsDB customval column

func (proc *Handle) loadReloadableConfig(defaultPayloadLimit int64, defaultMaxEventsToProcess int) {
	_ = "STUB: not implemented"
	return
}

// Capture event name as a tag in event level stats

// UserTransformation mirroring settings

type connection struct {
	sourceID, destinationID string
}

func (proc *Handle) backendConfigSubscriber(ctx context.Context) { _ = "STUB: not implemented"; return }

func (proc *Handle) getWorkspaceLibraries(workspaceID string) backendconfig.LibrariesT {
	_ = "STUB: not implemented"
	return *new(backendconfig.LibrariesT)
}

func (proc *Handle) getConnectionConfig(conn connection) backendconfig.Connection {
	_ = "STUB: not implemented"
	return *new(backendconfig.Connection)
}

func (proc *Handle) getSourceBySourceID(sourceId string) (*backendconfig.SourceT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (proc *Handle) getNonEventStreamSources() map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) getEnabledDestinations(sourceId, destinationName string) []backendconfig.DestinationT {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) getBackendEnabledDestinationTypes(sourceId string) map[string]backendconfig.DestinationDefinitionT {
	_ = "STUB: not implemented"
	return nil
}

func getTimestampFromEvent(event types.SingularEventT, field string, defaultTimestamp time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func enhanceWithTimeFields(event *types.TransformerEvent, singularEvent types.SingularEventT, receivedAt time.Time) {
	_ = "STUB: not implemented"
	// set timestamp skew based on timestamp fields from SDKs
	return
}

// use existing timestamp if it exists in the event, add new timestamp otherwise

// calculate new timestamp using the formula
// timestamp = receivedAt - (sentAt - originalTimestamp)

// set all timestamps in RFC3339 format

func (proc *Handle) singularEventMetadata(singularEvent types.SingularEventT, userID, partitionID string, jobId int64, receivedAt time.Time, source *backendconfig.SourceT, eventParams types.EventParams) *types.Metadata {
	_ = "STUB: not implemented"
	return nil
}

// job metadata

// event-related metadata

// source metadata

// retl metadata

// other metadata

func getKeyFromSourceAndDest(srcID, destID string) string { _ = "STUB: not implemented"; return "" }

func getSourceAndDestIDsFromKey(key string) (sourceID, destID string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (proc *Handle) recordEventDeliveryStatus(jobsByDestID map[string][]procErrorJob) {
	_ = "STUB: not implemented"
	return
}

func (proc *Handle) getTransformerEvents(
	response types.Response,
	commonMetaData *types.Metadata,
	eventsByMessageID map[string]types.SingularEventWithReceivedAt,
	destination *backendconfig.DestinationT,
	connection backendconfig.Connection,
	inPU, pu string,
) (
	[]types.TransformerEvent,
	[]*reportingtypes.PUReportedMetric,
	map[string]int64,
	map[string]MetricMetadata,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Update metrics maps

// Update metadata with updated event name before reporting

//
// fill non common metadata from response, effectively this strips out:
// - tracking plan information after tracking plan validation
// - user transformation information after user transformation (except for message IDs)

// job metadata

// event metadata

// retl metadata

// other metadata

// user transformation message ids

// REPORTING - START

// REPORTING - END

func (proc *Handle) updateMetricMaps(
	// countMetadataMap provides metadata context for getDiffMetrics to create detailed PUReportedMetric objects
	// storing rich context during event processing for diff metric reporting
	countMetadataMap map[string]MetricMetadata,

	// countMap accumulates event counts by unique key combinations, feeding into getDiffMetrics to capture diff metrics
	countMap map[string]int64,

	// connectionDetailsMap stores source-destination relationship data that becomes PUReportedMetric.ConnectionDetails
	// for constructing complete reporting metrics with workspace, source, and destination context
	connectionDetailsMap map[string]*reportingtypes.ConnectionDetails,

	// statusDetailsMap captures processing outcomes including error details, validation violations, and sample payloads
	// that become PUReportedMetric.StatusDetail
	statusDetailsMap map[string]map[string]*reportingtypes.StatusDetail,

	event *types.TransformerResponse,
	status, stage string,
	payload func() json.RawMessage,
	eventsByMessageID map[string]types.SingularEventWithReceivedAt,
) {
	_ = "STUB: not implemented"
	return
}

// create status details for each validation error
// single event can have multiple validation errors of same type

// create status details for a whole event

// this is called defensive programming... :(

func (proc *Handle) getNonSuccessfulMetrics(
	response types.Response,
	inputEvents []types.TransformerEvent,
	commonMetaData *types.Metadata,
	eventsByMessageID map[string]types.SingularEventWithReceivedAt,
	inPU, pu string,
) *NonSuccessfulTransformationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func procFilteredCountStat(destType, pu, statusCode string) { _ = "STUB: not implemented"; return }

func procErrorCountsStat(destType, pu, statusCode string) { _ = "STUB: not implemented"; return }

func (proc *Handle) getTransformationMetrics(
	transformerResponses []types.TransformerResponse,
	state string,
	commonMetaData *types.Metadata,
	eventsByMessageID map[string]types.SingularEventWithReceivedAt,
	metadataByMessageID map[string]*types.Metadata,
	inPU, pu string,
) ([]procErrorJob, []*reportingtypes.PUReportedMetric, map[string]int64) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// REPORTING - START

// REPORTING - END

func generateConnectionID(s1, s2 string) string { _ = "STUB: not implemented"; return "" }

func (proc *Handle) updateSourceEventStatsDetailed(event types.SingularEventT, sourceId string) {
	_ = "STUB: not implemented"
	// Any panics in this function are captured and ignore sending the stat
	return
}

// nolint:forbidigo

func getDiffMetrics(
	inPU, pu string,
	inCountMetadataMap map[string]MetricMetadata,
	successCountMetadataMap map[string]MetricMetadata,
	inCountMap, successCountMap, failedCountMap, filteredCountMap map[string]int64,
	statFactory stats.Stats,
) []*reportingtypes.PUReportedMetric {
	_ = "STUB: not implemented"
	return nil
}

// Helper function to create metric and record stats

// Process input metrics

// Process success metrics if enabled

// Skip if this key was already processed from input metrics

type dupStatKey struct {
	sourceID string
}

func (proc *Handle) eventAuditEnabled(workspaceID string) bool {
	_ = "STUB: not implemented"
	return false
}

type preTransformationMessage struct {
	partition                     string
	subJobs                       subJob
	eventSchemaJobsBySourceId     map[SourceIDT][]*jobsdb.JobT
	archivalJobs                  []*jobsdb.JobT
	connectionDetailsMap          map[string]*reportingtypes.ConnectionDetails
	statusDetailsMap              map[string]map[string]*reportingtypes.StatusDetail
	enricherStatusDetailsMap      map[string]map[string]*reportingtypes.StatusDetail
	botManagementStatusDetailsMap map[string]map[string]*reportingtypes.StatusDetail
	eventBlockingStatusDetailsMap map[string]map[string]*reportingtypes.StatusDetail
	destFilterStatusDetailMap     map[string]map[string]*reportingtypes.StatusDetail
	reportMetrics                 []*reportingtypes.PUReportedMetric
	totalEvents                   int
	groupedEventsBySourceId       map[SourceIDT][]types.TransformerEvent
	eventsByMessageID             map[string]types.SingularEventWithReceivedAt
	jobIDToSpecificDestMapOnly    map[int64]string
	statusList                    []*jobsdb.JobStatusT
	jobList                       []*jobsdb.JobT
	sourceDupStats                map[dupStatKey]int
	dedupKeys                     map[string]struct{}
	srcHydrationEnabledMap        map[SourceIDT]bool
}

func (proc *Handle) preprocessStage(partition string, subJobs subJob, delay time.Duration) (*srcHydrationMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Each block we receive from a client has a bunch of
// requests. We parse the block and take out individual
// requests, call the destination specific transformation
// function and create jobs for them.
// Transformation is called for a batch of jobs at a time
// to speed-up execution.

// Event count for performance stat monitoring

// map of jobID to destinationID: for messages that needs to be delivered to a specific destinations only

// dummy event for metrics purposes only

// REPORTING - BOT_MANAGEMENT metrics - START

// reset status code to 0 because transformerEvent is reused for other metrics

// REPORTING - BOT_MANAGEMENT metrics - END

// REPORTING - EVENT_BLOCKING metrics - START

// reset status code to 0 because transformerEvent is reused for other metrics

// REPORTING - EVENT_BLOCKING metrics - END

// schemas enabled

// TODO: could use source.SourceDefinition.Category instead?

// archival enabled&&

// REPORTING - GATEWAY metrics - START

// Pass nil for countMetadataMap and countMap as we don't want to capture diff metrics for bot enricher

// REPORTING - GATEWAY metrics - END
// Getting all the destinations which are enabled for this event.
// Event will be dropped if no valid destination is present
// if empty destinationID is passed in this fn all the destinations for the source are validated
// else only passed destinationID will be validated

// REPORTING - DESTINATION_FILTER filtered metrics - START

// REPORTING - DESTINATION_FILTER filtered metrics - END

// fill in tracking plan details

func (proc *Handle) pretransformStage(partition string, preTrans *preTransformationMessage) (*transformationMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// REPORTING - START

// REPORTING - END

// TRACKING PLAN - START
// Placing the trackingPlan validation filters here.
// Else further down events are duplicated by destId, so multiple validation takes places for same event

// Appending validatedReportMetrics to reportMetrics

// TRACKING PLAN - END

// The below part further segregates events by sourceID and DestinationID.

// Adding a singular event multiple times if there are multiple destinations of same type

// At the TP flow we are not having destination information, so adding it here.

// We have at-least one event so marking it good

func (proc *Handle) storeEventSchemaJobs(ctx context.Context, eventSchemaJobs []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) storeArchiveJobs(ctx context.Context, archivalJobs []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

type SourcePipelineSteps struct {
	srcHydration           bool
	trackingPlanValidation bool
}
type sourceIDPipelineSteps map[SourceIDT]SourcePipelineSteps

type transformationMessage struct {
	ctx           context.Context
	groupedEvents map[string][]types.TransformerEvent

	srcPipelineSteps             sourceIDPipelineSteps
	eventsByMessageID            map[string]types.SingularEventWithReceivedAt
	uniqueMessageIdsBySrcDestKey map[string]map[string]struct{}
	reportMetrics                []*reportingtypes.PUReportedMetric
	statusList                   []*jobsdb.JobStatusT
	sourceDupStats               map[dupStatKey]int
	dedupKeys                    map[string]struct{}

	totalEvents int

	hasMore             bool
	rsourcesStats       rsources.StatsCollector
	trackedUsersReports []*trackedusers.UsersReport
}

type userTransformData struct {
	ctx                           context.Context
	userTransformAndFilterOutputs map[string]userTransformAndFilterOutput
	reportMetrics                 []*reportingtypes.PUReportedMetric
	statusList                    []*jobsdb.JobStatusT
	sourceDupStats                map[dupStatKey]int
	dedupKeys                     map[string]struct{}
	trackedUsersReports           []*trackedusers.UsersReport

	totalEvents int
	start       time.Time

	hasMore       bool
	rsourcesStats rsources.StatsCollector
	traces        map[string]stats.Tags
}

func (proc *Handle) userTransformStage(partition string, in *transformationMessage) *userTransformData {
	_ = "STUB: not implemented"
	return nil
}

// Now do the actual transformation. We call it in batches, once
// for each destination ID

func (proc *Handle) destinationTransformStage(partition string, in *userTransformData) *storeMessage {
	_ = "STUB: not implemented"
	return nil
}

// Start worker goroutines

// Start a goroutine to close the channel after all workers are done

// Collect results

type storeMessage struct {
	ctx                 context.Context
	trackedUsersReports []*trackedusers.UsersReport
	statusList          []*jobsdb.JobStatusT
	destJobs            []*jobsdb.JobT
	batchDestJobs       []*jobsdb.JobT
	droppedJobs         []*jobsdb.JobT

	procErrorJobsByDestID map[string][]procErrorJob
	routerDestIDs         []string

	reportMetrics  []*reportingtypes.PUReportedMetric
	sourceDupStats map[dupStatKey]int
	dedupKeys      map[string]struct{}

	totalEvents int
	start       time.Time

	hasMore       bool
	rsourcesStats rsources.StatsCollector
	traces        map[string]stats.Tags
}

func (sm *storeMessage) merge(subJob *storeMessage) { _ = "STUB: not implemented"; return }

func (proc *Handle) sendRetryStoreStats(attempt int) { _ = "STUB: not implemented"; return }

func (proc *Handle) sendRetryUpdateStats(attempt int) { _ = "STUB: not implemented"; return }

func (proc *Handle) sendQueryRetryStats(attempt int) { _ = "STUB: not implemented"; return }

func (proc *Handle) storeStage(partition string, pipelineIndex int, in *storeMessage) {
	_ = "STUB: not implemented"
	return
}

// Use pipelineIndex in the lock pKey to avoid contention when multiple pipelines are running (each pipeline processes its own exclusive partition of userIDs)

// lock early to avoid deadlocks due to connection pool exhaustion

// XX: Need to do this in a transaction

// rsources stats

// Only one goroutine can store to a router destination at a time, otherwise we may have different transactions
// committing at different timestamps which can cause events with lower jobIDs to appear after events with higher ones.
// For that purpose, before storing, we lock the relevant destination IDs (in sorted order to avoid deadlocks).

// rsources stats

// this will publish stats for all sources involved in this batch

// this will publish rudder source stats only for dropped jobs involved in this batch.
// It needs to be called after rsourcesStats.Publish, otherwise, it may cause a deadlock
// e.g. the batch contains two keys and there are dropped jobs only for the second key but not for the first one

func getStoreSamplingUploader(conf *config.Config, log logger.Logger) (*filemanager.S3Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type destTransformOutput struct {
	reportMetrics   []*reportingtypes.PUReportedMetric
	destJobs        []*jobsdb.JobT
	batchDestJobs   []*jobsdb.JobT
	errorsPerDestID map[string][]procErrorJob
	routerDestIDs   map[string]struct{}
	droppedJobs     []*jobsdb.JobT
}

// userTransformAndFilterOutput holds the data passed between preprocessing and postprocessing steps
type userTransformAndFilterOutput struct {
	eventsToTransform     []types.TransformerEvent
	commonMetaData        *types.Metadata
	reportMetrics         []*reportingtypes.PUReportedMetric
	procErrorJobsByDestID map[string][]procErrorJob
	droppedJobs           []*jobsdb.JobT
	eventsByMessageID     map[string]types.SingularEventWithReceivedAt
	srcAndDestKey         string
	response              types.Response
	transformAt           string
}

func (proc *Handle) userTransformAndFilter(ctx context.Context, partition, srcAndDestKey string, eventList []types.TransformerEvent, srcPipelineSteps sourceIDPipelineSteps, eventsByMessageID map[string]types.SingularEventWithReceivedAt, uniqueMessageIdsBySrcDestKey map[string]map[string]struct{}) userTransformAndFilterOutput {
	_ = "STUB: not implemented"
	return *new(userTransformAndFilterOutput)
}

// REPORTING - START

// Grouping events by sourceid + destinationid + jobruniD + eventName + eventType to find the count

// REPORTING - END

// Send to custom transformer only if the destination has a transformer enabled

// mirroring go routine

// Let's create a copy of the response because it may be subject to changes later (see getTransformerEvents).
// We want to block while creating a copy to avoid a race condition.
// This shouldn't have a big impact on processor latency unless the sanity sampling is a very high percentage.

// filtered metric already bumped and transformation already cached as mirror-filtered
// in the mirroring go routine

// Apply the same Marshal→Unmarshal round-trip to the mirror response
// so both sides go through identical serialization. Without this,
// the primary is compared after a JSON round-trip (which can normalize
// typed nils, omitempty zero values, etc.) while the mirror is compared
// directly from HTTP deserialization — causing false-positive diffs.

// adding more data to help with debugging

// Cannot upload, we should just report the issue with no diff

// Upload the diff file and the eventListCopy via the file manager
// * eventListCopy is the original eventList that was passed to the transformer
// * diff contains the difference between the transformer response and the mirrored response

// REPORTING - START

// successCountMap will be inCountMap for filtering events based on supported event types

// REPORTING - END
// for the next step in the pipeline

// Check for overrides through env

// Filtering events based on the supported message types - START

// REPORTING - START

// successCountMap will be inCountMap for destination transform

// REPORTING - END

// Filtering events based on the supported message types - END

func (proc *Handle) destTransform(ctx context.Context, data userTransformAndFilterOutput) destTransformOutput {
	_ = "STUB: not implemented"
	return *new(destTransformOutput)
}

// Destination transformation - START
// Send to transformer only if is
// a. transformAt is processor
// OR
// b. transformAt is router and transformer doesn't support router transform

// REPORTING - PROCESSOR metrics - START

// Update metrics maps

// REPORTING - PROCESSOR metrics - END

// Save the JSON in DB. This is what the router uses

// Should be a valid JSON since it's our transformation, but we handle it anyway

// Need to replace UUID his with messageID from client

// read source_id from metadata that is replayed back from transformer
// in case of custom transformations metadata of first event is returned along with all events in session
// source_id will be same for all events belong to same user in a session

// If the response from the transformer does not have userID in metadata, setting userID to random-uuid.
// This is done to respect findWorker logic in router.

func (proc *Handle) saveDroppedJobs(ctx context.Context, droppedJobs []*jobsdb.JobT, tx *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// each dropped job should have a unique jobID in the scope of the batch

func (proc *Handle) isUserTransformMirroringEnabled(eventList []types.TransformerEvent, partition string) (bool, chan types.Response) {
	_ = "STUB: not implemented"
	return false, nil
}

// Mirroring is disabled.

// Check if this is a Python transformation and apply version-based filtering

// Check if the mirror URL is configured for the transformation language

// Mirroring is enabled in fire&forget mode and no sanity checks

// Determine if mirroring should be enabled based on sampling percentage.
// Sampling percentage precision can be with two decimals like 12.34%

// Sanity checks were enabled but the random value was less than the sampling percentage.
// Disabling mirroring altogether.

func ConvertToFilteredTransformerResponse(
	events []types.TransformerEvent,
	filter bool,
	drainFunc func(types.TransformerEvent) (bool, string),
) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// filter unsupported message types

// drain events

// filter unsupported message types

// filter unsupported message events

// add to FailedEvents

// allow event

func (proc *Handle) getJobsStage(ctx context.Context, partition string) jobsdb.JobsResult {
	_ = "STUB: not implemented"
	return *new(jobsdb.JobsResult)
}

// keep trying to get unprocessed jobs while no jobs are returned because ds limits are being reached

func (proc *Handle) markExecuting(ctx context.Context, partition string, jobs []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// Mark the jobs as executing

// handlePendingGatewayJobs is checking for any pending gateway jobs (failed and unprocessed), and routes them appropriately
// Returns true if any job is handled, otherwise returns false.
func (proc *Handle) handlePendingGatewayJobs(partition string) bool {
	_ = "STUB: not implemented"
	return false
}

// context is used for tracing

// `jobSplitter` func Splits the read Jobs into sub-batches after reading from DB to process.
// `subJobMerger` func merges the split jobs into a single batch before writing to DB.
// So, to keep track of sub-batch we have `hasMore` variable.
// each sub-batch has `hasMore`. If, a sub-batch is the last one from the batch it's marked as `false`, else `true`.
type subJob struct {
	ctx           context.Context
	subJobs       []*jobsdb.JobT
	hasMore       bool
	rsourcesStats rsources.StatsCollector
}

func (proc *Handle) jobSplitter(
	ctx context.Context, jobs []*jobsdb.JobT, rsourcesStats rsources.StatsCollector, //nolint:unparam
) []subJob {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) crashRecover() { _ = "STUB: not implemented"; return }

func (proc *Handle) updateSourceStats(sourceStats map[dupStatKey]int, bucket string) {
	_ = "STUB: not implemented"
	return
}

func (proc *Handle) isReportingEnabled() bool { _ = "STUB: not implemented"; return false }

func (proc *Handle) updateRudderSourcesStats(ctx context.Context, tx jobsdb.StoreSafeTx, jobs []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

func filterConfig(eventCopy *types.TransformerEvent) { _ = "STUB: not implemented"; return }

func (*Handle) getLimiterPriority(partition string) kitsync.LimiterPriorityValue {
	_ = "STUB: not implemented"
	return *new(kitsync.LimiterPriorityValue)
}

// check if event has eligible destinations to send to
//
// event will be dropped if no destination is found
func (proc *Handle) isDestinationAvailable(event types.SingularEventT, sourceId, destinationID string) bool {
	_ = "STUB: not implemented"
	return false
}

// pipelineDelayStats reports the delay of the pipeline as a range:
//
// - max - time elapsed since the first job was created
//
// - min - time elapsed since the last job was created
func (proc *Handle) pipelineDelayStats(partition string, first, last *jobsdb.JobT) {
	_ = "STUB: not implemented"
	return
}

func (proc *Handle) countPendingEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore context cancellation

// shouldSample sampling percentage precision can be with two decimals like 12.34%.
// To sample everything use 100. To sample nothing use 0.
func shouldSample(samplingPercentage float64) bool { _ = "STUB: not implemented"; return false }

// getUTSamplingUploader can be completely removed once we get rid of UT sampling
func getUTSamplingUploader(conf *config.Config, log logger.Logger) (*filemanager.S3Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
