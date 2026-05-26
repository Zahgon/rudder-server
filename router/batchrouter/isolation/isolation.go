package isolation

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type Mode string

const (
	ModeNone        Mode = "none"
	ModeWorkspace   Mode = "workspace"
	ModeDestination Mode = "destination"
)

// GetStrategy returns the strategy for the given isolation mode. An error is returned if the mode is invalid
func GetStrategy(mode Mode, customVal string, partitionFilter func(partition string) bool) (Strategy, error) {
	_ = "STUB: not implemented"
	return *new(Strategy), nil
}

// Strategy defines the operations that every different isolation strategy in processor must implement
type Strategy interface {
	// ActivePartitions returns the list of partitions that are active for the given strategy
	ActivePartitions(ctx context.Context, db jobsdb.JobsDB) ([]string, error)
	// AugmentQueryParams augments the given GetQueryParamsT with the strategy specific parameters
	AugmentQueryParams(partition string, params *jobsdb.GetQueryParams)
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

	// workspaceStrategy implements isolation at workspace level
	return
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

// destinationStrategy implements isolation at destination level
type destinationStrategy struct {
	destinationFilter func(destinationID string) bool
}

// ActivePartitions returns the list of active destinationIDs in jobsdb
func (ds destinationStrategy) ActivePartitions(ctx context.Context, db jobsdb.JobsDB) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AugmentQueryParams augments the given GetQueryParamsT by adding the partition as sourceID parameter filter
func (destinationStrategy) AugmentQueryParams(partition string, params *jobsdb.GetQueryParams) {
	_ = "STUB: not implemented"
	return
}
