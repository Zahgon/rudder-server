package throttler

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/router/throttler/internal/pickup/adaptive"
	"github.com/rudderlabs/rudder-server/router/throttler/internal/types"
)

const (
	throttlingAlgoTypeGCRA           = "gcra"
	throttlingAlgoTypeRedisGCRA      = "redis-gcra"
	throttlingAlgoTypeRedisSortedSet = "redis-sorted-set"
)

type (
	PickupThrottler   = types.PickupThrottler
	DeliveryThrottler = types.DeliveryThrottler
)

type Factory interface {
	// GetPickupThrottler returns a PickupThrottler for the given destination type, ID, and event type.
	GetPickupThrottler(destType, destID, eventType string) PickupThrottler
	// GetActivePickupThrottlers returns all instantiated PickupThrottlers for the given destination ID.
	GetActivePickupThrottlers(destinationID string) []PickupThrottler
	// GetDeliveryThrottler returns a DeliveryThrottler for the given destination type, ID, and endpoint path.
	GetDeliveryThrottler(destType, destID, endpointPath string) DeliveryThrottler
	// Shutdown gracefully shuts down the factory and all its throttlers.
	Shutdown()
}

// NewFactory constructs a new Throttler Factory
func NewFactory(config *config.Config, stats stats.Stats, log logger.Logger) (Factory, error) {
	_ = "STUB: not implemented"
	return *new(Factory), nil
}

type factory struct {
	config          *config.Config
	log             logger.Logger
	Stats           stats.Stats
	staticLimiter   limiter // limiter to use when static throttling is enabled
	adaptiveLimiter limiter // limiter to use when adaptive throttling is enabled

	mu                       sync.RWMutex                  // protects the resources below
	pickupThrottlers         *pickupThrottlers             // map key is the destinationID:eventType
	allEventTypesPickupAlgos map[string]adaptive.Algorithm // map key is the destinationID
	deliveryThrottlers       map[string]DeliveryThrottler  // map key is the destinationID:endpointPath
}

func (f *factory) GetPickupThrottler(destType, destinationID, eventType string) PickupThrottler {
	_ = "STUB: not implemented"
	// Use read lock first for common case
	return *new(PickupThrottler)
}

// Upgrade to write lock only when needed

// Double-check after acquiring write lock

// switching between static and adaptive throttling

func (f *factory) GetActivePickupThrottlers(destinationID string) []PickupThrottler {
	_ = "STUB: not implemented"
	return nil
}

func (f *factory) GetDeliveryThrottler(destType, destinationID, endpointPath string) DeliveryThrottler {
	_ = "STUB: not implemented"
	return *new(DeliveryThrottler)
}

// Use read lock first for common case

// Upgrade to write lock only when needed

// Double-check after acquiring write lock

// delivery throttler shall be using the static limiter exclusively (redis or in-memory)

func (f *factory) Shutdown() { _ = "STUB: not implemented"; return }

func (f *factory) initThrottlerFactory() error { _ = "STUB: not implemented"; return nil }

type NewNoOpFactory struct{}

func NewNoOpThrottlerFactory() Factory { _ = "STUB: not implemented"; return *new(Factory) }

func (f *NewNoOpFactory) GetPickupThrottler(destName, destID, eventType string) PickupThrottler {
	_ = "STUB: not implemented"
	return *new(PickupThrottler)
}

func (f *NewNoOpFactory) GetActivePickupThrottlers(destinationID string) []PickupThrottler {
	_ = "STUB: not implemented"
	return nil
}

func (f *NewNoOpFactory) GetDeliveryThrottler(destType, destID, endpointPath string) DeliveryThrottler {
	_ = "STUB: not implemented"
	return *new(DeliveryThrottler)
}

func (f *NewNoOpFactory) Shutdown() { _ = "STUB: not implemented"; return }

type noOpThrottler struct{}

func (t *noOpThrottler) CheckLimitReached(ctx context.Context, cost int64) (limited bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *noOpThrottler) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

func (t *noOpThrottler) Shutdown() { _ = "STUB: not implemented"; return }

func (t *noOpThrottler) GetLimitPerSecond() int64 { _ = "STUB: not implemented"; return 0 }

func (t *noOpThrottler) GetEventType() string { _ = "STUB: not implemented"; return "" }

func (t *noOpThrottler) GetLastUsed() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type noOpDeliveryThrottler struct{}

func (*noOpDeliveryThrottler) Wait(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

type limiter interface {
	Allow(ctx context.Context, cost, rate, window int64, key string) (bool, func(context.Context) error, error)
	AllowAfter(ctx context.Context, cost, rate, window int64, key string) (bool, time.Duration, func(context.Context) error, error)
}

type pickupThrottlers struct {
	all map[string]map[string]PickupThrottler
}

func (p *pickupThrottlers) Get(destID, eventType string) (PickupThrottler, bool) {
	_ = "STUB: not implemented"
	return *new(PickupThrottler), false
}

func (p *pickupThrottlers) Set(destID, eventType string, throttler PickupThrottler) {
	_ = "STUB: not implemented"
	return
}
