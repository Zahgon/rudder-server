package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
)

type stagingFilesRepo interface {
	Insert(ctx context.Context, stagingFile *model.StagingFileWithSchema) (int64, error)
}

type stagingFileSchemaSnapshotGetter interface {
	GetOrCreate(ctx context.Context, sourceID, destinationID, workspaceID string, schemaBytes json.RawMessage) (*model.StagingFileSchemaSnapshot, error)
}

// StagingFileSchemaSnapshotHandler Handler for schema snapshot logic
type StagingFileSchemaSnapshotHandler struct {
	Snapshots stagingFileSchemaSnapshotGetter
	PatchGen  func(original, modified json.RawMessage) (json.RawMessage, error)
}

func (h *StagingFileSchemaSnapshotHandler) Apply(ctx context.Context, stagingFile model.StagingFileWithSchema) (model.StagingFileWithSchema, error) {
	_ = "STUB: not implemented"
	return *new(model.StagingFileWithSchema), nil
}

type WarehouseAPI struct {
	Logger                logger.Logger
	Stats                 stats.Stats
	Repo                  stagingFilesRepo
	Multitenant           *multitenant.Manager
	SchemaSnapshotHandler *StagingFileSchemaSnapshotHandler
}

type destinationSchema struct {
	Source      backendconfig.SourceT
	Destination backendconfig.DestinationT
}

type stagingFileSchema struct {
	WorkspaceID           string
	Schema                map[string]map[string]any
	BatchDestination      destinationSchema
	Location              string
	FirstEventAt          time.Time
	LastEventAt           time.Time
	TotalEvents           int
	TotalBytes            int
	BytesPerTable         map[string]int64
	UseRudderStorage      bool
	DestinationRevisionID string
	// cloud sources specific info
	SourceTaskRunID  string
	SourceJobID      string
	SourceJobRunID   string
	TimeWindow       time.Time
	ServerInstanceID string
}

func mapStagingFile(payload *stagingFileSchema) (model.StagingFileWithSchema, error) {
	_ = "STUB: not implemented"
	return *new(model.StagingFileWithSchema), nil
}

// Handler returns a http handler for the warehouse API.
//
// Implemented routes:
// - POST /v1/process
func (api *WarehouseAPI) Handler() http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (api *WarehouseAPI) processHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
