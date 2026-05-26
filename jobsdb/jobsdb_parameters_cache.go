package jobsdb

import (
	"context"
	"sync"
	"time"

	"github.com/samber/lo"

	"github.com/rudderlabs/rudder-go-kit/config"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"
)

// NewCachingDistinctParameterValuesJobsdb creates a new JobsDB decorator that caches the results of GetDistinctParameterValues
// for a specified duration. It uses a partition locker to ensure that only one goroutine can access the cache for a
// specific parameter at a time. The cache is invalidated after the specified TTL.
// The cache is stored in memory and is not persistent across restarts.
// The cache is thread-safe and can be accessed by multiple goroutines concurrently.
func NewCachingDistinctParameterValuesJobsdb(
	ttl config.ValueLoader[time.Duration],
	jobsdb JobsDB,
) JobsDB {
	_ = "STUB: not implemented"
	return *new(JobsDB)
}

type cachingDpvJobsDB struct {
	ttl           config.ValueLoader[time.Duration]
	parameterLock *kitsync.PartitionLocker
	cacheMu       sync.RWMutex
	cache         map[string]lo.Tuple2[[]string, time.Time]
	JobsDB
}

func (c *cachingDpvJobsDB) GetDistinctParameterValues(ctx context.Context, parameter ParameterName, customVal string) (values []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only one goroutine can access the cache for a specific parameter at a time

// read the cache

// if not in cache or expired, fetch from DB
// and update the cache

// update the cache with the new values
