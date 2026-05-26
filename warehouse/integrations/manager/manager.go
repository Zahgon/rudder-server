package manager

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/client"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type Manager interface {
	Setup(ctx context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) error
	FetchSchema(ctx context.Context) (model.Schema, error)
	CreateSchema(ctx context.Context) (err error)
	CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error)
	AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error)
	AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error)
	LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error)
	LoadUserTables(ctx context.Context) map[string]error
	LoadIdentityMergeRulesTable(ctx context.Context) error
	LoadIdentityMappingsTable(ctx context.Context) error
	Cleanup(ctx context.Context)
	IsEmpty(ctx context.Context, warehouse model.Warehouse) (bool, error)
	DownloadIdentityRules(ctx context.Context, gzWriter *misc.GZipWriter) error
	Connect(ctx context.Context, warehouse model.Warehouse) (client.Client, error)
	SetConnectionTimeout(timeout time.Duration)
	ErrorMappings() []model.JobError

	TestConnection(ctx context.Context, warehouse model.Warehouse) error
	TestFetchSchema(ctx context.Context) error
	TestLoadTable(ctx context.Context, location, stagingTableName string, payloadMap map[string]any, loadFileFormat string) error
}

type WarehouseDelete interface {
	DropTable(ctx context.Context, tableName string) (err error)
	DeleteBy(ctx context.Context, tableName []string, params warehouseutils.DeleteByParams) error
}

type WarehouseOperations interface {
	Manager
	WarehouseDelete
}

// New is a Factory function that returns a Manager of a given destination-type
func New(destType string, conf *config.Config, logger logger.Logger, stats stats.Stats) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

func newManager(destType string, conf *config.Config, logger logger.Logger, stats stats.Stats) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

// NewWarehouseOperations is a Factory function that returns a WarehouseOperations of a given destination-type
func NewWarehouseOperations(destType string, conf *config.Config, logger logger.Logger, stats stats.Stats) (WarehouseOperations, error) {
	_ = "STUB: not implemented"
	return *new(WarehouseOperations), nil
}
