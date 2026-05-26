package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/samber/lo"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	routerutils "github.com/rudderlabs/rudder-server/router/utils"
)

const RouterUnMarshalErrorCode = 599

// RouterJobT holds the router job and its related metadata
type RouterJobT struct {
	Message     json.RawMessage            `json:"message"`
	JobMetadata JobMetadataT               `json:"metadata"`
	Destination backendconfig.DestinationT `json:"destination"`
	Connection  backendconfig.Connection   `json:"connection"`
}

type SourceDest struct {
	SourceID, DestinationID string
}

type ConnectionWithID struct {
	ConnectionID string
	Connection   backendconfig.Connection
}

type DestinationJobs []DestinationJobT

// Hydrate jobs in the destination jobs' job metadata array
func (djs DestinationJobs) Hydrate(jobs map[int64]lo.Tuple2[*jobsdb.JobT, routerutils.JobParameters]) {
	_ = "STUB: not implemented"
	return
}

// DestinationJobT holds the job to be sent to destination
// and metadata of all the router jobs from which this job is cooked up
type DestinationJobT struct {
	Message           json.RawMessage            `json:"batchedRequest"`
	JobMetadataArray  []JobMetadataT             `json:"metadata"` // multiple jobs may be batched in a single message
	Destination       backendconfig.DestinationT `json:"destination"`
	Connection        backendconfig.Connection   `json:"connection"`
	Batched           bool                       `json:"batched"`
	StatusCode        int                        `json:"statusCode"`
	Error             string                     `json:"error"`
	AuthErrorCategory string                     `json:"authErrorCategory"`
	StatTags          map[string]string          `json:"statTags"`
}

func (dj *DestinationJobT) MinJobID() int64 { _ = "STUB: not implemented"; return 0 }

// JobIDs returns the set of all job ids contained in the message
func (dj *DestinationJobT) JobIDs() map[int64]struct{} { _ = "STUB: not implemented"; return nil }

// JobMetadataT holds the job metadata
type JobMetadataT struct {
	UserID             string                    `json:"userId"`
	JobID              int64                     `json:"jobId"`
	SourceID           string                    `json:"sourceId"`
	SourceCategory     string                    `json:"sourceCategory"`
	DestinationID      string                    `json:"destinationId"`
	AttemptNum         int                       `json:"attemptNum"`
	ReceivedAt         string                    `json:"receivedAt"`
	CreatedAt          string                    `json:"createdAt"`
	FirstAttemptedAt   string                    `json:"firstAttemptedAt"`
	TransformAt        string                    `json:"transformAt"`
	WorkspaceID        string                    `json:"workspaceId"`
	Secret             json.RawMessage           `json:"secret"` // Populates transformer.ProxyRequestMetadata
	JobT               *jobsdb.JobT              `json:"jobsT,omitempty"`
	WorkerAssignedTime time.Time                 `json:"workerAssignedTime"`
	DestInfo           json.RawMessage           `json:"destInfo,omitempty"` // Populates transformer.ProxyRequestMetadata
	DontBatch          bool                      `json:"dontBatch"`
	TraceParent        string                    `json:"traceparent"`
	Parameters         routerutils.JobParameters `json:"-"`
}

// TransformMessageT is used to pass message to the transformer workers
type TransformMessageT struct {
	Data     []RouterJobT `json:"input"`
	DestType string       `json:"destType"`
}

func (tm *TransformMessageT) Compacted() *CompactedTransformMessageT {
	_ = "STUB: not implemented"
	return nil
}

type CompactedTransformMessageT struct {
	Data []struct {
		Message     json.RawMessage `json:"message"`
		JobMetadata JobMetadataT    `json:"metadata"`
	} `json:"input"`
	DestType     string                                `json:"destType"`
	Destinations map[string]backendconfig.DestinationT `json:"destinations"`
	Connections  map[string]backendconfig.Connection   `json:"connections"`
}

// Dehydrate JobT information from RouterJobT.JobMetadata returning the dehydrated message along with the jobs
func (tm *TransformMessageT) Dehydrate() (*TransformMessageT, map[int64]lo.Tuple2[*jobsdb.JobT, routerutils.JobParameters]) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JobIDs returns the set of all job ids of the jobs in the message
func (tm *TransformMessageT) JobIDs() map[int64]struct{} { _ = "STUB: not implemented"; return nil }

func NewEventTypeThrottlingCost(m map[string]any) (v EventTypeThrottlingCost) {
	_ = "STUB: not implemented"
	return *new(EventTypeThrottlingCost)
}

type EventTypeThrottlingCost map[string]any

func (e *EventTypeThrottlingCost) Cost(eventType string) (cost int64) {
	_ = "STUB: not implemented"
	return 0
}

var (
	// ErrContextCancelled is returned when the context is cancelled
	ErrContextCancelled = errors.New("context cancelled")
	// ErrJobOrderBlocked is returned when the job is blocked by another job discarded by the router in the same loop
	ErrJobOrderBlocked = errors.New("blocked")
	// ErrWorkerNoSlot is returned when the worker doesn't have an available slot
	ErrWorkerNoSlot = errors.New("no slot")
	// ErrJobBackoff is returned when the job is backoffed
	ErrJobBackoff = errors.New("backoff")
	// ErrDestinationThrottled is returned when the destination is being throttled
	ErrDestinationThrottled = errors.New("throttled")
	// ErrBarrierExists is returned when a job ordering barrier exists for the job's ordering key
	ErrBarrierExists = errors.New("barrier")
)
