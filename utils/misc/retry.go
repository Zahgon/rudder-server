package misc

import (
	"context"
	"time"
)

type Notify func(attempt int)

func RetryWith(parentContext context.Context, timeout time.Duration, maxAttempts int, f func(ctx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

// RetryWithNotify retries a function f with a timeout and a maximum number of attempts & calls notify on each failure.
func RetryWithNotify(parentContext context.Context, timeout time.Duration, maxAttempts int, f func(ctx context.Context) error, notify Notify) error {
	_ = "STUB: not implemented"
	return nil
}

// only retry if the child context's deadline was exceeded

func QueryWithRetries[T any](parentContext context.Context, timeout time.Duration, maxAttempts int, f func(ctx context.Context) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func QueryWithRetriesAndNotify[T any](parentContext context.Context, timeout time.Duration, maxAttempts int, f func(ctx context.Context) (T, error), notify Notify) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// only retry if the child context's deadline was exceeded
