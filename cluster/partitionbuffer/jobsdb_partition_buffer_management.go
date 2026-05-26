package partitionbuffer

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/maputil"

	"github.com/rudderlabs/rudder-server/utils/tx"
)

// BufferPartitions marks the provided partition ids to be buffered
func (b *jobsDBPartitionBuffer) BufferPartitions(ctx context.Context, partitionIds []string) error {
	_ = "STUB: not implemented"

	// dedup and sort partitionIds to avoid deadlocks
	return nil
}

func (b *jobsDBPartitionBuffer) RefreshBufferedPartitions(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// getBufferedPartitionsInTx fetches the list of buffered partitions from the database within the provided transaction
func (b *jobsDBPartitionBuffer) getBufferedPartitionsInTx(ctx context.Context, tx *tx.Tx) (*maputil.ReadOnlyMap[string, struct{}], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBufferedPartitionsVersionInTx fetches the version of buffered partitions from the database within the provided transaction
func (b *jobsDBPartitionBuffer) getBufferedPartitionsVersionInTx(ctx context.Context, tx *tx.Tx) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FOR SHARE used as a read lock

// removeBufferPartitions unmarks the provided partition ids as buffered. It assumes that the caller holds the necessary lock on bufferedPartitionsMu
func (b *jobsDBPartitionBuffer) removeBufferPartitions(ctx context.Context, tx *tx.Tx, partitionIds []string) error {
	_ = "STUB: not implemented"

	// dedup and sort partitionIds to avoid deadlocks
	return nil
}

// update in-memory buffered partitions only on successful commit

func (b *jobsDBPartitionBuffer) updateBufferedPartitonsGauge() { _ = "STUB: not implemented"; return }
