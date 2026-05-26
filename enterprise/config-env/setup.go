package configenv

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/types"
)

type Factory struct {
	EnterpriseToken string
	Log             logger.Logger
}

// Setup initializes Suppress User feature
func (m *Factory) Setup() types.ConfigEnvI {
	_ = "STUB: not implemented"
	return *new(types.ConfigEnvI)
}
