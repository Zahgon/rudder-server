package cache

import (
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

const (
	wildcard = "*"
)

type Option[T ParameterFilter] func(*NoResultsCache[T])

// WithWarnOnBranchInvalidation is a config option that enables logging of branch invalidations.
func WithWarnOnBranchInvalidation[T ParameterFilter](enabled config.ValueLoader[bool], logger logger.Logger) Option[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewNoResultsCache creates a new, properly initialised NoResultsCache.
func NewNoResultsCache[T ParameterFilter](supportedParams []string, ttlFn func() time.Duration, opts ...Option[T]) *NoResultsCache[T] {
	_ = "STUB: not implemented"
	return nil
}

type ParameterFilter interface {
	GetName() string
	GetValue() string
}

type NoResultsCache[T ParameterFilter] struct {
	ttl                      func() time.Duration // returns the time to live for a cache entry
	supportedParams          []string             // a list of parameters that are supported by the cache
	warnOnBranchInvalidation config.ValueLoader[bool]
	logger                   logger.Logger

	cacheTreeMu sync.RWMutex // protects the cacheTree
	cacheTree   cacheTree    // a hierarchical tree of cache entries
}

// Get returns true if the cache contains a valid entry for the provided dataset, partitions, workspace, customVals, states and parameters filters.
func (c *NoResultsCache[T]) Get(dataset string, partitions []string, workspace string, customVals, states []string, parameters []T) bool {
	_ = "STUB: not implemented"
	return false
}

// Invalidate invalidates all cache entries for the provided dataset, partitions, workspace, customVals, states and parameters.
func (c *NoResultsCache[T]) Invalidate(dataset string, partitions []string, workspace string, customVals, states []string, parameters []T) {
	_ = "STUB: not implemented"
	return
}

// if no partitions are provided, invalidate all by deleting the partitions's parent node

// if no workspace is provided, invalidate all by deleting the workspace's parent node

// TODO: invalidate entire branch in next release

// if no custom value is provided, invalidate all by deleting the customVal's parent node

// TODO: invalidate entire branch in next release

// if no state is provided, invalidate all by deleting the state's parent node

// if no parameter is provided, invalidate all by deleting the param's parent node

// if logging is enabled, log the invalidation of the leaf node at debug level, since this is the most granular level

// InvalidateDataset invalidates all cache entries for a given dataset.
func (c *NoResultsCache[T]) InvalidateDataset(dataset string) { _ = "STUB: not implemented"; return }

// InvalidatePartitions invalidates all cache entries for the given partitions.
func (c *NoResultsCache[T]) InvalidatePartitions(partitions []string) {
	_ = "STUB: not implemented"
	return
}

// StartNoResultTx prepares the cache for accepting new no result entries.
// The cache uses a special marker to prevent synchronisation issues between competing calls of Invalidate & SetNoResult.
func (c *NoResultsCache[T]) StartNoResultTx(dataset string, partitions []string, workspace string, customVals, states []string, parameters []T) (tx *NoResultTx[T]) {
	_ = "STUB: not implemented"
	return nil
}

// NoResultTx is a transaction for the NoResultsCache.
type NoResultTx[T ParameterFilter] struct {
	id                 string
	dataset            string
	partitions         []string
	workspace          string
	customVals, states []string
	parameters         []T
	c                  *NoResultsCache[T]
}

// Commit sets the necessary cache entries for the relevant dataset, workspace, states, customVals and parameters filters.
// It returns [true] if the cache was successfully updated for all cache entries, [false] otherwise.
func (tx *NoResultTx[T]) Commit() bool { _ = "STUB: not implemented"; return false }

// skipCache returns true if the cache should be skipped for the provided states and parameters.
func (c *NoResultsCache[T]) skipCache(states []string, parameters []T) bool {
	_ = "STUB: not implemented"
	// if no state filters are provided, we don't use the cache
	return false
}

// if not all parameter filters are a subset of the supported parameters, we don't use the cache

// filtersToCacheKeys returns the cache keys for the provided partition, workspace, states, customVals and parameters filters.
// Wildcards are used if empty parameters are provided.
func filtersToCacheKeys[T ParameterFilter](partitionFilter []string, workspaceFilter string, statesFilter, customValsFilter []string, parametersFilter []T) (partitionKeys []string, workspaceKey string, stateKeys, customValKeys, paramKeys []string) {
	_ = "STUB: not implemented"
	return nil, "", nil, nil, nil
}

// if no partition is provided, we use the wildcard

// if no workspace is provided, we use the wildcard

// if no custom value is provided, use the wildcard

// if no parameter is provided, we use the wildcard

// filtersToInvalidationKeys returns the cache keys that need to be invalidated for the provided workspace, states, customVals and parameters filters.
// Wildcard keys are also returned if needed. An empty slice is returned if all keys need to be invalidated at that level.
func (c *NoResultsCache[T]) filtersToInvalidationKeys(partitionFilter []string, workspaceFilter string, statesFilter, customValsFilter []string, parametersFilter []T) (partitionKeys, workspaceKeys, stateKeys, customValKeys, paramKeys []string) {
	_ = "STUB: not implemented"
	return nil,
		// include partitions along with the wildcard
		nil, nil, nil, nil
}

// include customVals along with the wildcard

// include params along with the wildcard

// String returns a string representation of the cache's tree contents.
func (c *NoResultsCache[T]) String() string { _ = "STUB: not implemented"; return "" }

type (
	datasetKey   = string
	partitionKey = string
	workspaceKey = string
	customValKey = string
	stateKey     = string
	paramKey     = string
	cacheTree    map[datasetKey]map[partitionKey]map[workspaceKey]map[customValKey]map[stateKey]map[paramKey]cacheEntry
)

type cacheEntry struct {
	noJobs bool
	tokens []string
	t      time.Time
}

// AddToken adds a token to the cache entry and removes the oldest one if there are more than 10.
func (ce *cacheEntry) AddToken(token string) { _ = "STUB: not implemented"; return }

// SetNoJobs sets the noJobs flag to true if the provided token is found in the cache entry.
func (ce *cacheEntry) SetNoJobs(token string) bool { _ = "STUB: not implemented"; return false }
