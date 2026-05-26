//go:generate mockgen -destination=../../mocks/services/dedup/mock_dedup.go -package mock_dedup github.com/rudderlabs/rudder-server/services/dedup/types Dedup

package dedup

import (
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/dedup/types"
)

type dedup struct {
	db            types.DB
	uncommittedMu sync.RWMutex
	uncommitted   map[string]struct{}
}

type BatchKey = types.BatchKey

// SingleKey creates a BatchKey with index 0
func SingleKey(key string) BatchKey { _ = "STUB: not implemented"; return *new(BatchKey) }

// New creates a new deduplication service. The service needs to be closed after use.
func New(conf *config.Config, stats stats.Stats, log logger.Logger) (types.Dedup, error) {
	_ = "STUB: not implemented"
	return *new(types.Dedup), nil
}

func (d *dedup) Allowed(batchKeys ...types.BatchKey) (map[types.BatchKey]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keys encountered for the first time
// keys already seen in the batch while iterating

// figure out which keys need to be checked against the DB
// keys to check in the DB

// if the key is already seen in the batch, skip it

// if the key is already in the uncommitted list , skip it

// if another goroutine managed to set this key, we should skip it

// mark this key as uncommitted

func (d *dedup) Commit(keys []string) error { _ = "STUB: not implemented"; return nil }

func (d *dedup) Close() { _ = "STUB: not implemented"; return }
