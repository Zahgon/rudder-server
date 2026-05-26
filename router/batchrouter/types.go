package batchrouter

import (
	stdjson "encoding/json"
	"time"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
	router_utils "github.com/rudderlabs/rudder-server/router/utils"
)

type Connection struct {
	Source      backendconfig.SourceT
	Destination backendconfig.DestinationT
}

type DestinationJobs struct {
	destWithSources router_utils.DestinationWithSources
	jobs            []*jobsdb.JobT
}

type ObjectStorageDefinition struct {
	Config          map[string]any
	Key             string
	Provider        string
	DestinationID   string
	DestinationType string
}

type batchRequestMetric struct {
	batchRequestSuccess int
	batchRequestFailed  int
}

type UploadResult struct {
	Config           map[string]any
	Key              string
	FileLocation     string
	LocalFilePaths   []string
	JournalOpID      int64
	Error            error
	FirstEventAt     string
	LastEventAt      string
	TotalEvents      int
	TotalBytes       int
	BytesPerTable    map[string]int64
	UseRudderStorage bool
}

type ErrorResponse struct {
	Error string
}

type WarningResponse struct {
	Remarks string
}

type BatchedJobs struct {
	Jobs       []*jobsdb.JobT
	Connection *Connection
	TimeWindow time.Time
	JobState   string // ENUM waiting, executing, succeeded, waiting_retry, filtered, failed, aborted, migrating, migrated, wont_migrate
}

type getReportMetricsParams struct {
	StatusList    []*jobsdb.JobStatusT
	ParametersMap map[int64]stdjson.RawMessage
	JobsList      []*jobsdb.JobT
}

type setMultipleJobStatusParams struct {
	asyncJobMetadata
	AsyncOutput common.AsyncUploadOutput
	Attempted   bool
	JobsList    []*jobsdb.JobT
}

// asyncJobMetadata holds metadata related to async jobs
type asyncJobMetadata struct {
	AttemptNums       map[int64]int                // jobID -> attemptNum
	FirstAttemptedAts map[int64]time.Time          // jobID -> first attempted at
	JobParameters     map[int64]stdjson.RawMessage // jobID -> job parameters
	PartitionIDs      map[int64]string             // jobID -> partitionID
}

// newAsyncJobMetadata creates asyncJobMetadata from a list of jobs
func newAsyncJobMetadata(jobList []*jobsdb.JobT) asyncJobMetadata {
	_ = "STUB: not implemented"
	return *new(asyncJobMetadata)
}

// newAsyncJobMetadataFromDestinationStruct creates asyncJobMetadata from an AsyncDestinationStruct
func newAsyncJobMetadataFromDestinationStruct(asyncDestinationStruct *common.AsyncDestinationStruct) asyncJobMetadata {
	_ = "STUB: not implemented"
	return *new(asyncJobMetadata)
}
