//go:generate mockgen -destination=mocks/upload.go -package=mocks -source=upload.go loadFilesRepo,stagingFilesRepo
package router

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/alerta"
	"github.com/rudderlabs/rudder-server/utils/types"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/manager"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/loadfiles"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/repo"
	"github.com/rudderlabs/rudder-server/warehouse/schema"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
	"github.com/rudderlabs/rudder-server/warehouse/validations"
)

const (
	GeneratingStagingFileFailedState = "generating_staging_file_failed"
	GeneratedStagingFileState        = "generated_staging_file"
	InternalProcessingFailed         = "internal_processing_failed"
)

const (
	cloudSourceCategory          = "cloud"
	singerProtocolSourceCategory = "singer-protocol"
)

type tableNameT string

type UploadJobFactory struct {
	reporting            types.Reporting
	db                   *sqlquerywrapper.DB
	destinationValidator validations.DestinationValidator
	loadFile             *loadfiles.LoadFileGenerator
	conf                 *config.Config
	logger               logger.Logger
	statsFactory         stats.Stats
	encodingFactory      *encoding.Factory
}

type loadFilesRepo interface {
	Get(ctx context.Context, uploadID int64) ([]model.LoadFile, error)
	Delete(ctx context.Context, uploadID int64) error
	TotalExportedEvents(ctx context.Context, uploadID int64, skipTables []string) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.LoadFile, error)
	DistinctTableName(ctx context.Context, sourceID, destinationID string, startID, endID int64) ([]string, error)
}

type stagingFilesRepo interface {
	TotalEventsForUploadID(ctx context.Context, uploadID int64) (int64, error)
	GetEventTimeRangesByUploadID(ctx context.Context, uploadID int64) ([]model.EventTimeRange, error)
}

type UploadJob struct {
	ctx                  context.Context
	db                   *sqlquerywrapper.DB
	reporting            types.Reporting
	destinationValidator validations.DestinationValidator
	loadfile             *loadfiles.LoadFileGenerator
	tableUploadsRepo     *repo.TableUploads
	uploadsRepo          *repo.Uploads
	stagingFileRepo      stagingFilesRepo
	loadFilesRepo        loadFilesRepo
	whSchemaRepo         *repo.WHSchema
	whManager            manager.Manager
	schemaHandle         schema.Handler
	conf                 *config.Config
	logger               logger.Logger
	statsFactory         stats.Stats

	upload       model.Upload
	warehouse    model.Warehouse
	stagingFiles []*model.StagingFile
	alertSender  alerta.AlertSender
	now          func() time.Time

	pendingTableUploads      []model.PendingTableUpload
	pendingTableUploadsRepo  pendingTableUploadsRepo
	pendingTableUploadsOnce  sync.Once
	pendingTableUploadsError error

	config struct {
		refreshPartitionBatchSize           int
		retryTimeWindow                     time.Duration
		minRetryAttempts                    int
		disableAlter                        bool
		minUploadBackoff                    time.Duration
		maxUploadBackoff                    time.Duration
		reportingEnabled                    bool
		maxParallelLoadsWorkspaceIDs        map[string]any
		columnsBatchSize                    int
		longRunningUploadStatThresholdInMin time.Duration
		skipPreviouslyFailedTables          bool
		// max number of parallel delete requests to filemanager (applies to GCS only)
		maxConcurrentObjDeleteRequests func(workspaceID string) int
		// batch size for parallel deletion of staging and loadfiles (applies to GCS only)
		objDeleteBatchSize func(workspaceID string) int
	}

	errorHandler       ErrorHandler
	encodingFactory    *encoding.Factory
	fileManagerFactory filemanager.Factory

	stats struct {
		uploadTime                         stats.Timer
		userTablesLoadTime                 stats.Timer
		identityTablesLoadTime             stats.Timer
		otherTablesLoadTime                stats.Timer
		loadFileGenerationTime             stats.Timer
		uploadFailed                       stats.Counter
		totalRowsSynced                    stats.Counter
		numStagedEvents                    stats.Counter
		uploadSuccess                      stats.Counter
		stagingLoadFileEventsCountMismatch stats.Gauge
		eventDeliveryTime                  stats.Timer
		objectsDeleted                     stats.Gauge
		objectsDeletionTime                stats.Timer
		consolidatedSchemaSize             stats.Histogram
	}
}

type pendingTableUploadsRepo interface {
	PendingTableUploads(ctx context.Context, destID, namespace string, priority int, firstEventAt time.Time, uploadID int64) ([]model.PendingTableUpload, error)
}

