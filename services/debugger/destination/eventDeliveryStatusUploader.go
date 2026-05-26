//go:generate mockgen -destination=../../../mocks/services/debugger/destination/eventDeliveryStatusUploader.go -package mock_destinationdebugger github.com/rudderlabs/rudder-server/services/debugger/destination/ DestinationDebugger

package destinationdebugger

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/debugger"
	"github.com/rudderlabs/rudder-server/services/debugger/cache"
)

// DeliveryStatusT is a structure to hold everything related to event delivery
type DeliveryStatusT struct {
	DestinationID string          `json:"destinationId"`
	SourceID      string          `json:"sourceId"`
	Payload       json.RawMessage `json:"payload"`
	AttemptNum    int             `json:"attemptNum"`
	JobState      string          `json:"jobState"`
	ErrorCode     string          `json:"errorCode"`
	ErrorResponse json.RawMessage `json:"errorResponse"`
	SentAt        string          `json:"sentAt"`
	EventName     string          `json:"eventName"`
	EventType     string          `json:"eventType"`
}

type DestinationDebugger interface {
	RecordEventDeliveryStatus(destinationID string, deliveryStatus *DeliveryStatusT) bool
	HasUploadEnabled(destID string) bool
	Stop()
}

type Handle struct {
	configBackendURL                  string
	log                               logger.Logger
	started                           bool
	disableEventDeliveryStatusUploads config.ValueLoader[bool]
	eventsDeliveryCache               cache.Cache[*DeliveryStatusT]
	uploader                          debugger.Uploader[*DeliveryStatusT]
	uploadEnabledDestinationIDs       map[string]bool
	uploadEnabledDestinationIDsMu     sync.RWMutex
	ctx                               context.Context
	cancel                            func()
	initialized                       chan struct{}
	done                              chan struct{}
}

func NewHandle(backendConfig backendconfig.BackendConfig) (DestinationDebugger, error) {
	_ = "STUB: not implemented"
	return *new(DestinationDebugger), nil
}

func (h *Handle) start(backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

func (h *Handle) Stop() { _ = "STUB: not implemented"; return }

func NewEventDeliveryStatusUploader(log logger.Logger) *EventDeliveryStatusUploader {
	_ = "STUB: not implemented"
	return nil
}

type EventDeliveryStatusUploader struct {
	log logger.Logger
}

// RecordEventDeliveryStatus is used to put the delivery status in the deliveryStatusesBatchChannel,
// which will be processed by handleJobs.
func (h *Handle) RecordEventDeliveryStatus(destinationID string, deliveryStatus *DeliveryStatusT) bool {
	_ = "STUB: not implemented"
	// if disableEventDeliveryStatusUploads is true, return;
	return false
}

// Check if destinationID part of enabled destinations, if not then push the job in cache to keep track

func (h *Handle) HasUploadEnabled(destID string) bool { _ = "STUB: not implemented"; return false }

func (e *EventDeliveryStatusUploader) Transform(deliveryStatusesBuffer []*DeliveryStatusT) ([]byte, error) {
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

func (h *Handle) recordHistoricEventsDelivery(destinationIDs []string) {
	_ = "STUB: not implemented"
	return
}

func NewNoOpService() DestinationDebugger {
	_ = "STUB: not implemented"
	return *new(DestinationDebugger)
}

type noopService struct{}

func (*noopService) RecordEventDeliveryStatus(_ string, _ *DeliveryStatusT) bool {
	_ = "STUB: not implemented"
	return false
}

func (*noopService) HasUploadEnabled(_ string) bool { _ = "STUB: not implemented"; return false }

func (*noopService) Start(_ backendconfig.BackendConfig) { _ = "STUB: not implemented"; return }

func (*noopService) Stop() { _ = "STUB: not implemented"; return }
