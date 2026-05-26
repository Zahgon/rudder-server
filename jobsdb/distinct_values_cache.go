package jobsdb

import (
	"sync"

	kitsync "github.com/rudderlabs/rudder-go-kit/sync"
)

func NewDistinctValuesCache() *distinctValuesCache { _ = "STUB: not implemented"; return nil }

type distinctValuesCache struct {
	cacheMu sync.RWMutex
	// key, dataset
	cache map[string]map[string][]string
	klock *kitsync.PartitionLocker
}

// GetDistinctValues returns the distinct values for the given key and datasets. If the values are
// already cached, it returns the cached values. If not, it loads the values from the given load
// function and caches them. The last dataset is never cached, so it is always loaded from the
// load function. The load function is called with the missing datasets and the last dataset.
func (dvc *distinctValuesCache) GetDistinctValues(key string, datasets []string, load func(datasets []string) (map[string][]string, error)) ([]string, error) {
	_ = "STUB: not implemented"
	// First check if we are missing any datasets from the cache.
	// If we are, we need to load them along with the last dataset
	// The last dataset is never cached, so we need to load it every time
	return nil, nil
}

// If we are missing any datasets, we need to lock the key, so that
// we don't load the same datasets multiple times for the same key.
// This lock needs to be retained until the datasets are loaded into the cache.

// Check again if we are missing any datasets, to deal with race conditions

// If we are not missing any datasets, we need to unlock the key

// Load all the missing datasets along with the last dataset

// if we were missing any datasets, we need to add them to the cache and unlock the key

// Now we need to get values for all the datasets requested so that we can calculate
// the distinct values.
// We already have some values in the results map (last dataset & missing), so we only need to fill in
// the rest of the datasets from the cache.

// Calculating distinct values is easy, we just need to
// iterate over all the datasets and add them to a map
// and then return the keys of the map.

// RemoveDataset removes the dataset from the cache for all keys.
func (dvc *distinctValuesCache) RemoveDataset(dataset string) { _ = "STUB: not implemented"; return }

func (dvc *distinctValuesCache) missing(key string, datasets []string) []string {
	_ = "STUB: not implemented"
	return nil
}
