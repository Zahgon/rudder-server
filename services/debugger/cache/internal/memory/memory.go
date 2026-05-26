package memory

import (
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

type cacheItem[E any] struct {
	data       []E
	lastAccess time.Time
}

/*
loadCacheConfig sets the properties of the cache after reading it from the config file.
This gives a feature of hot readability as well.
*/
func (c *Cache[E]) loadCacheConfig() { _ = "STUB: not implemented"; return }

// default keyTTL is 30 days

// default clearFreq is 15 seconds

/*
	Cache is an in-memory cache. Each key-value pair stored in this cache have a TTL and one goroutine removes the

key-value pair form the cache which is older than TTL time.
*/
type Cache[E any] struct {
	lock        sync.RWMutex
	keyTTL      config.ValueLoader[time.Duration] // Time after which the data will be expired and removed from the cache
	cleanupFreq config.ValueLoader[time.Duration] // This is the time at which a cleaner goroutines  checks whether the data is expired in cache
	size        config.ValueLoader[int]           // This is the size upto which this cache can store a value corresponding to any key
	cacheMap    map[string]*cacheItem[E]

	done   chan struct{}
	closed chan struct{}
}

/*
New method initiates the cache object. To initiate, this sets certain properties of the cache like keyTTL,
cleanupFreq, size, empty cacheMap
*/
func New[E any]() (*Cache[E], error) { _ = "STUB: not implemented"; return nil, nil }

/*
Update Inserts the data in the cache, This method expects a string as a key and []byte as the data
*/
func (c *Cache[E]) Update(key string, value E) error { _ = "STUB: not implemented"; return nil }

/*
	ReadAndPopData reads the data by taking a string key,

if there is any data available corresponding to the given key then it removes the data from the cache and returns it
in the form of []byte
*/
func (c *Cache[E]) Read(key string) ([]E, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cache[E]) Stop() error { _ = "STUB: not implemented"; return nil }
