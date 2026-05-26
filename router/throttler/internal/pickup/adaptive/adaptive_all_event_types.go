package adaptive

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// NewAllEventTypesThrottler constructs a new adaptive throttler for all event types of a destination
func NewAllEventTypesThrottler(destType, destinationID string, algorithm Algorithm, limiter Limiter, c *config.Config, stat stats.Stats, log Logger) *throttler {
	_ = "STUB: not implemented"
	return nil
}
