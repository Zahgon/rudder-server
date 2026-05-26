package types

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

var (
	ErrProcessorStopping           = errors.New("processor is stopping")
	ErrPermanentTransformerFailure = errors.New("transformer permanent failure")
)

// SingularEventT single event structure
type SingularEventT map[string]any

// GetRudderEventVal returns the value corresponding to the key in the message structure
func GetRudderEventVal(key string, rudderEvent SingularEventT) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

type SingularEventWithReceivedAt struct {
	SingularEvent SingularEventT
	ReceivedAt    time.Time
}

// GatewayBatchRequest batch request structure
type GatewayBatchRequest struct {
	Batch      []SingularEventT `json:"batch"`
	RequestIP  string           `json:"requestIP"`
	ReceivedAt time.Time        `json:"receivedAt"`
}

type TransformerEvent struct {
	Message     SingularEventT             `json:"message"`
	Metadata    Metadata                   `json:"metadata"`
	Destination backendconfig.DestinationT `json:"destination"`
	Connection  backendconfig.Connection   `json:"connection"`
	Libraries   []backendconfig.LibraryT   `json:"libraries"`
	Credentials []Credential               `json:"credentials"`
}

// UserTransformerEvent is the event sent to the user transformer, which is similar to a
// [TransformerEvent] but without the connection and with a simplified destination structure
type UserTransformerEvent struct {
	Message     SingularEventT `json:"message"`
	Metadata    Metadata       `json:"metadata"`
	Destination struct {
		Transformations []struct {
			VersionID string
		}
	} `json:"destination"`
	Libraries   []backendconfig.LibraryT `json:"libraries,omitempty"`
	Credentials []Credential             `json:"credentials,omitempty"`
}

// TrackingPlanValidationEvent is the event sent to the trackingplan transformer,
// whose fields are a subset of the [TransformerEvent]
type TrackingPlanValidationEvent struct {
	Message  SingularEventT `json:"message"`
	Metadata Metadata       `json:"metadata"`
}

// CompactedTransformerEvent is a subset of the [TransformerEvent] only containing the message and metadata fields
type CompactedTransformerEvent struct {
	Message  SingularEventT `json:"message"`
	Metadata Metadata       `json:"metadata"`
}

// CompactedTransformRequest is the request structure for a compacted transformer request payload.
// It contains a list of [CompactedTransformerEvent]s and two lookup maps,
// one for destinations and one for connections.
//
// The destinations and connections are used to look up the actual destination and connection
// objects based on their IDs, which are included in the [CompactedTransformerEvent]s.
//
// This allows for a more compact representation of the request payload, as the full destination
// and connection objects do not need to be included in each event.
type CompactedTransformRequest struct {
	Input        []CompactedTransformerEvent           `json:"input"`
	Destinations map[string]backendconfig.DestinationT `json:"destinations"`
	Connections  map[string]backendconfig.Connection   `json:"connections"`
}

func (ctr *CompactedTransformRequest) ToTransformerEvents() []TransformerEvent {
	_ = "STUB: not implemented"
	return nil
}

// ToUserTransformerEvent removes the connection from the event
// along with pruning the destination to only include the transformation ID and VersionID
// before sending it to the transformer thereby reducing the payload size
func (e *TransformerEvent) ToUserTransformerEvent() *UserTransformerEvent {
	_ = "STUB: not implemented"
	return nil
}

// ToTrackingPlanValidationEvent only keeps the message and metadata fields from the event
// before sending it to the trackingplan validator thereby reducing the payload size
func (e *TransformerEvent) ToTrackingPlanValidationEvent() *TrackingPlanValidationEvent {
	_ = "STUB: not implemented"
	return nil
}

