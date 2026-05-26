package sourcedebugger

//go:generate mockgen -destination=./mocks/mock.go -package=mocks github.com/rudderlabs/rudder-server/services/debugger/source SourceDebugger
import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/debugger"
	"github.com/rudderlabs/rudder-server/services/debugger/cache"
)

// GatewayEventBatchT is a structure to hold batch of events
type GatewayEventBatchT struct {
	WriteKey   string
	EventBatch []byte
}

// EventUploadT is a structure to hold actual event data
type EventUploadT map[string]any

// EventUploadBatchT is a structure to hold batch of events
type EventUploadBatchT struct {
	WriteKey   string
	ReceivedAt string
	Batch      []EventUploadT
}

type SourceDebugger interface {
	RecordEvent(writeKey string, eventBatch []byte) bool
	Stop()
}

type Handle struct {
	started             bool
	uploader            debugger.Uploader[*GatewayEventBatchT]
	configBackendURL    string
	disableEventUploads config.ValueLoader[bool]
	log                 logger.Logger
	eventsCache         cache.Cache[[]byte]

	uploadEnabledWriteKeysMu sync.RWMutex
	uploadEnabledWriteKeys   []string

	ctx         context.Context
	cancel      func()
	initialized chan struct{}
	done        chan struct{}
}

func NewHandle(backendConfig backendconfig.BackendConfig) (SourceDebugger, error) {
	_ = "STUB: not implemented"
	return *new(SourceDebugger), nil
}

// Start initializes this module
func (h *Handle) start(backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

func (h *Handle) Stop() { _ = "STUB: not implemented"; return }

// RecordEvent is used to put the event batch in the eventBatchChannel,
// which will be processed by handleEvents.
func (h *Handle) RecordEvent(writeKey string, eventBatch []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if writeKey part of enabled sources

func (h *Handle) updateConfig(config map[string]backendconfig.ConfigT) {
	_ = "STUB: not implemented"
	return
}

func (h *Handle) backendConfigSubscriber(backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

// recordHistoricEvents sends the events collected in cache as live events.
// This is called on config update.
// IMP: The function must be called before releasing configSubscriberLock lock to ensure the order of RecordEvent call
func (h *Handle) recordHistoricEvents(uploadEnabledWriteKeys []string) {
	_ = "STUB: not implemented"
	return
}

type EventUploader struct {
	log logger.Logger
}

func NewEventUploader(log logger.Logger) *EventUploader { _ = "STUB: not implemented"; return nil }

func (e *EventUploader) Transform(eventBuffer []*GatewayEventBatchT) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add the receivedAt time to each event

func NewNoOpService() SourceDebugger { _ = "STUB: not implemented"; return *new(SourceDebugger) }

type noopService struct{}

func (*noopService) Start(_ backendconfig.BackendConfig) { _ = "STUB: not implemented"; return }

func (*noopService) RecordEvent(_ string, _ []byte) bool { _ = "STUB: not implemented"; return false }

func (*noopService) Stop() { _ = "STUB: not implemented"; return }
