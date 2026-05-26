package repo

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const stagingTableName = warehouseutils.WarehouseStagingFilesTable

const stagingTableColumns = `
	id,
    location,
    source_id,
    destination_id,
    error,
    status,
    first_event_at,
    last_event_at,
    total_events,
	total_bytes,
    created_at,
    updated_at,
    metadata,
    workspace_id,
    bytes_per_table
`

// StagingFiles is a repository for inserting and querying staging files.
type StagingFiles struct {
	*repo
	conf *config.Config
}

type metadataSchema struct {
	UseRudderStorage              bool     `json:"use_rudder_storage"`
	SourceTaskRunID               string   `json:"source_task_run_id"`
	SourceJobID                   string   `json:"source_job_id"`
	SourceJobRunID                string   `json:"source_job_run_id"`
	TimeWindowYear                int      `json:"time_window_year"`
	TimeWindowMonth               int      `json:"time_window_month"`
	TimeWindowDay                 int      `json:"time_window_day"`
	TimeWindowHour                int      `json:"time_window_hour"`
	DestinationRevisionID         string   `json:"destination_revision_id"`
	ServerInstanceID              string   `json:"server_instance_id"`
	SnapshotPatchSize             *int     `json:"snapshot_patch_size,omitempty"`
	SnapshotPatchCompressionRatio *float64 `json:"snapshot_patch_compression_ratio,omitempty"`
}

func StagingFileIDs(stagingFiles []*model.StagingFile) []int64 {
	_ = "STUB: not implemented"
	return nil
}

func metadataFromStagingFile(stagingFile *model.StagingFile) metadataSchema {
	_ = "STUB: not implemented"
	return *new(metadataSchema)
}

func NewStagingFiles(db *sqlmiddleware.DB, conf *config.Config, opts ...Opt) *StagingFiles {
	_ = "STUB: not implemented"
	return nil
}

func (m *metadataSchema) SetStagingFile(stagingFile *model.StagingFile) {
	_ = "STUB: not implemented"
	return
}

// Insert inserts a staging file into the staging files table. It returns the ID of the inserted staging file.
//
// NOTE: The following fields are ignored and set by the database:
// - ID
// - Error
// - CreatedAt
// - UpdatedAt
func (sf *StagingFiles) Insert(ctx context.Context, stagingFile *model.StagingFileWithSchema) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// schema is not used for now

// praseRow is a helper for mapping a row of tableColumns to a model.StagingFile.
func parseStagingFiles(rows *sqlmiddleware.Rows) ([]*model.StagingFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetByID returns staging file with the given ID.
func (sf *StagingFiles) GetByID(ctx context.Context, ID int64) (model.StagingFile, error) {
	_ = "STUB: not implemented"
	return *new(model.StagingFile), nil
}

// GetSchemasByIDs returns staging file schemas for the given IDs.
func (sf *StagingFiles) GetSchemasByIDs(ctx context.Context, ids []int64) ([]model.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If both snapshot and patch are present, use them instead of the regular schema

// Fall back to regular schema only if snapshot or patch is missing

// GetForUploadID retrieves all the staging files associated with a specific upload ID.
func (sf *StagingFiles) GetForUploadID(ctx context.Context, uploadID int64) ([]*model.StagingFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *StagingFiles) Pending(ctx context.Context, sourceID, destinationID string) ([]*model.StagingFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lastStartStagingFileID is used as an optimization to avoid scanning the whole table.

func (sf *StagingFiles) CountPendingForSource(ctx context.Context, sourceID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sf *StagingFiles) CountPendingForDestination(ctx context.Context, destinationID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sf *StagingFiles) countPending(ctx context.Context, query string, value any) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sf *StagingFiles) TotalEventsForUploadID(ctx context.Context, uploadID int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sf *StagingFiles) GetEventTimeRangesByUploadID(ctx context.Context, uploadID int64) ([]model.EventTimeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *StagingFiles) DestinationRevisionIDsForUploadID(ctx context.Context, uploadID int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *StagingFiles) SetStatuses(ctx context.Context, ids []int64, status string) error {
	_ = "STUB: not implemented"
	return nil
}
