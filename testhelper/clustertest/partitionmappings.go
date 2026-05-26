package clustertest

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/partmap"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
)

type wpmh struct {
	workspaceID   string
	namespace     string
	client        etcdclient.Client
	numPartitions int
}

// NewWorkspacePartitionMappingHandler creates a handler which can read and write partition mappings for a workspace in etcd.
func NewWorkspacePartitionMappingHandler(client etcdclient.Client, numPartitions int, namespace, workspaceID string) *wpmh {
	_ = "STUB: not implemented"
	return nil
}

// SetWorkspacePartitionMappings sets the partition mappings for the workspace in etcd.
func (wpmh *wpmh) SetWorkspacePartitionMappings(ctx context.Context, mappings partmap.PartitionIndexMapping) error {
	_ = "STUB: not implemented"
	return nil
}

// GetWorkspacePartitionMappings gets the partition mappings for the workspace from etcd.
func (wpmh *wpmh) GetWorkspacePartitionMappings(ctx context.Context) (partmap.PartitionIndexMapping, error) {
	_ = "STUB: not implemented"
	return *new(partmap.PartitionIndexMapping), nil
}

func workspacePartitionMappingKey(workspaceID, namespace string) string {
	_ = "STUB: not implemented"
	return ""
}
