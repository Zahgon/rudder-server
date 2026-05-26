package workerpool

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

// WorkerPool manages a pool of workers and their lifecycle
type WorkerPool interface {
	// PingWorker instructs the pool to ping the worker for the given partition
	PingWorker(partition string)

	// Shutdown stops all workers in the pool and waits for them to stop
	Shutdown()

	// Size returns the number of workers in the pool
	Size() int
}

// WorkerSupplier is a function able to create a new worker for the given partition
type WorkerSupplier func(partition string) Worker

// Worker is a worker that can be pinged for work and stopped by the worker pool when it is idle
type Worker interface {
	// Work triggers the worker to work and returns true if it did work, false otherwise.
	// The worker's idle time is calculated based on the return value of this method
	Work() bool

	// SleepDurations returns the sleep durations for the worker (min, max), i.e. how long the worker should sleep if it has no work to do
	SleepDurations() (min, max time.Duration)

	// Stop stops the worker and waits until all its goroutines have stopped
	Stop()
}

// WithCleanupPeriod option sets the cleanup period for the worker pool
func WithCleanupPeriod(cleanupPeriod time.Duration) func(*workerPool) {
	_ = "STUB: not implemented"
	return nil
}

// WithIdleTimeout option sets the idle timeout for the worker pool
func WithIdleTimeout(idleTimeout time.Duration) func(*workerPool) {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new worker pool
func New(ctx context.Context, workerSupplier WorkerSupplier, logger logger.Logger, opts ...func(*workerPool)) WorkerPool {
	_ = "STUB: not implemented"
	return *new(WorkerPool)
}

// workerPool manages a pool of workers
type workerPool struct {
	logger   logger.Logger
	supplier WorkerSupplier

	cleanupPeriod time.Duration
	idleTimeout   time.Duration

	workersMu sync.RWMutex
	workers   map[string]*internalWorker

	lifecycle struct {
		ctx    context.Context
		cancel context.CancelFunc
		wg     sync.WaitGroup
	}
}

// PingWorker pings the worker for the given partition
func (wp *workerPool) PingWorker(partition string) { _ = "STUB: not implemented"; return }

// Shutdown stops all workers in the pull and waits for them to stop
func (wp *workerPool) Shutdown() { _ = "STUB: not implemented"; return }

// Size returns the number of workers in the pool
func (wp *workerPool) Size() int { _ = "STUB: not implemented"; return 0 }

// worker gets or creates a worker for the given partition
func (wp *workerPool) worker(partition string) *internalWorker {
	_ = "STUB: not implemented"
	return nil
}

// startCleanupLoop starts a loop that cleans up idle workers
func (wp *workerPool) startCleanupLoop() { _ = "STUB: not implemented"; return }
