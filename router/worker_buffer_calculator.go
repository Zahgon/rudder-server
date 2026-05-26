package router

import (
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-go-kit/stats/metric"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"
)

// bufferSizeCalculator is a function that calculates the buffer size for a worker
type bufferSizeCalculator func() int

// newStandardBufferSizeCalculator uses the maximum of number of jobs to batch in a worker and number of jobs per channel
// to calculate the buffer size for a worker
func newStandardBufferSizeCalculator(
	noOfJobsToBatchInAWorker config.ValueLoader[int], // number of jobs that a worker can batch together
	noOfJobsPerChannel int, // number of jobs per channel
) bufferSizeCalculator {
	_ = "STUB: not implemented"
	return *new(bufferSizeCalculator)
}

// newExperimentalBufferSizeCalculator calculates the buffer size for a worker based on the following algorithm (minimum value returned is 1):
//
//  1. Calculate the average throughput of jobs processed by the work loop per second (workLoopThroughput)
//  2. Calculate the number of jobs that are queried during pickup divided by the number of workers (jobQueryBatchSize / noOfWorkers)
//  3. Use the number of jobs to batch in a worker (noOfJobsToBatchInAWorker)
//  4. Take the maximum of the three metrics above and multiply it by a scaling factor (default: 2.0) to determine the buffer size
//
// Exceptional case:
//
//   - If throughput is less than 1 (workLoopThroughput < 1), set the buffer size to 1 so that we are forcing a slow buffer start & introducing backpressure in the buffer in case of slow processing.
func newExperimentalBufferSizeCalculator(
	jobQueryBatchSize config.ValueLoader[int], // number of jobs that are queried during pickup
	noOfWorkers int, // number of workers processing jobs
	noOfJobsToBatchInAWorker config.ValueLoader[int], // number of jobs that a worker can batch together
	workLoopThroughput metric.SimpleMovingAverage, // sliding average of work loop throughput
	scalingFactor config.ValueLoader[float64], // scaling factor to scale up the buffer size
	minBufferSize config.ValueLoader[int], // minimum buffer size
) bufferSizeCalculator {
	_ = "STUB: not implemented"
	return *new(bufferSizeCalculator)
}

// at least the average throughput of the work loop
// if there is no throughput yet, the throughput is less than 1 per second, set buffer to minBufferSize

// at least the average number of jobs per worker during pickup
// at least equal to the number of jobs to batch in a worker

// round up
// calculate the maximum of the three metrics to determine the buffer size

// scale up to provide some buffer
// ensure buffer size is at least one or the configured minimum

// newBufferSizeCalculatorSwitcher returns a function that switches between the standard and experimental calculators based on the
// enableExperimentalBufferSizeCalculator flag
func newBufferSizeCalculatorSwitcher(
	enableExperimentalBufferSizeCalculator config.ValueLoader[bool],
	jobQueryBatchSize config.ValueLoader[int], // number of jobs that are queried during pickup
	noOfWorkers int, // number of workers processing jobs
	noOfJobsToBatchInAWorker config.ValueLoader[int], // number of jobs that a worker can batch together
	workLoopThroughput metric.SimpleMovingAverage, // sliding average of work loop throughput
	scalingFactor config.ValueLoader[float64], // scaling factor to scale up the buffer size
	noOfJobsPerChannel int, // number of jobs per channel
	minBufferSize config.ValueLoader[int], // minimum buffer size
) bufferSizeCalculator {
	_ = "STUB: not implemented"
	return *new(bufferSizeCalculator)
}

// newSmaHistogram combines a SimpleMovingAverage with a stats.Histogram to periodically record the moving average into the histogram.
func newSmaHistogram(
	slidingAverage metric.SimpleMovingAverage,
	histogram stats.Histogram,
	onceEvery *kitsync.OnceEvery,
) stats.Histogram {
	_ = "STUB: not implemented"
	return *new(stats.Histogram)
}

type smaHistogram struct {
	slidingAverage metric.SimpleMovingAverage
	histogram      stats.Histogram
	onceEvery      *kitsync.OnceEvery
}

func (s *smaHistogram) Observe(v float64) { _ = "STUB: not implemented"; return }

type Gauge[T any] interface {
	// Gauge sets the gauge to the provided value
	Gauge(value T)
}
type GaugeWithLastValue[T any] interface {
	Gauge[T]
	config.ValueLoader[T]
}

// newGaugeWithLastValue wraps a stats.Gauge and keeps track of the last value set.
func newGaugeWithLastValue[T any](gauge stats.Gauge) GaugeWithLastValue[T] {
	_ = "STUB: not implemented"
	return nil
}

type gaugeWithLastValue[T any] struct {
	gauge     stats.Gauge
	mu        sync.RWMutex
	lastValue T
}

func (g *gaugeWithLastValue[T]) Gauge(value T) { _ = "STUB: not implemented"; return }

func (g *gaugeWithLastValue[T]) Load() T { _ = "STUB: not implemented"; return *new(T) }
