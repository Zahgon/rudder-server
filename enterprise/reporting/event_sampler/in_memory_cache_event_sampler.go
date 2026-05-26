package event_sampler

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/cachettl"
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type InMemoryCacheEventSampler struct {
	ctx    context.Context
	cancel context.CancelFunc
	cache  *cachettl.Cache[string, bool]
	ttl    config.ValueLoader[time.Duration]
	limit  config.ValueLoader[int]
	length int
	sc     *StatsCollector
}

func NewInMemoryCacheEventSampler(
	ctx context.Context,
	module string,
	ttl config.ValueLoader[time.Duration],
	limit config.ValueLoader[int],
	stats stats.Stats,
) (*InMemoryCacheEventSampler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *InMemoryCacheEventSampler) Get(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *InMemoryCacheEventSampler) Put(key string) error { _ = "STUB: not implemented"; return nil }

func (es *InMemoryCacheEventSampler) Close() { _ = "STUB: not implemented"; return }
