package v2

import (
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	oauth "github.com/rudderlabs/rudder-server/services/oauth/v2"
	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
	oauthexts "github.com/rudderlabs/rudder-server/services/oauth/v2/extensions"
)

type HttpClientOptionalArgs struct {
	Transport           http.RoundTripper
	Augmenter           oauthexts.Augmenter
	Locker              *sync.PartitionRWLocker
	OAuthHandler        oauth.OAuthHandler
	ExpirationTimeDiff  time.Duration
	Logger              logger.Logger
	OAuthBreakerOptions *oauth.OAuthBreakerOptions
}

// NewOAuthHttpClient returns a http client that will add the appropriate authorization information to oauth requests.
func NewOAuthHttpClient(client *http.Client, flowType common.RudderFlow, tokenCache *oauth.OauthTokenCache, backendConfig backendconfig.BackendConfig, getAuthErrorCategory func([]byte) (string, error), opArgs *HttpClientOptionalArgs) *http.Client {
	_ = "STUB: not implemented"
	return nil
}
