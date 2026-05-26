package common

import (
	"context"
	stdjson "encoding/json"
	"sync"
	"time"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

type AsyncUploadAndTransformManager interface {
	Upload(ctx context.Context, asyncDestStruct *AsyncDestinationStruct) AsyncUploadOutput
	Transform(job *jobsdb.JobT) (string, error)
}

type AsyncDestinationManager interface {
	AsyncUploadAndTransformManager
	Poll(ctx context.Context, pollInput AsyncPoll) PollStatusResponse
	GetUploadStats(UploadStatsInput GetUploadStatsInput) GetUploadStatsResponse
}

type SimpleAsyncDestinationManager struct {
	UploaderAndTransformer AsyncUploadAndTransformManager
}

func (m SimpleAsyncDestinationManager) Upload(ctx context.Context, asyncDestStruct *AsyncDestinationStruct) AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(AsyncUploadOutput)
}

func (m SimpleAsyncDestinationManager) Poll(_ context.Context, _ AsyncPoll) PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(PollStatusResponse)
}

func (m SimpleAsyncDestinationManager) GetUploadStats(GetUploadStatsInput) GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(GetUploadStatsResponse)
}

func (m SimpleAsyncDestinationManager) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type PollStatusResponse struct {
	Complete             bool
	InProgress           bool
	StatusCode           int
	HasFailed            bool
	HasWarning           bool
	FailedJobParameters  string
	WarningJobParameters string
	Error                string `json:"error"`
}

type AsyncUploadOutput struct {
	ImportingJobIDs     []int64
	ImportingParameters stdjson.RawMessage
	FailedJobIDs        []int64
	SucceededJobIDs     []int64
	SuccessResponse     string
	FailedReason        string
	AbortJobIDs         []int64
	AbortReason         string
	ImportingCount      int
	FailedCount         int
	AbortCount          int
	DestinationID       string
}

type AsyncPoll struct {
	ImportId    string `json:"importId"`
	ImportCount int    `json:"importCount"`
}

type AsyncJob struct {
	Message  map[string]any `json:"message"`
	Metadata map[string]any `json:"metadata"`
}

type AsyncUploadT struct {
	Config   map[string]any `json:"config"`
	Input    []AsyncJob     `json:"input"`
	DestType string         `json:"destType"`
}

type MetaDataT struct {
	CSVHeaders string `json:"csvHeader"`
}
type ImportParameters struct {
	ImportId    any `json:"importId"`
	ImportCount int `json:"importCount"`
}

type AsyncDestinationStruct struct {
	ImportingJobIDs      []int64
	FailedJobIDs         []int64
	Exists               bool
	Size                 int
	CreatedAt            time.Time
	FileName             string
	Count                int
	CanUpload            bool
	UploadInProgress     bool
	UploadMutex          sync.RWMutex
	DestinationUploadURL string
	Destination          *backendconfig.DestinationT
	Manager              AsyncDestinationManager
	PartFileNumber       int
	SourceJobRunID       string

	// Maps jobID to various metadata

	AttemptNums       map[int64]int
	FirstAttemptedAts map[int64]time.Time
	PartitionIDs      map[int64]string
	JobParameters     map[int64]stdjson.RawMessage
}

type GetUploadStatsInput struct {
	FailedJobParameters  string
	Parameters           stdjson.RawMessage
	ImportingList        []*jobsdb.JobT
	WarningJobParameters string
}

type EventStatMeta struct {
	FailedKeys     []int64
	AbortedKeys    []int64
	WarningKeys    []int64
	SucceededKeys  []int64
	FailedReasons  map[int64]string
	AbortedReasons map[int64]string
	WarningReasons map[int64]string
}

type GetUploadStatsResponse struct {
	StatusCode int           `json:"statusCode"`
	Metadata   EventStatMeta `json:"metadata"`
	Error      string        `json:"error"`
}

func GetMarshalledData(payload string, jobID int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetBatchRouterConfigInt64(key, destType string, defaultValue int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func GetBatchRouterConfigStringMap(key, destType string, defaultValue []string) []string {
	_ = "STUB: not implemented"
	return nil
}
