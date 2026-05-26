package flusher

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/enterprise/reporting/client"
	"github.com/rudderlabs/rudder-server/enterprise/reporting/flusher/aggregator"
)

type Flusher struct {
	log logger.Logger

	db                 *sql.DB
	maxOpenConnections int

	aggregator   aggregator.Aggregator
	commonClient *client.Client

	instanceId string
	table      string
	module     string

	sleepInterval                       config.ValueLoader[time.Duration]
	flushWindow                         config.ValueLoader[time.Duration]
	recentExclusionWindow               config.ValueLoader[time.Duration]
	batchSizeFromDB                     config.ValueLoader[int]
	aggressiveFlushEnabled              config.ValueLoader[bool]
	lagThresholdForAggresiveFlushInMins config.ValueLoader[time.Duration]
	vacuumThresholdDeletedRows          config.ValueLoader[int]
	vacuumThresholdBytes                config.ValueLoader[int64]
	deletedRows                         int
	lastVacuum                          time.Time
	vacuumFull                          config.ValueLoader[bool]
	vacuumInterval                      config.ValueLoader[time.Duration]

	minConcurrentRequests config.ValueLoader[int]
	maxConcurrentRequests config.ValueLoader[int]
	batchSizeToReporting  config.ValueLoader[int]

	stats                   stats.Stats
	minReportedAtQueryTimer stats.Measurement
	aggReportsTimer         stats.Measurement
	sendReportsTimer        stats.Measurement
	deleteReportsTimer      stats.Measurement
	vacuumReportsTimer      stats.Measurement
	concurrentRequests      stats.Measurement
	flushLag                stats.Measurement

	commonTags stats.Tags
}

func NewFlusher(db *sql.DB, log logger.Logger, stats stats.Stats, conf *config.Config, table string, commonClient *client.Client, aggregator aggregator.Aggregator, module string) (*Flusher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Flusher) CleanUp() error { _ = "STUB: not implemented"; return nil }

func (f *Flusher) initCommonTags() { _ = "STUB: not implemented"; return }

func (f *Flusher) initStats(tags map[string]string) { _ = "STUB: not implemented"; return }

func (f *Flusher) getStart(ctx context.Context) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (f *Flusher) getLag(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (f *Flusher) ShouldFlushAggressively(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// flush is the main logic for flushing data.
func (f *Flusher) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Emit the lag metric

// 1. Get the time range to flush

// 2. Aggregate reports. We have different aggregators for in-app and in-db aggregation

// 3. Flush aggregated reports

// 4. Delete reports

// 5. Vacuum the table

// Since we have hourly/daily/monthly aggregates on Reporting Service, we want the window to be within same hour
// Don't consider most recent data where there are inserts happening
// Always try to consider full window of flush interval or till current hour
func (f *Flusher) getRange(ctx context.Context, currentUTC time.Time) (start, end time.Time, isFullFlushWindow bool, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), false, nil
}

func (f *Flusher) aggregate(ctx context.Context, start, end time.Time) ([]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Flusher) getConcurrency(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

func (f *Flusher) send(ctx context.Context, aggReports []json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flusher) delete(ctx context.Context, minReportedAt, maxReportedAt time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flusher) vacuum(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
