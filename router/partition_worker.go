package router

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

// newPartitionWorker creates a worker that is responsible for picking up jobs for a single partition (none, workspace, destination).
// A partition worker uses multiple workers internally to process the jobs that are being picked up asynchronously.
func newPartitionWorker(ctx context.Context, rt *Handle, partition string) *partitionWorker {
	_ = "STUB: not implemented"
	return nil
}

// Sampler: refreshes each worker buffer's cached capacity once per second so the
// hot path (worker.AvailableSlots) doesn't have to invoke the calculator on every
// call. Stops when samplerCancel is called from Stop().

type partitionWorker struct {
	// dependencies
	rt     *Handle
	logger logger.Logger

	// configuration
	partition string

	// state
	ctx                  context.Context
	g                    *errgroup.Group         // group against which all the workers are spawned
	samplerCancel        context.CancelFunc      // stops the buffer-capacity sampler goroutine
	pickupBatchSizeGauge GaugeWithLastValue[int] // gauge to track the pickup batch size used in the last pickup iteration
	workers              []*worker               // workers that are responsible for processing the jobs

	pickupCount   int  // number of jobs picked up by the workers in the last iteration
	limitsReached bool // whether the limits were reached in the last iteration
}

// Work picks up jobs for the partitioned worker and returns whether it worked or not
func (pw *partitionWorker) Work() bool { _ = "STUB: not implemented"; return false }

// the following stats are used to track the total time taken for the pickup process and the number of jobs picked up

// sleep only if we worked and we didn't reach the limits

// SleepDurations returns the min and max sleep durations for the partitioned worker while not working
func (pw *partitionWorker) SleepDurations() (min, max time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

// Stop stops the partitioned worker by closing the input channel of all its internal workers and waiting for them to finish
func (pw *partitionWorker) Stop() { _ = "STUB: not implemented"; return }
