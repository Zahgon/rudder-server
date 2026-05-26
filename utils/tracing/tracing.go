package tracing

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

type Option func(*Tracer)

func WithNamePrefix(namePrefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

type traceConfig struct {
	timestamp time.Time
	tags      stats.Tags
	kind      stats.SpanKind
}

type TraceOption func(*traceConfig)

func WithTraceStart(t time.Time) TraceOption { _ = "STUB: not implemented"; return *new(TraceOption) }

func WithTraceTags(tags stats.Tags) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}

func WithTraceKind(kind stats.SpanKind) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}

type RecordSpanOption func(*traceConfig)

func WithRecordSpanTags(tags stats.Tags) RecordSpanOption {
	_ = "STUB: not implemented"
	return *new(RecordSpanOption)
}

func WithRecordSpanKind(kind stats.SpanKind) RecordSpanOption {
	_ = "STUB: not implemented"
	return *new(RecordSpanOption)
}

type Tracer struct {
	tracer     stats.Tracer
	namePrefix string
}

func New(tracer stats.Tracer, opts ...Option) *Tracer { _ = "STUB: not implemented"; return nil }

func (t *Tracer) Trace(ctx context.Context, name string, opts ...TraceOption) (context.Context, stats.TraceSpan) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(stats.TraceSpan)
}

func (t *Tracer) TraceFunc(ctx context.Context, name string, f func(ctx context.Context), opts ...TraceOption) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *Tracer) RecordSpan(ctx context.Context, name string, start time.Time, opts ...RecordSpanOption) {
	_ = "STUB: not implemented"
	return
}

func (t *Tracer) getOptions(opts ...TraceOption) (stats.SpanKind, []stats.SpanOption) {
	_ = "STUB: not implemented"
	return *new(stats.SpanKind), nil
}
