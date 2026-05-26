package forwarder

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type BaseForwarder struct {
	terminalErrFn func(error) // function to call when a terminal error occurs
	log           logger.Logger
	stat          stats.Stats
	jobsDB        jobsdb.JobsDB

	cancel context.CancelFunc // cancel function for the Start context (used to stop all goroutines during Stop)
	g      *errgroup.Group    // errgroup for the Start context (used to wait for all goroutines to exit)

	conf struct {
		pickupSize                int           // number of jobs to pickup in a single query
		loopSleepTime             time.Duration // time to sleep between each loop
		jobsDBQueryRequestTimeout time.Duration // timeout for jobsdb query
		jobsDBMaxRetries          int           // max retries for jobsdb query
		jobsDBPayloadSize         int64         // max payload size for jobsdb query
	}
}

// LoadMetaData loads the metadata required by the forwarders
func (bf *BaseForwarder) LoadMetaData(terminalErrFn func(error), schemaDB jobsdb.JobsDB, log logger.Logger, config *config.Config, stat stats.Stats) {
	_ = "STUB: not implemented"
	return
}

// GetJobs is an abstraction over the GetUnprocessed method of the jobsdb which includes retries
func (bf *BaseForwarder) GetJobs(ctx context.Context) ([]*jobsdb.JobT, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// MarkJobStatuses is an abstraction over the UpdateJobStatusInTx method of the jobsdb which includes retries
func (bf *BaseForwarder) MarkJobStatuses(ctx context.Context, statusList []*jobsdb.JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSleepTime returns the sleep time based on the limitReached flag
func (bf *BaseForwarder) GetSleepTime(limitReached bool) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (bf *BaseForwarder) sendQueryRetryStats(attempt int) { _ = "STUB: not implemented"; return }
