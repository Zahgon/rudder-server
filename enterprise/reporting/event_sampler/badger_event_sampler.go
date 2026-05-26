package event_sampler

import (
	"context"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type BadgerEventSampler struct {
	db     *badger.DB
	dbPath string
	module string
	mu     sync.Mutex
	ttl    config.ValueLoader[time.Duration]
	ctx    context.Context
	cancel context.CancelFunc
	logger badgerLogger
	wg     sync.WaitGroup
	sc     *StatsCollector
}

func GetPathName(module string) string { _ = "STUB: not implemented"; return "" }

func DefaultPath(pathName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewBadgerEventSampler(
	ctx context.Context,
	module string,
	ttl config.ValueLoader[time.Duration],
	conf *config.Config,
	log logger.Logger,
	stats stats.Stats,
) (*BadgerEventSampler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// corrupted or incompatible db, clean up the directory and retry

func (es *BadgerEventSampler) Get(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *BadgerEventSampler) Put(key string) error { _ = "STUB: not implemented"; return nil }

func (es *BadgerEventSampler) gcLoop() { _ = "STUB: not implemented"; return }

// One call would only result in removal of at max one log file.
// As an optimization, you could also immediately re-run it whenever it returns nil error
// (this is why `goto again` is used).

func (es *BadgerEventSampler) Close() { _ = "STUB: not implemented"; return }

type badgerLogger struct {
	logger.Logger
}

func (badgerLogger) Errorf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (badgerLogger) Warningf(format string, a ...any) { _ = "STUB: not implemented"; return }
