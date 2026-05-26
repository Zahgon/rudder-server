package sourcenode

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	etcdtypes "github.com/rudderlabs/rudder-schemas/go/cluster"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// Migrator defines the interface for a source node migrator
type Migrator interface {
	// Handle prepares the source node for starting a new partition migration
	Handle(ctx context.Context, migration *etcdtypes.PartitionMigration) error

	// Run watches for new migration jobs assigned to this source node
	Run(ctx context.Context, wg *errgroup.Group) error
}

type migrator struct {
	nodeIndex int
	nodeName  string

	etcdClient        etcdclient.Client
	readerJobsDBs     []jobsdb.JobsDB // reader jobsdbs for this source node
	config            *config.Config
	logger            logger.Logger
	stats             stats.Stats
	shutdown          func() // function to trigger a shutdown of the node
	targetURLProvider func(targetNodeIndex int) (string, error)

	c struct {
		readExcludeSleep         *config.Reloadable[time.Duration] // duration to wait after marking partitions as read-excluded
		waitForInProgressTimeout *config.Reloadable[time.Duration] // timeout for waiting for in-progress jobs to complete
		inProgressPollSleep      *config.Reloadable[time.Duration] // sleep duration between checks for in-progress jobs
	}

	// state
	pendingMigrationJobsMu sync.Mutex
	pendingMigrationJobs   map[string]struct{}
}

// Handle prepares the source node for starting a new partition migration
// by marking the source partitions as read-excluded and waiting for in-progress jobs to complete
// on those partitions.
//
// If there are no source partitions assigned to this node, it returns immediately.
// Otherwise, it performs the following steps:
//  1. Marks the source partitions as read-excluded in all jobsdbs.
//  2. Waits for a configured sleep duration to allow in-flight queries to complete.
//  3. Waits until there are no in-progress job statuses in all jobsdbs for the source partitions.
//
// If any step above fails, it returns an error. If the wait for in-progress jobs times out, it triggers a shutdown of the node before returning an error.
func (m *migrator) Handle(ctx context.Context, migration *etcdtypes.PartitionMigration) error {
	_ = "STUB: not implemented"
	return nil
}

// no partitions to handle

// wait for sleep time so that in-flight queries can complete

// wait until there are no in-progress job statuses in all jobsdbs for the source partitions

// addReadExcludedPartitions marks the given partitions as excluded from reading in all jobsdbs
func (m *migrator) addReadExcludedPartitions(ctx context.Context, sourcePartitions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// removeReadExcludedPartitions unmarks the given partitions as excluded from reading in all jobsdbs
func (m *migrator) removeReadExcludedPartitions(ctx context.Context, sourcePartitions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// waitForNoInProgressJobs waits until there are no in-progress job statuses in all jobsdbs for the given source partitions.
func (m *migrator) waitForNoInProgressJobs(ctx context.Context, sourcePartitions []string) error {
	_ = "STUB: not implemented"
	return nil
}

// found in-progress jobs

// no in-progress jobs and no more jobs to fetch

// sleep for a short duration before checking again

// if timed out waiting for no in-progress jobs, trigger a shutdown

// Run watches for new migration jobs assigned to this source node and handles them asynchronously.
// All go routines are added to the provided errgroup.Group. It returns an error if the watcher cannot be created.
func (m *migrator) Run(ctx context.Context, wg *errgroup.Group) error {
	_ = "STUB: not implemented"
	return nil
}

// use priority pool for migration job handling
// reset pending migration jobs map

// create a watcher for partition migration jobs

// only watch for new migration jobs

// where this node is a source node

// It keeps retrying on errors using an exponential backoff until the context is done.

// Watch for new partition migration job events and handle them

// skip if migration job is already being processed,
// otherwise add it to pending migration jobs

// handle new migration event asynchronously

// onNewJob handles a new partition migration job assigned to this source node:
//
// 1. Moves jobs for the specified partitions from this source node to the target node.
// 2. Removes the read-excluded status for the migrated partitions.
// 3. Marks the migration job status as "moved" in etcd.
// 4. Removes the job from the pending migration jobs map.
//
// It retries on errors with an exponential backoff until the context is done.
func (m *migrator) onNewJob(ctx context.Context, key string, job *etcdtypes.PartitionMigrationJob) {
	_ = "STUB: not implemented"
	return
}

// Keep retrying errors with a backoff until context is done

// move jobs from all source jobsdbs concurrently

// remove read-excluded partitions from all jobsdbs

// mark partition migration job status as [moved] in etcd

// remove from pending migration jobs

func (m *migrator) statsTags() stats.Tags { _ = "STUB: not implemented"; return *new(stats.Tags) }

// getSourcePartitions returns the list of partitions assigned to the given source node index in the migration jobs
func getSourcePartitions(migration *etcdtypes.PartitionMigration, nodeIndex int) []string {
	_ = "STUB: not implemented"
	return nil
}
