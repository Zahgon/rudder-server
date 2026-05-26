package targetnode

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	etcdtypes "github.com/rudderlabs/rudder-schemas/go/cluster"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/cluster/partitionbuffer"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// Migrator defines the interface for a target node migrator
type Migrator interface {
	// Handle prepares the target node for starting a new partition migration
	Handle(ctx context.Context, migration *etcdtypes.PartitionMigration) error

	// Run starts a gRPC server for accepting jobs from source nodes and watches for moved migration jobs assigned to this target node
	Run(ctx context.Context, wg *errgroup.Group) error
}

type migrator struct {
	nodeIndex int
	nodeName  string

	etcdClient        etcdclient.Client
	bufferedJobsDBs   [][]partitionbuffer.JobsDBPartitionBuffer // hierarchy of partition buffer jobsdbs
	unbufferedJobsDBs []jobsdb.JobsDB                           // jobsdbs used for writing migrated jobs to (unbuffered)
	config            *config.Config
	logger            logger.Logger
	stats             stats.Stats

	// state
	pendingMigrationJobsMu sync.Mutex
	pendingMigrationJobs   map[string]struct{}
}

// Handle marks the partitions assigned to this target node as buffered in all buffered jobsDBs.
func (m *migrator) Handle(ctx context.Context, migration *etcdtypes.PartitionMigration) error {
	_ = "STUB: not implemented"
	return nil
}

// no partitions assigned to this target node

// Mark the partitions as buffered in all bufferedJobsDBs

// Run watches for moved migration jobs assigned to this target node and handles them asynchronously.
// It also starts a gRPC server for accepting jobs from source nodes.
// All go routines are added to the provided errgroup.Group.
// It returns an error if the watcher cannot be created, or if the gRPC server fails to start.
func (m *migrator) Run(ctx context.Context, wg *errgroup.Group) error {
	_ = "STUB: not implemented"
	return nil
}

// use priority pool for migration job handling
// reset pending migration jobs map

// create a watcher for partition migration jobs

// only watch for moved migration jobs

// where this node is a target node

// start gRPC server for accepting jobs from source nodes

// start watching for moved migration jobs

// It keeps retrying on errors using an exponential backoff until the context is done.

// Watch for moved partition migration job events and handle them

// skip if migration job is already being processed,
// otherwise add it to pending migration jobs

// handle moved migration event asynchronously

// stop gRPC server when context is done

// onNewJob handles a moved partition migration job assigned to this target node:
// For each group of buffered JobsDBPartitionBuffers, it flushes the buffered partitions for
// this job concurrently in all partition buffers of that group. Then it marks the job as completed
//
// It retries on errors with an exponential backoff until the context is done.
func (m *migrator) onNewJob(ctx context.Context, key string, job *etcdtypes.PartitionMigrationJob) {
	_ = "STUB: not implemented"
	return
}

// Keep retrying errors with a backoff until context is done

// flush jobs in sequence

// flush buffered partitions for this job in all partition buffers of this group concurrently

// mark partition migration job status as [completed] in etcd

// remove from pending migration jobs

func (m *migrator) statsTags() stats.Tags { _ = "STUB: not implemented"; return *new(stats.Tags) }

// getTargetPartitions returns the list of partitions assigned to the given target node index in the migration jobs
func getTargetPartitions(migration *etcdtypes.PartitionMigration, nodeIndex int) []string {
	_ = "STUB: not implemented"
	return nil
}
