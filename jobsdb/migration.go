package jobsdb

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb/internal/dsindex"
	"github.com/rudderlabs/rudder-server/jobsdb/internal/lock"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
)

// startMigrateDSLoop migrates jobs from src dataset (srcDS) to destination dataset (dest_ds)
// First all the unprocessed jobs are copied over. Then all the jobs which haven't
// completed (state is failed or waiting or waiting_retry or executiong) are copied
// over. Then the status (only the latest) is set for those jobs
func (jd *Handle) startMigrateDSLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (jd *Handle) migrateDSLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (jd *Handle) doMigrateDS(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Fast path: route datasets that have no pending jobs through the async
// dropDSLoop instead of the in-TX postMigrateHandleDS path. This skips the
// dsMigrationLock for empty-source drops so concurrent readers are not blocked.

// addToDropDSList republishes dsList without the completed datasets;
// reuse that snapshot so the second getMigrationList pass inside the
// TX does not re-encounter datasets we just queued for async drop.

// if we reached the probe limit, we know that all datasets before lastProbed are ineligible,
// so we can use lastProbed as the resume point for the next iteration

// no eligible datasets found, nothing to migrate

// cannot run while schema migration is running
// Take the lock and run actual migration

// repeat the check after the dsMigrationLock is acquired to get correct pending jobs count.
// the pending jobs count cannot change after the dsMigrationLock is acquired.
// skip datasets before the first eligible one found in the optimistic check.

// migrate incomplete jobs

// acquire an async lock, as this needs to be released after the transaction commits

type dsWithPendingJobCount struct {
	ds             dataSetT
	numJobsPending int
}

// migrationListResult holds the output of getMigrationList
type migrationListResult struct {
	// migrateFrom is the list of datasets eligible for migration,
	// along with the number of pending (non-terminal) jobs in each.
	migrateFrom []dsWithPendingJobCount
	// pendingJobsCount is the total number of unfinished jobs across all
	// datasets in migrateFrom. When > 0, these jobs need to be copied
	// to a new destination dataset.
	pendingJobsCount int
	// insertBeforeDS is the dataset before which the new (migrated)
	// destination dataset should be created.
	insertBeforeDS dataSetT
	// firstEligible is the parsed dsindex of the first dataset found eligible
	// for migration. nil if no eligible datasets were found.
	firstEligible *dsindex.Index
	// lastProbed is the parsed dsindex of the last dataset that was checked
	// (had checkIfMigrateDS called on it). Used for cross-invocation resumption
	// so the next migration loop iteration can skip already-probed datasets.
	// nil if no datasets were probed.
	lastProbed *dsindex.Index
	// probeLimitReached is true when the loop stopped because migrateDSProbeCount
	// exceeded maxMigrateDSProbe. Only when this is true should lastProbed be used
	// as a cross-invocation resume point, since other exit conditions (idxCheck,
	// found eligible datasets, etc.) don't guarantee that skipped datasets won't
	// become eligible before the next iteration.
	probeLimitReached bool
}

// based on size of given DSs, gives a list of DSs for us to vacuum full status tables
func (jd *Handle) getVacuumFullCandidates(ctx context.Context, dsList []dataSetT) ([]string, error) {
	_ = "STUB: not implemented"
	// get name and it's size of all tables
	return nil, nil
}

// based on an estimate of the rows in DSs, gives a list of DSs for us to cleanup status tables
func (jd *Handle) getCleanUpCandidates(ctx context.Context, dsList []dataSetT) ([]dataSetT, error) {
	_ = "STUB: not implemented"
	// get analyzer estimates for the number of rows(jobs, statuses) in each DS
	return nil, nil
}

// using max ds size if we have no stats for the number of jobs

// based on an estimate cleans up the status tables
func (jd *Handle) cleanupStatusTables(ctx context.Context, dsList []dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

// clean up and vacuum if not present in toVacuumFullMap

// vacuum full

// vacuum analyze

// cleanStatusTable deletes all rows except for the latest status for each job
func (jd *Handle) cleanStatusTable(ctx context.Context, tx *Tx, table string, canBeVacuumed bool) (vacuum bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// getMigrationList returns the list of datasets to migrate from,
// the number of unfinished jobs contained in these datasets
// and the dataset before which the new (migrated) dataset that will hold these jobs needs to be created.
//
// If skipBefore is non-nil, datasets whose parsed index is Less than skipBefore
// are skipped (not checked). This avoids redundant checkIfMigrateDS calls for
// datasets already known to be ineligible.
func (jd *Handle) getMigrationList(dsList []dataSetT, skipBefore *dsindex.Index) (result migrationListResult, err error) {
	_ = "STUB: not implemented"
	return *new(migrationListResult), nil
}

// we don't want `maxDSSize` value to change, during dsList loop

// exempting the last dataset from migration since it is the one being currently written to.

// have another dataset waiting for a pair

// we already know that we'll be migrating another dataset with pending jobs, so can add this one too

// adding this dataset would exceed maxDSSize, leave it for the next migration cycle

// add the current DS as waiting for the next iteration to pickup

// if there was a DS waiting, we should remove it since its next dataset is not eligible for migration

func getColumnConversion(srcType, destType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (jd *Handle) migrateJobsInTx(ctx context.Context, tx *Tx, srcDS, destDS dataSetT) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// find column types first - to differentiate between `text`, `bytea` and `jsonb`

func (jd *Handle) computeNewIdxForIntraNodeMigration(l lock.LockToken, insertBeforeDS dataSetT) (string, error) {
	_ = "STUB: not implemented" // Within the node
	return "", nil
}

func (jd *Handle) postMigrateHandleDS(tx *Tx, migrateFrom []dataSetT) error {
	_ = "STUB: not implemented"
	// Rename datasets before dropping them, so that they can be uploaded to s3
	return nil
}

func computeInsertIdx(beforeIndex, afterIndex string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// checkIfMigrateDS checks when DB is full or DB needs to be migrated.
// We migrate the DB ONCE most of the jobs have been processed (succeeded/aborted)
// Or when the job_status table gets too big because of lots of retries/failures
func (jd *Handle) checkIfMigrateDS(ds dataSetT) (
	migrate, needsPair bool, recordsLeft int, err error,
) {
	_ = "STUB: not implemented"
	return false, false, 0, nil
}
