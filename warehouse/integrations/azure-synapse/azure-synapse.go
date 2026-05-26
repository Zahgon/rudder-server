package azuresynapse

import (
	"context"
	"database/sql"
	"time"

	"github.com/samber/lo"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/client"
	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/service/loadfiles/downloader"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	varcharDefaultLength = 512
	varcharMaxLength     = -1
	provider             = warehouseutils.AzureSynapse
	tableNameLimit       = 127
)

var errorsMappings []model.JobError

var rudderDataTypesMapToAzureSynapse = map[string]string{
	"int":      "bigint",
	"float":    "decimal(28,10)",
	"string":   "varchar(512)",
	"datetime": "datetimeoffset",
	"boolean":  "bit",
	"json":     "jsonb",
}

var azureSynapseDataTypesMapToRudder = map[string]string{
	"integer":                  "int",
	"smallint":                 "int",
	"bigint":                   "int",
	"tinyint":                  "int",
	"double precision":         "float",
	"numeric":                  "float",
	"decimal":                  "float",
	"real":                     "float",
	"float":                    "float",
	"text":                     "string",
	"varchar":                  "string",
	"nvarchar":                 "string",
	"ntext":                    "string",
	"nchar":                    "string",
	"char":                     "string",
	"datetimeoffset":           "datetime",
	"date":                     "datetime",
	"datetime2":                "datetime",
	"timestamp with time zone": "datetime",
	"timestamp":                "datetime",
	"jsonb":                    "json",
	"bit":                      "boolean",
}

var stringColumns = lo.Keys(lo.PickBy(azureSynapseDataTypesMapToRudder, func(_, value string) bool {
	return value == "string"
}))

type AzureSynapse struct {
	db                 *sqlmw.DB
	namespace          string
	objectStorage      string
	warehouse          model.Warehouse
	uploader           warehouseutils.Uploader
	connectTimeout     time.Duration
	loadFileDownLoader downloader.Downloader

	stats  stats.Stats
	conf   *config.Config
	logger logger.Logger

	config struct {
		numWorkersDownloadLoadFiles int
		slowQueryThreshold          time.Duration
	}
}

