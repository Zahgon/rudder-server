package cache

import (
	"sync"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

// StatsCacheKey interface defines the requirements for cache keys
type StatsCacheKey interface {
	comparable
	ToStatTags() stats.Tags
}

// StatsCache is a generic cache for stats measurements using sync.Map
type StatsCache[T StatsCacheKey] struct {
	cache    sync.Map // stores map[T]stats.Measurement
	producer func(T) stats.Measurement
}

// NewStatsCache creates a new stats cache instance
func NewStatsCache[T StatsCacheKey](producer func(T) stats.Measurement) *StatsCache[T] {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a measurement from cache, creating it if it doesn't exist
func (c *StatsCache[T]) Get(key T) stats.Measurement {
	_ = "STUB: not implemented"
	// Try to load the value from the map
	return *new(stats.Measurement)
}

// Value not found—create it

// Store and possibly get actual value (if stored by another goroutine in the meantime)
