package snapshots

import (
	"time"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

// StagingFileSchemaExpiryStrategy defines the interface for cache expiration strategies.
type StagingFileSchemaExpiryStrategy interface {
	IsExpired(snapshot *model.StagingFileSchemaSnapshot) bool
}

// StagingFileSchemaTimeBasedExpiryStrategy expires snapshots after a fixed duration from CreatedAt.
type StagingFileSchemaTimeBasedExpiryStrategy struct {
	duration time.Duration
}

func (t *StagingFileSchemaTimeBasedExpiryStrategy) IsExpired(snapshot *model.StagingFileSchemaSnapshot) bool {
	_ = "STUB: not implemented"
	return false
}

// NewStagingFileSchemaTimeBasedExpiryStrategy returns a time-based expiry strategy for the cache.
func NewStagingFileSchemaTimeBasedExpiryStrategy(duration time.Duration) StagingFileSchemaExpiryStrategy {
	_ = "STUB: not implemented"
	return *new(StagingFileSchemaExpiryStrategy)
}
