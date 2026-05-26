package partitionbuffer

import (
	"context"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

// StoreEachBatchRetry stores the provided job batches into the appropriate JobsDBs based on their partition buffering status
// It fails all batches with ErrStoreNotSupported if the JobsDBPartitionBuffer does not support Store operations
func (b *jobsDBPartitionBuffer) StoreEachBatchRetry(ctx context.Context, jobBatches [][]*jobsdb.JobT) map[uuid.UUID]string {
	_ = "STUB: not implemented"
	return nil
}

// StoreEachBatchRetryInTx stores the provided job batches into the appropriate JobsDBs based on their partition buffering status within the provided StoreSafeTx
// It fails all batches with ErrStoreNotSupported if the JobsDBPartitionBuffer does not support Store operations
func (b *jobsDBPartitionBuffer) StoreEachBatchRetryInTx(ctx context.Context, tx jobsdb.StoreSafeTx, jobBatches [][]*jobsdb.JobT) (map[uuid.UUID]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no buffered partitions

// we can only include the first job's UUID from each batch in the result map in case of a failure

// store failed, skip buffered jobs store

// store failed, skip buffered jobs store
