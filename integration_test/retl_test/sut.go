package retltest

import (
	"context"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/rudderlabs/rudder-server/services/rsources"
)

type dst interface {
	ID() string
	Name() string
	TypeName() string
	Config() map[string]any

	Start(t *testing.T)
	Shutdown(t *testing.T)

	Count() int
}

type src interface {
	ID() string
}

// SUT is a System Under Test, running a rudder-server instance and a set of destinations.
type SUT struct {
	cancel context.CancelFunc
	done   chan struct{}

	URL string

	workspaceID string
	Sources     []srcWithDst
}

type srcWithDst struct {
	source       src
	destinations []dst
}

type batch struct {
	Batch []record `json:"batch"`
}

type record struct {
	Context recordContext `json:"context"`

	Type      string    `json:"type"`
	MessageID string    `json:"messageId"`
	UserID    string    `json:"userId"`
	SentAt    time.Time `json:"sentAt"`
	Timestamp time.Time `json:"timestamp"`
}

type rudderSource struct {
	JobID     string `json:"job_id"`
	JobRunID  string `json:"job_run_id"`
	TaskRunID string `json:"task_run_id"`
}

type recordContext struct {
	Sources rudderSource `json:"sources"`
}

func Connect(s src, d ...dst) srcWithDst { _ = "STUB: not implemented"; return *new(srcWithDst) }

// Start rudder-server, its dependencies and the destinations. It blocks until rudder-server and destinations are ready.
func (s *SUT) Start(t *testing.T) { _ = "STUB: not implemented"; return }

// uses a sensible default on windows (tcp/http) and linux/osx (socket)

// quick looops

func (s *SUT) generateConfig() map[string]any { _ = "STUB: not implemented"; return nil }

func (s *SUT) Shutdown(t *testing.T) { _ = "STUB: not implemented"; return }

// SendRETL sends a batch of records to the rETL endpoint.
func (s *SUT) SendRETL(t *testing.T, sourceID, destinationID string, payload batch) {
	_ = "STUB: not implemented"
	return
}

// JobStatus hits the job-status endpoint and returns the status for a sourceID, jobRunID, jobTaskID.
// If the job is not found, the second return value is false. Any other error is logged and the test is failed.
func (s *SUT) JobStatus(t *testing.T, sourceID, jobRunID, jobTaskID string) (rsources.JobStatus, bool) {
	_ = "STUB: not implemented"
	return *new(rsources.JobStatus), false
}

// job_run_id
