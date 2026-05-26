package adaptive

import (
	"context"
	"time"

	"go.uber.org/atomic"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"
)

type throttler struct {
	destinationID string
	eventType     string
	key           string

	lastUsed  atomic.Time
	limiter   Limiter
	algorithm Algorithm
	log       Logger

	window     config.ValueLoader[time.Duration]
	minLimit   config.ValueLoader[int64]
	maxLimit   config.ValueLoader[int64]
	staticCost config.ValueLoader[bool]

	everyStats       *kitsync.OnceEvery
	limitFactorGauge stats.Gauge
	rateLimitGauge   stats.Gauge
}

func (t *throttler) enabled() bool { _ = "STUB: not implemented"; return false }

func (t *throttler) CheckLimitReached(ctx context.Context, cost int64) (limited bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *throttler) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

func (t *throttler) Shutdown() { _ = "STUB: not implemented"; return }

func (t *throttler) getLimitFactor() float64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) getMinLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) getMaxLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) getLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) GetLimitPerSecond() int64 { _ = "STUB: not implemented"; return 0 }

// ceiling division

func (t *throttler) GetEventType() string { _ = "STUB: not implemented"; return "" }

func (t *throttler) GetLastUsed() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (t *throttler) getTimeWindowInSeconds() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) updateStats() { _ = "STUB: not implemented"; return }

func (t *throttler) costFn(input int64) int64 { _ = "STUB: not implemented"; return 0 }
