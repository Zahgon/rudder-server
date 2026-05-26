package preprocessdelay

import (
	"context"
	"time"
)

// NewHandle creates a new Handle that ensures that at least 'delay' time has
// passed since the most recent JobReceivedAt time before Sleep returns.
// If delay is less than or equal to zero, a no-op Handle is returned.
func NewHandle(delay time.Duration, sleeper Sleeper) Handle {
	_ = "STUB: not implemented"
	return *new(Handle)
}

type Handle interface {
	// JobReceivedAt records the time a job was received. It should be called
	// for each job that is part of the same processing batch.
	// If called multiple times, the most recent time is kept.
	JobReceivedAt(receivedAt time.Time)
	// Sleep sleeps until at least 'delay' time has passed since the most recent
	// JobReceivedAt time. If 'delay' time has already passed, it returns
	// immediately. If JobReceivedAt was never called, it returns immediately.
	// It returns any error returned by the Sleeper function.
	// The context can be used to cancel the sleep.
	Sleep(ctx context.Context) error
}

type Sleeper func(ctx context.Context, d time.Duration) error

type handle struct {
	delay                time.Duration
	sleeper              Sleeper
	mostRecentReceivedAt time.Time
}

func (h *handle) JobReceivedAt(receivedAt time.Time) { _ = "STUB: not implemented"; return }

func (h *handle) Sleep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type nopHandle struct{}

func (h *nopHandle) JobReceivedAt(receivedAt time.Time) { _ = "STUB: not implemented"; return }

func (h *nopHandle) Sleep(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
