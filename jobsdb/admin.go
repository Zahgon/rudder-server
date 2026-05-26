package jobsdb

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb/internal/lock"
)

/*
Ping returns health check for pg database
*/
func (jd *Handle) Ping() error { _ = "STUB: not implemented"; return nil }

/*
DeleteExecuting deletes events whose latest job state is executing.
This is only done during recovery, which happens during the server start.
*/
func (jd *Handle) DeleteExecuting() { _ = "STUB: not implemented"; return }

// deleteJobStatus deletes the latest status of a batch of jobs
func (jd *Handle) deleteJobStatus() { _ = "STUB: not implemented"; return }

func (jd *Handle) deleteJobStatusDSInTx(txHandler transactionHandler, ds dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

/*
FailExecuting fails events whose latest job state is executing.

This is only done during recovery, which happens during the server start.
*/
func (jd *Handle) FailExecuting() { _ = "STUB: not implemented"; return }

// failExecuting sets the state of the executing jobs to failed
func (jd *Handle) failExecuting() { _ = "STUB: not implemented"; return }

func (jd *Handle) failExecutingDSInTx(txHandler transactionHandler, ds dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

// doCleanup performs cleanup of old jobs and journal entries. It requires a ds list lock token to ensure that the list does not change during the cleanup process.
func (jd *Handle) doCleanup(ctx context.Context, l lock.LockToken) error {
	_ = "STUB: not implemented"
	return nil
}

// 2. cleanup journal

func (jd *Handle) abortOldJobs(ctx context.Context, dsList []dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}
