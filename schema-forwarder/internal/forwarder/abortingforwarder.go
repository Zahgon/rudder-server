package forwarder

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

// AbortingForwarder is a forwarder which aborts all jobs instead of forwarding them
type AbortingForwarder struct {
	BaseForwarder
}

// NewAbortingForwarder returns a new, properly initialized, AbortingForwarder
func NewAbortingForwarder(terminalErrFn func(error), schemaDB jobsdb.JobsDB, config *config.Config, log logger.Logger, stat stats.Stats) *AbortingForwarder {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the forwarder which reads jobs from the database and aborts them
func (nf *AbortingForwarder) Start() error { _ = "STUB: not implemented"; return nil }

// we are shutting down
//nolint:nilerr

// we are signaling to shutdown the app

// we are shutting down
//nolint:nilerr

// we are signaling to shutdown the app

// Stop stops the forwarder
func (nf *AbortingForwarder) Stop() { _ = "STUB: not implemented"; return }
