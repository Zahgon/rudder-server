package error_index

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	"github.com/rudderlabs/rudder-server/jobsdb"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	"github.com/rudderlabs/rudder-server/utils/types"
)

type ErrorIndexReporter struct {
	ctx              context.Context
	cancel           context.CancelFunc
	g                *errgroup.Group
	log              logger.Logger
	conf             *config.Config
	configSubscriber configSubscriber
	now              func() time.Time
	dbsMu            sync.RWMutex
	dbs              map[string]*handleWithSqlDB

	trigger func() <-chan time.Time

	limiterGroup sync.WaitGroup
	limiter      struct {
		fetch  kitsync.Limiter
		upload kitsync.Limiter
		update kitsync.Limiter
	}

	concurrency config.ValueLoader[int]

	statsFactory stats.Stats
	stats        struct {
		partitionTime stats.Timer
		partitions    stats.Gauge
	}
}

type handleWithSqlDB struct {
	*jobsdb.Handle
	sqlDB *sql.DB
}

func NewErrorIndexReporter(ctx context.Context, log logger.Logger, configSubscriber configSubscriber, conf *config.Config, statsFactory stats.Stats) *ErrorIndexReporter {
	_ = "STUB: not implemented"
	return nil
}

// Report reports the metrics to the errorIndex JobsDB
func (eir *ErrorIndexReporter) Report(ctx context.Context, metrics []*types.PUReportedMetric, tx *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (eir *ErrorIndexReporter) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	_ = "STUB: not implemented"
	return *new(types.ReportingSyncer)
}

// returning a no-op syncer since another go routine has already started syncing

func (eir *ErrorIndexReporter) mainLoop(ctx context.Context, errIndexDB *jobsdb.Handle) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nilerr

func (eir *ErrorIndexReporter) Stop() { _ = "STUB: not implemented"; return }

// resolveJobsDB returns the jobsdb that matches the current transaction (using system information functions)
// https://www.postgresql.org/docs/11/functions-info.html
func (eir *ErrorIndexReporter) resolveJobsDB(tx *Tx) (jobsdb.JobsDB, error) {
	_ = "STUB: not implemented"
	return *new(jobsdb.JobsDB), nil
}

// optimisation, if there is only one jobsdb, return this. If it is the wrong one, it will fail anyway
