package processor

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/cluster/migrator/processor/sourcenode"
	"github.com/rudderlabs/rudder-server/cluster/migrator/processor/targetnode"
)

// NewProcessorPartitionMigratorBuilder creates a new builder for ProcessorPartitionMigrator
func NewProcessorPartitionMigratorBuilder(nodeIndex int, nodeName string) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ProcessorPartitionMigratorBuilder is a builder for ProcessorPartitionMigrator
type ProcessorPartitionMigratorBuilder struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config         *config.Config
	logger         logger.Logger
	stats          stats.Stats
	etcdClient     etcdclient.Client
	sourceMigrator sourcenode.Migrator
	targetMigrator targetnode.Migrator
}

// WithConfig sets the configuration for the ProcessorPartitionMigrator
func (b *ProcessorPartitionMigratorBuilder) WithConfig(config *config.Config) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithLogger sets the logger for the ProcessorPartitionMigrator
func (b *ProcessorPartitionMigratorBuilder) WithLogger(logger logger.Logger) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithStats sets the stats collector for the ProcessorPartitionMigrator
func (b *ProcessorPartitionMigratorBuilder) WithStats(stats stats.Stats) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithEtcdClient sets the etcd client for the ProcessorPartitionMigrator
}

func (b *ProcessorPartitionMigratorBuilder) WithEtcdClient(etcdClient etcdclient.Client) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithSourceMigrator sets the source migrator for the ProcessorPartitionMigrator
func (b *ProcessorPartitionMigratorBuilder) WithSourceMigrator(sourceMigrator sourcenode.Migrator) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithTargetMigrator sets the target migrator for the ProcessorPartitionMigrator
func (b *ProcessorPartitionMigratorBuilder) WithTargetMigrator(targetMigrator targetnode.Migrator) *ProcessorPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs the ProcessorPartitionMigrator with the provided dependencies
func (b *ProcessorPartitionMigratorBuilder) Build() (PartitionMigrator, error) {
	_ = "STUB: not implemented"
	return *new(PartitionMigrator), nil
}
