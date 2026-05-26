package config

import (
	"github.com/rudderlabs/rudder-go-kit/config"
)

// ThrottlerPerEventTypeEnabled returns a reloadable boolean indicating whether per event type throttling is enabled for the given destination type and ID.
func ThrottlerPerEventTypeEnabled(c *config.Config, destType, destinationID string) *config.Reloadable[bool] {
	_ = "STUB: not implemented"
	return nil
}
