// Package snapshots provides a memory-cached, pluggable-expiry schema snapshot lookup for staging files.
package snapshots

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-go-kit/cachettl"
	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

// StagingFileSchemaDBRepo defines the interface for DB operations needed by the cache.
type StagingFileSchemaDBRepo interface {
	Insert(ctx context.Context, sourceID, destinationID, workspaceID string, schemaBytes json.RawMessage) (uuid.UUID, error)
	GetLatest(ctx context.Context, sourceID, destinationID string) (*model.StagingFileSchemaSnapshot, error)
}

// StagingFileSchema provides a memory-cached, pluggable-expiry schema snapshot lookup.
type StagingFileSchema struct {
	dbRepo               StagingFileSchemaDBRepo
	cache                *cachettl.Cache[string, *model.StagingFileSchemaSnapshot] // cache for the latest schema snapshot to avoid DB lookups
	cacheRefreshInterval config.ValueLoader[time.Duration]                         // interval at which to refresh the cache
	cacheRefreshJitter   func() time.Duration                                      // jitter to add to the cache refresh interval
	now                  func() time.Time                                          // function to get the current time
	expiryStrategy       StagingFileSchemaExpiryStrategy                           // strategy to determine if the cache is expired
}

// NewStagingFileSchema creates a new cache with the given DB repo and expiration strategy.
func NewStagingFileSchema(conf *config.Config, dbRepo StagingFileSchemaDBRepo, expiryStrategy StagingFileSchemaExpiryStrategy) *StagingFileSchema {
	_ = "STUB: not implemented"
	return nil
}

// GetOrCreate returns the latest schema snapshot for the given IDs, using cache and DB as needed. If not found or expired, inserts a new snapshot.
func (c *StagingFileSchema) GetOrCreate(ctx context.Context, sourceID, destinationID, workspaceID string, schemaBytes json.RawMessage) (*model.StagingFileSchemaSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check cache

// Not expired: return the snapshot

// Expired: insert new and fetch

// Cache miss: fetch from DB

// Expired in DB: insert new and fetch

// Only insert if the error is ErrNoSchemaSnapshot (no entry)

// insertAndCache inserts a new snapshot into the DB, and caches the result.
func (c *StagingFileSchema) insertAndCache(ctx context.Context, sourceID, destinationID, workspaceID string, schemaBytes json.RawMessage) (*model.StagingFileSchemaSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the snapshot and insert it into the cache

func (c *StagingFileSchema) cacheRefreshTTL() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func cacheKey(sourceID, destinationID string) string { _ = "STUB: not implemented"; return "" }
