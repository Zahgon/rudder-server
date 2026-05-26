package badger

import (
	"time"

	"github.com/dgraph-io/badger/v4"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

/*
loadCacheConfig sets the properties of the cache after reading it from the config file.
This gives a feature of hot readability as well.
*/
func (e *Cache[E]) loadCacheConfig() { _ = "STUB: not implemented"; return }

// Using the maximum value threshold: (1 << 20) == 1048576 (1MB)

/*
Cache is an in-memory cache. Each key-value pair stored in this cache have a TTL and one goroutine removes the
key-value pair form the cache which is older than TTL time.
*/
type Cache[E any] struct {
	limiter                 config.ValueLoader[int]
	path                    string
	origin                  string
	done                    chan struct{}
	closed                  chan struct{}
	ticker                  time.Duration
	queryTimeout            time.Duration
	ttl                     config.ValueLoader[time.Duration]
	gcDiscardRatio          float64
	numMemtables            int
	numLevelZeroTables      int
	numLevelZeroTablesStall int
	valueThreshold          int64
	syncWrites              bool
	cleanupOnStartup        bool
	db                      *badger.DB
	logger                  logger.Logger
	stats                   stats.Stats
}

type badgerLogger struct {
	logger.Logger
}

func (l badgerLogger) Errorf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (l badgerLogger) Warningf(format string, a ...any) { _ = "STUB: not implemented"; return }

// Update writes the entries into badger db with a TTL
func (e *Cache[E]) Update(key string, value E) error { _ = "STUB: not implemented"; return nil }

// Read fetches all the entries for a given key from badgerDB
func (e *Cache[E]) Read(key string) ([]E, error) { _ = "STUB: not implemented"; return nil, nil }

// ignore unmarshal errors (old version of the data)

func New[E any](origin string, log logger.Logger, stats stats.Stats, opts ...func(Cache[E])) (*Cache[E], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO : Remove this after badgerdb v2 is completely removed

// 16mb

func (e *Cache[E]) gcBadgerDB() { _ = "STUB: not implemented"; return }

// One call would only result in removal of at max one log file.
// As an optimization, you could also immediately re-run it whenever it returns nil error
// (this is why `goto again` is used).

// see https://dgraph.io/docs/badger/get-started/#garbage-collection

func (e *Cache[E]) Stop() error { _ = "STUB: not implemented"; return nil }
