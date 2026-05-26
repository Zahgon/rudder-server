package eventorder

import (
	"errors"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

var ErrUnsupportedState = errors.New("unsupported state")

type OptFn func(b *Barrier)

// WithMetadata includes the provided metadata in the error messages
func WithMetadata(metadata map[string]string) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

// WithEventOrderKeyThreshold sets the maximum number of concurrent jobs for a given key. After this limit is reached, the barrier will be disabled for this key.
func WithEventOrderKeyThreshold(eventOrderKeyThreshold config.ValueLoader[int]) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

// WithDisabledStateDuration sets the duration for which the barrier will remain in the disabled state after the concurrency limit has been reached
func WithDisabledStateDuration(disabledStateDuration config.ValueLoader[time.Duration]) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

// WithHalfEnabledStateDuration sets the duration for which the barrier will remain in the half-enabled state
func WithHalfEnabledStateDuration(halfEnabledStateDuration config.ValueLoader[time.Duration]) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

// WithDrainConcurrencyLimit sets the maximum number of concurrent jobs for a given key when the limiter is enabled (after a failed job has been drained, i.e. aborted)
func WithDrainConcurrencyLimit(drainLimit config.ValueLoader[int]) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

// WithDebugInfoProvider sets the debug info provider for the barrier (used for debugging purposes in case an illegal job sequence is detected)
func WithDebugInfoProvider(debugInfoProvider func(key BarrierKey) string) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

func WithOrderingDisabledCheckForBarrierKey(orderingDisabledForKey func(key BarrierKey) bool) OptFn {
	_ = "STUB: not implemented"
	return *new(OptFn)
}

// NewBarrier creates a new properly initialized Barrier
func NewBarrier(fns ...OptFn) *Barrier { _ = "STUB: not implemented"; return nil }

// Barrier is an abstraction for applying event ordering guarantees in the router.
//
// Events for the same key need to be processed in order, thus when an event fails but will be retried later by the router,
// we need to put all subsequent events for that key on hold until the failed event succeeds or fails
// in a terminal way.
//
// A barrier controls the concurrency of events in two places:
//
// 1. At entrance, before the event enters the pipeline.
//
// 2. Before actually trying to send the event, since events after being accepted by the router, they are
// processed asynchronously through buffered channels by separate goroutine(s) aka workers.
type Barrier struct {
	mu       sync.RWMutex // mutex to synchronize concurrent access to the barrier's methods
	queue    []command
	barriers map[BarrierKey]*barrierInfo
	metadata map[string]string

	eventOrderKeyThreshold   config.ValueLoader[int] // maximum number of concurrent jobs for a given key (0 means no threshold)
	disabledStateDuration    config.ValueLoader[time.Duration]
	halfEnabledStateDuration config.ValueLoader[time.Duration]

	drainLimit config.ValueLoader[int] // maximum number of concurrent jobs to accept after a previously failed job has been aborted

	debugInfo func(key BarrierKey) string

	orderingDisabledForKey func(key BarrierKey) bool
}

type BarrierKey struct {
	DestinationID, UserID, WorkspaceID string
}

func (bk *BarrierKey) String() string { _ = "STUB: not implemented"; return "" }

// Enter the barrier for this key and jobID. If there is not already a barrier for this key
// returns true, otherwise false along with the previous failed jobID if this is the cause of the barrier.
// Another scenario where a barrier might exist for a key is when the previous job has failed in an unrecoverable manner and the drain limiter is enabled.
func (b *Barrier) Enter(key BarrierKey, jobID int64) (accepted bool, previousFailedJobID *int64) {
	_ = "STUB: not implemented"
	return false, nil
}

// if no key threshold is set accept the job

// create a new barrier with the concurrency limiter enabled

// if the barrier is in a disabled state, accept the job

// if key threshold is reached, disable the barrier and accept the job

// if drain limit is reached, don't accept the job

// if there is a failed job in the barrier, only this job can enter the barrier

// if the job is finally accepted, add it to the active limiters

// Leave the barrier for this key and jobID. Leave acts as an undo operation for Enter, i.e.
// when a previously-entered job leaves the barrier it is as if this key and jobID didn't enter the barrier.
// Calling Leave is idempotent.
func (b *Barrier) Leave(key BarrierKey, jobID int64) { _ = "STUB: not implemented"; return }

// remove the job from the active limiters

// Peek returns the previously failed jobID for the given key, if any
func (b *Barrier) Peek(key BarrierKey) (previousFailedJobID *int64) {
	_ = "STUB: not implemented"
	return nil
}

// Wait returns true if the job for this key shouldn't continue, but wait (transition to a waiting state)
func (b *Barrier) Wait(key BarrierKey, jobID int64) (wait bool, previousFailedJobID *int64) {
	_ = "STUB: not implemented"
	return false, nil
}

// no barrier, don't wait

// wait if this is not the failed job

// no failed job, don't wait

// StateChanged must be called at the end, after the job state change has been persisted.
// The only exception to this rule is when a job has failed in a retryable manner, in this scenario you should notify the barrier immediately after the failure.
// An [ErrUnsupportedState] error will be returned if the state is not supported.
func (b *Barrier) StateChanged(key BarrierKey, jobID int64, state string) error {
	_ = "STUB: not implemented"
	return nil
}

// Sync applies any enqueued commands to the barrier's state. It should be called at the beginning of every new iteration of the main loop
func (b *Barrier) Sync() int { _ = "STUB: not implemented"; return 0 }

// Disabled returns [true] if the barrier is disabled for this key, [false] otherwise
func (b *Barrier) Disabled(key BarrierKey) bool { _ = "STUB: not implemented"; return false }

// Size returns the number of active barriers
func (b *Barrier) Size() int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the barrier
func (b *Barrier) String() string { _ = "STUB: not implemented"; return "" }

// updateState applies the state transitions for the barrier if necessary.
//
// 1. Disabled: transitions to half-enabled after disabledStateDuration
// 2. Half-enabled: transitions to enabled after halfEnabledStateDuration
func (b *Barrier) updateState(barrier *barrierInfo, key BarrierKey) {
	_ = "STUB: not implemented"
	return
}

type barrierState int

const (
	stateEnabled barrierState = iota
	stateDisabled
	stateHalfEnabled
)

type barrierInfo struct {
	state     barrierState
	stateTime time.Time

	failedJobID        *int64 // nil if no failed job
	concurrencyLimiter map[int64]struct{}
	drainLimiter       map[int64]struct{} // nil if limiter is off
}

// Enter adds the jobID to the barrier's active limiter(s)
func (bi *barrierInfo) Enter(jobID int64) { _ = "STUB: not implemented"; return }

// Leave removes the jobID from the barrier's limiter(s)
func (bi *barrierInfo) Leave(jobID int64) { _ = "STUB: not implemented"; return }

// concurrencyLimitReached returns true if the barrier's concurrency limit has been reached
func (bi *barrierInfo) concurrencyLimitReached(jobID int64, limit int) bool {
	_ = "STUB: not implemented"
	return false
}

// DrainLimitReached returns true if the barrier's drain limit has been reached
func (bi *barrierInfo) DrainLimitReached(jobID int64, limit int) bool {
	_ = "STUB: not implemented"
	return false
}

// Inactive returns true if the barrier is enabled, there isn't a failed job and both limiters are inactive
func (bi *barrierInfo) Inactive() bool { _ = "STUB: not implemented"; return false }

// barrier is enabled
// no failed job
// no concurrent jobs
// drain limiter is off

type command interface {
	enqueue(b *Barrier) bool
	execute(b *Barrier)
}

type cmd struct {
	key   BarrierKey
	jobID int64
}

// default behaviour is to try and remove the jobID from the concurrent jobs map
func (c *cmd) execute(b *Barrier) { _ = "STUB: not implemented"; return }

// default behaviour is to enqueue the command if a barrier for this key already exists
func (c *cmd) enqueue(b *Barrier) bool { _ = "STUB: not implemented"; return false }

// jobFailedCommand is a command that is executed when a job has failed.
type jobFailedCommand struct {
	*cmd
}

// If no failed jobID is in the barrier make this jobID the failed job for this key. Removes the job from the concurrent jobs map if it exists there
func (c *jobFailedCommand) execute(b *Barrier) { _ = "STUB: not implemented"; return }

// don't do anything if the barrier is disabled

// turn off drain limiter

// a failed command never gets enqueued
func (*jobFailedCommand) enqueue(_ *Barrier) bool {
	_ = "STUB: not implemented"

	// jobSucceededCmd is a command that is executed when a job has succeeded.
	return false
}

type jobSucceededCmd struct {
	*cmd
}

// removes the barrier for this key, if it exists
func (c *jobSucceededCmd) execute(b *Barrier) { _ = "STUB: not implemented"; return }

// out-of-sync command (failed commands get executed immediately)

// jobFilteredCmd is a command that is executed when a job is filtered.
type jobFilteredCmd struct {
	*cmd
}

// removes the barrier for this key, if it exists
func (c *jobFilteredCmd) execute(b *Barrier) { _ = "STUB: not implemented"; return }

// out-of-sync command (failed commands get executed immediately)

// jobAbortedCommand is a command that is executed when a job has aborted.
type jobAbortedCommand struct {
	*cmd
}

// Creates a concurrent jobs map if none exists. Also removes the jobID from the concurrent jobs map if it exists there
func (c *jobAbortedCommand) execute(b *Barrier) { _ = "STUB: not implemented"; return }

// out-of-sync command (failed commands get executed immediately)

// previouslyFailed indicates whether the job that was aborted was previouly a failed job, which is the condition for enabling the drain limiter

// remove the failed job
// turn off drain limiter
// enable the drain limiter only if a previously failed job has been aborted
