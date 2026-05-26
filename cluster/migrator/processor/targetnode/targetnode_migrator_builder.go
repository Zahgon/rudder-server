package targetnode

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/cluster/partitionbuffer"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// NewMigratorBuilder creates a new builder for targetnode Migrator
func NewMigratorBuilder(nodeIndex int, nodeName string) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// MigratorBuilder is a builder for targetnode Migrator
type MigratorBuilder struct {
	nodeIndex int
	nodeName  string

	// dependencies
	config            *config.Config
	logger            logger.Logger
	stats             stats.Stats
	etcdClient        etcdclient.Client
	bufferedJobsDBs   [][]partitionbuffer.JobsDBPartitionBuffer
	unbufferedJobsDBs []jobsdb.JobsDB
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

// WithBufferedJobsDBs sets the buffered jobsdbs for the Migrator
func (b *MigratorBuilder) WithBufferedJobsDBs(bufferedJobsDBs [][]partitionbuffer.JobsDBPartitionBuffer) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithUnbufferedJobsDBs sets the unbuffered jobsdbs for the Migrator
func (b *MigratorBuilder) WithUnbufferedJobsDBs(unbufferedJobsDBs []jobsdb.JobsDB) *MigratorBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs the Migrator with the provided dependencies
func (b *MigratorBuilder) Build() (Migrator, error) {
	_ = "STUB: not implemented"
	return *new(Migrator), nil
}

// build a map of unbuffered jobsdb identifiers for validation

// validate buffered jobsdbs: flatten and check identifiers match unbuffered
