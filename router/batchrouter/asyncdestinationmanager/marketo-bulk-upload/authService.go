package marketobulkupload

import (
	"net/http"
)

type MarketoAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	FetchedAt   int64
	Scope       string `json:"scope"`
}

type MarketoAuthServiceInterface interface {
	GetAccessToken() (string, error)
}

type MarketoAuthService struct {
	httpCLient   *http.Client
	munchkinId   string
	clientId     string
	clientSecret string
	accessToken  MarketoAccessToken
}

func (m *MarketoAuthService) fetchOrUpdateAccessToken() error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketoAuthService) GetAccessToken() (string, error) {
	_ = "STUB: not implemented"
	// keeping simple logic for now
	return "", nil
}
