package migrator

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
)

// NewGatewayPartitionMigratorBuilder creates a new builder for GatewayPartitionMigrator
func NewGatewayPartitionMigratorBuilder(nodeIndex int, nodeName string) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// GatewayPartitionMigratorBuilder is a builder for GatewayPartitionMigrator
type GatewayPartitionMigratorBuilder struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config             *config.Config
	logger             logger.Logger
	stats              stats.Stats
	etcdClient         etcdclient.Client
	partitionRefresher PartitionRefresher
}

// WithConfig sets the configuration for the GatewayPartitionMigrator
func (b *GatewayPartitionMigratorBuilder) WithConfig(config *config.Config) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithLogger sets the logger for the GatewayPartitionMigrator
func (b *GatewayPartitionMigratorBuilder) WithLogger(logger logger.Logger) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithStats sets the stats collector for the GatewayPartitionMigrator
func (b *GatewayPartitionMigratorBuilder) WithStats(stats stats.Stats) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithEtcdClient sets the etcd client for the GatewayPartitionMigrator
}

func (b *GatewayPartitionMigratorBuilder) WithEtcdClient(etcdClient etcdclient.Client) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithPartitionRefresher sets the partition refresher for the GatewayPartitionMigrator
func (b *GatewayPartitionMigratorBuilder) WithPartitionRefresher(partitionRefresher PartitionRefresher) *GatewayPartitionMigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs the GatewayPartitionMigrator with the provided dependencies
func (b *GatewayPartitionMigratorBuilder) Build() (PartitionMigrator, error) {
	_ = "STUB: not implemented"
	return *new(PartitionMigrator), nil
}
