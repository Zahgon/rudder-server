package schemarepository

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type Uploader interface {
	GetLocalSchema(ctx context.Context) (model.Schema, error)
	UpdateLocalSchema(ctx context.Context, schema model.Schema) error
}
type LocalSchemaRepository struct {
	warehouse model.Warehouse
	uploader  Uploader

	// mu protects the read-modify-write pattern for schema operations
	// Operations like CreateTable, AddColumns, AlterColumn:
	// 1. Read schema from PostgreSQL via GetLocalSchema
	// 2. Modify the schema in memory
	// 3. Write back to PostgreSQL via UpdateLocalSchema
	// This prevents race conditions when multiple goroutines modify the schema concurrently
	mu sync.RWMutex
}

func NewLocalSchemaRepository(warehouse model.Warehouse, uploader Uploader) (*LocalSchemaRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSchemaRepository) FetchSchema(ctx context.Context, _ model.Warehouse) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (*LocalSchemaRepository) CreateSchema(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ls *LocalSchemaRepository) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

// fetch schema from local db

// check if table already exists

// add table to schema

// update schema

func (ls *LocalSchemaRepository) AddColumns(ctx context.Context, tableName string, columnsInfo []whutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// fetch schema from local db

// check if table exists

// update schema

func (ls *LocalSchemaRepository) AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// fetch schema from local db

// check if table exists

// check if column exists

// update column type

// update schema

func (*LocalSchemaRepository) RefreshPartitions(context.Context, string, []whutils.LoadFile) error {
	_ = "STUB: not implemented"
	return nil
}
