package audience

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	oauthv2 "github.com/rudderlabs/rudder-server/services/oauth/v2"
)

func newManagerInternal(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT, oauthHandler oauthv2.OAuthHandler) (*BingAdsBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewManager(
	conf *config.Config,
	logger logger.Logger,
	statsFactory stats.Stats,
	destination *backendconfig.DestinationT,
	backendConfig backendconfig.BackendConfig,
) (*BingAdsBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
