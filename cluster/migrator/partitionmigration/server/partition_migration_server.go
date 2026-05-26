// GRPC service API implementation for partition migration server
package server

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
	proto "github.com/rudderlabs/rudder-server/proto/cluster"
)

type Opt func(*partitionMigrationServer)

// WithLogger sets the logger for the PartitionMigrationServer.
func WithLogger(logger logger.Logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithStats sets the stats for the PartitionMigrationServer.
func WithStats(stats stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithDedupEnabled sets whether deduplication is enabled for the PartitionMigrationServer (enabled by default).
func WithDedupEnabled(dedupEnabled bool) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithStreamTimeout sets the timeout for stream receive/send operations (default: 5 minutes).
func WithStreamTimeout(timeout config.ValueLoader[time.Duration]) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// NewPartitionMigrationServer creates a new PartitionMigrationServer instance.
func NewPartitionMigrationServer(ctx context.Context, jobsdbs []jobsdb.JobsDB, opts ...Opt) proto.PartitionMigrationServer {
	_ = "STUB: not implemented"
	return *new(proto.PartitionMigrationServer)
}

type partitionMigrationServer struct {
	proto.UnimplementedPartitionMigrationServer
	lifecycleCtx  context.Context
	logger        logger.Logger
	stats         stats.Stats
	jobsdbs       map[string]jobsdb.JobsDB          // map of table prefix to jobsdb instance
	dedupEnabled  bool                              // whether deduplication is enabled when receiving jobs
	streamTimeout config.ValueLoader[time.Duration] // timeout for stream receive/send operations

	activeMigrationsMu sync.Mutex          // protects activeMigrations
	activeMigrations   map[string]struct{} // set of active migration keys (migrationJobID_tablePrefix)
}

// StreamJobs handles the streaming of jobs from the client for migration
// to the target jobsdb:
//
//   - The client first sends a metadata message, followed by multiple job chunk messages.
//   - The server receives the job chunks, reconstructs the job batches, and stores them in the target jobsdb.
//   - It acknowledges each batch back to the client after successful storage.
//   - While a batch is being stored, the server continues to receive subsequent chunks to maximize throughput.
//   - At most 2 batches are kept in memory at any time to control memory usage.
//   - Deduplication is performed based on job IDs to avoid storing duplicate jobs.
func (s *partitionMigrationServer) StreamJobs(stream proto.PartitionMigration_StreamJobsServer) (streamErr error) {
	_ = "STUB: not implemented"
	return nil
}

// merge the stream context with the server lifecycle context

// use priority pool for migration job handling (values from lifecycleCtx are not copied to the merged context)

// receive the first message containing metadata

// resolve the jobsdb for the given table prefix

// construct the migration key, which is a combination of migration job ID and table prefix

// only allow one active migration per migration key

// clean up active migration entry

// Deduplication, if enabled, will keep deduplicating jobs until it finds a job with JobID > lastJobID.
// This allows us to resume interrupted migrations without re-storing already stored jobs.
// Once we find a job with JobID > lastJobID, we disable deduplication for the rest of the migration,
// as we assume that the client will not resend any previously sent jobs after that point.

// Load lastJobID from persistent storage to resume interrupted migrations

// unbuffered channel for backpressure

// goroutine for storing complete batches and sending acknowledgments

// helper to get the store error if any

// store only non-empty batches (we may have empty batches after deduplication)

// store the batch

// store jobs

// update lastJobID for deduplication

// acknowledge the batch

// done with for loop, close storeErr to signal completion

// if context is not cancelled (i.e. successful completion), cleanup dedup state

// helper to stop the store goroutine, wait for it to finish and return any errors

// figure out the cause (either the passed error or the store error),
// we give priority to any store error that may have occurred before the passed error

// cancel the context in case of an error to abort ongoing operations

// close the batches channel to stop the storing goroutine

// wait for the store goroutine to finish

// return the cause error if any

// even if there was no cause error, after we closed the channel, store might have failed, thus we need to check again
// no need to lock again as the goroutine has finished

// process incoming stream, collect job chunks into batches and send them to the store goroutine

// client has finished sending

// any other error should be considered fatal

// first chunk of a new batch

// deduplicate jobs if needed, i.e. ignore jobs with JobID <= lastJobID

// disable dedup as soon as we find at least one job in the batch that is not a duplicate

// batch sent for storing

// context cancelled

// streamingBatch represents a batch of jobs received from the client for migration
type streamingBatch struct {
	index int64          // index of the batch, used for acknowledgments
	jobs  []*jobsdb.JobT // jobs contained in the batch
}

// receiveWithTimeout wraps stream.Recv with a timeout and context cancellation support
func (s *partitionMigrationServer) receiveWithTimeout(ctx context.Context, stream proto.PartitionMigration_StreamJobsServer) (*proto.StreamJobsRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// known issue - goroutine leaks on timeout - acceptable.

// sendWithTimeout wraps stream.Send with a timeout and context cancellation support
func (s *partitionMigrationServer) sendWithTimeout(ctx context.Context, stream proto.PartitionMigration_StreamJobsServer, ack *proto.JobsBatchAck) error {
	_ = "STUB: not implemented"
	return nil
}
