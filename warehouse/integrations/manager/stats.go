package manager

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type statsManager struct {
	Manager

	statsFactory stats.Stats
	stats        struct {
		addColumnsCount   stats.Counter
		addTablesCount    stats.Counter
		alterColumnsCount stats.Counter
	}
}

func newStatsManager(manager Manager, statsFactory stats.Stats) *statsManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *statsManager) Setup(ctx context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *statsManager) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *statsManager) AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

func (m *statsManager) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}
