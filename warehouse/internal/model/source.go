package model

import (
	"encoding/json"
	"time"
)

type SourceJobType interface {
	String() string
	sourceJobTypeProtected()
}

type sourceJobType string

func (s sourceJobType) String() string          { _ = "STUB: not implemented"; return "" }
func (s sourceJobType) sourceJobTypeProtected() { _ = "STUB: not implemented"; return }

var SourceJobTypeDeleteByJobRunID SourceJobType = sourceJobType("deletebyjobrunid")

func FromSourceJobType(jobType string) (SourceJobType, error) {
	_ = "STUB: not implemented"
	return *new(SourceJobType), nil
}

type SourceJobStatus interface {
	String() string
	sourceJobStatusProtected()
}

type sourceJobStatus string

func (s sourceJobStatus) String() string            { _ = "STUB: not implemented"; return "" }
func (s sourceJobStatus) sourceJobStatusProtected() { _ = "STUB: not implemented"; return }

var (
	SourceJobStatusWaiting   SourceJobStatus = sourceJobStatus("waiting")
	SourceJobStatusExecuting SourceJobStatus = sourceJobStatus("executing")
	SourceJobStatusFailed    SourceJobStatus = sourceJobStatus("failed")
	SourceJobStatusAborted   SourceJobStatus = sourceJobStatus("aborted")
	SourceJobStatusSucceeded SourceJobStatus = sourceJobStatus("succeeded")
)

func FromSourceJobStatus(status string) (SourceJobStatus, error) {
	_ = "STUB: not implemented"
	return *new(SourceJobStatus), nil
}

type SourceJob struct {
	ID int64

	SourceID      string
	DestinationID string
	WorkspaceID   string

	TableName string

	Status  SourceJobStatus
	Error   error
	JobType SourceJobType

	Metadata json.RawMessage
	Attempts int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
