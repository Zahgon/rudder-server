package static

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// NewAllEventTypesThrottler constructs a new static throttler for all event types of a destination
func NewAllEventTypesThrottler(destType, destinationID string, limiter Limiter, c *config.Config, stat stats.Stats, log Logger) *throttler {
	_ = "STUB: not implemented"
	return nil
}

// key is destinationID
