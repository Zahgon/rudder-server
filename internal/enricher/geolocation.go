package enricher

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
	"github.com/rudderlabs/rudder-server/services/geolocation"
)

const (
	ERR_INVALID_IP    = "invalid_ip"
	ERR_EMPTY_IP      = "empty_ip"
	ERR_LOCATE_FAILED = "locate_failed"
)

type Geolocation struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Postal   string `json:"postal"`
	Location string `json:"location"`
	Timezone string `json:"timezone"`
}

type geoEnricher struct {
	fetcher geolocation.GeoFetcher
	logger  logger.Logger
	stats   stats.Stats
}

func NewGeoEnricher(conf *config.Config, log logger.Logger, statClient stats.Stats) (PipelineEnricher, error) {
	_ = "STUB: not implemented"
	return *new(PipelineEnricher), nil
}

// Enrich function runs on a request of GatewayBatchRequest which contains
// multiple singular events from a source. The enrich function augments the
// geolocation information per event based on IP address.
func (e *geoEnricher) Enrich(source *backendconfig.SourceT, request *types.GatewayBatchRequest, _ *types.EventParams) error {
	_ = "STUB: not implemented"
	return nil
}

// if the context section is missing on the event
// set it with default as map[string]interface{}

// if the context is other than map[string]interface{}, add error and continue

// if the `geo` key already present on the event, continue

// `emptyIP` even though it's invalid is treated differently
// to get better stats about how many non-empty values are coming in which are invalid.

// empty / invalidIP's are not terminal errors but
// any error except that mean the database failed for lookup

// Set the empty data on the context nonetheless

func (e *geoEnricher) Close() error { _ = "STUB: not implemented"; return nil }

// downloadMaxmindDB downloads database file from upstream s3 and stores it in
// a specified location. Download is skipped if the file already exists in the expected path.
func downloadMaxmindDB(ctx context.Context, conf *config.Config, log logger.Logger) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the filepath exists return

// before renaming, we need to sync data to the disk

// Finally move the downloaded file from previous temp location to new location

func extractGeolocationData(ip string, geoCity geolocation.GeoInfo) Geolocation {
	_ = "STUB: not implemented"
	return *new(Geolocation)
}

// default values of latitude and longitude can give
// incorrect result, so we have casted them in pointers so we know
// when the value is missing.
