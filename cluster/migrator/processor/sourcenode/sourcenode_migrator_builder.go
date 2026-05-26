package sourcenode

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// NewMigratorBuilder creates a new builder for sourcenode Migrator
func NewMigratorBuilder(nodeIndex int, nodeName string) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// MigratorBuilder is a builder for sourcenode Migrator
type MigratorBuilder struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config            *config.Config
	logger            logger.Logger
	stats             stats.Stats
	etcdClient        etcdclient.Client
	readerJobsDBs     []jobsdb.JobsDB
	shutdown          func()
	targetURLProvider func(targetNodeIndex int) (string, error)
}

// WithConfig sets the configuration for the Migrator
func (b *MigratorBuilder) WithConfig(config *config.Config) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithLogger sets the logger for the Migrator
func (b *MigratorBuilder) WithLogger(logger logger.Logger) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithStats sets the stats collector for the Migrator
func (b *MigratorBuilder) WithStats(stats stats.Stats) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithEtcdClient sets the etcd client for the Migrator
}

func (b *MigratorBuilder) WithEtcdClient(etcdClient etcdclient.Client) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithReaderJobsDBs sets the reader jobsdbs for the Migrator
func (b *MigratorBuilder) WithReaderJobsDBs(readerJobsDBs []jobsdb.JobsDB) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithShutdown sets the shutdown function for the Migrator
func (b *MigratorBuilder) WithShutdown(shutdown func()) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithTargetURLProvider sets the target URL provider for the Migrator
func (b *MigratorBuilder) WithTargetURLProvider(targetURLProvider func(targetNodeIndex int) (string, error)) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs the Migrator with the provided dependencies
func (b *MigratorBuilder) Build() (Migrator, error) {
	_ = "STUB: not implemented"
	return *new(Migrator), nil
}
