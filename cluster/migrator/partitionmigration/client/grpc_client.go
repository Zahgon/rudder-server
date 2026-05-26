// GRPC client for partition migration
package client

import (
	"google.golang.org/grpc"

	"github.com/rudderlabs/rudder-go-kit/config"

	proto "github.com/rudderlabs/rudder-server/proto/cluster"
)

// NewPartitionMigrationClient creates a new client for the PartitionMigration gRPC service
func NewPartitionMigrationClient(target string, conf *config.Config) (PartitionMigrationClient, error) {
	_ = "STUB: not implemented"
	return *new(PartitionMigrationClient), nil
}

// PartitionMigrationClient is a wrapper around the generated gRPC client
type PartitionMigrationClient interface {
	proto.PartitionMigrationClient
	Close() error
}

type partitionMigrator struct {
	conn *grpc.ClientConn
	proto.PartitionMigrationClient
}

func (c *partitionMigrator) Close() error { _ = "STUB: not implemented"; return nil }
