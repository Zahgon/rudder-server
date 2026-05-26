package reporting

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/enterprise/reporting/client"
	"github.com/rudderlabs/rudder-server/enterprise/reporting/event_sampler"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	"github.com/rudderlabs/rudder-server/utils/types"
)

const (
	ErrorDetailReportsTable = "error_detail_reports"
	groupKeyDelimitter      = "$::$"
)

var ErrorDetailReportsColumns = []string{
	"workspace_id",
	"namespace",
	"instance_id",
	"source_definition_id",
	"source_id",
	"destination_definition_id",
	"destination_id",
	"dest_type",
	"pu", // reportedBy
	"reported_at",
	"count",
	"status_code",
	"event_type",
	"error_code",
	"error_message",
	"sample_response",
	"sample_event",
	"event_name",
}

// ErrorReportingStats manages all stats for error reporting
type ErrorReportingStats struct {
	// Basic stats
	ReportTime                   stats.Measurement
	ReportingLag                 stats.Measurement
	ErrorDetailReportingFailures stats.Measurement
	HttpRequest                  stats.Measurement
	VacuumDuration               stats.Measurement

	// Error normalizer stats
	NormalizerCleanupTime stats.Measurement
}

// NewErrorReportingStats creates a new stats manager
func NewErrorReportingStats(statsInstance stats.Stats) *ErrorReportingStats {
	_ = "STUB: not implemented"
	return nil
}

type ErrorDetailReporter struct {
	ctx              context.Context
	cancel           context.CancelFunc
	g                *errgroup.Group
	configSubscriber *configSubscriber
	syncersMu        sync.RWMutex
	syncers          map[string]*types.SyncSource
	log              logger.Logger
	namespace        string

	instanceID            string
	sleepInterval         config.ValueLoader[time.Duration]
	mainLoopSleepInterval config.ValueLoader[time.Duration]
	maxConcurrentRequests config.ValueLoader[int]
	maxOpenConnections    int
	vacuumFull            config.ValueLoader[bool]

	errorDetailExtractor *ExtractorHandle
	errorNormalizer      ErrorNormalizer

	// Stats management
	statsManager *ErrorReportingStats

	// Tagged stats (created dynamically with tags)
	minReportedAtQueryTime      stats.Measurement
	errorDetailReportsQueryTime stats.Measurement
	eventSamplingEnabled        config.ValueLoader[bool]
	eventSamplingDuration       config.ValueLoader[time.Duration]
	eventSampler                event_sampler.EventSampler
	groupingThreshold           config.ValueLoader[float64]

	stats  stats.Stats
	config *config.Config

	commonClient *client.Client
}

func NewErrorDetailReporter(
	ctx context.Context,
	configSubscriber *configSubscriber,
	statsInstance stats.Stats,
	conf *config.Config,
) *ErrorDetailReporter {
	_ = "STUB: not implemented"
	return nil
}

// Initialize stats manager

func (edr *ErrorDetailReporter) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	_ = "STUB: not implemented"
	return *new(types.ReportingSyncer)
}

// returning a no-op syncer since another go routine has already started syncing

func (edr *ErrorDetailReporter) emitLagMetric(ctx context.Context, lastReportedAtTime *atomic.Time) error {
	_ = "STUB: not implemented"
	// for monitoring reports pileups
	return nil
}

func (edr *ErrorDetailReporter) GetSyncer(syncerKey string) *types.SyncSource {
	_ = "STUB: not implemented"
	return nil
}

func shouldReport(metric types.PUReportedMetric) bool { _ = "STUB: not implemented"; return false }

func (edr *ErrorDetailReporter) Report(ctx context.Context, metrics []*types.PUReportedMetric, txn *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract error details and filter metrics that should be reported

// Early exit if no metrics to report

// Group errors by connection details

// Normalize errors after grouping (modifies connectionGroups in place)

// Merge metrics by error messages within the same connection group

// Write grouped errors to database

func (edr *ErrorDetailReporter) extractErrorDetailsAndFilterMetrics(metrics []*types.PUReportedMetric) []*types.EDReportsDB {
	_ = "STUB: not implemented"
	return nil
}

// EXTRACT ERROR DETAILS - This is where error messages are created from SampleResponse

// Convert to EDReportsDB instead of deep copying

func (edr *ErrorDetailReporter) normalizeErrors(ctx context.Context, connectionGroups map[types.ErrorDetailGroupKey][]*types.EDReportsDB) map[types.ErrorDetailGroupKey][]*types.EDReportsDB {
	_ = "STUB: not implemented"
	return nil
}

// Update all metrics in the group with the normalized error message

func (edr *ErrorDetailReporter) writeGroupedErrors(ctx context.Context, groups map[types.ErrorDetailGroupKey][]*types.EDReportsDB, txn *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Calculate total count for the group

// Record stats

func (ed *ErrorDetailReporter) IsPIIReportingDisabled(workspaceID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (edr *ErrorDetailReporter) migrate(c types.SyncerConfig) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: shall we use separate env ?

func (edr *ErrorDetailReporter) extractErrorDetails(sampleResponse string, statTags map[string]string, destType string) types.ErrorDetails {
	_ = "STUB: not implemented"
	return *new(types.ErrorDetails)
}

func (edr *ErrorDetailReporter) getDBHandle(syncerKey string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (edr *ErrorDetailReporter) getTags(label string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

// Sending metrics to Reporting service --- STARTS
func (edr *ErrorDetailReporter) mainLoop(ctx context.Context, c types.SyncerConfig) {
	_ = "STUB: not implemented"
	return
}

// In infinite loop
// Get Reports
// Aggregate
// Send in a separate go-routine
// Delete in a separate go-routine

// if any of errGroup's goroutines fail - don't send anymore requests for this batch

// sqlStatement := fmt.Sprintf(`DELETE FROM %s WHERE reported_at = %d`, ErrorDetailReportsTable, reportedAt)

// vacuum error_reports_details table

func (edr *ErrorDetailReporter) vacuum(ctx context.Context, dbHandle *sql.DB, _ stats.Tags) error {
	_ = "STUB: not implemented"
	return nil
}

func (edr *ErrorDetailReporter) getReports(ctx context.Context, currentMs int64, syncerKey string) ([]*types.EDReportsDB, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

/*
	"workspace_id",
	"namespace",
	"instance_id",
	"source_definition_id",
	"source_id",
	"destination_definition_id",
	"destination_id",
	"pu",
	"reported_at",
	"count",
	"status_code",
	"event_type",
	"error_code",
	"error_message",
	"dest_type",
	"sample_response",
	"sample_event",
	"event_name",
*/

func (edr *ErrorDetailReporter) aggregate(reports []*types.EDReportsDB) []*types.EDMetric {
	_ = "STUB: not implemented"
	return nil
}

func (edr *ErrorDetailReporter) Stop() { _ = "STUB: not implemented"; return }
