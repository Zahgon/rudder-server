package transientsource

import (
	"context"
	"encoding/json"
	"sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

// Service provides services related to transient source ids
type Service interface {
	// SourceIdsSupplier provides an up-to-date supplier of transient source ids
	SourceIdsSupplier() func() []string

	// Apply performs transient source_id filtering logic against a single source_id.
	// If it corresponds to a transient source_id, the result will be true, otherwise false.
	Apply(sourceId string) bool

	// ApplyJob performs transient source_id filtering logic against a single job.
	// If the job corresponds to a transient source_id, the result will be true, otherwise false.
	ApplyJob(job *jobsdb.JobT) bool

	// ApplyParams performs transient source_id filtering logic against a job's parameters.
	// If the parameters contain a transient source_id, the result will be true, otherwise false.
	ApplyParams(params json.RawMessage) bool
}

// NewService creates a new service that updates its transient source ids while
// backend configuration gets updated.
func NewService(ctx context.Context, config backendconfig.BackendConfig) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

// NewEmptyService creates a new service that operates against an empty list of transient source ids
// Useful for tests, when you are not interested in testing for transient sources.
func NewEmptyService() Service { _ = "STUB: not implemented"; return *new(Service) }

// NewStaticService creates a new service that operates against a predefined list of transient source ids.
// Useful for tests.
func NewStaticService(sourceIds []string) Service { _ = "STUB: not implemented"; return *new(Service) }

type service struct {
	onceInit     sync.Once
	init         chan struct{}
	sourceIds    []string
	sourceIdsMap map[string]struct{}
}

func (r *service) SourceIdsSupplier() func() []string { _ = "STUB: not implemented"; return nil }

func (r *service) Apply(sourceId string) bool { _ = "STUB: not implemented"; return false }

func (r *service) ApplyParams(params json.RawMessage) bool { _ = "STUB: not implemented"; return false }

func (r *service) ApplyJob(job *jobsdb.JobT) bool { _ = "STUB: not implemented"; return false }

// updateLoop uses backend config to retrieve & keep up-to-date the list of transient source ids
func (r *service) updateLoop(ctx context.Context, config backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

// transientSourceIds scans a backend configuration and extracts
// source ids which have the following configuration option
//
//	transient : true
func transientSourceIds(c *backendconfig.ConfigT) []string { _ = "STUB: not implemented"; return nil }

// asMap converts a slice of strings to a set, i.e. a map of strings to empty structs
func asMap(arr []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }
