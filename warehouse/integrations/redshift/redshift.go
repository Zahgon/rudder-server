package redshift

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
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var errorsMappings = []model.JobError{
	{
		Type:   model.AlterColumnError,
		Format: regexp.MustCompile(`pq: cannot alter type of a column used by a view or rule`),
	},
	{
		Type:   model.InsufficientResourceError,
		Format: regexp.MustCompile(`pq: Disk Full`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`redshift set query_group error : EOF`),
	},
	{
		Type:   model.ConcurrentQueriesError,
		Format: regexp.MustCompile(`pq: 1023`),
	},
	{
		Type:   model.ColumnSizeError,
		Format: regexp.MustCompile(`pq: Value too long for character type`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`pq: permission denied for database`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`pq: must be owner of relation`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`pq: Cannot execute write query because system is in resize mode`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`pq: SSL is not enabled on the server`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`Bucket .* not found`),
	},
	{
		Type:   model.ColumnCountError,
		Format: regexp.MustCompile(`pq: tables can have at most 1600 columns`),
	},
}

const (
	rudderStringLength = 512
	provider           = warehouseutils.RS
	tableNameLimit     = 127
)

var dataTypesMap = map[string]string{
	"boolean":  "boolean encode runlength",
	"int":      "bigint",
	"bigint":   "bigint",
	"float":    "double precision",
	"string":   "varchar(65535)",
	"text":     "varchar(65535)",
	"datetime": "timestamp",
	"json":     "super",
}

var dataTypesMapToRudder = map[string]string{
	"int":                         "int",
	"int2":                        "int",
	"int4":                        "int",
	"int8":                        "int",
	"bigint":                      "int",
	"float":                       "float",
	"float4":                      "float",
	"float8":                      "float",
	"numeric":                     "float",
	"double precision":            "float",
	"boolean":                     "boolean",
	"bool":                        "boolean",
	"text":                        "string",
	"character varying":           "string",
	"nchar":                       "string",
	"bpchar":                      "string",
	"character":                   "string",
	"nvarchar":                    "string",
	"string":                      "string",
	"date":                        "datetime",
	"timestamp without time zone": "datetime",
	"timestamp with time zone":    "datetime",
	"super":                       "json",
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

type Redshift struct {
	DB             *sqlmiddleware.DB
	Namespace      string
	Warehouse      model.Warehouse
	Uploader       warehouseutils.Uploader
	connectTimeout time.Duration
	conf           *config.Config
	logger         logger.Logger
	stats          stats.Stats

	config struct {
		allowMerge                    bool
		slowQueryThreshold            time.Duration
		dedupWindow                   bool
		dedupWindowInHours            time.Duration
		skipDedupDestinationIDs       []string
		skipComputingUserLatestTraits bool
		enableDeleteByJobs            bool
		loadByFolderPath              bool
	}
}

type s3ManifestEntryMetadata struct {
	ContentLength int64 `json:"content_length"`
}

type s3ManifestEntry struct {
	Url       string                  `json:"url"`
	Mandatory bool                    `json:"mandatory"`
	Metadata  s3ManifestEntryMetadata `json:"meta"`
}

type s3Manifest struct {
	Entries []s3ManifestEntry `json:"entries"`
}

func New(conf *config.Config, log logger.Logger, stat stats.Stats) *Redshift {
	_ = "STUB: not implemented"
	return nil
}

// getRSDataType gets datatype for rs which is mapped with RudderStack datatype
func getRSDataType(columnType string) string { _ = "STUB: not implemented"; return "" }

func ColumnsWithDataTypes(columns model.TableSchema, prefix string) string {
	_ = "STUB: not implemented"
	// TODO: do we need sorted order here?
	return ""
}

func (rs *Redshift) CreateTable(ctx context.Context, tableName string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) DropTable(ctx context.Context, tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) schemaExists(ctx context.Context) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rs *Redshift) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckAndIgnoreColumnAlreadyExistError(err error) bool { _ = "STUB: not implemented"; return false }

