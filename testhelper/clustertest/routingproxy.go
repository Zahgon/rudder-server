package clustertest

import (
	"net/http/httptest"
	"net/http/httputil"
	"sync"
	"testing"

	"github.com/rudderlabs/rudder-go-kit/partmap"
)

type PartitionRoutingProxy interface {
	// Close stops the routing proxy server
	Close()
	// UpdatePartitionMapping updates the mapping of a partition to a node index
	UpdatePartitionMapping(partitionIdx partmap.PartitionIndex, nodeIndex partmap.NodeIndex)
	// SetPartitionMappings sets the entire partition to node index mapping
	SetPartitionMappings(partitionMappings map[partmap.PartitionIndex]partmap.NodeIndex)
}

// NewRoutingProxy creates a new routing proxy that routes requests to different backends based on partition mappings.
// The proxy expects requests to have an "X-Partition-Key" header, which it uses to determine the partition index.
//
// Parameters:
// - numPartitions: Total number of partitions.
// - mappings: Initial mapping of partition indices to backend node indices.
// - backendUrls: URLs of the backend servers to route requests to. Order matters and should correspond to node indices.
//
// Returns:
// - A PartitionRoutingProxy instance that can be used to manage the routing proxy.
func NewRoutingProxy(t *testing.T, numPartitions int, mappings partmap.PartitionIndexMapping, backendUrls ...string) *routingProxy {
	_ = "STUB: not implemented"
	return nil
}

// unlock only after request is processed

type routingProxy struct {
	*httptest.Server
	numPartitions       int
	partitionMappingsMu sync.RWMutex
	partitionMappings   map[partmap.PartitionIndex]partmap.NodeIndex
	backends            []*httputil.ReverseProxy
}

// SetPartitionMappings sets the entire partition to node index mapping. This method returns only after acquiring a write lock,
// ensuring that:
// 1. All ongoing requests are processed with the old mapping before the new mapping takes effect.
// 2. Post-return, any new incoming requests will be routed based on the updated mapping.
func (rp *routingProxy) SetPartitionMappings(partitionMappings map[partmap.PartitionIndex]partmap.NodeIndex) {
	_ = "STUB: not implemented"
	return
}
