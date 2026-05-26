package pubsub

import (
	"context"
	"sync"
)

type DataEvent struct {
	Data  any
	Topic string
}

// DataChannel is a channel which can accept an DataEvent
type DataChannel <-chan DataEvent

// PublishSubscriber stores the information about subscribers interested for a particular topic
type PublishSubscriber struct {
	lastEventMutex sync.RWMutex
	// lastEvent holds the last event for each topic so that we can send it to new subscribers
	lastEvent map[string]*DataEvent

	subscriptionsMutex sync.RWMutex
	// subscriptions keep the list of subscription publishers per topic
	subscriptions map[string]subPublishers
}

func New() *PublishSubscriber { _ = "STUB: not implemented"; return nil }

func (eb *PublishSubscriber) Publish(topic string, data any) { _ = "STUB: not implemented"; return }

func (eb *PublishSubscriber) Subscribe(ctx context.Context, topic string) DataChannel {
	_ = "STUB: not implemented"
	return *new(DataChannel)
}

func (eb *PublishSubscriber) removePubSub(topic string, r *listener) {
	_ = "STUB: not implemented"
	return
}

func (eb *PublishSubscriber) Close() { _ = "STUB: not implemented"; return }

// listener is a slice of subPublisher pointers
type subPublishers []*listener

// listener is responsible to publish events to a single subscription (channel).
type listener struct {
	// the channel of the subscription where events are published
	channel chan DataEvent

	lastValueLock sync.Mutex
	// the last value waiting to be published to the channel
	lastValue *DataEvent

	// channel for signaling the loop to read a new value
	ping chan struct{}
}

func newListener(channel chan DataEvent) *listener { _ = "STUB: not implemented"; return nil }

// publish sets the publisher's lastValue and starts the
// internal goroutine if it is not already started
func (r *listener) publish(data *DataEvent) { _ = "STUB: not implemented"; return }

// signals the startLoop that it has to read the value

// do nothing - leaky bucket

// startLoop publishes lastValues to the subscription's channel until there is no other lastValue to publish
func (r *listener) startLoop() { _ = "STUB: not implemented"; return }

func (r *listener) close() { _ = "STUB: not implemented"; return }
