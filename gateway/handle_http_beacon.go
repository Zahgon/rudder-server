package gateway

import (
	"net/http"
)

// beaconBatchHandler can handle beacon batch requests where writeKey is passed as a query param
func (gw *Handle) beaconBatchHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// beaconInterceptor reads the writeKey from the query params and sets it in the request Authorization header
func (gw *Handle) beaconInterceptor(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// set basic auth header

// send req to webHandler
