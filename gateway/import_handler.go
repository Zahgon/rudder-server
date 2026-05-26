package gateway

import (
	"net/http"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

// ImportRequestHandler is an empty struct to capture import specific request handling functionality
type ImportRequestHandler struct {
	*Handle
}

// ProcessRequest on ImportRequestHandler splits payload by user and throws them into the webrequestQ and waits for all their responses before returning
func (irh *ImportRequestHandler) ProcessRequest(w *http.ResponseWriter, r *http.Request, _ string, payload []byte, arctx *gwtypes.AuthRequestContext) string {
	_ = "STUB: not implemented"
	return ""
}

// getPayloadFromRequest reads the request body and returns event payloads grouped by user id
// for performance see: https://github.com/rudderlabs/rudder-server/pull/2040
func getUsersPayload(requestPayload []byte) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
