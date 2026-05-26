package source

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type Manager struct {
	logger           logger.Logger
	sourceRepo       sourceRepo
	tableUploadsRepo tableUploadsRepo
	publisher        publisher

	config struct {
		maxBatchSizeToProcess int64
		maxAttemptsPerJob     int
	}

	trigger struct {
		processingTimeout       func() <-chan time.Time
		processingSleepInterval func() <-chan time.Time
	}
}

func New(conf *config.Config, log logger.Logger, statsFactory stats.Stats, db *sqlmw.DB, publisher publisher) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) InsertJobs(ctx context.Context, payload insertJobRequest) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There is no need to create source jobs for discards and identity resolution tables.
// Source jobs are basically used for deleting old data in case of Google Sheets for full sync and discards and identity resolution tables are de

func (m *Manager) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) process(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processPendingJobs
// 1. Prepare and publish claims to notifier
// 2. Mark source jobs as executing
// 3. Mark source jobs as failed if notifier returned timeout
// 4. Mark source jobs as failed if notifier returned error else mark as succeeded
func (m *Manager) processPendingJobs(ctx context.Context, pendingJobs []model.SourceJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) markFailed(ctx context.Context, ids []int64, failError error) error {
	_ = "STUB: not implemented"
	return nil
}

type Uploader struct{}

func (*Uploader) IsWarehouseSchemaEmpty() bool { _ = "STUB: not implemented"; return false }
func (*Uploader) UpdateLocalSchema(context.Context, model.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
func (*Uploader) GetTableSchemaInUpload(string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}
func (*Uploader) ShouldOnDedupUseNewRecord() bool { _ = "STUB: not implemented"; return false }
func (*Uploader) UseRudderStorage() bool          { _ = "STUB: not implemented"; return false }
func (*Uploader) CanAppend() bool                 { _ = "STUB: not implemented"; return false }
func (*Uploader) GetLoadFileType() string         { _ = "STUB: not implemented"; return "" }
func (*Uploader) GetLocalSchema(context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}
func (*Uploader) GetTableSchemaInWarehouse(string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}
func (*Uploader) GetSampleLoadFileLocation(context.Context, string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
func (*Uploader) GetLoadFilesMetadata(context.Context, whutils.GetLoadFilesOptions) ([]whutils.LoadFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Uploader) GetSingleLoadFile(context.Context, string) (whutils.LoadFile, error) {
	_ = "STUB: not implemented"
	return *new(whutils.LoadFile), nil
}
