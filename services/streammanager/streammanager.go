package streammanager

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

// NewProducer delegates the call to the appropriate based on parameter destination for creating producer
func NewProducer(destination *backendconfig.DestinationT, opts common.Opts) (common.StreamProducer, error) {
	_ = "STUB: not implemented"
	return *new(common.StreamProducer), nil
}

// 404, "No provider configured for StreamManager", ""
