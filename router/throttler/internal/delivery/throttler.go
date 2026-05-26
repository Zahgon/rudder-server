package delivery

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"
)

// NewThrottler constructs a new static delivery throttler for a destination endpoint
func NewThrottler(destType, destinationID, endpointPath string, limiter Limiter, config *config.Config, stat stats.Stats, log logger.Logger) *throttler {
	_ = "STUB: not implemented"
	return nil
}

// key is destinationID:endpointPath

type throttler struct {
	destinationID string
	endpointPath  string
	key           string

	limiter Limiter
	log     logger.Logger
	limit   config.ValueLoader[int64]
	window  config.ValueLoader[time.Duration]

	onceEveryGauge   *kitsync.OnceEvery
	rateLimitGauge   stats.Gauge
	waitTimerSuccess stats.Timer
	waitTimerFailure stats.Timer
}

func (t *throttler) Wait(ctx context.Context) (dur time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (t *throttler) enabled() bool { _ = "STUB: not implemented"; return false }

func (t *throttler) getLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) getTimeWindowInSeconds() int64 { _ = "STUB: not implemented"; return 0 }

func (t *throttler) updateGauges() { _ = "STUB: not implemented"; return }
