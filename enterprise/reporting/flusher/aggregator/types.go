package aggregator

import (
	"time"

	"github.com/segmentio/go-hll"
)

type TrackedUsersReport struct {
	ReportedAt                  time.Time `json:"reportedAt"`
	WorkspaceID                 string    `json:"workspaceId"`
	SourceID                    string    `json:"sourceId"`
	InstanceID                  string    `json:"instanceId"`
	UserIDHLL                   *hll.Hll  `json:"-"`
	AnonymousIDHLL              *hll.Hll  `json:"-"`
	IdentifiedAnonymousIDHLL    *hll.Hll  `json:"-"`
	UserIDHLLHex                string    `json:"userIdHLL"`
	AnonymousIDHLLHex           string    `json:"anonymousIdHLL"`
	IdentifiedAnonymousIDHLLHex string    `json:"identifiedAnonymousIdHLL"`
}

func (t *TrackedUsersReport) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
