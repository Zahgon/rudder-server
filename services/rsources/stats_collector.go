package rsources

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

const rsourcesPublishTime = "rsources_publish_time_second"

// StatsPublisher publishes stats
type StatsPublisher interface {
	// Publish publishes statistics
	Publish(ctx context.Context, tx *sql.Tx) error
}

// StatsCollector collects and publishes stats as jobs are
// being created, processed and their statuses are being updated.
type StatsCollector interface {
	StatsPublisher
	// JobsStored captures incoming job statistics
	JobsStored(jobs []*jobsdb.JobT)

	// JobsStoredWithErrors captures incoming job statistics
	JobsStoredWithErrors(jobs []*jobsdb.JobT, failedJobs map[uuid.UUID]string)

	// BeginProcessing prepares the necessary indices in order to
	// be ready for capturing JobStatus statistics
	BeginProcessing(jobs []*jobsdb.JobT)

	// CollectStats captures outgoing job statistics.
	// A call to BeginProcessing must precede a call to this method,
	// so that all necessary indices can be created, since a JobStatus
	// doesn't carry all necessary job metadata such as jobRunId, taskRunId, etc.
	CollectStats(jobStatuses []*jobsdb.JobStatusT)

	// CollectFailedRecords captured `recordId`s for the jobs that were aborted.
	// A call to BeginProcessing must precede a call to this method,
	// so that all necessary indices can be created, since a JobStatus
	// doesn't carry all necessary job metadata such as jobRunId, taskRunId, etc.
	CollectFailedRecords(jobStatuses []*jobsdb.JobStatusT)
}

// FailedJobsStatsCollector collects stats for failed jobs
type FailedJobsStatsCollector interface {
	StatsPublisher
	JobsDropped(jobs []*jobsdb.JobT)
}

// NewStatsCollector creates a new stats collector
func NewStatsCollector(jobservice JobService, component string, statFactory stats.Stats, opts ...OptFunc) StatsCollector {
	_ = "STUB: not implemented"
	return *new(StatsCollector)
}

// NewDroppedJobsCollector creates a new stats collector for publishing failed job stats and records
func NewDroppedJobsCollector(jobservice JobService, component string, statFactory stats.Stats, opts ...OptFunc) FailedJobsStatsCollector {
	_ = "STUB: not implemented"
	return *new(FailedJobsStatsCollector)
}

type statKey struct {
	jobRunId string
	JobTargetKey
}

func (sk statKey) String() string { _ = "STUB: not implemented"; return "" }

var _ StatsCollector = (*statsCollector)(nil)

type statsCollector struct {
	processing            bool
	jobService            JobService
	jobIdsToStatKeyIndex  map[int64]statKey
	jobIdsToRecordIdIndex map[int64]json.RawMessage
	statsIndex            map[statKey]*Stats
	failedRecordsIndex    map[statKey][]FailedRecord
	parametersParser      parametersParser
	stats                 struct {
		publishTime stats.Timer
	}
}

func (r *statsCollector) orderedStatMapKeys() []statKey { _ = "STUB: not implemented"; return nil }

func (r *statsCollector) orderedFailedRecordsKeys() []statKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *statsCollector) JobsStored(jobs []*jobsdb.JobT) { _ = "STUB: not implemented"; return }

func (r *statsCollector) JobsDropped(jobs []*jobsdb.JobT) { _ = "STUB: not implemented"; return }

func (r *statsCollector) JobsStoredWithErrors(jobs []*jobsdb.JobT, failedJobs map[uuid.UUID]string) {
	_ = "STUB: not implemented"
	return
}

func (r *statsCollector) BeginProcessing(jobs []*jobsdb.JobT) { _ = "STUB: not implemented"; return }

func (r *statsCollector) CollectStats(jobStatuses []*jobsdb.JobStatusT) {
	_ = "STUB: not implemented"
	return
}

// Filtered state is being considered as a success. If we want to report them separately, we can add a new field in stats

func (r *statsCollector) CollectFailedRecords(jobStatuses []*jobsdb.JobStatusT) {
	_ = "STUB: not implemented"
	return
}

func (r *statsCollector) Publish(ctx context.Context, tx *sql.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// sort the maps to avoid deadlocks

// sort the records as well to avoid deadlocks

func (r *statsCollector) buildStats(jobs []*jobsdb.JobT, failedJobs map[uuid.UUID]string, incrementIn bool) {
	_ = "STUB: not implemented" // skipcq: RVV-A0005
	return
}

type parametersParser func(jp json.RawMessage) (jobRunID, recordID string, target JobTargetKey)

type OptFunc func(*statsCollector)

// IgnoreDestinationID ignores the destinationID parameter of the job and while capturing statistics
func IgnoreDestinationID() OptFunc { _ = "STUB: not implemented"; return *new(OptFunc) }

func defaultParametersParser(jobParams json.RawMessage) (jobRunID, recordID string, target JobTargetKey) {
	_ = "STUB: not implemented"
	return "", "", *new(JobTargetKey)
}
