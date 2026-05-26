package repo

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	tableUploadTableName            = warehouseutils.WarehouseTableUploadsTable
	tableUploadUniqueConstraintName = "unique_table_upload_wh_upload"

	tableUploadColumns = `
		id,
		wh_upload_id,
		table_name,
		status,
		error,
		last_exec_time,
		total_events,
		created_at,
		updated_at,
		location
	`
)

// TableUploads is a repository for table uploads
type TableUploads struct {
	*repo
}

type TableUploadSetOptions struct {
	Status       *string
	Error        *string
	LastExecTime *time.Time
	Location     *string
	TotalEvents  *int64
}

func NewTableUploads(db *sqlmiddleware.DB, conf *config.Config, opts ...Opt) *TableUploads {
	_ = "STUB: not implemented"
	return nil
}

func (tu *TableUploads) WithTx(ctx context.Context, f func(tx *sqlmiddleware.Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (tu *TableUploads) Insert(ctx context.Context, uploadID int64, tableNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (tu *TableUploads) GetByUploadID(ctx context.Context, uploadID int64) ([]model.TableUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tu *TableUploads) GetByUploadIDAndTableName(ctx context.Context, uploadID int64, tableName string) (model.TableUpload, error) {
	_ = "STUB: not implemented"
	return *new(model.TableUpload), nil
}

func scanTableUploads(rows *sqlmiddleware.Rows) ([]model.TableUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanTableUpload(scan scanFn, tableUpload *model.TableUpload) error {
	_ = "STUB: not implemented"
	return nil
}

// PopulateTotalEventsWithTx Update the 'total_events' field in the Table Uploads table
// by summing the 'total_events' from load files associated with specific staging file IDs.
func (tu *TableUploads) PopulateTotalEventsWithTx(ctx context.Context, tx *sqlmiddleware.Tx, uploadId int64, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (tu *TableUploads) TotalExportedEvents(ctx context.Context, uploadId int64, skipTables []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tu *TableUploads) Set(ctx context.Context, uploadId int64, tableName string, options TableUploadSetOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// remove trailing comma

func (tu *TableUploads) ExistsForUploadID(ctx context.Context, uploadId int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (tu *TableUploads) SyncsInfo(ctx context.Context, uploadID int64) ([]model.TableUploadInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tu *TableUploads) GetByJobRunTaskRun(
	ctx context.Context,
	sourceID,
	destinationID,
	jobRunID,
	taskRunID string,
) ([]model.TableUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
