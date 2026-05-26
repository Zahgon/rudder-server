package augmenter

import (
	"encoding/json"
	"errors"
	"net/http"

	oauthv2extensions "github.com/rudderlabs/rudder-server/services/oauth/v2/extensions"
)

type requestAugmenter struct{}

func NewRequestAugmenter() oauthv2extensions.Augmenter {
	_ = "STUB: not implemented"
	return *new(oauthv2extensions.Augmenter)
}

var (
	ErrSecretNil        = errors.New("secret is nil")
	ErrAccessTokenEmpty = errors.New("access token is empty")
	ErrInstanceURLEmpty = errors.New("instance URL is empty")
)

// Custom augmenter for Salesforce which sets token to Authorization header and instance URL to the request URL
func (s *requestAugmenter) Augment(r *http.Request, body []byte, secret json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// format -> Authorization : OAuth <accessToken>

// GetAuthErrorCategoryForSalesforce returns the error category for Salesforce authentication errors
func GetAuthErrorCategoryForSalesforce(responseBody []byte) string {
	_ = "STUB: not implemented"
	return ""
}
