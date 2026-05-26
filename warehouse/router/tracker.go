package router

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

// cronTracker Track the status of the staging file whether it has reached the terminal state or not for every warehouse
// we pick the staging file which is oldest within the range NOW() - 2 * syncFrequency and NOW() - 3 * syncFrequency
// and checks if the corresponding upload has reached the terminal state or not.
// If the upload has not reached the terminal state, then we send a gauge metric with value 1 else 0
func (r *Router) cronTracker(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:nilerr

func (r *Router) retryTrackSync(ctx context.Context, warehouse *model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) trackSync(ctx context.Context, warehouse *model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) isWithinExcludeWindow(warehouse *model.Warehouse) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Router) getOldestStagingFile(ctx context.Context, warehouse *model.Warehouse) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (r *Router) calculateTimeWindow(warehouse *model.Warehouse) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *Router) checkUploadStatus(ctx context.Context, warehouse *model.Warehouse, createdAt time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Router) recordUploadMissingMetric(warehouse *model.Warehouse, exists bool) {
	_ = "STUB: not implemented"
	return
}
