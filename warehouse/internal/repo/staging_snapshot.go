// Package repo provides repository implementations for warehouse entities.
package repo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	stagingFileSchemaSnapshotTableName    = whutils.WarehouseStagingFileSchemaSnapshotTable
	stagingFileSchemaSnapshotTableColumns = `id, schema, source_id, destination_id, workspace_id, created_at`
)

var ErrNoSchemaSnapshot = errors.New("no schema snapshot found")

type StagingFileSchemaSnapshots repo

// NewStagingFileSchemaSnapshots creates a new StagingFileSchemaSnapshots using the given DB connection.
func NewStagingFileSchemaSnapshots(db *sqlmiddleware.DB, opts ...Opt) *StagingFileSchemaSnapshots {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a new schema snapshot into the database and returns its auto-generated ID.
func (r *StagingFileSchemaSnapshots) Insert(ctx context.Context, sourceID, destinationID, workspaceID string, schemaBytes json.RawMessage) (uuid.UUID, error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

// GetLatest returns the most recent schema snapshot for the given source and destination.
// Returns ErrNoSchemaSnapshot if not found.
func (r *StagingFileSchemaSnapshots) GetLatest(ctx context.Context, sourceID, destinationID string) (*model.StagingFileSchemaSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
