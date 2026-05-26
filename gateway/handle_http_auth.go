package gateway

import (
	"net/http"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

// writeKeyAuth middleware to authenticate writeKey in the Authorization header.
// If the writeKey is valid and the source is enabled, the source auth info is added to the request context.
// If the writeKey is invalid, the request is rejected.
func (gw *Handle) writeKeyAuth(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webhookAuth middleware to authenticate webhook requests.
// The writeKey can be passed in the Authorization header or as a query param.
// If the writeKey is valid, corresponds to a webhook source and the source is enabled, the source auth info is added to the request context.
// If the writeKey is invalid, the request is rejected.
func (gw *Handle) webhookAuth(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// sourceIDAuth middleware to authenticate sourceID in the X-Rudder-Source-Id header.
// If the sourceID is valid and the source is enabled, the source auth info is added to the request context.
// If the sourceID is invalid, the request is rejected.
func (gw *Handle) sourceIDAuth(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// authDestIDForSource middleware to authenticate destinationId in the X-Rudder-Destination-Id header.
// If the destinationId is invalid, the request is rejected.
// destinationID authentication should be performed only after source is authenticated and source is present in context
// Following validations are performed
//  1. Destination should be present for source config
//  2. Destination should be enabled for the source
func (gw *Handle) authDestIDForSource(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// TODO: make default value true once rETL team migrates to sending destination ID in header

// replaySourceIDAuth middleware to authenticate sourceID in the X-Rudder-Source-Id header.
// If the sourceID is valid, i.e. it is a replay source and enabled, the source auth info is added to the request context.
// If the sourceID is invalid, the request is rejected.
func (gw *Handle) replaySourceIDAuth(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// sourceDestIDAuth middleware to authenticate sourceID and destinationID
// in the X-Rudder-Source-Id and X-Rudder-Destination-Id header respectively.
// If the sourceID or destinationID is invalid, the request is rejected.
func (gw *Handle) sourceDestIDAuth(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// augmentAuthRequestContext adds source job run id and task run id from the request to the authentication context.
func augmentAuthRequestContext(arctx *gwtypes.AuthRequestContext, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// authRequestContextForSourceID gets request context for a given sourceID. If the sourceID is invalid, returns nil.
func (gw *Handle) authRequestContextForSourceID(sourceID string) *gwtypes.AuthRequestContext {
	_ = "STUB: not implemented"
	return nil
}

// authRequestContextForWriteKey gets request context for a given writeKey. If the writeKey is invalid, returns nil.
func (gw *Handle) authRequestContextForWriteKey(writeKey string) *gwtypes.AuthRequestContext {
	_ = "STUB: not implemented"
	return nil
}

// sourceToRequestContext converts a source to request context.
func sourceToRequestContext(s backendconfig.SourceT) *gwtypes.AuthRequestContext {
	_ = "STUB: not implemented"
	return nil
}

func (gw *Handle) handleHttpError(w http.ResponseWriter, r *http.Request, errorMessage string) {
	_ = "STUB: not implemented"
	return
}

func (gw *Handle) handleFailureStats(errorMessage, reqType string, arctx *gwtypes.AuthRequestContext) {
	_ = "STUB: not implemented"
	return
}
