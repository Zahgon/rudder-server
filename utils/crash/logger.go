package crash

import (
	"net/http"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

type panicLogger struct {
	notifyOnce sync.Once
	logger     logger.Logger

	opts PanicWrapperOpts
}

// UsingLogger uses the provided logger to log panics.
func UsingLogger(logger logger.Logger, opts PanicWrapperOpts) *panicLogger {
	_ = "STUB: not implemented"
	return nil
}

// Notify returns a function that should be deferred to capture panics.
// It will do a Fatal log with the stack trace and panic value.
func (p *panicLogger) Notify(team string) func() { _ = "STUB: not implemented"; return nil }

// nolint:forbidigo

// Handler creates an http Handler that captures any panics that
// happen. It then repanics so that the default http Server panic handler can
// handle the panic too.
func (p *panicLogger) Handler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
