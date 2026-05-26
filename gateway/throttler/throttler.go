package throttler

//go:generate mockgen -destination=../../mocks/gateway/throttler.go -package=mocks_gateway github.com/rudderlabs/rudder-server/gateway/throttler Throttler

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

const (
	throttlingAlgoTypeGCRA = "gcra"
)

type Limiter interface {
	// Allow returns true if the limit is not exceeded, false otherwise.
	Allow(ctx context.Context, cost, rate, window int64, key string) (bool, func(context.Context) error, error)
}

type Throttler interface {
	CheckLimitReached(context context.Context, workspaceId string, eventCount int64) (bool, error)
}

type Factory struct {
	Stats        stats.Stats
	limiter      Limiter
	throttlers   map[string]*throttler // map key is the workspaceId
	throttlersMu sync.Mutex
}

// New constructs a new Throttler Factory
func New(stats stats.Stats) (*Factory, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Factory) CheckLimitReached(context context.Context, workspaceId string, eventCount int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *Factory) get(workspaceId string) *throttler { _ = "STUB: not implemented"; return nil }

func (f *Factory) initThrottlerFactory() error { _ = "STUB: not implemented"; return nil }

type throttler struct {
	limiter Limiter
	config  throttlingConfig
}

// checkLimitReached returns true if we're not allowed to process the number of event
func (t *throttler) checkLimitReached(ctx context.Context, key string, count int64) (limited bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

// no token to return when limited

type throttlingConfig struct {
	limit  int64
	window time.Duration
}

func (c *throttlingConfig) readThrottlingConfig(workspaceID string) {
	_ = "STUB: not implemented"
	return
}

func getWindowInSecs(d time.Duration) int64 { _ = "STUB: not implemented"; return 0 }
