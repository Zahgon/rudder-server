package eloqua

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

func NewManager(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT) (*EloquaBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEloquaBulkUploader(logger logger.Logger, statsFactory stats.Stats, destinationName, authorization, baseEndpoint string, eloqua EloquaService) *EloquaBulkUploader {
	_ = "STUB: not implemented"
	return nil
}
