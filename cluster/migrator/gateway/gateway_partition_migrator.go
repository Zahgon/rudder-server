// Migrator component for rudder-server operating in gateway mode
package migrator

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	etcdtypes "github.com/rudderlabs/rudder-schemas/go/cluster"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
)

// PartitionRefresher is the interface for refreshing buffered partitions
type PartitionRefresher interface {
	// RefreshBufferedPartitions refreshes the list of buffered partitions from the database
	RefreshBufferedPartitions(ctx context.Context) error
}

// PartitionMigrator handles partition migrations for a gateway node
type PartitionMigrator interface {
	// Start starts the partition migrator
	Start() error

	// Stop stops the partition migrator
	Stop()
}

type gatewayPartitionMigrator struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config             *config.Config
	logger             logger.Logger
	stats              stats.Stats
	etcdClient         etcdclient.Client
	partitionRefresher PartitionRefresher

	// component lifecycle
	wg              *errgroup.Group
	lifecycleCtx    context.Context
	lifecycleCancel func()

	// state
	pendingReloadsMu sync.Mutex
	pendingReloads   map[string]struct{}
}

func (gpm *gatewayPartitionMigrator) Start() error { _ = "STUB: not implemented"; return nil }

// start watching for reload requests

// watchReloadRequests watches for gateway reload requests and handles them.
// It keeps retrying on errors using an exponential backoff until the context is done.
func (gpm *gatewayPartitionMigrator) watchReloadRequests(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// create a watcher for gateway reload requests

// only handle reload commands for this node

// if ack key already exists, no need to process again

// watch for reload events and handle them

// skip if reload is already being processed,
// otherwise add it to pending reloads

// handle reload event asynchronously

// onReloadRequest handles a gateway reload request. It keeps retrying until the reload is successful or the context is done.
func (gpm *gatewayPartitionMigrator) onReloadRequest(ctx context.Context, requestKey string, cmd *etcdtypes.ReloadGatewayCommand) {
	_ = "STUB: not implemented"
	return
}

// keep retrying until context is done

// refresh buffered partitions from database

// send ack for reload completed

// ensure the reload request still exists before acknowledging

// remove from pending reloads

func (gpm *gatewayPartitionMigrator) Stop() { _ = "STUB: not implemented"; return }

func (gpm *gatewayPartitionMigrator) statsTags() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
