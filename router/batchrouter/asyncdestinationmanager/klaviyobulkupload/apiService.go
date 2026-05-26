package klaviyobulkupload

import (
	"net/http"

	"golang.org/x/time/rate"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

const (
	KlaviyoAPIURL = "https://a.klaviyo.com/api/profile-bulk-import-jobs/"
)

type RateLimiterHTTPClient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

func (c *RateLimiterHTTPClient) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type KlaviyoAPIServiceImpl struct {
	client        *RateLimiterHTTPClient
	PrivateAPIKey string
	logger        logger.Logger
	statsFactory  stats.Stats
	statLabels    stats.Tags
}

func newRateLimiterClient() *RateLimiterHTTPClient { _ = "STUB: not implemented"; return nil }

// Doc: https://developers.klaviyo.com/en/reference/bulk_import_profiles

func setRequestHeaders(req *http.Request, apiKey string) { _ = "STUB: not implemented"; return }

func (k *KlaviyoAPIServiceImpl) UploadProfiles(profiles Payload) (*UploadResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KlaviyoAPIServiceImpl) GetUploadStatus(importId string) (*PollResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KlaviyoAPIServiceImpl) GetUploadErrors(importId string) (*UploadStatusResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKlaviyoAPIService(destination *backendconfig.DestinationT, logger logger.Logger, statsFactory stats.Stats) (KlaviyoAPIService, error) {
	_ = "STUB: not implemented"
	return *new(KlaviyoAPIService), nil
}
