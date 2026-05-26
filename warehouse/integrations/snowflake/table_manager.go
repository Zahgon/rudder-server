package snowflake

import (
	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type tableManager interface {
	createTableQuery(schemaIdentifier, tableName string, columns model.TableSchema) string
	addColumnsQuery(schemaIdentifier, tableName string, columnsInfo []whutils.ColumnInfo) (string, error)
}

func newTableManager(config *config.Config, warehouse model.Warehouse) tableManager {
	_ = "STUB: not implemented"
	return *new(tableManager)
}

// for standard Snowflake tables
type standardTableManager struct {
	dataTypesMap map[string]string
}

func newStandardTableManager() tableManager { _ = "STUB: not implemented"; return *new(tableManager) }

func (m *standardTableManager) createTableQuery(schemaIdentifier, tableName string, columns model.TableSchema) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *standardTableManager) addColumnsQuery(schemaIdentifier, tableName string, columnsInfo []whutils.ColumnInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// for snowflake managed Iceberg tables
type icebergTableManager struct {
	dataTypesMap   map[string]string
	externalVolume string
}

func newIcebergTableManager(externalVolume string) tableManager {
	_ = "STUB: not implemented"
	return *new(tableManager)
}

func (m *icebergTableManager) createTableQuery(schemaIdentifier, tableName string, columns model.TableSchema) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *icebergTableManager) addColumnsQuery(schemaIdentifier, tableName string, columnsInfo []whutils.ColumnInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func columnsWithDataTypes(columns model.TableSchema, dataTypesMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}