type credentials struct {
	host     string
	dbName   string
	user     string
	password string
	port     string
	sslMode  string
	timeout  time.Duration
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

func New(conf *config.Config, log logger.Logger, stats stats.Stats) *AzureSynapse {
	_ = "STUB: not implemented"
	return nil
}

// connect to the azure synapse database
// if TrustServerCertificate is set to true, all options(disable, false, true) works.
// if forceSSL is set to 1, disable option doesn't work.
// If forceSSL is set to true or false, it works alongside with TrustServerCertificate=true
// more about combinations in here: https://docs.microsoft.com/en-us/sql/connect/odbc/linux-mac/connection-string-keywords-and-data-source-names-dsns?view=sql-server-ver15
func (as *AzureSynapse) connect() (*sqlmw.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (as *AzureSynapse) connectionCredentials() *credentials { _ = "STUB: not implemented"; return nil }

func columnsWithDataTypes(columns model.TableSchema, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*AzureSynapse) IsEmpty(_ context.Context, _ model.Warehouse) (empty bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (as *AzureSynapse) loadTable(
	ctx context.Context,
	tableName string,
	tableSchemaInUpload model.TableSchema,
	skipTempTableDelete bool,
) (*types.LoadTableStats, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// The use of prepared statements for creating temporary tables is not suitable in this context.
// Temporary tables in SQL Server have a limited scope and are automatically purged after the transaction commits.
// Therefore, creating normal tables is chosen as an alternative.
//
// For more information on this behavior:
// - See the discussion at https://github.com/denisenkom/go-mssqldb/issues/149 regarding prepared statements.
// - Refer to Microsoft's documentation on temporary tables at
//   https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008-r2/ms175528(v=sql.105)?redirectedfrom=MSDN.

// getVarcharLengthMap retrieves the maximum allowed length for varchar columns in a given table.
// A `CHARACTER_MAXIMUM_LENGTH` of `-1` indicates that the column has the maximum possible length (i.e., `varchar(max)`).
func (as *AzureSynapse) getVarcharLengthMap(ctx context.Context, tableName string) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (as *AzureSynapse) loadDataIntoStagingTable(
	ctx context.Context,
	log logger.Logger,
	stmt *sql.Stmt,
	fileName string,
	sortedColumnKeys []string,
	extraColumns []string,
	tableSchemaInUpload model.TableSchema,
	varcharLengthMap map[string]int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// To ensure the successful execution of the 'copyIn' operation in Azure Synapse,
// it is necessary to handle the scenario where new columns are added to the target table.
// Without this adjustment, attempting to perform 'copyIn' when the column count in the
// target table does not match the column count specified in the input data will result
// in an error like:
//
//   mssql: Column count in target table does not match column count specified in input.
//
// If this error is encountered, it is important to verify that the column structure in
// the source data matches the destination table's structure. If you are using the BCP command,
// ensure that the format file's column count matches the destination table. For SSIS data imports,
// double-check that the column mappings are consistent with the target table.

// ProcessColumnValue processes the input string `value` based on its specified `valueType`.
// It converts the value to the appropriate type and ensures it adheres to the constraints
// such as `varcharLength` for string types.
func ProcessColumnValue(
	value string,
	valueType string,
	varcharLength int,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// If the varchar length is set to the maximum allowed, return the string as is.

func (as *AzureSynapse) deleteFromLoadTable(
	ctx context.Context,
	txn *sqlmw.Tx,
	tableName string,
	stagingTableName string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (as *AzureSynapse) insertIntoLoadTable(
	ctx context.Context,
	txn *sqlmw.Tx,
	tableName string,
	stagingTableName string,
	sortedColumnKeys []string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Taken from https://github.com/denisenkom/go-mssqldb/blob/master/tds.go
func str2ucs2(s string) []byte { _ = "STUB: not implemented"; return nil }

func hasDiacritics(str string) bool { _ = "STUB: not implemented"; return false }

func (as *AzureSynapse) loadUserTables(ctx context.Context) (errorMap map[string]error) {
	_ = "STUB: not implemented"
	return nil
}

// IGNORE NULLS only supported in Azure SQL edge, in which case the query can be shortened to below
// https://docs.microsoft.com/en-us/sql/t-sql/functions/first-value-transact-sql?view=sql-server-ver15
// caseSubQuery := fmt.Sprintf(`FIRST_VALUE(%[1]s) IGNORE NULLS OVER (PARTITION BY id ORDER BY received_at DESC ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) AS "%[1]s"`, colName)

// TODO: skipped top level temporary table for now

// BEGIN TRANSACTION

func (*AzureSynapse) DeleteBy(context.Context, []string, warehouseutils.DeleteByParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) dropStagingTable(ctx context.Context, stagingTableName string) {
	_ = "STUB: not implemented"
	return
}

func (as *AzureSynapse) createTable(ctx context.Context, name string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	// Search paths doesn't exist unlike Postgres, default is dbo. Hence, use namespace wherever possible
	return nil
}

func (as *AzureSynapse) DropTable(ctx context.Context, tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*AzureSynapse) AlterColumn(_ context.Context, _, _, _ string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

func (as *AzureSynapse) TestConnection(ctx context.Context, _ model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) Setup(_ context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) dropDanglingStagingTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchSchema returns the schema of the warehouse
func (as *AzureSynapse) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (as *AzureSynapse) LoadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (as *AzureSynapse) Cleanup(ctx context.Context) {
	_ = "STUB: not implemented"

	// extra check aside dropStagingTable(table)
	return
}

func (*AzureSynapse) LoadIdentityMergeRulesTable(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*AzureSynapse) LoadIdentityMappingsTable(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*AzureSynapse) DownloadIdentityRules(context.Context, *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) Connect(_ context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (as *AzureSynapse) TestLoadTable(ctx context.Context, _, tableName string, payloadMap map[string]any, _ string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (as *AzureSynapse) SetConnectionTimeout(timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (*AzureSynapse) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }
