package common

import (
	"time"

	"golang.org/x/oauth2"

	oauthv2 "github.com/rudderlabs/rudder-server/services/oauth/v2"
)

type SecretStruct struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	DeveloperToken string `json:"developer_token"`
	ExpirationDate string `json:"expirationDate"`
}

type TokenSource struct {
	WorkspaceID        string
	DestinationDefName string
	AccountID          string
	OauthHandler       oauthv2.OAuthHandler
	DestinationID      string
	CurrentTime        func() time.Time
}

// authentication related utils

func (ts *TokenSource) GenerateTokenV2() (*SecretStruct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts *TokenSource) Token() (*oauth2.Token, error) { _ = "STUB: not implemented"; return nil, nil }

// skipping error check as similar check is already done on the previous function
