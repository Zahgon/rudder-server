package jobsdb

import (
	"context"
)

// Management interface for read excluded partitions
type ReadExcludedPartitionsManager interface {
	// AddReadExcludedPartitionIDs adds partition IDs to the excluded read list
	AddReadExcludedPartitionIDs(ctx context.Context, partitionIDs []string) error
	// RemoveReadExcludedPartitionIDs removes partition IDs from the excluded read list
	RemoveReadExcludedPartitionIDs(ctx context.Context, partitionIDs []string) error
}

func (jd *Handle) AddReadExcludedPartitionIDs(ctx context.Context, partitionIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Sort to avoid deadlocks

func (jd *Handle) RemoveReadExcludedPartitionIDs(ctx context.Context, partitionIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Sort to avoid deadlocks

// loadReadExcludedPartitions loads the excluded read partitions from the database
// and populates the excludedReadPartitions map in the JobsDB handle.
func (jd *Handle) loadReadExcludedPartitions() error { _ = "STUB: not implemented"; return nil }

// Partitioning is not enabled; nothing to read.
