package gateway

import (
	"net/http"
)

// webAudienceListHandler - handler for audience list requests
func (gw *Handle) webAudienceListHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webExtractHandler - handler for extract requests
func (gw *Handle) webExtractHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webBatchHandler - handler for batch requests
func (gw *Handle) webBatchHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (gw *Handle) internalBatchHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webIdentifyHandler - handler for identify requests
func (gw *Handle) webIdentifyHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webTrackHandler - handler for track requests
func (gw *Handle) webTrackHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webPageHandler - handler for page requests
func (gw *Handle) webPageHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webScreenHandler - handler for screen requests
func (gw *Handle) webScreenHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webAliasHandler - handler for alias requests
func (gw *Handle) webAliasHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webMergeHandler - handler for merge requests
func (gw *Handle) webMergeHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webGroupHandler - handler for group requests
func (gw *Handle) webGroupHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// robotsHandler prevents robots from crawling the gateway endpoints
func (*Handle) robotsHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// webHandler - regular web request handler
func (gw *Handle) webHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// webRequestHandler - handles web requests containing rudder events as payload.
// It parses the payload and calls the request handler to process the request.
func (gw *Handle) webRequestHandler(rh RequestHandler, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// callType middleware sets the call type in the request context
func (gw *Handle) callType(callType string, delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// withContentType sets the content type of the response to the given value
func withContentType(contentType string, delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented" // nolint: unparam
	return *new(http.HandlerFunc)
}
