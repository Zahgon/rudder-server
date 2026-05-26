package v2

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
)

const OAUTH_V2_STAT_PREFIX = "oauth_action"

type OAuthStats struct {
	stats           stats.Stats
	id              string // destinationId -> for action == auth_status_inactive, accountId -> for action == refresh_token/fetch_token
	workspaceID     string
	errorMessage    string
	rudderCategory  string // destination
	statName        string
	isCallToCpApi   bool   // is a call being made to control-plane APIs
	authErrCategory string // for action=refresh_token -> REFRESH_TOKEN, for action=fetch_token -> "", for action=auth_status_inactive -> auth_status_inactive
	destType        string
	flowType        common.RudderFlow // delivery, delete
	action          string            // refresh_token, fetch_token, auth_status_inactive
}

type OAuthStatsHandler struct {
	stats       stats.Stats
	defaultTags stats.Tags
}

func (oauthStats *OAuthStats) ToStatsTags() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

func NewStatsHandlerFromOAuthStats(oauthStats *OAuthStats) OAuthStatsHandler {
	_ = "STUB: not implemented"
	return *new(OAuthStatsHandler)
}

func (m *OAuthStatsHandler) Increment(statSuffix string, tags stats.Tags) {
	_ = "STUB: not implemented"
	return
}

func (m *OAuthStatsHandler) SendTiming(startTime time.Time, statSuffix string, tags stats.Tags) {
	_ = "STUB: not implemented"
	return
}
