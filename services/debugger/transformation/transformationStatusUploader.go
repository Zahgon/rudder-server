package transformationdebugger

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
	"github.com/rudderlabs/rudder-server/services/debugger"
	"github.com/rudderlabs/rudder-server/services/debugger/cache"
)

type TransformationStatusT struct {
	SourceID              string
	DestID                string
	Destination           *backendconfig.DestinationT
	UserTransformedEvents []types.TransformerEvent
	EventsByMessageID     map[string]types.SingularEventWithReceivedAt
	FailedEvents          []types.TransformerResponse
	UniqueMessageIds      map[string]struct{}
}

// TransformStatusT is a structure to hold transformation status
type TransformStatusT struct {
	TransformationID string                `json:"transformationId"`
	SourceID         string                `json:"sourceId"`
	DestinationID    string                `json:"destinationId"`
	EventBefore      *EventBeforeTransform `json:"eventBefore"`
	EventsAfter      *EventsAfterTransform `json:"eventsAfter"`
	IsError          bool                  `json:"error"`
}

type EventBeforeTransform struct {
	EventName  string               `json:"eventName"`
	EventType  string               `json:"eventType"`
	ReceivedAt string               `json:"receivedAt"`
	Payload    types.SingularEventT `json:"payload"`
}

type EventPayloadAfterTransform struct {
	EventName string               `json:"eventName"`
	EventType string               `json:"eventType"`
	Payload   types.SingularEventT `json:"payload"`
}

type EventsAfterTransform struct {
	ReceivedAt    string                        `json:"receivedAt"`
	IsDropped     bool                          `json:"isDropped"`
	Error         string                        `json:"error"`
	StatusCode    int                           `json:"statusCode"`
	EventPayloads []*EventPayloadAfterTransform `json:"payload"`
}

type UploadT struct {
	Payload []*TransformStatusT `json:"payload"`
}

type Handle struct {
	configBackendURL               string
	started                        bool
	disableTransformationUploads   config.ValueLoader[bool]
	limitEventsInMemory            config.ValueLoader[int]
	uploader                       debugger.Uploader[*TransformStatusT]
	log                            logger.Logger
	transformationCacheMap         cache.Cache[TransformationStatusT]
	uploadEnabledTransformations   map[string]bool
	uploadEnabledTransformationsMu sync.RWMutex
	ctx                            context.Context
	cancel                         func()
	initialized                    chan struct{}
	done                           chan struct{}
}

type TransformationDebugger interface {
	IsUploadEnabled(id string) bool
	RecordTransformationStatus(transformStatus *TransformStatusT)
	UploadTransformationStatus(tStatus *TransformationStatusT) bool
	Stop()
}

func NewHandle(backendConfig backendconfig.BackendConfig) (TransformationDebugger, error) {
	_ = "STUB: not implemented"
	return *new(TransformationDebugger), nil
}

type TransformationStatusUploader struct {
	log logger.Logger
}

func (h *Handle) IsUploadEnabled(id string) bool { _ = "STUB: not implemented"; return false }

// Start initializes this module
func (h *Handle) start(backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

func (h *Handle) Stop() { _ = "STUB: not implemented"; return }

// RecordTransformationStatus is used to put the transform event in the eventBatchChannel,
// which will be processed by handleEvents.
func (h *Handle) RecordTransformationStatus(transformStatus *TransformStatusT) {
	_ = "STUB: not implemented"
	// if disableTransformationUploads is true, return;
	return
}

func (t *TransformationStatusUploader) Transform(eventBuffer []*TransformStatusT) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handle) updateConfig(config map[string]backendconfig.ConfigT) {
	_ = "STUB: not implemented"
	return
}

func (h *Handle) backendConfigSubscriber(backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

// limit the number of stored events
func (ts *TransformationStatusT) Limit(
	limit int,
	transformation backendconfig.TransformationT,
) *TransformationStatusT {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) UploadTransformationStatus(tStatus *TransformationStatusT) bool {
	_ = "STUB: not implemented"
	return false
}

// nolint:forbidigo

// if disableTransformationUploads is true, return;

func getEventBeforeTransform(singularEvent types.SingularEventT, receivedAt time.Time) *EventBeforeTransform {
	_ = "STUB: not implemented"
	return nil
}

func getEventAfterTransform(singularEvent types.SingularEventT) *EventPayloadAfterTransform {
	_ = "STUB: not implemented"
	return nil
}

func getEventsAfterTransform(singularEvent types.SingularEventT, receivedAt time.Time) *EventsAfterTransform {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) recordHistoricTransformations(tIDs []string) { _ = "STUB: not implemented"; return }

func (h *Handle) processRecordTransformationStatus(tStatus *TransformationStatusT, tID string) {
	_ = "STUB: not implemented"
	return
}

func NewNoOpService() TransformationDebugger {
	_ = "STUB: not implemented"
	return *new(TransformationDebugger)
}

type noopService struct{}

func (*noopService) Start(_ backendconfig.BackendConfig) { _ = "STUB: not implemented"; return }

func (*noopService) IsUploadEnabled(_ string) bool { _ = "STUB: not implemented"; return false }

func (*noopService) RecordTransformationStatus(_ *TransformStatusT) {
	_ = "STUB: not implemented"
	return
}

func (*noopService) UploadTransformationStatus(_ *TransformationStatusT) bool {
	_ = "STUB: not implemented"
	return false
}

func (*noopService) Stop() { _ = "STUB: not implemented"; return }
