package reporting

import (
	"context"
	"database/sql"
	"errors"
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

const ReportsTable = "reports"

var errBlockedByActiveTransactions = errors.New("blocked by active transactions")

const (
	StatReportingMainLoopTime                                = "reporting_client_main_loop_time"
	StatReportingGetReportsTime                              = "reporting_client_get_reports_time"
	StatReportingGetReportsCount                             = "reporting_client_get_reports_count"
	StatReportingGetAggregatedReportsTime                    = "reporting_client_get_aggregated_reports_time"
	StatReportingGetAggregatedReportsCount                   = "reporting_client_get_aggregated_reports_count"
	StatReportingGetMinReportedAtQueryTime                   = "reporting_client_get_min_reported_at_query_time"
	StatReportingGetReportsQueryTime                         = "reporting_client_get_reports_query_time"
	StatReportingVacuumDuration                              = "reporting_vacuum_duration"
	StatReportingMainLoopBlockedDueToActiveTransactionsCount = "reporting_main_loop_blocked_due_to_active_transactions_count"
)

type DefaultReporter struct {
	ctx              context.Context
	cancel           context.CancelFunc
	g                *errgroup.Group
	configSubscriber *configSubscriber
	syncersMu        sync.RWMutex
	syncers          map[string]*types.SyncSource
	log              logger.Logger
	namespace        string

	instanceID                           string
	whActionsOnly                        bool
	sleepInterval                        config.ValueLoader[time.Duration]
	mainLoopSleepInterval                config.ValueLoader[time.Duration]
	dbQueryTimeout                       *config.Reloadable[time.Duration]
	sourcesWithEventNameTrackingDisabled []string
	maxOpenConnections                   int
	maxConcurrentRequests                config.ValueLoader[int]
	vacuumFull                           config.ValueLoader[bool]

	getMinReportedAtQueryTime stats.Measurement
	getReportsQueryTime       stats.Measurement
	stats                     stats.Stats
	maxReportsCountInARequest config.ValueLoader[int]

	eventSamplingEnabled  config.ValueLoader[bool]
	eventSamplingDuration config.ValueLoader[time.Duration]
	eventSampler          event_sampler.EventSampler

	eventNamePrefixLength config.ValueLoader[int]
	eventNameSuffixLength config.ValueLoader[int]
	commonClient          *client.Client

	activeTransactionsByReportedAtMutex sync.Mutex
	activeTransactionsByReportedAt      map[int64]int64
}

func NewDefaultReporter(ctx context.Context, conf *config.Config, log logger.Logger, configSubscriber *configSubscriber, stats stats.Stats) *DefaultReporter {
	_ = "STUB: not implemented"
	return nil
}

// only send reports for wh actions sources if whActionsOnly is configured

func (r *DefaultReporter) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	_ = "STUB: not implemented"
	return *new(types.ReportingSyncer)
}

// returning a no-op syncer since another go routine has already started syncing

func (r *DefaultReporter) GetSyncer(syncerKey string) *types.SyncSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *DefaultReporter) getDBHandle(syncerKey string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *DefaultReporter) getReports(currentMs, aggregationIntervalMin int64, syncerKey string) (reports []*types.ReportByStatus, reportedAt int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// we don't want to flush partial buckets, so we wait for the current bucket to be complete

// Handle rows error

func (r *DefaultReporter) getAggregatedReports(reports []*types.ReportByStatus) []*types.Metric {
	_ = "STUB: not implemented"
	return nil
}

// send reportedAt in milliseconds

func (r *DefaultReporter) emitLagMetric(ctx context.Context, c types.SyncerConfig, lastReportedAtTime *atomic.Time) error {
	_ = "STUB: not implemented"
	// for monitoring reports pileups
	return nil
}

func (r *DefaultReporter) mainLoop(ctx context.Context, c types.SyncerConfig) {
	_ = "STUB: not implemented"
	return
}

// Values should be a factor of 60 or else we will panic, for example 1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60

// if whActionsOnly is true, we only send reports for wh actions sources
// we silently drop all other reports

// if any of errGroup's goroutines fail - don't send anymore requests for this batch

// Use the same aggregationIntervalMin value that was used to query the reports in getReports()

// vacuum the table

func (r *DefaultReporter) vacuum(ctx context.Context, db *sql.DB, tags stats.Tags) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *DefaultReporter) Report(ctx context.Context, metrics []*types.PUReportedMetric, txn *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup when transaction completes (success or failure)

// This should never happen - indicates a bug in transaction tracking

func (r *DefaultReporter) getTags(label string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

func (r *DefaultReporter) Stop() { _ = "STUB: not implemented"; return }
