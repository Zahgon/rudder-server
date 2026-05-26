// Migration Job Executor, i.e. moving a set of partitions from one node's jobsdb to another
package client

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
	proto "github.com/rudderlabs/rudder-server/proto/cluster"
)

type Opt func(*migrationJobExecutor)

// WithConfig sets the config to be used by the executor
func WithConfig(conf *config.Config) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithLogger sets the logger to be used by the executor
func WithLogger(logger logger.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithStats sets the stats instance to be used by the executor
func WithStats(stats stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// NewMigrationJobExecutor creates a new MigrationJobExecutor
func NewMigrationJobExecutor(migrationJobID string, nodeIndex int, partitionIDs []string, sourceDB jobsdb.JobsDB, target string, opts ...Opt) MigrationJobExecutor {
	_ = "STUB: not implemented"
	return *new(MigrationJobExecutor)
}

// 100 MB

type MigrationJobExecutor interface {
	// Run executes the migration job
	Run(ctx context.Context) error
}

type migrationJobExecutor struct {
	migrationJobID string   // unique identifier for the migration job
	nodeIndex      int      // source node index
	partitionIDs   []string // partitions to be migrated

	sourceDB jobsdb.JobsDB // source jobsdb
	target   string        // target address of the partition migrator server

	conf   *config.Config
	logger logger.Logger
	stats  stats.Stats

	batchSize        config.ValueLoader[int]           // number of jobs to fetch in each batch
	payloadSizeLimit config.ValueLoader[int64]         // maximum total payload size of jobs to fetch
	chunkSize        config.ValueLoader[int]           // number of jobs to send in each chunk to the remote server
	progressPeriod   config.ValueLoader[time.Duration] // period for logging progress
	sendTimeout      config.ValueLoader[time.Duration] // timeout for stream send operations
	receiveTimeout   config.ValueLoader[time.Duration] // timeout for stream receive operations
}

// Run executes the migration job
func (mpe *migrationJobExecutor) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// refresh the DS list to ensure that we'll move partition data from all datasets

// mark any executing jobs as failed to handle previous interrupted migrations

// create a partition migrator client

// close the client on exit

// group for sender and receiver goroutines
// open the stream

// counter for total sent jobs
// counter for total acknowledged jobs

// map of batch index to unacked jobs in case of failure, for marking them as failed.
// channel for sending unacked batch info to the receiver goroutine to acknowledge (one at a time). channel is closed when sender is done

// sender goroutine

// start by sending the metadata

// send job batches in chunks

// signals the end of jobs to be sent
// As the last batch we are always sending an empty chunk, in order to avoid a race condition in deduplication, if this is enabled on server:
// The server, after sending an acknowledgement for the last batch, it then resets its deduplication state for this migration job id.
// If client crashes before marking the jobs as migrated, upon restart it will resend the last batch which will not be deduplicated by server.

// we sent an empty chunk in the previous iteration to signal end of jobs
// we are done sending

// get next batch of jobs to process

// no more jobs to process, we are done sending

// mark them as executing

// add to unacked batches

// split jobs into chunks and send them

// we will send one empty chunk if there are no jobs

// send batch info to receiver only just before the last chunk

// added successfully

// receiver goroutine

// context cancelled

// new unacked batch to ack
// channel is closed, just wait for an eof from the server

// wait for the acknowledgement

// mark jobs as migrated

// remove from unacked batches

// periodic logger

// wait for sender and receiver to finish

// no need to lock unackedBatches, we are done with sender and receiver goroutines

// markExecutingJobsAsFailed marks any executing jobs as failed to handle previous interrupted migrations
func (mpe *migrationJobExecutor) markExecutingJobsAsFailed(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// if there are no more executing jobs and DS limits are not reached, we are done

// updateJobStatus updates the status of the given jobs to the given state
func (mpe *migrationJobExecutor) updateJobStatus(ctx context.Context, jobs []*jobsdb.JobT, state string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// receiveWithTimeout wraps stream.Recv with a timeout and context cancellation support
func (mpe *migrationJobExecutor) receiveWithTimeout(ctx context.Context, stream proto.PartitionMigration_StreamJobsClient) (*proto.JobsBatchAck, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// known issue - goroutine leaks on timeout - acceptable.

// sendWithTimeout wraps stream.Send with a timeout and context cancellation support
func (mpe *migrationJobExecutor) sendWithTimeout(ctx context.Context, stream proto.PartitionMigration_StreamJobsClient, request *proto.StreamJobsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (mpe *migrationJobExecutor) statsTags() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
