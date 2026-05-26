package static

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// NewPerEventTypeThrottler constructs a new static throttler for a specific event type of a destination
func NewPerEventTypeThrottler(destType, destinationID, eventType string, limiter Limiter, c *config.Config, stat stats.Stats, log Logger) *throttler {
	_ = "STUB: not implemented"
	return nil
}

// key is destinationID + ":" + eventType

// static cost for per-event-type throttler: cost was originally introduced to address rate limit differences between different event types, so not needed when using per-event-type throttler
