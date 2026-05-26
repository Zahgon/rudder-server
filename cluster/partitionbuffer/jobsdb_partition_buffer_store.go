package partitionbuffer

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-server/jobsdb"
	utilstx "github.com/rudderlabs/rudder-server/utils/tx"
)

// sentinel error to indicate stale buffered partitions that need refreshing
var errStaleBufferedPartitions = errors.New("stale buffered partitions")

// Store stores the provided jobs into the appropriate JobsDBs based on their partition buffering status
// It returns ErrStoreNotSupported if the JobsDBPartitionBuffer does not support Store operations
func (b *jobsDBPartitionBuffer) Store(ctx context.Context, jobList []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// WithStoreSafeTx acquires a read lock on buffered partitions early and starts a StoreSafeTx on the primary write JobsDB.
// It returns ErrStoreNotSupported if the JobsDBPartitionBuffer does not support Store operations
func (b *jobsDBPartitionBuffer) WithStoreSafeTx(ctx context.Context, fn func(tx jobsdb.StoreSafeTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to check for stale version

// get the version difference

// stale version

// refresh buffered partitions and retry

// StoreInTx stores the provided jobs into the appropriate JobsDBs based on their partition buffering status within the provided StoreSafeTx
func (b *jobsDBPartitionBuffer) StoreInTx(ctx context.Context, tx jobsdb.StoreSafeTx, jobList []*jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// no buffered partitions

// split jobs and store accordingly
// make sure we have the latest version of buffered partitions if that is necessary

// splitJobs is a convenience wrapper around spitJobBatches for a single batch of jobs
func (b *jobsDBPartitionBuffer) splitJobs(jobList []*jobsdb.JobT) (primary, buffered []*jobsdb.JobT) {
	_ = "STUB: not implemented"
	return nil, nil
}

// versionDiff returns the difference between the version of buffered partitions in the database and the current version in memory
func (b *jobsDBPartitionBuffer) versionDiff(ctx context.Context, tx *utilstx.Tx) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// failAllBatches returns a map of job UUIDs to the provided error message for the first job of each batch
func failAllBatches(jobBatches [][]*jobsdb.JobT, err error) map[uuid.UUID]string {
	_ = "STUB: not implemented"
	return nil
}
