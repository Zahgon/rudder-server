package loadfiles

import (
	"context"
	stdjson "encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	defaultPublishBatchSize = 100
)

var warehousesToVerifyLoadFilesFolder = []string{warehouseutils.SNOWFLAKE}

type Notifier interface {
	Publish(ctx context.Context, payload *notifier.PublishRequest) (ch <-chan *notifier.PublishResponse, err error)
}

type StageFileRepo interface {
	SetStatuses(ctx context.Context, ids []int64, status string) (err error)
}

type LoadFileRepo interface {
	Insert(ctx context.Context, loadFiles []model.LoadFile) error
	Delete(ctx context.Context, uploadID int64) error
	Get(ctx context.Context, uploadID int64) ([]model.LoadFile, error)
}

type ControlPlaneClient interface {
	DestinationHistory(ctx context.Context, revisionID string) (backendconfig.DestinationT, error)
}

type LoadFileGenerator struct {
	Conf     *config.Config
	Logger   logger.Logger
	Notifier Notifier

	StageRepo StageFileRepo
	LoadRepo  LoadFileRepo

	ControlPlaneClient ControlPlaneClient

	publishBatchSize             int
	publishBatchSizePerWorkspace map[string]int
}

type WorkerJobResponseV2 struct {
	Output []LoadFileUpload `json:"output"`
}

type LoadFileUpload struct {
	TableName             string
	Location              string
	TotalRows             int
	ContentLength         int64
	DestinationRevisionID string
	UseRudderStorage      bool
}

// baseWorkerJobRequest contains common fields for both v1 and v2 job requests
type baseWorkerJobRequest struct {
	BatchID                      string         `json:"batch_id"`
	UploadID                     int64          `json:"upload_id"`
	WorkspaceID                  string         `json:"workspace_id"`
	SourceID                     string         `json:"source_id"`
	SourceName                   string         `json:"source_name"`
	DestinationID                string         `json:"destination_id"`
	DestinationName              string         `json:"destination_name"`
	DestinationType              string         `json:"destination_type"`
	DestinationNamespace         string         `json:"destination_namespace"`
	DestinationRevisionID        string         `json:"destination_revision_id"`
	StagingDestinationRevisionID string         `json:"staging_destination_revision_id"`
	DestinationConfig            map[string]any `json:"destination_config"`
	StagingDestinationConfig     any            `json:"staging_destination_config"`
	UseRudderStorage             bool           `json:"use_rudder_storage"`
	StagingUseRudderStorage      bool           `json:"staging_use_rudder_storage"`
	UniqueLoadGenID              string         `json:"unique_load_gen_id"`
	RudderStoragePrefix          string         `json:"rudder_storage_prefix"`
	LoadFilePrefix               string         `json:"load_file_prefix"` // prefix for the load file name
	LoadFileType                 string         `json:"load_file_type"`
}

type StagingFileInfo struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
}

type WorkerJobRequestV2 struct {
	baseWorkerJobRequest
	StagingFiles []StagingFileInfo `json:"staging_files"`
}

func WithConfig(ld *LoadFileGenerator, config *config.Config) { _ = "STUB: not implemented"; return }

func (lf *LoadFileGenerator) CreateLoadFiles(ctx context.Context, job *model.UploadJob) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Delete previous load files for the upload

// Set staging file status to executing

// ensure that if there is an error, we set the staging file status to failed

// Always use V2 job creation (no more V1 vs V2 branching)

// Assign it to "err" to ensure that the defer function is called for the error

func (lf *LoadFileGenerator) getLoadFileIDs(ctx context.Context, job *model.UploadJob, uniqueLoadGenID string) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// verify if all load files are in same folder in object storage

func (lf *LoadFileGenerator) prepareJobRequestV2(
	job *model.UploadJob,
	uniqueLoadGenID string,
	stagingFiles []*model.StagingFile,
	destinationRevisionIDMap map[string]backendconfig.DestinationT,
) *WorkerJobRequestV2 {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LoadFileGenerator) publishToNotifier(
	ctx context.Context,
	job *model.UploadJob,
	messages []stdjson.RawMessage,
	jobType notifier.JobType,
) (<-chan *notifier.PublishResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lf *LoadFileGenerator) processNotifierResponseV2(ctx context.Context, ch <-chan *notifier.PublishResponse, job *model.UploadJob, chunk []*model.StagingFile) error {
	_ = "STUB: not implemented"
	return nil
}

func (lf *LoadFileGenerator) createUploadV2Jobs(ctx context.Context, job *model.UploadJob, stagingFiles []*model.StagingFile, publishBatchSize int, uniqueLoadGenID string) error {
	_ = "STUB: not implemented"
	return nil
}

// capture for goroutine

func (lf *LoadFileGenerator) destinationRevisionIDMap(ctx context.Context, job *model.UploadJob) (map[string]backendconfig.DestinationT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No need to make config backend api call for the current config

// No need to make config backend api call for the same revision ID

func (lf *LoadFileGenerator) GetLoadFilePrefix(timeWindow time.Time, warehouse model.Warehouse) string {
	_ = "STUB: not implemented"
	return ""
}

func toLoadFile(output LoadFileUpload, job *model.UploadJob) model.LoadFile {
	_ = "STUB: not implemented"
	return *new(model.LoadFile)
}

// The fields in this struct are determined by the fields being used in the prepareBaseJobRequest
type stagingFileGroupKey struct {
	UseRudderStorage             bool
	StagingUseRudderStorage      bool
	DestinationRevisionID        string
	StagingDestinationRevisionID string
	TimeWindow                   time.Time
}

// GroupStagingFiles groups staging files based on their key characteristics
// and then applies size constraints within each group. The maxSizeMB parameter controls the maximum size of any table within a group.
func (lf *LoadFileGenerator) GroupStagingFiles(files []*model.StagingFile, maxSizeMB int) [][]*model.StagingFile {
	_ = "STUB: not implemented"
	return nil
}

// For each group, apply size constraints

type tableSizeResult struct {
	sizes map[string]int64
	name  string
	size  int64
}

// groupBySize splits a group of staging files based on size constraints
func (lf *LoadFileGenerator) groupBySize(files []*model.StagingFile, maxSizeMB int) [][]*model.StagingFile {
	_ = "STUB: not implemented"
	return nil
}

// Convert MB to bytes

// Find the table with the maximum total size

// Sorting ensures that minimum batches are created

// Assuming that there won't be any overflows

// Start a new batch

// Try to add files to the current batch

// Check if adding this file would exceed size limit for any table

// Add file to batch and update table sizes

// If this is the first file in this iteration and it exceeds limits,
// add it to its own batch

// This condition will be false if a file was already added to the batch because it exceeded the size limit
// In that case, it should not be added again to result