type Credential struct {
	ID       string `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsSecret bool   `json:"isSecret"`
}

type Metadata struct {
	// job metadata
	JobID       int64  `json:"jobId"`
	WorkspaceID string `json:"workspaceId"`
	RudderID    string `json:"rudderId,omitempty"` // used for ordering events (e.g. can be userId or anonymousId)
	ReceivedAt  string `json:"receivedAt,omitempty"`
	PartitionID string `json:"partitionId,omitempty"`

	// event metadata
	MessageID string `json:"messageId"`
	EventName string `json:"eventName,omitempty"`
	EventType string `json:"eventType,omitempty"`

	// source metadata
	SourceID             string `json:"sourceId"`
	OriginalSourceID     string `json:"originalSourceId"` // for replayed events
	SourceDefinitionID   string `json:"sourceDefinitionId,omitempty"`
	SourceName           string `json:"sourceName"`
	SourceType           string `json:"sourceType"`
	SourceCategory       string `json:"sourceCategory"`
	SourceDefinitionType string `json:"-"`

	// retl metadata
	SourceJobID     string `json:"sourceJobId,omitempty"`
	SourceJobRunID  string `json:"sourceJobRunId,omitempty"`
	SourceTaskRunID string `json:"sourceTaskRunId,omitempty"`
	RecordID        any    `json:"recordId,omitempty"`

	// other metadata
	InstanceID  string `json:"instanceId,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	TraceParent string `json:"traceparent,omitempty"`

	// tracking plan metadata (available only during tracking plan transformations)
	TrackingPlanID      string                    `json:"trackingPlanId,omitempty"`
	TrackingPlanVersion int                       `json:"trackingPlanVersion,omitempty"`
	SourceTpConfig      map[string]map[string]any `json:"sourceTpConfig,omitempty"`
	MergedTpConfig      map[string]any            `json:"mergedTpConfig,omitempty"`

	// destination metadata (available after tracking plan)
	DestinationID           string `json:"destinationId"`
	DestinationName         string `json:"destinationName"`
	DestinationType         string `json:"destinationType"`
	DestinationDefinitionID string `json:"destinationDefinitionId,omitempty"`

	// user transformation metadata (available only during user transformation)
	TransformationID        string   `json:"transformationId,omitempty"`
	TransformationVersionID string   `json:"transformationVersionId,omitempty"`
	MessageIDs              []string `json:"messageIds,omitempty"` // set only by user transformer to indicate transformed event is part of group indicated by messageIDs
}

func (m Metadata) GetMessagesIDs() []string { _ = "STUB: not implemented"; return nil }

// CommonMetadata creates a new metadata instance keeping only common fields across events
//
//   - workspace id
//   - source metadata
//   - other metadata (instance id, namespace)
//   - destination metadata
func (m Metadata) CommonMetadata() *Metadata {
	_ = "STUB: not implemented"

	// job metadata
	return nil
}

// source metadata

// other metadata

// destination metadata (available after tracking plan)

type TransformerResponse struct {
	// Not marking this Singular Event, since this not a RudderEvent
	Output           map[string]any    `json:"output"`
	Metadata         Metadata          `json:"metadata"`
	StatusCode       int               `json:"statusCode"`
	Error            string            `json:"error"`
	ValidationErrors []ValidationError `json:"validationErrors"`
	StatTags         map[string]string `json:"statTags"`
}

type ValidationError struct {
	Type     string            `json:"type"`
	Message  string            `json:"message"`
	Meta     map[string]string `json:"meta"`
	Property string            `json:"property"`
}

// Response represents a Transformer response
type Response struct {
	Events         []TransformerResponse
	FailedEvents   []TransformerResponse
	MirrorFiltered bool
}

var responseDatetimePattern = regexp.MustCompile(
	`^\d{4}-\d{2}-\d{2}[T ]\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?(?:Z|[+-]\d{2}:\d{2})?$`,
)

// EqualResult holds the result of a Response comparison.
type EqualResult struct {
	Diff             string
	Equal            bool
	DatetimeForgiven bool // true when responses matched only because datetime strings were treated as equal
}

