// Migrator component for rudder-server operating in processor mode
package processor

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	etcdtypes "github.com/rudderlabs/rudder-schemas/go/cluster"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/cluster/migrator/processor/sourcenode"
	"github.com/rudderlabs/rudder-server/cluster/migrator/processor/targetnode"
)

// PartitionMigrator handles partition migrations for a processor node
type PartitionMigrator interface {
	// Start starts the partition migrator
	Start() error

	// Stop stops the partition migrator
	Stop()
}

type processorPartitionMigrator struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config         *config.Config
	logger         logger.Logger
	stats          stats.Stats
	etcdClient     etcdclient.Client
	sourceMigrator sourcenode.Migrator
	targetMigrator targetnode.Migrator

	// component lifecycle
	wg              *errgroup.Group
	lifecycleCtx    context.Context
	lifecycleCancel func()

	// state
	pendingMigrationsMu sync.Mutex
	pendingMigrations   map[string]struct{}
}

func (ppm *processorPartitionMigrator) Start() error { _ = "STUB: not implemented"; return nil }

// start watching for new partition migrations

// start running source migrator

// start running target migrator

// watchNewMigrations watches for new partition migration requests and handles them
// It keeps retrying on errors using an exponential backoff until the context is done.
func (ppm *processorPartitionMigrator) watchNewMigrations(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// create a watcher for partition migration requests

// only watch for new migrations

// where this node is a source node, a target node, or both

// if ack key already exists, skip processing this migration

// watch for new partition migration events and handle them

// skip if migration is already being processed,
// otherwise add it to pending migrations

// handle new migration event asynchronously

// onNewMigration handles a new partition migration event. It keeps retrying until the migration is successfully started or the context is done.
func (ppm *processorPartitionMigrator) onNewMigration(ctx context.Context, pm *etcdtypes.PartitionMigration) {
	_ = "STUB: not implemented"
	return
}

// keep retrying until context is done

// use errgroup to run source and target migrations concurrently

// if this node is a source node, start source migration

// if this node is a target node, start target migration

// send ack for migration started

// remove from pending migrations

func (ppm *processorPartitionMigrator) Stop() { _ = "STUB: not implemented"; return }

func (ppm *processorPartitionMigrator) statsTags() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
