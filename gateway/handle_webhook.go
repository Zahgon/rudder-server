package gateway

import (
	"net/http"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
	"github.com/rudderlabs/rudder-server/gateway/webhook/model"
)

func (gw *Handle) webhookHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// ProcessTransformedWebhookRequest is an interface wrapper for webhook
func (gw *Handle) ProcessTransformedWebhookRequest(w *http.ResponseWriter, r *http.Request, reqType string, payload []byte, arctx *gwtypes.AuthRequestContext) string {
	_ = "STUB: not implemented"
	return ""
}

func (gw *Handle) SaveWebhookFailures(reqs []*model.FailedWebhookPayload) error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
