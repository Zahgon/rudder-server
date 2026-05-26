package v2

//go:generate mockgen -destination=../../../mocks/services/oauthV2/mock_oauthV2.go -package=mock_oauthV2 github.com/rudderlabs/rudder-server/services/oauth/v2 AuthIdentityProvider

import (
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
)

// OAuthToken is the access token returned by the oauth server
type OAuthToken struct {
	ExpirationDate string          `json:"expirationDate"`
	Secret         json.RawMessage `json:"secret"`
}

func (at *OAuthToken) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Expires returns true if the token expiration date is not empty, it is a valid RFC3339 timestamp and will expire after the given duration
func (at *OAuthToken) Expires(after time.Duration) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type AuthIdentityProvider interface {
	Identity() identity.Identifier
}

type OAuthResponse struct {
	OauthToken   OAuthToken
	Err          string
	ErrorMessage string
}
type OAuthTokenParams struct {
	AccountID     string
	WorkspaceID   string
	DestType      string
	DestinationID string
}

type OauthTokenRequestBody struct {
	HasExpired    bool            `json:"hasExpired"`
	ExpiredSecret json.RawMessage `json:"expiredSecret"`
}

type StatusRequestParams struct {
	AccountID     string
	WorkspaceID   string
	DestType      string
	DestinationID string

	Status string
}

type OAuthInterceptorResponse struct {
	StatusCode int    `json:"statusCode"`         // This is non-zero when the OAuth interceptor, upon completing its functions, intends to pass on the status code to the caller.
	Response   string `json:"response,omitempty"` // This is non-empty when the OAuth interceptor, upon completing its functions, intends to pass on the response body to the caller.
}

type TransportResponse struct {
	OriginalResponse    string                   `json:"originalResponse"`
	InterceptorResponse OAuthInterceptorResponse `json:"interceptorResponse"`
}
