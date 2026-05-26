package gateway

import (
	"net/http"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

// RegularRequestHandler is an empty struct to capture non-import specific request handling functionality
type RegularRequestHandler struct {
	*Handle
}

// ProcessRequest throws a webRequest into the queue and waits for the response before returning
func (rrh *RegularRequestHandler) ProcessRequest(w *http.ResponseWriter, r *http.Request, reqType string, payload []byte, arctx *gwtypes.AuthRequestContext) string {
	_ = "STUB: not implemented"
	return ""
}
