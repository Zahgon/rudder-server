package isolation

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type Mode string

const (
	ModeNone        Mode = "none"
	ModeWorkspace   Mode = "workspace"
	ModeDestination Mode = "destination"
)

// GetStrategy returns the strategy for the given isolation mode. An error is returned if the mode is invalid
func GetStrategy(mode Mode, destType string, partitionFilter func(destinationID string) bool, c *config.Config) (Strategy, error) {
	_ = "STUB: not implemented"
	return *new(Strategy), nil
}

// Strategy defines the operations that every different isolation strategy in processor must implement
type Strategy interface {
	// ActivePartitions returns the list of partitions that are active for the given strategy
	ActivePartitions(ctx context.Context, db jobsdb.JobsDB) ([]string, error)
	// AugmentQueryParams augments the given GetQueryParamsT with the strategy specific parameters
	AugmentQueryParams(partition string, params *jobsdb.GetQueryParams)
	// StopIteration returns true if the iteration should be stopped for the given error
	StopIteration(err error, destinationID string) bool
	// StopQueries returns true if the iterator should stop fetching more jobs from jobsDB
	StopQueries(err error, destinationID string) bool
	// SupportsPickupQueryThrottling returns true if the strategy supports pickup query throttling, i.e., if it can throttle queries to jobsDB based on the throttling limits set at destination level
	SupportsPickupQueryThrottling() bool
}

// noneStrategy implements isolation at no level
type noneStrategy struct{}

func (noneStrategy) ActivePartitions(_ context.Context, _ jobsdb.JobsDB) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (noneStrategy) AugmentQueryParams(_ string, _ *jobsdb.GetQueryParams) {
	_ = "STUB: not implemented"
	// no-op
	return
}

func (noneStrategy) StopIteration(_ error, _ string) bool { _ = "STUB: not implemented"; return false }

func (noneStrategy) StopQueries(_ error, _ string) bool { _ = "STUB: not implemented"; return false }

func (noneStrategy) SupportsPickupQueryThrottling() bool {
	_ = "STUB: not implemented"

	// workspaceStrategy implements isolation at workspace level
	return false
}

type workspaceStrategy struct {
	customVal string
}

// ActivePartitions returns the list of active workspaceIDs in jobsdb
func (ws workspaceStrategy) ActivePartitions(ctx context.Context, db jobsdb.JobsDB) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (workspaceStrategy) AugmentQueryParams(partition string, params *jobsdb.GetQueryParams) {
	_ = "STUB: not implemented"
	return
}

func (workspaceStrategy) StopIteration(_ error, _ string) bool {
	_ = "STUB: not implemented"
	return false
}

func (workspaceStrategy) StopQueries(_ error, _ string) bool {
	_ = "STUB: not implemented"
	return false
}

func (workspaceStrategy) SupportsPickupQueryThrottling() bool {
	_ = "STUB: not implemented"

	// destinationStrategy implements isolation at destination level
	return false
}

type destinationStrategy struct {
	config                        *config.Config
	pickupQueryThrottlingEnabled  config.ValueLoader[bool]
	destinationFilter             func(destinationID string) bool
	destType                      string
	throttlerPerEventTypeConfigMu sync.RWMutex
	throttlerPerEventTypeConfig   map[string]config.ValueLoader[bool]
}

// ActivePartitions returns the list of active destinationIDs in jobsdb
func (ds *destinationStrategy) ActivePartitions(ctx context.Context, db jobsdb.JobsDB) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AugmentQueryParams augments the given GetQueryParamsT by adding the partition as sourceID parameter filter
func (*destinationStrategy) AugmentQueryParams(partition string, params *jobsdb.GetQueryParams) {
	_ = "STUB: not implemented"
	return
}

// StopIteration returns true if the error is ErrDestinationThrottled
func (ds *destinationStrategy) StopIteration(err error, destinationID string) bool {
	_ = "STUB: not implemented"
	return false
}

// StopQueries returns true if the error is ErrDestinationThrottled and throttlerPerEventType is enabled for the destination
func (ds *destinationStrategy) StopQueries(err error, destinationID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (ds *destinationStrategy) SupportsPickupQueryThrottling() bool {
	_ = "STUB: not implemented"
	return false
}

func (ds *destinationStrategy) hasDestinationThrottlerPerEventType(destinationID string) bool {
	_ = "STUB: not implemented"
	return false
}
