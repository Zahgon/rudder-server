package gateway

import (
	"context"
)

// TrackRequestMetrics updates the track counters (success and failure counts)
func (gw *Handle) TrackRequestMetrics(errorMessage string) { _ = "STUB: not implemented"; return }

// collectMetrics collects the gateway metrics and sends them using diagnostics
func (gw *Handle) collectMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }
