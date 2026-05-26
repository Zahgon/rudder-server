package v2

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
)

var (
	ErrorCategoriesMap          = map[string]struct{}{common.CategoryRefreshToken: {}, common.CategoryAuthStatusInactive: {}}
	ErrPermissionOrTokenRevoked = errors.New("problem with user permission or access/refresh token have been revoked")
)

// isOauthTokenExpired checks if the token is expired or is about to expire within the refreshBeforeExpiry duration.
// If the token is not expired, but its secret is the same as the previous secret, it is also considered to be expired.
func isOauthTokenExpired(oauthToken OAuthToken, previousSecret json.RawMessage, refreshBeforeExpiry time.Duration, statsHandler OAuthStatsHandler) bool {
	_ = "STUB: not implemented"
	return false
}

func IsValidAuthErrorCategory(category string) bool { _ = "STUB: not implemented"; return false }
