package marketobulkupload

import (
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

// MarketoBulkUploaderOptions contains all dependencies needed for the uploader
type MarketoBulkUploaderOptions struct {
	DestinationName   string
	DestinationConfig MarketoConfig
	Logger            logger.Logger
	StatsFactory      stats.Stats
	APIService        MarketoAPIServiceInterface
}

func NewManager(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT) (*MarketoBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMarketoBulkUploader(destinationName string, log logger.Logger, statsFactory stats.Stats, httpClient *http.Client, destConfig MarketoConfig) *MarketoBulkUploader {
	_ = "STUB: not implemented"
	return nil
}

// NewMarketoBulkUploaderWithOptions creates a new MarketoBulkUploader with the given options
func NewMarketoBulkUploaderWithOptions(options MarketoBulkUploaderOptions) *MarketoBulkUploader {
	_ = "STUB: not implemented"
	return nil
}
