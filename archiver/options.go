package archiver

import (
	"time"

	"github.com/rudderlabs/rudder-server/utils/payload"
)

type Option func(*archiver)

func WithAdaptiveLimit(limiter payload.AdaptiveLimiterFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithArchiveTrigger(trigger func() <-chan time.Time) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithArchiveFrom(archiveFrom string) Option { _ = "STUB: not implemented"; return *new(Option) }
