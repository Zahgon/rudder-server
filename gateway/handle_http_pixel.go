package gateway

import (
	"net/http"
	"net/url"
)

// pixelPageHandler can handle pixel page requests where everything is passed as query params.
// it also writes a pixel response to the client regardless of the actual result of the request
func (gw *Handle) pixelPageHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// pixelTrackHandler can handle pixel track requests where everything is passed as query params.
// it also writes a pixel response to the client regardless of the actual result of the request
func (gw *Handle) pixelTrackHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// pixelInterceptor reads information from the query parameters to fill in the request's body and authorization header before passing it to the next handler
// It also writes a pixel response to the client regardless of the next handler's response
func (gw *Handle) pixelInterceptor(reqType string, next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// write pixel response even if there is an error

// make a new request

// set basic auth header

// set X-Forwarded-For header

// convert the pixel request(r) to a web request(req)

// create a new writer since the pixel is going to be written to the client regardless of the next handler's response

// preparePixelPayload reads a pixel GET request and maps it to a proper payload in the request's body
func (gw *Handle) preparePixelPayload(r *http.Request, qp url.Values, reqType string) error {
	_ = "STUB: not implemented"
	// add default fields to body
	return nil
}

// make sure anonymousId is in correct format

// add queryParams to body

// add request specific fields to body

// add body to request

// newPixelWriter returns a new, properly initialized pixel writer
// it is used to capture the status code and body of the response without writing it to the client
func newPixelWriter() *pixelHttpWriter { _ = "STUB: not implemented"; return nil }

// pixelHttpWriter captures the status code and body of the response
type pixelHttpWriter struct {
	status int
	body   []byte
}

func (w *pixelHttpWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (w *pixelHttpWriter) WriteHeader(status int) { _ = "STUB: not implemented"; return }

func (w *pixelHttpWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
