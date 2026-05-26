package httputil

import (
	"net/http"
)

// RetriableStatus returns true if the HTTP status code should be retried.
//
//	We consider retriable status code:
//		* 5xx - all server errors
//		* 408 - request timeout
//		* 429 - too many requests.
func RetriableStatus(statusCode int) bool {
	_ = "STUB: not implemented"
	// 1xx, 2xx, 3xx -- we assume no error and thus no retry should happen
	return false
}

// 5xx

// 4xx codes:

// CloseResponse closes the response's body. But reads at least some of the body so if it's
// small the underlying TCP connection will be re-used. No need to check for errors: if it
// fails, the Transport won't reuse it anyway.
func CloseResponse(resp *http.Response) { _ = "STUB: not implemented"; return }

// 2KB
