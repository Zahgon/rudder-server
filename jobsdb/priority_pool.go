package jobsdb

import "context"

// priorityPoolKey is the context key for priority pool requests
type priorityPoolKey struct{}

// WithPriorityPool returns a context that signals JobsDB operations
// should use the priority pool (if configured) instead of the regular
// connection pool and should bypass reader/writer queues.
func WithPriorityPool(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// usePriorityPool checks if the context requests priority pool usage
func usePriorityPool(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
