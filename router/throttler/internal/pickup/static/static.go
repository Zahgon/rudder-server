package static

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/router/throttler/internal/types"
)

// NewThrottler constructs a new static throttler that can switch between all event types and per event type throttling.
func NewThrottler(destType, destinationID, eventType string, limiter Limiter, config *config.Config, stat stats.Stats, log logger.Logger) types.PickupThrottler {
	_ = "STUB: not implemented"
	return *new(types.PickupThrottler)
}
