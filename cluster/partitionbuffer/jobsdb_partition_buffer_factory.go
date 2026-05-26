package partitionbuffer

import (
	"context"
	"errors"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type Opt func(*jobsDBPartitionBuffer)

// ErrInvalidJobsDBPartitionBufferConfig is returned when the configuration for JobsDBPartitionBuffer is invalid
var ErrInvalidJobsDBPartitionBufferConfig = errors.New("invalid jobsdb partition buffer configuration, need to use WithReadWriteJobsDBs, WithWriterOnlyJobsDBs or WithReaderOnlyAndFlushJobsDBs")

// WithReadWriteJobsDBs sets both read and write JobsDBs for primary and buffer
func WithReadWriteJobsDBs(primary, buffer jobsdb.JobsDB) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithWriterOnlyJobsDBs sets only the writer JobsDBs for primary and buffer
func WithWriterOnlyJobsDBs(primaryWriter, bufferWriter jobsdb.JobsDB) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithReaderOnlyAndFlushJobsDBs sets only the reader JobsDBs for primary and buffer, and writer as primary
func WithReaderOnlyAndFlushJobsDBs(primaryReader, bufferReader, primaryWriter jobsdb.JobsDB) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithSeparateReaderAndWriterPrimaryJobsDBs uses the primary writer as a delegate, however does not manage its lifecycle
func WithSeparateReaderAndWriterPrimaryJobsDBs(primaryReader, primaryWriter, buffer jobsdb.JobsDB) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// primaryWriter's and buffer's lifecycle are externally managed

// WithLogger sets the logger for the JobsDBPartitionBuffer
func WithLogger(logger logger.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithStats sets the stats collector for the JobsDBPartitionBuffer
func WithStats(stats stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithNumPartitions sets the number of partitions for the JobsDBPartitionBuffer
func WithNumPartitions(numPartitions int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithFlushBatchSize sets the flush batch size for the JobsDBPartitionBuffer
func WithFlushBatchSize(flushBatchSize config.ValueLoader[int]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithFlushPayloadSize sets the flush payload size for the JobsDBPartitionBuffer
func WithFlushPayloadSize(flushPayloadSize config.ValueLoader[int64]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithFlushMoveTimeout sets the flush move timeout for the JobsDBPartitionBuffer
func WithFlushMoveTimeout(flushMoveTimeout config.ValueLoader[time.Duration]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithFlushMoveConcurrency sets the number of goroutines that distribute partitions during the move phase of a flush
func WithFlushMoveConcurrency(flushMoveConcurrency config.ValueLoader[int]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithWatchdogInterval sets the watchdog interval for the JobsDBPartitionBuffer
func WithWatchdogInterval(watchdogInterval config.ValueLoader[time.Duration]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// NewJobsDBPartitionBuffer creates a new JobsDBPartitionBuffer with the given options
func NewJobsDBPartitionBuffer(ctx context.Context, opts ...Opt) (JobsDBPartitionBuffer, error) {
	_ = "STUB: not implemented"
	return *new(JobsDBPartitionBuffer), nil
}
