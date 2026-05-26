package types

import (
	"database/sql"
	"encoding/json"
	"time"
)

type SyncerConfig struct {
	ConnInfo string
	Label    string
}

const (
	CoreReportingLabel      = "core"
	WarehouseReportingLabel = "warehouse"

	SupportedTransformerApiVersion = 2

	DefaultReportingEnabled = true
	DefaultReplayEnabled    = false
)

const (
	DiffStatus        = "diff"
	BotFlaggedStatus  = "bot_flagged"
	BotDetectedStatus = "bot_detected"

	// Module names
	BOT_MANAGEMENT         = "bot_management"
	EVENT_BLOCKING         = "event_blocking"
	GATEWAY                = "gateway"
	DESTINATION_FILTER     = "destination_filter"
	SOURCE_HYDRATION       = "source_hydration"
	TRACKINGPLAN_VALIDATOR = "tracking_plan_validator"
	USER_TRANSFORMER       = "user_transformer"
	EVENT_FILTER           = "event_filter"
	DEST_TRANSFORMER       = "dest_transformer"
	ROUTER                 = "router"
	BATCH_ROUTER           = "batch_router"
	WAREHOUSE              = "warehouse"
)

var (
	// Tracking plan validation states
	SUCCEEDED_WITHOUT_VIOLATIONS = "succeeded_without_violations"
	SUCCEEDED_WITH_VIOLATIONS    = "succeeded_with_violations"
)

type SyncSource struct {
	SyncerConfig
	DbHandle *sql.DB
}

type StatusDetail struct {
	Status         string            `json:"state"`
	Count          int64             `json:"count"`
	StatusCode     int               `json:"statusCode"`
	SampleResponse string            `json:"sampleResponse"`
	SampleEvent    json.RawMessage   `json:"sampleEvent"`
	EventName      string            `json:"eventName"`
	EventType      string            `json:"eventType"`
	ErrorType      string            `json:"errorType"`
	ViolationCount int64             `json:"violationCount"`
	StatTags       map[string]string `json:"-"`
	FailedMessages []*FailedMessage  `json:"-"`
}

type ErrorDetails struct {
	Code    string
	Message string
}

type FailedMessage struct {
	MessageID  string    `json:"messageId"`
	ReceivedAt time.Time `json:"receivedAt"`
}

type ReportByStatus struct {
	InstanceDetails
	ConnectionDetails
	PUDetails
	ReportMetadata
	StatusDetail *StatusDetail
}

type InstanceDetails struct {
	WorkspaceID string `json:"workspaceId"`
	Namespace   string `json:"namespace"`
	InstanceID  string `json:"instanceId"`
}

type ReportMetadata struct {
	ReportedAt        int64 `json:"reportedAt"`
	SampleEventBucket int64 `json:"bucket"`
}

type Metric struct {
	InstanceDetails
	ConnectionDetails
	PUDetails
	ReportMetadata
	StatusDetails []*StatusDetail `json:"reports"`
}

// EDMetric => ErrorDetailMetric
type EDConnectionDetails struct {
	SourceID                string `json:"sourceId"`
	DestinationID           string `json:"destinationId"`
	SourceDefinitionId      string `json:"sourceDefinitionId"`
	DestinationDefinitionId string `json:"destinationDefinitionId"`
	DestType                string `json:"destinationDefinitionName"`
}

type EDInstanceDetails struct {
	WorkspaceID string `json:"workspaceId"`
	Namespace   string `json:"namespace"`
	InstanceID  string `json:"-"`
}

type EDErrorDetailsKey struct {
	StatusCode   int    `json:"statusCode"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	EventType    string `json:"eventType"`
	EventName    string `json:"eventName"`
}

type EDErrorDetails struct {
	EDErrorDetailsKey
	SampleResponse string          `json:"sampleResponse"`
	SampleEvent    json.RawMessage `json:"sampleEvent"`
	ErrorCount     int64           `json:"count"`
}

type EDReportsDB struct {
	EDErrorDetails
	EDInstanceDetails
	EDConnectionDetails
	ReportMetadata

	PU    string `json:"reportedBy"`
	Count int64  `json:"-"`
}

type EDReportMapValue struct {
	Count          int64
	SampleResponse string
	SampleEvent    json.RawMessage
}

// EDMetric The structure in which the error detail data is being sent to reporting service
type EDMetric struct {
	EDInstanceDetails
	PU string `json:"reportedBy"`

	ReportMetadata

	EDConnectionDetails

	Errors []EDErrorDetails `json:"errors"`

	Count int64 `json:"-"`
}

type ConnectionDetails struct {
	SourceID                string `json:"sourceId"`
	DestinationID           string `json:"destinationId"`
	SourceTaskRunID         string `json:"sourceTaskRunId"`
	SourceJobID             string `json:"sourceJobId"`
	SourceJobRunID          string `json:"sourceJobRunId"`
	SourceDefinitionID      string `json:"sourceDefinitionId"`
	DestinationDefinitionID string `json:"DestinationDefinitionId"`
	SourceCategory          string `json:"sourceCategory"`
	TransformationID        string `json:"transformationId"`
	TransformationVersionID string `json:"transformationVersionId"`
	TrackingPlanID          string `json:"trackingPlanId"`
	TrackingPlanVersion     int    `json:"trackingPlanVersion"`
}
type PUDetails struct {
	InPU       string `json:"inReportedBy"`
	PU         string `json:"reportedBy"`
	TerminalPU bool   `json:"terminalState"`
	InitialPU  bool   `json:"initialState"`
}

type PUReportedMetric struct {
	ConnectionDetails
	PUDetails
	StatusDetail *StatusDetail
}

// ErrorMetricParams holds the parameters needed for converting PUReportedMetric to EDReportsDB
type ErrorMetricParams struct {
	WorkspaceID             string
	Namespace               string
	InstanceID              string
	DestType                string
	DestinationDefinitionID string
	ErrorDetails            ErrorDetails
}

// PUReportedMetricToEDReportsDB converts PUReportedMetric to EDReportsDB
// This avoids deep copying by extracting only the fields needed for error reporting
func PUReportedMetricToEDReportsDB(
	metric *PUReportedMetric,
	params ErrorMetricParams,
) *EDReportsDB {
	_ = "STUB: not implemented"
	return nil
}

func CreatePUDetails(inPU, pu string, terminalPU, initialPU bool) *PUDetails {
	_ = "STUB: not implemented"
	return nil
}

func AssertSameKeys[V1, V2 any](m1 map[string]V1, m2 map[string]V2) {
	_ = "STUB: not implemented"
	return
}

// TODO improve msg

// TODO improve msg

// AssertKeysSubset checks that all keys from the subset map are a subset of keys in the superset map
// The superset map can have additional keys that are not in the subset map
func AssertKeysSubset[V1, V2 any](superset map[string]V1, subset map[string]V2) {
	_ = "STUB: not implemented"
	return
}

// TODO improve msg

// ErrorDetailGroupKey represents the key for grouping error detail reports
type ErrorDetailGroupKey struct {
	SourceID      string
	DestinationID string
	PU            string
	EventType     string
}
