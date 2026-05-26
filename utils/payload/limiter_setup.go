package payload

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type AdaptiveLimiterFunc func(int64) int64

// SetupAdaptiveLimiter creates a new AdaptiveLimiter, starts its RunLoop in a goroutine and periodically collects statistics.
func SetupAdaptiveLimiter(ctx context.Context, g *errgroup.Group) AdaptiveLimiterFunc {
	_ = "STUB: not implemented"
	return *new(AdaptiveLimiterFunc)
}

// run tick periodically

// collect statistics
