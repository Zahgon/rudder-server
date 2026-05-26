package source

import (
	"net/http"
	"regexp"
)

// emptyRegex matches empty strings
var emptyRegex = regexp.MustCompile(`^\s*$`)

// InsertJobHandler adds a job to the warehouse_jobs table
func (m *Manager) InsertJobHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// StatusJobHandler The following handler gets called for getting the status of the async job
func (m *Manager) StatusJobHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func validatePayload(payload *insertJobRequest) error { _ = "STUB: not implemented"; return nil }
