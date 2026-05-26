package misc

import (
	"context"
	"sync"
	"sync/atomic"
)

// NewAsyncInit returns a new AsyncInit object with the given expected initialization events count.
func NewAsyncInit(count int64) *AsyncInit { _ = "STUB: not implemented"; return nil }

// AsyncInit is a helper object to wait for multiple asynchronous initialization events.
type AsyncInit struct {
	mu    sync.Mutex
	count atomic.Int64
	c     chan struct{}
}

// Done decrements the initialization events count
func (ia *AsyncInit) Done() { _ = "STUB: not implemented"; return }

// Wait returns the channel that will be closed when the initialization events count reaches zero.
func (ia *AsyncInit) Wait() chan struct{} { _ = "STUB: not implemented"; return nil }

// WaitContext returns no error if initialization events happen before the provided context is done. It returns the context's error otherwise
func (ia *AsyncInit) WaitContext(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (ia *AsyncInit) channel() chan struct{} { _ = "STUB: not implemented"; return nil }
