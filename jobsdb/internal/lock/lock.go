package lock

import (
	"context"

	golock "github.com/viney-shih/go-lock"
)

// LockToken represents proof that a list lock has been acquired
type LockToken interface {
	listLockToken()
}

// Locker
type Locker struct {
	m *golock.CASMutex
}

func NewLocker() *Locker { _ = "STUB: not implemented"; return nil }

// RLock acquires a read lock
func (r *Locker) RLock() {
	_ = "STUB: not implemented"

	// RTryLockWithCtx tries to acquires a read lock with context and returns false if context is done, otherwise returns true.
	return
}

func (r *Locker) RTryLockWithCtx(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// RUnlock releases a read lock
func (r *Locker) RUnlock() {
	_ = "STUB: not implemented"

	// TryLockWithCtx tries to acquires a lock with context and returns false if context is done, otherwise returns true.
	return
}

func (r *Locker) TryLockWithCtx(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// Unlock releases a lock
func (r *Locker) Unlock() {
	_ = "STUB: not implemented"

	// WithLock acquires a lock for the duration that the provided function
	// is being executed. A token as proof of the lock is passed to the function.
	return
}

func (r *Locker) WithLock(f func(l LockToken)) { _ = "STUB: not implemented"; return }

// WithLockInCtx tries to acquires a lock until it succeeds or context times out. If it fails, return value is false otherwise true. And, executes the function `f`, if lock is acquired.
// A token as proof of the lock is passed to the function.
func (r *Locker) WithLockInCtx(ctx context.Context, f func(l LockToken) error) error {
	_ = "STUB: not implemented"
	return nil
}

// AsyncLock acquires a lock until the token is returned to the receiving channel
func (r *Locker) AsyncLockWithCtx(ctx context.Context) (LockToken, chan<- LockToken, error) {
	_ = "STUB: not implemented"
	return *new(LockToken), nil, nil
}

type lockToken struct{}

func (*lockToken) listLockToken() {
	_ = "STUB: not implemented"
	// no-op
	return
}
