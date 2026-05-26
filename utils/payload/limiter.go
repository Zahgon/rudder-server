package payload

import (
	"context"
	"sync"
	"time"
)

// LimiterState represents the LimiterState of the adaptive payload limiter algorithm
type LimiterState int

const (
	// LimiterStateNormal is the default state and the state when free memory is above the threshold
	LimiterStateNormal LimiterState = iota
	// LimiterStateThreshold is the state when free memory is below the LimiterStateThreshold but above the critical LimiterStateThreshold
	LimiterStateThreshold
	// LimiterStateCritical is the state when free memory is below the LimiterStateCritical threshold
	LimiterStateCritical
)

type LimiterStats struct {
	State           LimiterState
	ThresholdFactor int
}

// Limit is a function that returns the current payload limit in bytes
type Limiter interface {
	RunLoop(ctx context.Context, frequency func() <-chan time.Time)
	Limit(maxLimit int64) int64
	Stats() LimiterStats
	tick()
}

// NewAdaptiveLimiter creates a PayloadLimit function following an adaptive payload limiting algorithm
func NewAdaptiveLimiter(config AdaptiveLimiterConfig) Limiter {
	_ = "STUB: not implemented"
	return *new(Limiter)
}

type adaptivePayloadLimitAlgorithm struct {
	state           LimiterState
	config          AdaptiveLimiterConfig
	thresholdFactor int
	freeMem         float64
	mu              sync.Mutex
}

func (r *adaptivePayloadLimitAlgorithm) RunLoop(ctx context.Context, frequency func() <-chan time.Time) {
	_ = "STUB: not implemented"
	return
}

func (r *adaptivePayloadLimitAlgorithm) Limit(maxLimit int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// during normal state we return the max limit

// during threshold state we return the max limit decremented by 10% times the threshold factor

// during critical state we return 1 byte as a limit, since 0 bytes is interpreted as unlimited

func (r *adaptivePayloadLimitAlgorithm) Stats() LimiterStats {
	_ = "STUB: not implemented"
	return *new(LimiterStats)
}

func (r *adaptivePayloadLimitAlgorithm) tick() { _ = "STUB: not implemented"; return }

func (r *adaptivePayloadLimitAlgorithm) stateChanged(newState LimiterState) {
	_ = "STUB: not implemented"
	return
}
