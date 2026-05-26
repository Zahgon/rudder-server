package partition

import (
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats/metric"
)

// NewStats returns a new, initialised partition stats
func NewStats() *Stats { _ = "STUB: not implemented"; return nil }

// Stats keeps track of throughput and error rates for each partition
type Stats struct {
	pstatsMu sync.RWMutex
	pstats   map[string]*pstat
}

// Update updates the stats for the given partition in terms of throughput, total requests and errors
func (s *Stats) Update(partition string, duration time.Duration, total, errors int) {
	_ = "STUB: not implemented"
	return
}

// Score returns a score for the given partition. The score is a number between 0 and 100. Scores are calculated
// comparatively to other partitions, so that the partition with the highest throughput and the lowest error ratio
// will receive the highest score.
func (s *Stats) Score(partition string) int { _ = "STUB: not implemented"; return 0 }

type PartitionStats struct {
	Partition  string
	Throughput float64
	Successes  float64
	Errors     float64

	NormalizedThroughput float64
	Score                int
}

// All returns a map containing stats for all available partitions
func (s *Stats) All() map[string]PartitionStats { _ = "STUB: not implemented"; return nil }

// error ratio decreases throughput

// highest throughput wins

type pstat struct {
	throughput metric.MovingAverage // successful events per millisecond
	successes  metric.MovingAverage // number of successful events
	errors     metric.MovingAverage // number of errors
}
