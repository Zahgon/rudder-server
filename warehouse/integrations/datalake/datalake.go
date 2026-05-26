package datalake

import (
	"context"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/client"
	schemarepository "github.com/rudderlabs/rudder-server/warehouse/integrations/datalake/schema-repository"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var errorsMappings = []model.JobError{
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`AccessDeniedException: Insufficient Lake Formation permission.*: Required Create Database on Catalog`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`AccessDeniedException: User: .* is not authorized to perform: .* on resource: .*`),
	},
}

type Datalake struct {
	SchemaRepository schemarepository.SchemaRepository
	Warehouse        model.Warehouse
	Uploader         warehouseutils.Uploader
	conf             *config.Config
	logger           logger.Logger
}

func New(conf *config.Config, log logger.Logger) *Datalake { _ = "STUB: not implemented"; return nil }

func (d *Datalake) Setup(_ context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (d *Datalake) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Datalake) DropTable(context.Context, string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

func (d *Datalake) LoadTable(_ context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Datalake) DeleteBy(context.Context, []string, warehouseutils.DeleteByParams) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) LoadUserTables(context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// return map with nil error entries for identifies and users(if any) tables
// this is so that they are marked as succeeded

func (d *Datalake) LoadIdentityMergeRulesTable(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) LoadIdentityMappingsTable(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Datalake) Cleanup(context.Context) { _ = "STUB: not implemented"; return }

func (*Datalake) IsEmpty(context.Context, model.Warehouse) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (*Datalake) TestConnection(context.Context, model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Datalake) DownloadIdentityRules(context.Context, *misc.GZipWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Datalake) Connect(context.Context, model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (*Datalake) TestLoadTable(context.Context, string, string, map[string]any, string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Datalake) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Datalake) SetConnectionTimeout(_ time.Duration) { _ = "STUB: not implemented"; return }

func (*Datalake) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }
