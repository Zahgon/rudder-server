package v2

//go:generate mockgen -destination=../../../mocks/services/oauthV2/mock_oauthhandler.go -package=mock_oauthV2 github.com/rudderlabs/rudder-server/services/oauth/v2 OAuthHandler

import (
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
	"github.com/rudderlabs/rudder-server/services/oauth/v2/controlplane"
)

// OAuthHandler is the interface that wraps the methods to fetch and refresh OAuth tokens. Handler is using in-memory cache to store the tokens
// and fetches from control plane only when the token is expired or not present in the cache.
type OAuthHandler interface {
	// FetchToken fetches the OAuth token for a given account ID
	FetchToken(params *OAuthTokenParams) (json.RawMessage, StatusCodeError)
	// RefreshToken refreshes the OAuth token for a given account ID. The previousSecret is used to check if the token has already been rotated or not.
	// Token refresh is skipped if the token has already been rotated and previous secret is returned instead.
	RefreshToken(params *OAuthTokenParams, previousSecret json.RawMessage) (json.RawMessage, StatusCodeError)
}

type refreshTokenResponseError struct {
	errType    string
	message    string
	statusCode int
}

// WithCache sets the cache for the OAuthHandler
func WithCache(cache OauthTokenCache) func(*oauthHandler) { _ = "STUB: not implemented"; return nil }

// WithLocker sets the locker for the OAuthHandler
func WithLocker(lock *kitsync.PartitionRWLocker) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

// WithRefreshBeforeExpiry sets the duration before expiry to refresh the token
func WithRefreshBeforeExpiry(refreshBeforeExpiry time.Duration) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

// WithLogger sets the logger for the OAuthHandler
func WithLogger(parentLogger logger.Logger) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

// WithCPClientTimeout sets the control plane client timeout for the OAuthHandler
func WithCPClientTimeout(cpClientTimeout time.Duration) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

// WithStats sets the stats client for the OAuthHandler
func WithStats(stats stats.Stats) func(*oauthHandler) { _ = "STUB: not implemented"; return nil }

// WithCpClient sets the control plane client for the OAuthHandler
func WithCpClient(cpClient controlplane.Connector) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

// WithOauthBreakerOptions sets the OAuth breaker options for the OAuthHandler
func WithOauthBreakerOptions(options *OAuthBreakerOptions) func(*oauthHandler) {
	_ = "STUB: not implemented"
	return nil
}

func WithConfigBackendURL(url string) func(*oauthHandler) { _ = "STUB: not implemented"; return nil }

// NewOAuthHandler returns a new instance of OAuthHandler
func NewOAuthHandler(provider AuthIdentityProvider, options ...func(*oauthHandler)) OAuthHandler {
	_ = "STUB: not implemented"
	return *new(OAuthHandler)
}

// default logger

// default stats

// default control plane client

// default in-memory cache

// default locker

// default refresh before expiry duration

// If breaker options are provided, wrap the handler with the breaker decorator

// Implementation of OAuthHandler interface
type oauthHandler struct {
	stats stats.Stats
	AuthIdentityProvider
	logger         logger.Logger
	rudderFlowType common.RudderFlow
	cpClient       controlplane.Connector

	cacheMu *kitsync.PartitionRWLocker
	cache   OauthTokenCache

	refreshBeforeExpiry time.Duration
	cbeURL              string
	cpClientTimeout     time.Duration
	breakerOptions      *OAuthBreakerOptions
}

// FetchToken fetches the OAuth token for a given account ID
//
//   - If the token is present in the in-memory cache and is not expired, it returns the token from the cache
//   - If the token has expired, it refreshes the token from  control plane
//   - If the token is not present in the cache, it fetches the token from the control plane using a hint to refresh the token if it is expired
func (h *oauthHandler) FetchToken(params *OAuthTokenParams) (json.RawMessage, StatusCodeError) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), *new(StatusCodeError)
}

// RefreshToken refreshes the OAuth token for a given account ID. The previousSecret is used to check if the token has already been rotated or not.
// Token refresh is skipped if the token has already been rotated and previous secret is returned instead.
//
//   - If the token is present in the in-memory cache and is not expired and it doesn't match the previous secret it returns the token from the cache
//   - If the token has expired, or it matches the previous secret it refreshes the token from control plane
//   - If the token is not present in the cache, it fetches the token from the control plane using a hint to refresh the token if it is expired
func (h *oauthHandler) RefreshToken(params *OAuthTokenParams, previousSecret json.RawMessage) (json.RawMessage, StatusCodeError) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), *new(StatusCodeError)
}

func (h *oauthHandler) getToken(params *OAuthTokenParams, previousSecret json.RawMessage, action string, statsHandler OAuthStatsHandler) (json.RawMessage, StatusCodeError) {
	_ = "STUB: not implemented"
	// Create a logger with all context fields
	return *new(json.RawMessage), *new(StatusCodeError)
}

// Add a unique request ID to trace this specific request through logs

// setting this to true as a hint so that the upstream service refreshes the token if it is past its expiry time
// we are not forcing upstream service to refresh the token by default, which will only happen if hasExpired is true and expiredSecret is equal to the upstream service's stored secret

// forcing refresh of token as it is expired or about to expire, or previousSecret matches the cached secret

// getRefreshTokenFromResponse parses the response from control plane and returns either the token or an error
func getRefreshTokenFromResponse(response string, l logger.Logger) (oauthToken *OAuthToken, err *refreshTokenResponseError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: is this indeed a valid case?

// Some problem with AccessToken unmarshalling

// no error, return the token

// handle error type scenarios

// User (or) AccessToken (or) RefreshToken has been revoked

// Fallback to generic message if body.message is missing

// This method hits the Control Plane to get the token
func (h *oauthHandler) getTokenFromAPI(params *OAuthTokenParams, body OauthTokenRequestBody, statsHandler OAuthStatsHandler, action string) (*OAuthToken, StatusCodeError) {
	_ = "STUB: not implemented"
	return nil, *new(StatusCodeError)
}

// Stat for counting number of Refresh Token endpoint calls

// Empty Refresh token response

// Setting empty accessToken value into in-memory auth info map(cache)
