package router

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/internal/eventorder"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
)

func (rt *Handle) trackRequestMetrics(reqMetric requestMetric) { _ = "STUB: not implemented"; return }

func (rt *Handle) collectMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

// This lock will ensure we don't send out Track Request while filling up the
// failureMetric struct

func (rt *Handle) updateRudderSourcesStats(
	ctx context.Context,
	tx jobsdb.UpdateSafeTx,
	jobs []*jobsdb.JobT,
	jobStatuses []*jobsdb.JobStatusT,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (rt *Handle) sendRetryStoreStats(attempt int) { _ = "STUB: not implemented"; return }

func (rt *Handle) sendRetryUpdateStats(attempt int) { _ = "STUB: not implemented"; return }

func (rt *Handle) sendQueryRetryStats(attempt int) { _ = "STUB: not implemented"; return }

// pipelineDelayStats reports the delay of the pipeline as a range:
//
// - max - time elapsed since the first job was created
//
// - min - time elapsed since the last job was created
func (rt *Handle) pipelineDelayStats(partition string, first, last *jobsdb.JobT) {
	_ = "STUB: not implemented"
	return
}

// eventOrderDebugInfo provides some debug information for the given orderKey in case of a panic.
// Top 100 job statuses for the given orderKey are returned.
func (rt *Handle) eventOrderDebugInfo(orderKey eventorder.BarrierKey) (res string) {
	_ = "STUB: not implemented"
	return ""
}
