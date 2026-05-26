package suppression

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
)

// newHandler creates a new handler for the suppression feature
func newHandler(r Repository, log logger.Logger) *handler { _ = "STUB: not implemented"; return nil }

// handler is a handle to this object
type handler struct {
	log logger.Logger
	r   Repository
}

func (h *handler) GetSuppressedUser(workspaceID, userID, sourceID string) *model.Metadata {
	_ = "STUB: not implemented"
	return nil
}
