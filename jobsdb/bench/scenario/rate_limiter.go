package scenario

import (
	"context"
)

func rateLimiter(limit int) (func(ctx context.Context, key string, cost int) error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
