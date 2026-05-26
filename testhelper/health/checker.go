package health

import (
	"context"
	"testing"
	"time"
)

func WaitUntilReady(
	ctx context.Context, t testing.TB, endpoint string, atMost, interval time.Duration, caller string,
) {
	_ = "STUB: not implemented"
	return
}

func IsReady(t testing.TB, endpoint string) bool { _ = "STUB: not implemented"; return false }
