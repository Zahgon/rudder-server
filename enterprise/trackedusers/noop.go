package trackedusers

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/jobsdb"
	txn "github.com/rudderlabs/rudder-server/utils/tx"
)

type NoopDataCollector struct{}

func NewNoopDataCollector() *NoopDataCollector { _ = "STUB: not implemented"; return nil }

func (n *NoopDataCollector) ReportUsers(context.Context, []*UsersReport, *txn.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NoopDataCollector) GenerateReportsFromJobs([]*jobsdb.JobT, map[string]bool) []*UsersReport {
	_ = "STUB: not implemented"
	return nil
}

func (n *NoopDataCollector) MigrateDatabase(string, *config.Config) error {
	_ = "STUB: not implemented"
	return nil
}
