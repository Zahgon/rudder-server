package lyticsBulkUpload

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func NewLyticsBulkUploader(logger logger.Logger, statsFactory stats.Stats, destinationName, authorization, endpoint string, lytics LyticsService) common.AsyncUploadAndTransformManager {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadAndTransformManager)
}

func NewManager(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT) (common.AsyncDestinationManager, error) {
	_ = "STUB: not implemented"
	return *new(common.AsyncDestinationManager), nil
}
