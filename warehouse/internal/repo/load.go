package repo

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	loadTableName    = warehouseutils.WarehouseLoadFilesTable
	loadTableColumns = `
		id,
		location,
		source_id,
		destination_id,
		destination_type,
		table_name,
		total_events,
		metadata,
		created_at,
		upload_id
`
)

type LoadFiles struct {
	*repo
}

func NewLoadFiles(db *sqlmiddleware.DB, conf *config.Config, opts ...Opt) *LoadFiles {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes load files associated with the uploadID.
func (lf *LoadFiles) Delete(ctx context.Context, uploadID int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Insert loadFiles into the database.
func (lf *LoadFiles) Insert(ctx context.Context, loadFiles []model.LoadFile) error {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LoadFiles) Get(ctx context.Context, uploadID int64) ([]model.LoadFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanLoadFiles(rows *sqlmiddleware.Rows) ([]model.LoadFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanLoadFile(scan scanFn, loadFile *model.LoadFile) error {
	_ = "STUB: not implemented"
	return nil
}

// GetByID returns the load file matching the id.
func (lf *LoadFiles) GetByID(ctx context.Context, id int64) (*model.LoadFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalExportedEvents returns the total number of events exported by the corresponding staging files.
// It excludes the tables present in skipTables.
func (lf *LoadFiles) TotalExportedEvents(
	ctx context.Context,
	uploadID int64,
	skipTables []string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DistinctTableName returns the distinct table names for the given parameters.
func (lf *LoadFiles) DistinctTableName(
	ctx context.Context,
	sourceID string,
	destinationID string,
	startID int64,
	endID int64,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
