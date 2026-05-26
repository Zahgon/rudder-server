package event_sampler

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

const (
	StatReportingEventSamplerRequestsTotal   = "reporting_event_sampler_requests_total"
	StatReportingEventSamplerRequestDuration = "reporting_event_sampler_request_duration_seconds"
	StatReportingBadgerDBSize                = "reporting_badger_db_size_bytes"
)

type StatsCollector struct {
	module      string
	stats       stats.Stats
	getCounter  stats.Measurement
	putCounter  stats.Measurement
	getDuration stats.Measurement
	putDuration stats.Measurement
}

func NewStatsCollector(eventSamplerType, module string, statsFactory stats.Stats) *StatsCollector {
	_ = "STUB: not implemented"
	return nil
}

func (sc *StatsCollector) RecordGet() { _ = "STUB: not implemented"; return }

func (sc *StatsCollector) RecordPut() { _ = "STUB: not implemented"; return }

func (sc *StatsCollector) RecordGetDuration(start time.Time) { _ = "STUB: not implemented"; return }

func (sc *StatsCollector) RecordPutDuration(start time.Time) { _ = "STUB: not implemented"; return }

func (sc *StatsCollector) RecordBadgerDBSize(usageType string, size int64) {
	_ = "STUB: not implemented"
	return
}

func getTags(eventSamplerType, module, operation string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
