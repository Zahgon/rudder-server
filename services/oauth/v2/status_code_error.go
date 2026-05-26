package v2

// Helper to construct a StatusCodeError
func NewStatusCodeError(code int, err error) StatusCodeError {
	_ = "STUB: not implemented"
	return *new(StatusCodeError)
}

// StatusCodeError wraps an error with an HTTP-like status code.
type StatusCodeError interface {
	StatusCode() int
	error
}

type statusCodeError struct {
	Code int
	Err  error
}

func (e *statusCodeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *statusCodeError) StatusCode() int {
	_ = "STUB: not implemented"

	// Unwrap allows errors.Is and errors.As to work with the wrapped error.
	return 0
}

func (e *statusCodeError) Unwrap() error { _ = "STUB: not implemented"; return nil }
