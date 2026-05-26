package event_sampler

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

const (
	BadgerTypeEventSampler        = "badger"
	InMemoryCacheTypeEventSampler = "in_memory_cache"
	MetricsReporting              = "metrics-reporting"
	ErrorsReporting               = "errors-reporting"
)

//go:generate mockgen -destination=../../../mocks/enterprise/reporting/event_sampler/mock_event_sampler.go -package=mocks github.com/rudderlabs/rudder-server/enterprise/reporting/event_sampler EventSampler
type EventSampler interface {
	Put(key string) error
	Get(key string) (bool, error)
	Close()
}

func NewEventSampler(
	ctx context.Context,
	ttl config.ValueLoader[time.Duration],
	eventSamplerType config.ValueLoader[string],
	eventSamplingCardinality config.ValueLoader[int],
	module string,
	conf *config.Config,
	log logger.Logger,
	stats stats.Stats,
) (es EventSampler, err error) {
	_ = "STUB: not implemented"
	return *new(EventSampler), nil
}
