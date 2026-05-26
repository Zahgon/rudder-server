package batchrouter

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/router/batchrouter/circuitbreaker"
)

// newWorker creates a new worker for the provided partition.
func newWorker(partition string, logger logger.Logger, brt *Handle) *worker {
	_ = "STUB: not implemented"
	return nil
}

type worker struct {
	partition string
	logger    logger.Logger
	brt       *Handle
	cb        circuitbreaker.CircuitBreaker
	pw        *PartitionWorker
}

// Work retrieves jobs from batch router for the worker's partition and processes them,
// grouped by destination and in parallel.
// The function returns when processing completes and the return value is true if at least 1 job was processed,
// false otherwise.
func (w *worker) Work() bool {
	_ = "STUB: not implemented"
	// Check if circuit breaker is open before doing any work
	return false
}

// routeJobsToBuffer sends jobs to appropriate channels in the job buffer
func (w *worker) scheduleJobs(destinationJobs *DestinationJobs) { _ = "STUB: not implemented"; return }

// Organize jobs by destination and source

// Process jobs and check for drain conditions

// Check if source exists for the destination

// Check standard drain conditions

// Consolidate drain checks: either source not found OR drainer says so

// Job is not drained, prepare it for buffering

// Mark the drainList jobs as Aborted

// Mark jobs as executing in a single batch operation

// Now that all statuses are updated, we can safely send jobs to channels

// SleepDurations returns the min and max sleep durations for the worker when idle, i.e when [Work] returns false.
func (w *worker) SleepDurations() (min, max time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

// Stop is no-op for this worker since the worker is not running any goroutine internally.
func (w *worker) Stop() { _ = "STUB: not implemented"; return }
