package router

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/router/types"
)

type workerBatchLoop struct {
	ctx                      context.Context                                             // context for managing the lifecycle of the loop
	jobsBatchTimeout         config.ValueLoader[time.Duration]                           // timeout for processing jobs in a batch
	noOfJobsToBatchInAWorker config.ValueLoader[int]                                     // maximum number of jobs to batch in a worker before processing
	inputCh                  <-chan *workerJob                                           // channel to receive jobs for processing
	enableBatching           bool                                                        // whether to enable batching of jobs
	batchTransform           func(routerJobs []types.RouterJobT) []types.DestinationJobT // function to transform router jobs into destination jobs in batch mode
	transform                func(routerJobs []types.RouterJobT) []types.DestinationJobT // function to transform router jobs into destination jobs in non-batch mode
	process                  func(destinationJobs []types.DestinationJobT)               // function to process the transformed destination jobs
	acceptWorkerJob          func(workerJob workerJob) *types.RouterJobT                 // function to accept a worker job and return a router job if applicable
	throughputStat           stats.Histogram                                             // stat to record throughput of the loop
}

// runLoop processes jobs from the input channel, batching them if it is enabled.
func (wl *workerBatchLoop) runLoop() { _ = "STUB: not implemented"; return }

// prevent division by zero

// reset routerJobs for the next batch

// reset the timeout

// process any remaining jobs in the batch
// input channel is closed, exit the loop

// Context is done, we are stopping the worker
// it is fine to ignore any pending jobs in the input channel (executing state)
// since when the router restarts, it will mark them as failed and reprocess them.
// What is important is to stop as soon as possible.

// process the current batch if batching is not enabled, transform for the current job is not at router and there are pending jobs in the batch
// (scenario where we are switching from router to processor transformation)

// process the batch if it reaches the limit

// job was not accepted to enter the batch, but was processed, so we need to capture its throughput

// process any remaining jobs in the batch
