package aggregator

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/segmentio/go-hll"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

const tableName = `tracked_users_reports`

type TrackedUsersInAppAggregator struct {
	db    *sql.DB
	stats stats.Stats

	reportsCounter    stats.Measurement
	aggReportsCounter stats.Measurement
}

func NewTrackedUsersInAppAggregator(db *sql.DB, s stats.Stats, conf *config.Config, module string) *TrackedUsersInAppAggregator {
	_ = "STUB: not implemented"
	return nil
}

func (a *TrackedUsersInAppAggregator) Aggregate(ctx context.Context, start, end time.Time) (jsonReports []json.RawMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *TrackedUsersInAppAggregator) decodeHLL(encoded string) (*hll.Hll, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalReports(aggReportsMap map[string]*TrackedUsersReport) ([]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
