package forwarder

import (
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type sampler[K comparable] struct {
	period time.Duration            // the sampling period
	cache  *lru.Cache[K, time.Time] // the cache of keys and their last sampled time
	now    func() time.Time         //	now returns the current time. It is used to make the sampler testable.
}

// Sample returns true if the key should be sampled. It returns false otherwise.
func (s *sampler[K]) Sample(key K) bool { _ = "STUB: not implemented"; return false }

// newSampler returns a new, properly initialized sampler.
func newSampler[K comparable](period time.Duration, cacheSize int) *sampler[K] {
	_ = "STUB: not implemented"
	return nil
}
