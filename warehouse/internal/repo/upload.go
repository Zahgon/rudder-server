package repo

import (
	"context"
	"database/sql"
	"time"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var syncStatusMap = map[string]string{
	"success": model.ExportedData,
	"waiting": model.Waiting,
	"aborted": model.Aborted,
	"failed":  "%failed%",
}

const (
	defaultPriority  = "100"
	uploadsTableName = warehouseutils.WarehouseUploadsTable
	uploadColumns    = `
		id,
		status,
		schema,
		namespace,
		workspace_id,
		source_id,
		destination_id,
		destination_type,
		start_staging_file_id,
		end_staging_file_id,
		start_load_file_id,
		end_load_file_id,
		error,
		metadata,
		timings->0 as firstTiming,
		timings->-1 as lastTiming,
		timings,
		COALESCE(metadata->>'priority', '` + defaultPriority + `')::int,
		first_event_at,
		last_event_at
	`
)

var (
	UploadFieldStatus          UpdateField = func(v any) UpdateKeyValue { return keyValue{"status", v} }
	UploadFieldStartLoadFileID UpdateField = func(v any) UpdateKeyValue { return keyValue{"start_load_file_id", v} }
	UploadFieldEndLoadFileID   UpdateField = func(v any) UpdateKeyValue { return keyValue{"end_load_file_id", v} }
	UploadFieldUpdatedAt       UpdateField = func(v any) UpdateKeyValue { return keyValue{"updated_at", v} }
	UploadFieldTimings         UpdateField = func(v any) UpdateKeyValue { return keyValue{"timings", v} }
	UploadFieldSchema          UpdateField = func(v any) UpdateKeyValue { return keyValue{"schema", v} }
	UploadFieldLastExecAt      UpdateField = func(v any) UpdateKeyValue { return keyValue{"last_exec_at", v} }
	UploadFieldInProgress      UpdateField = func(v any) UpdateKeyValue { return keyValue{"in_progress", v} }
	UploadFieldMetadata        UpdateField = func(v any) UpdateKeyValue { return keyValue{"metadata", v} }
	UploadFieldError           UpdateField = func(v any) UpdateKeyValue { return keyValue{"error", v} }
	UploadFieldErrorCategory   UpdateField = func(v any) UpdateKeyValue { return keyValue{"error_category", v} }
)

type Uploads repo

type scanFn func(dest ...any) error

type ProcessOptions struct {
	SkipIdentifiers                   []string
	SkipWorkspaces                    []string
	AllowMultipleSourcesForJobsPickup bool
}

type UploadMetadata struct {
	UseRudderStorage bool      `json:"use_rudder_storage"`
	SourceTaskRunID  string    `json:"source_task_run_id"`
	SourceJobID      string    `json:"source_job_id"`
	SourceJobRunID   string    `json:"source_job_run_id"`
	LoadFileType     string    `json:"load_file_type"`
	Retried          bool      `json:"retried"`
	Priority         int       `json:"priority"`
	NextRetryTime    time.Time `json:"nextRetryTime"`
}

func NewUploads(db *sqlmiddleware.DB, opts ...Opt) *Uploads { _ = "STUB: not implemented"; return nil }

func ExtractUploadMetadata(upload model.Upload) UploadMetadata {
	_ = "STUB: not implemented"
	return *new(UploadMetadata)
}

func (u *Uploads) CreateWithStagingFiles(ctx context.Context, upload model.Upload, files []*model.StagingFile) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type FilterBy struct {
	Key       string
	Value     any
	NotEquals bool
}

func (u *Uploads) Count(ctx context.Context, filters ...FilterBy) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (u *Uploads) Get(ctx context.Context, id int64) (model.Upload, error) {
	_ = "STUB: not implemented"
	return *new(model.Upload), nil
}

func (u *Uploads) GetToProcess(ctx context.Context, destType string, limit int, opts ProcessOptions) ([]model.Upload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Uploads) UploadJobsStats(ctx context.Context, destType string, opts ProcessOptions) (model.UploadJobsStats, error) {
	_ = "STUB: not implemented"
	return *new(model.UploadJobsStats), nil
}

// UploadTimings returns the timings for an upload.
func (u *Uploads) UploadTimings(ctx context.Context, uploadID int64) (model.Timings, error) {
	_ = "STUB: not implemented"
	return *new(model.Timings), nil
}

func (u *Uploads) DeleteWaiting(ctx context.Context, uploadID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func scanUpload(scan scanFn, upload *model.Upload) error { _ = "STUB: not implemented"; return nil }

// PendingTableUploads returns a list of pending table uploads for a given upload.
// Filtering conditions neeeds to be in sync with GetToProcess partitioning and pickup condition.
func (u *Uploads) PendingTableUploads(ctx context.Context, destID, namespace string, priority int, firstEventAt time.Time, uploadID int64) ([]model.PendingTableUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Uploads) ResetInProgress(ctx context.Context, destType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) LastCreatedAt(ctx context.Context, sourceID, destinationID string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (u *Uploads) SyncsInfoForMultiTenant(ctx context.Context, limit, offset int, opts model.SyncUploadOptions) ([]model.UploadInfo, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Getting the syncs count in case, we were not able to get the count

func (u *Uploads) SyncsInfoForNonMultiTenant(ctx context.Context, limit, offset int, opts model.SyncUploadOptions) ([]model.UploadInfo, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (u *Uploads) syncsInfo(ctx context.Context, limit, offset int, opts model.SyncUploadOptions, countOver bool) ([]model.UploadInfo, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// set duration as time between updatedAt and lastExec recorded timings for ongoing/retrying uploads
// set diff between lastExec and current time

// set error only for failed uploads. skip for retried and then successful uploads

func syncUploadQueryArgs(suo *model.SyncUploadOptions) (string, []any) {
	_ = "STUB: not implemented"
	return "", nil
}

func (u *Uploads) syncsCount(ctx context.Context, opts model.SyncUploadOptions) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (u *Uploads) TriggerUpload(ctx context.Context, uploadID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) Retry(ctx context.Context, opts model.RetryOptions) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func retryQueryArgs(ro *model.RetryOptions) (string, []any) {
	_ = "STUB: not implemented"
	return "", nil
}

func (u *Uploads) RetryCount(ctx context.Context, opts model.RetryOptions) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (u *Uploads) GetLatestUploadInfo(ctx context.Context, sourceID, destinationID string) (*model.LatestUploadInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Uploads) RetrieveFailedBatches(
	ctx context.Context,
	req model.RetrieveFailedBatchesRequest,
) ([]model.RetrieveFailedBatchesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Uploads) RetryFailedBatches(
	ctx context.Context,
	req model.RetryFailedBatchesRequest,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (u *Uploads) WithTx(ctx context.Context, f func(tx *sqlmiddleware.Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) Update(ctx context.Context, id int64, fields []UpdateKeyValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) UpdateWithTx(ctx context.Context, tx *sqlmiddleware.Tx, id int64, fields []UpdateKeyValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) update(
	ctx context.Context,
	exec func(context.Context, string, ...any) (sql.Result, error),
	id int64,
	fields []UpdateKeyValue,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *Uploads) GetFirstAbortedUploadInContinuousAbortsByDestination(ctx context.Context, workspaceID string, start time.Time) ([]model.FirstAbortedUploadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Uploads) GetSyncLatencies(ctx context.Context, request model.SyncLatencyRequest) ([]model.LatencyTimeSeriesDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLatencyAggregationSQL(aggType model.LatencyAggregationType) string {
	_ = "STUB: not implemented"
	return ""
}

// Default to max
