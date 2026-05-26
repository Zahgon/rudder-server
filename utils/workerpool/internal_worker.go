package workerpool

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

func newInternalWorker(partition string, log logger.Logger, delegate Worker) *internalWorker {
	_ = "STUB: not implemented"
	return nil
}

type internalWorker struct {
	partition string
	delegate  Worker
	logger    logger.Logger

	ping      chan struct{} // ping channel triggers the worker to start working
	lifecycle struct {      // worker lifecycle related fields
		stoppedMu sync.Mutex
		stopped   bool
		ctx       context.Context    // worker context
		cancel    context.CancelFunc // worker context cancel function
		wg        sync.WaitGroup     // worker wait group

		idleMu    sync.RWMutex // idle mutex
		idleSince time.Time    // idle since
	}
}

// start starts the various worker goroutines
func (w *internalWorker) start() {
	_ = "STUB: not implemented"
	// ping loop
	return
}

func (w *internalWorker) setIdleSince(t time.Time) { _ = "STUB: not implemented"; return }

// Ping triggers the worker to pick more jobs
func (w *internalWorker) Ping() { _ = "STUB: not implemented"; return }

// IdleSince returns the time when the worker was last idle. If the worker is not idle, it returns a zero time.
func (w *internalWorker) IdleSince() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Stop stops the worker and waits until all its goroutines have stopped
func (w *internalWorker) Stop() { _ = "STUB: not implemented"; return }