// EqualDetailed compares two Response structs using a two-pass strategy:
//  1. Strict comparison (exact match). If it passes, returns Equal with DatetimeForgiven=false.
//  2. Lax comparison (datetime strings forgiven). If it passes, returns Equal with DatetimeForgiven=true.
//  3. If both fail, returns the diff from the lax pass.
func (r *Response) EqualDetailed(v *Response) EqualResult {
	_ = "STUB: not implemented"
	return *new(EqualResult)
}

// First pass: strict comparison (no datetime forgiveness)

// Second pass: lax comparison (datetime strings forgiven)

// Equal compares two Response structs and returns true if they are equal
// regardless of the order of elements in the Events and FailedEvents slices.
// Matching datetime values in Output are treated as equal when both values
// match responseDatetimePattern.
func (r *Response) Equal(v *Response) (string, bool) { _ = "STUB: not implemented"; return "", false }

type EventParams struct {
	SourceJobRunId      string `json:"source_job_run_id"`
	SourceId            string `json:"source_id"`
	SourceTaskRunId     string `json:"source_task_run_id"`
	TraceParent         string `json:"traceparent"`
	DestinationID       string `json:"destination_id"`
	IsBot               bool   `json:"is_bot,omitempty"`
	BotName             string `json:"bot_name,omitempty"`
	BotURL              string `json:"bot_url,omitempty"`
	BotIsInvalidBrowser bool   `json:"bot_is_invalid_browser,omitempty"`
	BotAction           string `json:"bot_action,omitempty"`
	IsEventBlocked      bool   `json:"is_event_blocked,omitempty"`
}

type TransformerMetricLabels struct {
	Endpoint         string // hostname of the service
	DestinationType  string // BQ, etc.
	SourceType       string // webhook
	Language         string // js, python
	Stage            string // processor, router, gateway
	WorkspaceID      string // workspace identifier
	SourceID         string // source identifier
	DestinationID    string // destination identifier
	TransformationID string // transformation identifier
	Mirroring        bool
}

// ToStatsTag converts transformerMetricLabels to stats.Tags and includes legacy tags for backwards compatibility
func (t TransformerMetricLabels) ToStatsTag() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

// Legacy tags: to be removed

// ToLoggerFields converts the metric labels to a slice of logger.Fields
func (t TransformerMetricLabels) ToLoggerFields() []logger.Field {
	_ = "STUB: not implemented"
	return nil
}

// SrcHydrationEvent represents a single event in the hydration request/response
type SrcHydrationEvent struct {
	ID    string         `json:"id" required:"true"` // JobID of the event we are hydrating
	Event map[string]any `json:"event" required:"true"`
}

// SrcHydrationRequest represents the request format for source hydration API
type SrcHydrationRequest struct {
	Batch  []SrcHydrationEvent `json:"batch"`
	Source SrcHydrationSource  `json:"source"`
}

type SrcHydrationSource struct {
	ID               string                          `json:"id"`
	Config           json.RawMessage                 `json:"config"`
	InternalSecret   json.RawMessage                 `json:"internalSecret"`
	WorkspaceID      string                          `json:"workspaceId"`
	SourceDefinition backendconfig.SourceDefinitionT `json:"sourceDefinition"`
}

// SrcHydrationResponse represents the response format from source hydration API
type SrcHydrationResponse struct {
	// Batch is a required field containing hydration events
	Batch []SrcHydrationEvent `json:"batch" required:"true"`
}

func diffLists(listA, listB any, equalFn func(a, b any) bool) (extraA, extraB []any) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mark indexes in bValue that we already used

func strictObjectsEqual(left, right any) bool { _ = "STUB: not implemented"; return false }

func responseObjectsEqual(left, right any) bool { _ = "STUB: not implemented"; return false }

func transformerResponsesEqual(left, right TransformerResponse) bool {
	_ = "STUB: not implemented"
	return false
}

func responseValuesEqual(left, right any) bool { _ = "STUB: not implemented"; return false }

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
	MaxDepth:                10,
}

func formatListDiff(listA, listB any, extraA, extraB []any) string {
	_ = "STUB: not implemented"
	return ""
}
