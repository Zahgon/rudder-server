package algorithm

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

type AdaptiveAlgorithm interface {
	// ResponseCodeReceived processes the response code and updates the limit factor accordingly
	ResponseCodeReceived(code int)
	// Shutdown stops the algorithm and waits for all goroutines to finish
	Shutdown()
	// LimitFactor returns a factor that is supposed to be used to multiply the limit, a number between 0 and 1
	LimitFactor() float64
}

type adaptiveAlgorithm struct {
	limitFactor          *limitFactor
	increaseLimitCounter *increaseLimitCounter
	decreaseLimitCounter *decreaseLimitCounter
	cancel               context.CancelFunc
	wg                   *sync.WaitGroup
}

// NewAdaptiveAlgorithm creates a new adaptive algorithm instance.
//
// An adaptive algorithm dynamically adjusts the limit factor based on the response codes received within a certain time window.
//
// Configurable parameters include:
// - increaseWindowMultiplier: Multiplier for the increase limit counter window.
// - increasePercentage: Percentage to increase the limit factor when no 429s are received.
// - decreaseWaitWindowMultiplier: Multiplier for the wait window after a decrease.
// - decreasePercentage: Percentage to decrease the limit factor when 429s are received.
// - throttleTolerancePercentage: Percentage of throttled requests that triggers a decrease in the limit factor
func NewAdaptiveAlgorithm(destType string, config *config.Config, window config.ValueLoader[time.Duration]) AdaptiveAlgorithm {
	_ = "STUB: not implemented"
	return *new(AdaptiveAlgorithm)
}

// LimitFactor returns the current limit factor
func (a *adaptiveAlgorithm) LimitFactor() float64 { _ = "STUB: not implemented"; return 0 }

// ResponseCodeReceived processes the response code and updates the limit factor accordingly
func (a *adaptiveAlgorithm) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

// Shutdown stops the algorithm and waits for all goroutines to finish
func (a *adaptiveAlgorithm) Shutdown() { _ = "STUB: not implemented"; return }

type limitFactor struct {
	mu    sync.RWMutex
	value float64
}

// Add adds value to the current value of the limit factor, and clamps it between 0 and 1
func (l *limitFactor) Add(value float64) { _ = "STUB: not implemented"; return }

func (l *limitFactor) Get() float64 { _ = "STUB: not implemented"; return 0 }
