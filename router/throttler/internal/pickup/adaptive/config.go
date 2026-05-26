package adaptive

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

func GetAllEventsWindowConfig(config *config.Config, destType, destinationID string) config.ValueLoader[time.Duration] {
	_ = "STUB: not implemented"
	return nil
}

func GetPerEventWindowConfig(config *config.Config, destType, destinationID, eventType string) config.ValueLoader[time.Duration] {
	_ = "STUB: not implemented"
	return nil
}
