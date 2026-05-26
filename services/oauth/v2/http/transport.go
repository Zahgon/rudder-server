package v2

//go:generate mockgen -destination=../../../../mocks/services/oauthV2/mock_roundtripper.go -package=mock_oauthV2 net/http RoundTripper

import (
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	v2 "github.com/rudderlabs/rudder-server/services/oauth/v2"
	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
	oauthexts "github.com/rudderlabs/rudder-server/services/oauth/v2/extensions"
)

// TransportArgs is a struct that contains the required parameters to create a new Oauth2Transport.
type TransportArgs struct {
	BackendConfig backendconfig.BackendConfig
	FlowType      common.RudderFlow
	// TokenCache is a cache for storing OAuth tokens.
	TokenCache *v2.OauthTokenCache
	// Locker provides synchronization mechanisms.
	Locker *kitsync.PartitionRWLocker
	// GetAuthErrorCategory is a function to get the auth error category from the response body. It can be REFRESH_TOKEN or AUTH_STATUS_INACTIVE.
	GetAuthErrorCategory func([]byte) (string, error)
	// Augmenter is an interface for augmenting requests with OAuth tokens.
	oauthexts.Augmenter
	// OAuthHandler handles refreshToken and fetchToken requests.
	OAuthHandler v2.OAuthHandler
	// OriginalTransport is the underlying HTTP transport.
	OriginalTransport http.RoundTripper
	logger            logger.Logger
	stats             stats.Stats
}

// OAuthTransport is a http.RoundTripper that adds the appropriate authorization information to oauth requests.
// Also, it makes the calls to the actual endpoint and handles the response by refreshing the token if required or by updating the authStatus to "inactive".
type OAuthTransport struct {
	oauthHandler v2.OAuthHandler
	oauthexts.Augmenter
	Transport            http.RoundTripper
	log                  logger.Logger
	flow                 common.RudderFlow
	getAuthErrorCategory func([]byte) (string, error)
	stats                stats.Stats
}

// This struct is used to transport common information across the pre and post round trip methods.
type roundTripState struct {
	destination *backendconfig.DestinationT
	accountID   string
	tokenParams *v2.OAuthTokenParams
	res         *http.Response
	req         *http.Request
}

func httpResponseCreator(statusCode int, body []byte) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

func NewOAuthTransport(args *TransportArgs) *OAuthTransport { _ = "STUB: not implemented"; return nil }

func (t *OAuthTransport) preRoundTrip(rts *roundTripState) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

func (t *OAuthTransport) postRoundTrip(rts *roundTripState) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// Create a new response with a 500 status code

// internal function

// Instead of returning an error, set a 500 status code in the interceptor response
// This will make the error retryable instead of causing a panic

// since same token that was used to make the http call needs to be refreshed, we need the current token information

// Instead of returning an error, set a 500 status code in the interceptor response

// use the message from the underlying TypeMessageError if possible

// token refresh was successful, so we return a 500 to the caller to retry the request

// when error is not nil, the response sent will be ignored(downstream)

func (rts *roundTripState) getAccountID(flow common.RudderFlow) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *OAuthTransport) fireTimerStats(statName string, tags stats.Tags, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (t *OAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a 500 error response instead of propagating the error
