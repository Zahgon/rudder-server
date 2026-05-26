package rmetrics

import (
	"sync"

	"github.com/rudderlabs/rudder-go-kit/stats/metric"
)

const (
	JobsdbPendingEventsCount = "jobsdb_%s_pending_events_count"
	All                      = "ALL"
)

type (
	DecreasePendingEventsFunc func(tablePrefix, workspaceID, destType, destinationID string, value float64)
	IncreasePendingEventsFunc func(tablePrefix, workspaceID, destType, destinationID string, value float64)
)

// PendingEventsRegistry is a registry for pending events metrics
type PendingEventsRegistry interface {
	// IncreasePendingEvents increments three gauges, the dest & workspace-specific gauge, plus two aggregate (global) gauges
	IncreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64)
	// DecreasePendingEvents decrements three gauges, the dest & workspace-specific gauge, plus two aggregate (global) gauges
	DecreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64)
	// PendingEvents gets the measurement for pending events metric
	PendingEvents(tablePrefix, workspaceID, destType, destinationID string) metric.Gauge
	// Publish publishes the metrics to the global published metrics registry
	Publish()
	// Reset resets the registry to a new, non published one and clears the global published metrics registry
	Reset()
}

type Option func(*pendingEventsRegistry)

// WithPublished creates a registry that writes metrics to the global published metrics registry, without having to call Publish first.
func WithPublished() Option { _ = "STUB: not implemented"; return *new(Option) }

// NewPendingEventsRegistry creates a new PendingEventsRegistry. By default, metrics are not published to the global published metrics registry, until [Publish] is called.
func NewPendingEventsRegistry(opts ...Option) PendingEventsRegistry {
	_ = "STUB: not implemented"
	return *new(PendingEventsRegistry)
}

type pendingEventsRegistry struct {
	registryMu sync.RWMutex
	published  bool
	registry   metric.Registry
}

// IncreasePendingEvents increments three gauges, the dest & workspace-specific gauge, plus two aggregate (global) gauges
func (pem *pendingEventsRegistry) IncreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64) {
	_ = "STUB: not implemented"
	return
}

// DecreasePendingEvents decrements three gauges, the dest & workspace-specific gauge, plus two aggregate (global) gauges
func (pem *pendingEventsRegistry) DecreasePendingEvents(tablePrefix, workspaceID, destType, destinationID string, value float64) {
	_ = "STUB: not implemented"
	return
}

// PendingEvents gets the measurement for pending events metric
func (pem *pendingEventsRegistry) PendingEvents(tablePrefix, workspaceID, destType, destinationID string) metric.Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Gauge)
}

// Publish publishes the metrics to the global published metrics registry
func (pem *pendingEventsRegistry) Publish() { _ = "STUB: not implemented"; return }

// copy all gauge metrics to the published registry

// Reset resets the registry to a new, non published one and clears the global published metrics registry
func (pem *pendingEventsRegistry) Reset() { _ = "STUB: not implemented"; return }

func newPendingEventsMeasurement(tablePrefix, workspaceID, destType, destinationID string) metric.Measurement {
	_ = "STUB: not implemented"
	return *new(metric.Measurement)
}

type pendingEventsMeasurement struct {
	tablePrefix   string
	workspaceID   string
	destType      string
	destinationID string
}

func (r pendingEventsMeasurement) GetName() string { _ = "STUB: not implemented"; return "" }

func (r pendingEventsMeasurement) GetTags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type pendingEventsMeasurementAll struct {
	tablePrefix string
	destType    string
}

func (r pendingEventsMeasurementAll) GetName() string { _ = "STUB: not implemented"; return "" }

func (r pendingEventsMeasurementAll) GetTags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}
