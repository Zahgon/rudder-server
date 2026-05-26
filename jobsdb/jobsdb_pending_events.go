package jobsdb

import (
	"context"
)

type PendingEventsRegistry interface {
	IncreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64)
	DecreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64)
}

// NewPendingEventsJobsDB wraps a JobsDB with pending events metrics collection
func NewPendingEventsJobsDB(jobsDB JobsDB, registry PendingEventsRegistry) JobsDB {
	_ = "STUB: not implemented"
	return *new(JobsDB)
}

type pendingEventsJobsDB struct {
	JobsDB
	registry PendingEventsRegistry
}

func (pejdb *pendingEventsJobsDB) Store(ctx context.Context, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

func (pejdb *pendingEventsJobsDB) StoreInTx(ctx context.Context, tx StoreSafeTx, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// workspaceID -> destType -> destinationID -> count

func (pejdb *pendingEventsJobsDB) UpdateJobStatus(ctx context.Context, statusList []*JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

func (pejdb *pendingEventsJobsDB) UpdateJobStatusInTx(ctx context.Context, tx UpdateSafeTx, statusList []*JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

// workspaceID -> destType -> destinationID -> count
