//go:generate mockgen -destination=../../mocks/services/debugger/uploader.go -package mock_debugger github.com/rudderlabs/rudder-server/services/debugger TransformerAny

package debugger

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/sysUtils"
)

var (
	pkgLogger logger.Logger
	Http      sysUtils.HttpI = sysUtils.NewHttp()
)

type Uploader[E any] interface {
	Start()
	Stop()
	RecordEvent(data E)
}

type TransformerAny interface {
	Transformer[any]
}
type Transformer[E any] interface {
	Transform(data []E) ([]byte, error)
}

type uploaderImpl[E any] struct {
	url                                    string
	transformer                            Transformer[E]
	eventBatchChannel                      chan E
	eventBufferLock                        sync.RWMutex
	eventBuffer                            []E
	Client                                 sysUtils.HTTPClientI
	maxBatchSize, maxRetry, maxESQueueSize config.ValueLoader[int]
	batchTimeout, retrySleep               config.ValueLoader[time.Duration]
	region                                 string
	authorizer                             identity.Authorizer

	bgWaitGroup sync.WaitGroup
}

func init() {
	pkgLogger = logger.NewLogger().Child("debugger")
}

func (uploader *uploaderImpl[E]) Setup() {
	_ = "STUB: not implemented"
	// Number of events that are batched before sending events to control plane
	return
}

func New[E any](url string, authorizer identity.Authorizer, transformer Transformer[E]) Uploader[E] {
	_ = "STUB: not implemented"
	return nil
}

func (uploader *uploaderImpl[E]) Start() { _ = "STUB: not implemented"; return }

func (uploader *uploaderImpl[E]) Stop() { _ = "STUB: not implemented"; return }

// RecordEvent is used to put the event batch in the eventBatchChannel,
// which will be processed by handleEvents.
func (uploader *uploaderImpl[E]) RecordEvent(data E) { _ = "STUB: not implemented"; return }

func (uploader *uploaderImpl[E]) uploadEvents(eventBuffer []E) {
	_ = "STUB: not implemented"
	// Upload to a Config Backend
	return
}

// Sending live events to Config Backend

// Refresh the connection

func (uploader *uploaderImpl[E]) handleEvents() { _ = "STUB: not implemented"; return }

// If eventBuffer size is more than maxESQueueSize, Delete oldest.

// Append to request buffer

func (uploader *uploaderImpl[E]) flushEvents(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}
