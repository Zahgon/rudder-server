package batchrouter

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

func (brt *Handle) collectMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

func sendDestStatusStats(batchDestination *Connection, jobStateCounts map[string]int, destType string, isWarehouse bool) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) recordAsyncDestinationDeliveryStatus(sourceID, destinationID string, statusList []*jobsdb.JobStatusT) {
	_ = "STUB: not implemented"
	return
}

// Emit event_delivery_time metric for successful async destination deliveries

// Payload and AttemptNum don't make sense in recording batch router delivery status,
// So they are set to default values.

// emitAsyncEventDeliveryTimeMetrics emits event_delivery_time metrics for successful async destination deliveries
func (brt *Handle) emitAsyncEventDeliveryTimeMetrics(sourceID, destinationID string, statusList []*jobsdb.JobStatusT) {
	_ = "STUB: not implemented"
	// Get the async destination struct to access original job parameters
	return
}

// Process each successful job status to emit event_delivery_time metric

// Only emit metrics for successful deliveries

// Get original job parameters for this job

// Extract receivedAt from original job parameters

// Parse receivedAt time

// Extract source category from original job parameters

// Create and emit the event_delivery_time metric

// Send the timing metric (time from received to delivered)

func (brt *Handle) recordDeliveryStatus(batchDestination Connection, output UploadResult, isWarehouse bool) {
	_ = "STUB: not implemented"
	return
}

// Payload and AttemptNum don't make sense in recording batch router delivery status,
// So they are set to default values.

func (brt *Handle) trackRequestMetrics(batchReqDiagnostics batchRequestMetric) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) recordUploadStats(destination Connection, output UploadResult) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) sendRetryUpdateStats(attempt int) { _ = "STUB: not implemented"; return }

func (brt *Handle) sendQueryRetryStats(attempt int) { _ = "STUB: not implemented"; return }

func (brt *Handle) updateRudderSourcesStats(
	ctx context.Context,
	tx jobsdb.UpdateSafeTx,
	jobs []*jobsdb.JobT,
	jobStatuses []*jobsdb.JobStatusT,
) error {
	_ = "STUB: not implemented"
	return nil
}

// pipelineDelayStats reports the delay of the pipeline as a range:
//
// - max - time elapsed since the first job was created
//
// - min - time elapsed since the last job was created
func (brt *Handle) pipelineDelayStats(partition string, first, last *jobsdb.JobT) {
	_ = "STUB: not implemented"
	return
}
