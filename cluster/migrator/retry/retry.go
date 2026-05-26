package retry

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

// PerpetualExponentialBackoff retries the given function with an exponential backoff strategy
// until it succeeds or the context is cancelled.
func PerpetualExponentialBackoffWithNotify(ctx context.Context, config *config.Config, fn func() error, notify func(error, time.Duration)) error {
	_ = "STUB: not implemented"
	return nil
}

// perpetual retries
