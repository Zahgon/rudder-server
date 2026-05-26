package utils

import (
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

var EmptyPayload = []byte(`{}`)

const (
	DRAIN_ERROR_CODE = "410"
	// transformation(router or batch)
	ERROR_AT_TF = "transformation"
	// event delivery
	ERROR_AT_DEL = "delivery"
	// custom destination manager
	ERROR_AT_CUST = "custom"

	DrainReasonDestNotFound      = "destination is not available in the config"
	DrainReasonDestDisabled      = "destination is disabled"
	DrainReasonDestAbort         = "destination configured to abort"
	DrainReasonJobRunIDCancelled = "cancelled jobRunID"
	DrainReasonJobExpired        = "job expired"
)

type DestinationWithSources struct {
	Destination backendconfig.DestinationT
	Sources     []backendconfig.SourceT
}

type DrainStats struct {
	Count     int
	Reasons   []string
	Workspace string
}

type SendPostResponse struct {
	StatusCode          int
	ResponseContentType string
	ResponseBody        []byte
}

type JobParameters struct {
	SourceID                string `json:"source_id"`
	DestinationID           string `json:"destination_id"`
	ReceivedAt              string `json:"received_at"`
	TransformAt             string `json:"transform_at"`
	SourceTaskRunID         string `json:"source_task_run_id"`
	SourceJobID             string `json:"source_job_id"`
	SourceJobRunID          string `json:"source_job_run_id"`
	SourceDefinitionID      string `json:"source_definition_id"`
	DestinationDefinitionID string `json:"destination_definition_id"`
	SourceCategory          string `json:"source_category"`
	RecordID                any    `json:"record_id"`
	MessageID               string `json:"message_id"`
	EventName               string `json:"event_name"`
	EventType               string `json:"event_type"`
	WorkspaceID             string `json:"workspaceId"`
	RudderAccountID         string `json:"rudderAccountId"`
	DontBatch               bool   `json:"dontBatch"`
	TraceParent             string `json:"traceparent"`
}

// ParseReceivedAtTime parses the [ReceivedAt] field and returns the parsed time or a zero value time if parsing fails
func (jp *JobParameters) ParseReceivedAtTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// rawMsg passed must be a valid JSON
func EnhanceJSON(rawMsg []byte, key, val string) []byte { _ = "STUB: not implemented"; return nil }

func EnhanceJsonWithTime(t time.Time, key string, resp []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func IsNotEmptyString(s string) bool { _ = "STUB: not implemented"; return false }

type Drainer interface {
	Drain(
		createdAt time.Time,
		destID string,
		sourceJobRunID string,
	) (bool, string)
}

func NewDrainer(
	conf *config.Config,
	destDrainFunc func(string) (*DestinationWithSources, bool),
) Drainer {
	_ = "STUB: not implemented"
	return *new(Drainer)
}

type drainer struct {
	destinationIDs config.ValueLoader[[]string]
	jobRunIDs      config.ValueLoader[[]string]

	destinationResolver func(string) (*DestinationWithSources, bool)
	retentionTimesMu    sync.Mutex
	retentionTimes      map[string]config.ValueLoader[time.Duration]
}

func (d *drainer) Drain(
	createdAt time.Time,
	destID string,
	sourceJobRunID string,
) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (d *drainer) getRetentionTimeForDestination(destID string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func UpdateProcessedEventsMetrics(statsHandle stats.Stats, module, destType string, statusList []*jobsdb.JobStatusT, jobIDConnectionDetailsMap map[int64]jobsdb.ConnectionDetails) {
	_ = "STUB: not implemented"
	return
}
