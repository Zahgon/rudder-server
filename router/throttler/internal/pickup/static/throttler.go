package static

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

	lastUsed   atomic.Time
	limiter    Limiter
	log        Logger
	limit      config.ValueLoader[int64]
	window     config.ValueLoader[time.Duration]
	staticCost config.ValueLoader[bool]

	everyStats     *kitsync.OnceEvery
	rateLimitGauge stats.Gauge
}

func (t *throttler) CheckLimitReached(ctx context.Context, cost int64) (limited bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *throttler) enabled() bool { _ = "STUB: not implemented"; return false }

func (t *throttler) getLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) GetLimitPerSecond() int64 { _ = "STUB: not implemented"; return 0 }

// ceiling division

func (t *throttler) GetEventType() string { _ = "STUB: not implemented"; return "" }

func (t *throttler) GetLastUsed() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (t *throttler) getTimeWindowInSeconds() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) ResponseCodeReceived(code int) {
	_ = "STUB: not implemented"
	// no-op
	return
}

func (t *throttler) Shutdown() {
	_ = "STUB: not implemented"
	// no-op
	return
}

func (t *throttler) updateStats() { _ = "STUB: not implemented"; return }

func (t *throttler) costFn(input int64) int64 { _ = "STUB: not implemented"; return 0 }
