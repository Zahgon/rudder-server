package jobiterator

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type IteratorOptFn func(*Iterator)

// WithMaxQueries sets the maximum number of queries that can be made to the jobsDB
// for fetching more jobs.
func WithMaxQueries(maxQueries int) IteratorOptFn {
	_ = "STUB: not implemented"
	return *new(IteratorOptFn)
}

// WithDiscardedPercentageTolerance sets the discarded percentage tolerance,
// i.e. the maximum percentage of discarded jobs that can be tolerated without further querying jobsDB.
func WithDiscardedPercentageTolerance(discardedPercentageTolerance int) IteratorOptFn {
	_ = "STUB: not implemented"
	return *new(IteratorOptFn)
}

// Iterator is a job iterator with support for fetching more than the original set of jobs requested,
// in case some of these jobs get discarded, according to the configured discarded percentage tolerance.
type Iterator struct {
	stopped                      bool
	params                       jobsdb.GetQueryParams
	maxQueries                   int
	discardedPercentageTolerance int
	getJobsFn                    func(context.Context, jobsdb.GetQueryParams, jobsdb.MoreToken) (*jobsdb.MoreJobsResult, error)
	state                        struct {
		// running iterator state
		jobs        []*jobsdb.JobT
		idx         int
		previousJob *jobsdb.JobT

		// closed indicates whether the iterator has reached the end or not
		closed bool

		stats IteratorStats

		// for the next query
		discarded         int
		jobsLimit         int
		continuationToken jobsdb.MoreToken
	}
}

// IteratorStats holds statistics about an iterator
type IteratorStats struct {
	// QueryCount is the number of queries made to the jobsDB
	QueryCount int
	// TotalJobs is the total number of jobs queried
	TotalJobs int
	// DiscardedJobs is the number of jobs discarded
	DiscardedJobs int
	// LimitsReached indicates whether the iterator reached the limits of the jobsDB while querying
	LimitsReached bool
}

// New returns a new job iterator
//   - params: jobsDB query parameters
//   - getJobsFn: the function to fetch jobs from jobsDB
//   - opts: optional iterator options
func New(params jobsdb.GetQueryParams, getJobsFn func(context.Context, jobsdb.GetQueryParams, jobsdb.MoreToken) (*jobsdb.MoreJobsResult, error), opts ...IteratorOptFn) *Iterator {
	_ = "STUB: not implemented"
	return nil
}

// HasNext returns true when there are more jobs to be fetched by Next(), false otherwise.
func (ji *Iterator) HasNext() bool { _ = "STUB: not implemented"; return false }

// we have more jobs in the current batch

// iterator has reached the end

// nothing left to fetch

// don't continue if discarded jobs are within tolerance limits

// try to fetch some more jobs

// for getting the first page, keep trying to get jobs while no jobs are returned because ds limits are being reached

// reset state

// no more jobs, iterator has reached the end

// Next returns the next job to be processed.
// Never call Next() without checking HasNext() first.
func (ji *Iterator) Next() *jobsdb.JobT { _ = "STUB: not implemented"; return nil }

// Discard is called when a job is not processed.
// By discarding a job we are allowing the iterator to fetch more jobs from jobsDB.
func (ji *Iterator) Discard(_ *jobsdb.JobT) { _ = "STUB: not implemented"; return }

// Stop marks the iterator as stopped and prevents it from fetching more jobs.
// Any jobs that are not yet processed will be discarded. The number of discarded jobs is returned.
func (ji *Iterator) Stop() int { _ = "STUB: not implemented"; return 0 }

// StopQueries forces the iterator to stop fetching more jobs from jobsDB by setting the maximum number of queries to the current query count.
func (ji *Iterator) StopQueries() { _ = "STUB: not implemented"; return }

// Stats returns the statistics of the iterator, including the number of queries made, total jobs fetched, discarded jobs, and whether limits were reached.
func (ji *Iterator) Stats() IteratorStats { _ = "STUB: not implemented"; return *new(IteratorStats) }
