package schemarepository

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

// glue specific config
var (
	glueSerdeName             = "ParquetHiveSerDe"
	glueSerdeSerializationLib = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
	glueParquetInputFormat    = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
	glueParquetOutputFormat   = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"
)

var (
	partitionFolderRegex = regexp.MustCompile(`.*/(?P<name>.*)=(?P<value>.*)$`)
	partitionWindowRegex = regexp.MustCompile(`^(?P<name>.*)=(?P<value>.*)$`)
)

type GlueSchemaRepository struct {
	GlueClient *glue.Client
	conf       *config.Config
	logger     logger.Logger
	s3bucket   string
	s3prefix   string
	Warehouse  model.Warehouse
	Namespace  string
}

func NewGlueSchemaRepository(conf *config.Config, logger logger.Logger, wh model.Warehouse) (*GlueSchemaRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GlueSchemaRepository) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlueSchemaRepository) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// create table request

// add storage descriptor to create table request

func (g *GlueSchemaRepository) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlueSchemaRepository) AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

func (g *GlueSchemaRepository) RefreshPartitions(ctx context.Context, tableName string, loadFiles []warehouseutils.LoadFile) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip if time window layout is not defined

// Skip if we are already going to process this locationFolder

// Check for existing partitions. We do not want to generate unnecessary (for already existing
// partitions) changes in Glue tables (since the number of versions of a Glue table
// is limited)

// Updating table partitions with empty columns to create partition keys if not created

func (g *GlueSchemaRepository) FetchSchema(ctx context.Context, warehouse model.Warehouse) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

// add nextToken to the request if there are multiple list segments

// break out of the loop if there are no more list segments

func (g *GlueSchemaRepository) updateTable(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// get table schema

// add new columns to table schema

// add storage descriptor to update table request

func (g *GlueSchemaRepository) partitionColumns() (columns []types.Column, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GlueSchemaRepository) getStorageDescriptor(tableName string, columnMap model.TableSchema) *types.StorageDescriptor {
	_ = "STUB: not implemented"
	return nil
}

// add columns to storage descriptor

func (g *GlueSchemaRepository) getS3LocationForTable(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}
