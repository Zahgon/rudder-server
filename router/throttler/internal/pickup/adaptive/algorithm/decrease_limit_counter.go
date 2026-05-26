package algorithm

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

type decreaseLimitCounter struct {
	window                      func() time.Duration
	waitWindow                  func() time.Duration
	decreasePercentage          config.ValueLoader[int64]
	throttleTolerancePercentage func() int64
	limitFactor                 *limitFactor

	counterMu      sync.Mutex
	throttledCount int64
	totalCount     int64
}

func (c *decreaseLimitCounter) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

func (c *decreaseLimitCounter) run(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// wait waits for the waitWindow duration and resets the throttledCount and totalCount to 0
func (c *decreaseLimitCounter) wait(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
