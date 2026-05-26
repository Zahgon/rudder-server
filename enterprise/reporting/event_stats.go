package reporting

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/stats"

	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	"github.com/rudderlabs/rudder-server/utils/types"
)

type EventStatsReporter struct {
	stats            stats.Stats
	configSubscriber *configSubscriber
}

const EventStream = "event-stream"

func NewEventStatsReporter(configSubscriber *configSubscriber, stats stats.Stats) *EventStatsReporter {
	_ = "STUB: not implemented"
	return nil
}

const EventsProcessedMetricName = "events_processed_total"

func (es *EventStatsReporter) Record(metrics []*types.PUReportedMetric) {
	_ = "STUB: not implemented"
	return
}

func (es *EventStatsReporter) Report(_ context.Context, metrics []*types.PUReportedMetric, tx *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *EventStatsReporter) Stop() { _ = "STUB: not implemented"; return }

func (es *EventStatsReporter) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	_ = "STUB: not implemented"
	return *new(types.ReportingSyncer)
}
