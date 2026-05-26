package algorithm

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

type increaseLimitCounter struct {
	window             func() time.Duration
	increasePercentage config.ValueLoader[int64]
	limitFactor        *limitFactor
	throttledCountMu   sync.Mutex
	throttledCount     int64
}

func (c *increaseLimitCounter) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

func (c *increaseLimitCounter) run(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}