var (
	alwaysMarkExported     = []string{whutils.DiscardsTable}
	mergeSourceCategoryMap = map[string]struct{}{
		"cloud":           {},
		"singer-protocol": {},
	}
)

func (f *UploadJobFactory) NewUploadJob(ctx context.Context, dto *model.UploadJob, whManager manager.Manager) *UploadJob {
	_ = "STUB: not implemented"
	return nil
}

// 24h

func (job *UploadJob) trackLongRunningUpload() chan struct{} { _ = "STUB: not implemented"; return nil }

// do nothing

func (job *UploadJob) run() (err error) { _ = "STUB: not implemented"; return nil }

// schema is being checked only for error in ExportedData and not in other cases
// to prevent unnecessary calls to warehouse

// This sets the schema expiry to now, so the next attempt will fetch the latest schema

// If unknown state, start again

// record metric for time taken by the current state

func (job *UploadJob) cleanupObjectStorageFiles() error { _ = "STUB: not implemented"; return nil }

// GCS doesn't support batch delete, so we need to delete files in chunks to speed up the deletion

// CanAppend returns true if:
// * the source is not an ETL source
// * the source is not a replay source
// * the source category is not in "mergeSourceCategoryMap"
// * the job is not a retry
func (job *UploadJob) CanAppend() bool { _ = "STUB: not implemented"; return false }

// getNewTimings appends current status with current time to timings column
// e.g. status: exported_data, timings: [{exporting_data: 2020-04-21 15:16:19.687716}] -> [{exporting_data: 2020-04-21 15:16:19.687716, exported_data: 2020-04-21 15:26:34.344356}]
func (job *UploadJob) getNewTimings(status string) ([]byte, model.Timings, error) {
	_ = "STUB: not implemented"
	return nil, *new(model.Timings), nil
}

func (job *UploadJob) getUploadFirstAttemptTime() (timing time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

type UploadStatusOpts struct {
	Status          string
	ReportingMetric types.PUReportedMetric
}

func (job *UploadJob) setUploadStatus(statusOpts UploadStatusOpts) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: fetch upload model instead of just timings

// extractAndUpdateUploadErrorsByState extracts and augment errors in format
// { "internal_processing_failed": { "errors": ["account-locked", "account-locked"] }}
// from a particular upload.
func extractAndUpdateUploadErrorsByState(message json.RawMessage, state string, statusError error) (map[string]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// increment attempts for errored stage

// append errors for errored stage

// Aborted returns true if the job has been aborted
func (job *UploadJob) Aborted(attempts int, startTime time.Time) bool {
	_ = "STUB: not implemented"
	// Defensive check to prevent garbage startTime
	return false
}

func (job *UploadJob) setUploadError(statusError error, state string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Reset the state as aborted if max retries
// exceeded.

// TODO: Change this to error specific code

// TODO: Change this to error specific code

// On aborted state, validate credentials to allow
// us to differentiate between user caused abort vs platform issue.

// base tag to be sent as stat

func (job *UploadJob) durationBeforeNextAttempt(attempt int64) time.Duration {
	_ = "STUB: not implemented" // Add state(retryable/non-retryable) as an argument to decide backoff etc.
	return *new(time.Duration)
}

func (job *UploadJob) validateDestinationCredentials() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (job *UploadJob) GetLoadFilesMetadata(ctx context.Context, options whutils.GetLoadFilesOptions) (loadFiles []whutils.LoadFile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (job *UploadJob) getLoadFilesMetadataQuery(tableFilterSQL, limitSQL string) string {
	_ = "STUB: not implemented"
	return ""
}

func (job *UploadJob) GetSampleLoadFileLocation(ctx context.Context, tableName string) (location string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (job *UploadJob) IsWarehouseSchemaEmpty() bool { _ = "STUB: not implemented"; return false }

func (job *UploadJob) GetTableSchemaInWarehouse(tableName string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}

func (job *UploadJob) GetTableSchemaInUpload(tableName string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}

func (job *UploadJob) GetSingleLoadFile(ctx context.Context, tableName string) (whutils.LoadFile, error) {
	_ = "STUB: not implemented"
	return *new(whutils.LoadFile), nil
}

func (job *UploadJob) ShouldOnDedupUseNewRecord() bool { _ = "STUB: not implemented"; return false }

func (job *UploadJob) UseRudderStorage() bool { _ = "STUB: not implemented"; return false }

func (job *UploadJob) GetLoadFileType() string { _ = "STUB: not implemented"; return "" }

func (job *UploadJob) DTO() *model.UploadJob { _ = "STUB: not implemented"; return nil }

func (job *UploadJob) GetLocalSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (job *UploadJob) UpdateLocalSchema(ctx context.Context, schema model.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