func (rs *Redshift) DeleteBy(ctx context.Context, tableNames []string, params warehouseutils.DeleteByParams) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) createSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) generateManifest(ctx context.Context, tableName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// add contentLength to manifest entry if it exists

func (rs *Redshift) dropStagingTables(ctx context.Context, stagingTableNames []string) {
	_ = "STUB: not implemented"
	return
}

func (rs *Redshift) createStagingTable(ctx context.Context, sourceTableName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rs *Redshift) loadTable(
	ctx context.Context,
	tableName string,
	tableSchemaInUpload,
	tableSchemaAfterUpload model.TableSchema,
	skipTempTableDelete bool,
) (*types.LoadTableStats, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Users table still needs to be deduped by partition key.
// In case of users table, staging table should be created, so that users table can be deduped by partition key

type sqlExecer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func (rs *Redshift) copyIntoLoadTable(
	ctx context.Context,
	sqlExecer sqlExecer,
	tableName string,
	stagingTableName string,
	strKeys []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) deleteFromLoadTable(
	ctx context.Context,
	txn *sqlmiddleware.Tx,
	tableName string,
	stagingTableName string,
	tableSchemaAfterUpload model.TableSchema,
) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (rs *Redshift) insertIntoLoadTable(
	ctx context.Context,
	txn *sqlmiddleware.Tx,
	tableName string,
	stagingTableName string,
	sortedColumnKeys []string,
) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (rs *Redshift) loadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// do not reference uuid in queries as it can be an autoincrement field set by segment compatible tables

func (rs *Redshift) connect(ctx context.Context) (*sqlmiddleware.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *Redshift) useIAMForAuth() bool { _ = "STUB: not implemented"; return false }

func (rs *Redshift) connectUsingIAMRole() (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *Redshift) connectUsingPassword() (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *Redshift) dropDanglingStagingTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) AlterColumn(ctx context.Context, tableName, columnName, columnType string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// Begin a transaction

// creating staging column

// populating staging column

// renaming original column to deprecated column

// renaming staging column to original column

// Commit the transaction

// dropping deprecated column
// Since dropping the column can fail, we need to do it outside the transaction
// Because if it will fail during the commit of the transaction
// https://github.com/lib/pq/blob/d5affd5073b06f745459768de35356df2e5fd91d/conn.go#L600

// FetchSchema queries redshift and returns the schema associated with provided namespace
func (rs *Redshift) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func calculateDataType(columnType string, charLength sql.NullInt64) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (rs *Redshift) Setup(ctx context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) TestConnection(ctx context.Context, _ model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) Cleanup(ctx context.Context) { _ = "STUB: not implemented"; return }

func (*Redshift) IsEmpty(context.Context, model.Warehouse) (empty bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rs *Redshift) LoadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Redshift) LoadIdentityMergeRulesTable(context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Redshift) LoadIdentityMappingsTable(context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Redshift) DownloadIdentityRules(context.Context, *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) Connect(ctx context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (rs *Redshift) TestLoadTable(ctx context.Context, location, tableName string, _ map[string]any, format string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// copy statement for parquet load files

// copy statement for csv load files

func (rs *Redshift) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *Redshift) SetConnectionTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (rs *Redshift) ShouldMerge(tableName string) bool { _ = "STUB: not implemented"; return false }

// If we are here it's because skipComputingUserLatestTraits is true.
// preferAppend doesn't apply to the users table, so we are just checking skipDedupDestinationIDs for
// backwards compatibility.

// It's important to check the ability to append after skipDedup to make sure that if both
// skipDedupDestinationIDs and skipComputingUserLatestTraits are set, we still merge.
// see hyperverge user table use case for more details.

func (*Redshift) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }

func normalizeError(err error) error { _ = "STUB: not implemented"; return nil }
