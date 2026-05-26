package rsources

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

//go:generate mockgen -source=rsources.go -destination=mock_rsources.go -package=rsources github.com/rudderlabs/rudder-server/services/rsources JobService

type JobFilter struct {
	TaskRunID []string
	SourceID  []string
}

type JobTargetKey struct {
	TaskRunID     string `json:"source_task_run_id"`
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
}

func (k JobTargetKey) String() string { _ = "STUB: not implemented"; return "" }

type Stats struct {
	In     uint `json:"in"`
	Out    uint `json:"out"`
	Failed uint `json:"failed"`
}

func (r *Stats) completed() bool { _ = "STUB: not implemented"; return false }

func (r *Stats) corrupted() bool { _ = "STUB: not implemented"; return false }

func (r *Stats) fixCorrupted() { _ = "STUB: not implemented"; return }

type JobStatus struct {
	ID          string       `json:"id"`
	TasksStatus []TaskStatus `json:"tasks"`
}

func (js *JobStatus) FixCorruptedStats(log logger.Logger) { _ = "STUB: not implemented"; return }

type TaskStatus struct {
	ID            string         `json:"id"`
	SourcesStatus []SourceStatus `json:"sources"`
}

type SourceStatus struct {
	ID                 string              `json:"id"`
	Completed          bool                `json:"completed"`
	Stats              Stats               `json:"stats"`
	DestinationsStatus []DestinationStatus `json:"destinations"`
}

func (sourceStatus *SourceStatus) calculateCompleted() { _ = "STUB: not implemented"; return }

type DestinationStatus struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Stats     Stats  `json:"stats"`
}

type PagingInfo struct {
	Size          int    `json:"size"`
	NextPageToken string `json:"next"`
}

func NextPageTokenFromString(v string) (NextPageToken, error) {
	_ = "STUB: not implemented"
	return *new(NextPageToken), nil
}

type NextPageToken struct {
	ID       string `json:"id"`
	RecordID string `json:"record_id"`
}

func (npt *NextPageToken) String() string { _ = "STUB: not implemented"; return "" }

type (
	JobFailedRecordsV2      JobFailedRecords[FailedRecord]
	JobFailedRecordsV1      JobFailedRecords[json.RawMessage]
	JobFailedRecords[R any] struct {
		ID     string                 `json:"id"`
		Tasks  []TaskFailedRecords[R] `json:"tasks"`
		Paging *PagingInfo            `json:"paging,omitempty"`
	}
)

type TaskFailedRecords[R any] struct {
	ID      string                   `json:"id"`
	Sources []SourceFailedRecords[R] `json:"sources"`
}

type SourceFailedRecords[R any] struct {
	ID           string                        `json:"id"`
	Records      []R                           `json:"records"`
	Destinations []DestinationFailedRecords[R] `json:"destinations"`
}

type DestinationFailedRecords[R any] struct {
	ID      string `json:"id"`
	Records []R    `json:"records"`
}
type FailedRecord struct {
	Record json.RawMessage `json:"record"`
	Code   int             `json:"code"`
}

// ErrStatusNotFound sentinel error indicating that status cannot be found
var ErrStatusNotFound = errors.New("status not found")

// ErrSourceNotCompleted sentinel error indicating that a source is not completed
var ErrSourceNotCompleted = errors.New("source not completed")

// ErrFailedRecordsNotFound sentinel error indicating that failed records cannot be found
var ErrFailedRecordsNotFound = errors.New("failed records not found")

// StatsIncrementer increments stats
type StatsIncrementer interface {
	// IncrementStats increments the existing statistic counters
	// for a specific job measurement.
	IncrementStats(ctx context.Context, tx *sql.Tx, jobRunId string, key JobTargetKey, stats Stats) error
}

type JobServiceConfig struct {
	LocalHostname                string
	LocalConn                    string
	MaxPoolSize                  int
	MinPoolSize                  int
	SharedConn                   string
	SubscriptionTargetConn       string
	SkipFailedRecordsCollection  bool
	FailedRecordsInsertBatchSize config.ValueLoader[int]
	Log                          logger.Logger
	ShouldSetupSharedDB          bool
}

// JobService manages information about jobs created by rudder-sources
type JobService interface {
	StatsIncrementer

	// Delete deletes all relevant information for a given jobRunId
	Delete(ctx context.Context, jobRunId string, filter JobFilter) error

	// DeleteJobStatus deletes the status for a given jobRunId
	DeleteJobStatus(ctx context.Context, jobRunId string, filter JobFilter) error

	// DeleteFailedRecords deletes all failed records for a given jobRunId
	DeleteFailedRecords(ctx context.Context, jobRunId string, filter JobFilter) error

	// GetStatus gets the current status of a job
	GetStatus(ctx context.Context, jobRunId string, filter JobFilter) (JobStatus, error)

	// AddFailedRecords adds failed records to the database as part of a transaction
	AddFailedRecords(ctx context.Context, tx *sql.Tx, jobRunId string, key JobTargetKey, records []FailedRecord) error

	// GetFailedRecords gets the failed records for a jobRunID, with filters on taskRunId and sourceId
	GetFailedRecords(ctx context.Context, jobRunId string, filter JobFilter, paging PagingInfo) (JobFailedRecordsV2, error)

	// GetFailedRecordsV1 gets the failed records for a jobRunID, with filters on taskRunId and sourceId
	GetFailedRecordsV1(ctx context.Context, jobRunId string, filter JobFilter, paging PagingInfo) (JobFailedRecordsV1, error)

	// CleanupLoop starts the cleanup loop in the background which will stop upon context termination or in case of an error
	CleanupLoop(ctx context.Context) error

	// Monitor monitors the logical replication slot and lag when a shared database is configured
	Monitor(ctx context.Context, lagGauge, replicationSlotGauge Gauger)
}

type Gauger interface {
	Gauge(any)
}

func NewJobService(ctx context.Context, jobServiceConfig JobServiceConfig, stats stats.Stats) (JobService, error) {
	_ = "STUB: not implemented"
	return *new(JobService), nil
}

// minimum 2 connections in the pool for proper startup

func NewNoOpService() JobService { _ = "STUB: not implemented"; return *new(JobService) }

type noopService struct{}

func (*noopService) Delete(_ context.Context, _ string, _ JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopService) DeleteJobStatus(_ context.Context, _ string, _ JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopService) DeleteFailedRecords(_ context.Context, _ string, _ JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopService) GetStatus(_ context.Context, _ string, _ JobFilter) (JobStatus, error) {
	_ = "STUB: not implemented"
	return *new(JobStatus), nil
}

func (*noopService) IncrementStats(_ context.Context, _ *sql.Tx, _ string, _ JobTargetKey, _ Stats) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopService) AddFailedRecords(_ context.Context, _ *sql.Tx, _ string, _ JobTargetKey, _ []FailedRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopService) GetFailedRecords(_ context.Context, _ string, _ JobFilter, _ PagingInfo) (JobFailedRecordsV2, error) {
	_ = "STUB: not implemented"
	return *new(JobFailedRecordsV2), nil
}

func (*noopService) GetFailedRecordsV1(_ context.Context, _ string, _ JobFilter, _ PagingInfo) (JobFailedRecordsV1, error) {
	_ = "STUB: not implemented"
	return *new(JobFailedRecordsV1), nil
}

func (*noopService) CleanupLoop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*noopService) Monitor(_ context.Context, _, _ Gauger) { _ = "STUB: not implemented"; return }
