package partitionbuffer

import (
	"context"
)

// FlushBufferedPartitions flushes the buffered data for the provided partition ids to the database and unmarks them as buffered.
func (b *jobsDBPartitionBuffer) FlushBufferedPartitions(ctx context.Context, partitions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// use priority pool for flush operations

// block for validation and marking partitions as flushing

// only keep partitions that are actually buffered

// ensure we unmark partitions as flushing at the end regardless of success or failure

// move in batches until we stop reaching limits

// timeout reached, break out to switchover

// switchover

// moveBufferedPartitionsConcurrently distributes the given partition IDs across the requested number of goroutines, each calling moveBufferedPartitions on its share.
// It returns the total moved count, whether any goroutine reached its limits, and the first encountered error.
// All goroutines must finish before this returns, so the caller may safely re-read the concurrency setting between iterations.
func (b *jobsDBPartitionBuffer) moveBufferedPartitionsConcurrently(ctx context.Context, partitionIDs []string, batchSize int, payloadSize int64, concurrency int) (count int, limitsReached bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// moveBufferedPartitions moves a batch of buffered jobs to the primary JobsDB for the given partition IDs. It returns whether any limits were reached during the fetch.
// If limits were reached, the caller should call this method again to move more data.
func (b *jobsDBPartitionBuffer) moveBufferedPartitions(ctx context.Context, partitionIDs []string, batchSize int, payloadSize int64) (count int, limitsReached bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// create job statuses

func (b *jobsDBPartitionBuffer) switchoverBufferedPartitions(ctx context.Context, partitionIDs []string, batchSize int, payloadSize int64) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// disable idle_in_transaction_session_timeout for the duration of this transaction, since it may take long to move all remaining data

// mark partitions as unbuffered in the database early, for holding the global lock

// refresh DS list in the buffer read JobsDB so that we are confident that we are going to be moving all remaining data from the buffer

// move any remaining buffered data

// mark partitions as unbuffered in the database late
