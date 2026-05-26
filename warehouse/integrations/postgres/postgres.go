package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/client"
	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/tunnelling"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/service/loadfiles/downloader"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	verifyCA = "verify-ca"
)

const (
	provider       = warehouseutils.POSTGRES
	tableNameLimit = 127
)

var errorsMappings = []model.JobError{
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`dial tcp: lookup .*: no such host`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`dial tcp .* connect: connection refused`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: database .* does not exist`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: the database system is starting up`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: the database system is shutting down`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: relation .* does not exist`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: cannot set transaction read-write mode during recovery`),
	},
	{
		Type:   model.ColumnCountError,
		Format: regexp.MustCompile(`pq: tables can have at most 1600 columns`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`pq: password authentication failed for user`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`pq: permission denied`),
	},
}

var rudderDataTypesMapToPostgres = map[string]string{
	"int":      "bigint",
	"float":    "numeric",
	"string":   "text",
	"datetime": "timestamptz",
	"boolean":  "boolean",
	"json":     "jsonb",
}

var postgresDataTypesMapToRudder = map[string]string{
	"integer":                  "int",
	"smallint":                 "int",
	"bigint":                   "int",
	"double precision":         "float",
	"numeric":                  "float",
	"real":                     "float",
	"text":                     "string",
	"varchar":                  "string",
	"char":                     "string",
	"timestamptz":              "datetime",
	"timestamp with time zone": "datetime",
	"timestamp":                "datetime",
	"boolean":                  "boolean",
	"jsonb":                    "json",
}

type Postgres struct {
	DB                 *sqlmiddleware.DB
	Namespace          string
	ObjectStorage      string
	Warehouse          model.Warehouse
	Uploader           warehouseutils.Uploader
	connectTimeout     time.Duration
	conf               *config.Config
	logger             logger.Logger
	stats              stats.Stats
	LoadFileDownloader downloader.Downloader

	config struct {
		allowMerge                                bool
		enableDeleteByJobs                        bool
		numWorkersDownloadLoadFiles               int
		slowQueryThreshold                        time.Duration
		txnRollbackTimeout                        time.Duration
		skipDedupDestinationIDs                   []string
		skipComputingUserLatestTraits             bool
		skipComputingUserLatestTraitsWorkspaceIDs []string
	}
}

type credentials struct {
	host       string
	database   string
	user       string
	password   string
	port       string
	sslMode    string
	sslDir     string
	tunnelInfo *tunnelling.TunnelInfo
	timeout    time.Duration
}

var primaryKeyMap = map[string]string{
	warehouseutils.UsersTable:      "id",
	warehouseutils.IdentifiesTable: "id",
	warehouseutils.DiscardsTable:   "row_id",
}

var partitionKeyMap = map[string]string{
	warehouseutils.UsersTable:      "id",
	warehouseutils.IdentifiesTable: "id",
	warehouseutils.DiscardsTable:   "row_id, column_name, table_name",
}

func New(conf *config.Config, log logger.Logger, stat stats.Stats) *Postgres {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) getNewMiddleWare(db *sql.DB) *sqlmiddleware.DB {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) connect() (*sqlmiddleware.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pg *Postgres) getConnectionCredentials() credentials {
	_ = "STUB: not implemented"
	return *new(credentials)
}

func ColumnsWithDataTypes(columns model.TableSchema, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*Postgres) IsEmpty(context.Context, model.Warehouse) (empty bool, err error) {
	_ = "STUB: not implemented"

	// DeleteBy Need to create a structure with delete parameters instead of simply adding a long list of params
	return false, nil
}

func (pg *Postgres) DeleteBy(ctx context.Context, tableNames []string, params warehouseutils.DeleteByParams) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) schemaExists(ctx context.Context, _ string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pg *Postgres) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) createTable(ctx context.Context, name string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	// set the schema in search path. so that we can query table with unqualified name which is just the table name rather than using schema.table in queries
	return nil
}

func (pg *Postgres) DropTable(ctx context.Context, tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// set the schema in search path. so that we can query table with unqualified name which is just the table name rather than using schema.table in queries

func (*Postgres) AlterColumn(context.Context, string, string, string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

func (pg *Postgres) TestConnection(ctx context.Context, warehouse model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) Setup(_ context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FetchSchema queries postgres and returns the schema associated with provided namespace
func (pg *Postgres) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (pg *Postgres) Cleanup(context.Context) { _ = "STUB: not implemented"; return }

func (*Postgres) LoadIdentityMergeRulesTable(context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Postgres) LoadIdentityMappingsTable(context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Postgres) DownloadIdentityRules(context.Context, *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) Connect(_ context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (pg *Postgres) TestLoadTable(ctx context.Context, _, tableName string, payloadMap map[string]any, _ string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) SetConnectionTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (*Postgres) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }
