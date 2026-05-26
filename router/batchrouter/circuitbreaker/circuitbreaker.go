package circuitbreaker

import (
	"time"

	"github.com/sony/gobreaker"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

type CircuitBreaker interface {
	IsOpen() bool
	Success()
	Failure()
}

type Opt func(*cfg)

func WithMaxRequests(maxRequests int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithConsecutiveFailures(consecutiveFailures int) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithLogger(logger logger.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func NewCircuitBreaker(name string, opts ...Opt) CircuitBreaker {
	_ = "STUB: not implemented"
	return *new(CircuitBreaker)
}

// Allow 1 request to pass through when in half-open state
// Doesn't count failures when time between requests > interval
// Time after which to transition from Open to Half-Open

type circuitBreaker struct {
	cb *gobreaker.CircuitBreaker
}

func (cb *circuitBreaker) IsOpen() bool { _ = "STUB: not implemented"; return false }

func (cb *circuitBreaker) Success() { _ = "STUB: not implemented"; return }

func (cb *circuitBreaker) Failure() { _ = "STUB: not implemented"; return }

type cfg struct {
	name                string
	maxRequests         int
	timeout             time.Duration
	consecutiveFailures int
	logger              logger.Logger
}
