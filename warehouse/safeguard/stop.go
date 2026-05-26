package safeguard

import (
	"context"
	"sync"
	"time"
)

var defaultGuard = &Guard{}

type Guard struct {
	Go   func(func())
	once sync.Once
}

func (g *Guard) MustStop(ctx context.Context, timeout time.Duration) func() {
	_ = "STUB: not implemented"
	return nil
}

func MustStop(ctx context.Context, timeout time.Duration) func() {
	_ = "STUB: not implemented"
	return nil
}
