package deltalake

import (
	"context"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/misc"
	warehouseclient "github.com/rudderlabs/rudder-server/warehouse/client"
	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	provider       = warehouseutils.DELTALAKE
	tableNameLimit = 127 // Maximum table name length in rudder-transformer
)

const (
	schemaNotFound = "[SCHEMA_NOT_FOUND]"
)

const (
	rudderStagingTableRegex    = "^rudder_staging_.*$"       // matches rudder_staging_* tables
	nonRudderStagingTableRegex = "^(?!rudder_staging_.*$).*" // matches tables that do not start with rudder_staging_
)

// dataTypesMap maps rudder data types to delta lake data types
var dataTypesMap = map[string]string{
	"boolean":  "BOOLEAN",
	"int":      "BIGINT",
	"float":    "DOUBLE",
	"string":   "STRING",
	"datetime": "TIMESTAMP",
	"date":     "DATE",
}

// dataTypesMapToRudder maps delta lake data types to rudder data types
// Reference: https://docs.databricks.com/sql/language-manual/sql-ref-datatype-rules.html
var dataTypesMapToRudder = map[string]string{
	"TINYINT":   "int",
	"SMALLINT":  "int",
	"INT":       "int",
	"BIGINT":    "int",
	"DECIMAL":   "float",
	"FLOAT":     "float",
	"DOUBLE":    "float",
	"BOOLEAN":   "boolean",
	"STRING":    "string",
	"DATE":      "date",
	"TIMESTAMP": "datetime",
	"tinyint":   "int",
	"smallint":  "int",
	"int":       "int",
	"bigint":    "int",
	"decimal":   "float",
	"float":     "float",
	"double":    "float",
	"boolean":   "boolean",
	"string":    "string",
	"date":      "date",
	"timestamp": "datetime",
}

var semiStructuredDataTypes = []string{
	"array",
	"map",
	"struct",
}

// excludeColumnsMap Columns you need to exclude
// Since event_date is an auto generated column in order to support partitioning.
// We need to ignore it during query generation.
var excludeColumnsMap = map[string]struct{}{
	"event_date": {},
}

var primaryKeyMap = map[string]string{
	warehouseutils.UsersTable:      "id",
	warehouseutils.IdentifiesTable: "id",
	warehouseutils.DiscardsTable:   "row_id",
}

var errorsMappings = []model.JobError{
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`UnauthorizedAccessException: PERMISSION_DENIED: User does not have READ FILES on External Location`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`SecurityException: User does not have permission CREATE on CATALOG`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`ENDPOINT_NOT_FOUND`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`RESOURCE_DOES_NOT_EXIST`),
	},
}

type Deltalake struct {
	DB             *sqlmiddleware.DB
	Namespace      string
	ObjectStorage  string
	Warehouse      model.Warehouse
	Uploader       warehouseutils.Uploader
	connectTimeout time.Duration
	conf           *config.Config
	logger         logger.Logger
	stats          stats.Stats

	config struct {
		allowMerge             bool
		enablePartitionPruning bool
		slowQueryThreshold     time.Duration
		maxRetries             int
		retryMinWait           time.Duration
		retryMaxWait           time.Duration
		maxErrorLength         int
	}
}

func New(conf *config.Config, log logger.Logger, stat stats.Stats) *Deltalake {
	_ = "STUB: not implemented"
	return nil
}

// 64 KB

// Setup sets up the warehouse
func (d *Deltalake) Setup(_ context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) error {
	_ = "STUB: not implemented"
	return nil
}

// connect connects to the warehouse
func (d *Deltalake) connect() (*sqlmiddleware.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dropDanglingStagingTables drops dangling staging tables
func (d *Deltalake) dropDanglingStagingTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// fetchTables fetches tables from the database
func (d *Deltalake) fetchTables(ctx context.Context, regex string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dropStagingTables drops all the staging tables
func (d *Deltalake) dropStagingTables(ctx context.Context, stagingTables []string) error {
	_ = "STUB: not implemented"
	return nil
}

// DropTable drops a table from the warehouse
func (d *Deltalake) dropTable(ctx context.Context, table string) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchSchema fetches the schema from the warehouse
func (d *Deltalake) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	// Since error handling is not so good with the Databricks driver we need to verify the exact string in the error.
	// Therefore, creating the schema every time before we fetch it. Also, creating the schema is idempotent.
	return *new(model.Schema), nil
}

// For each table, fetch the attributes

func (d *Deltalake) sendStatForMissingDatatype(missingDatatype string) {
	_ = "STUB: not implemented"
	return
}

// fetchTableAttributes fetches the attributes of a table
func (d *Deltalake) fetchTableAttributes(ctx context.Context, tableName string) (model.TableSchema, error) {
	_ = "STUB: not implemented"
	return *new(model.TableSchema), nil
}

// CreateSchema creates a schema in the warehouse if it does not exist.
func (d *Deltalake) CreateSchema(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// schemaExists checks if a schema exists in the warehouse.
func (d *Deltalake) schemaExists(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// createSchema creates a schema in the warehouse.
func (d *Deltalake) createSchema(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CreateTable creates a table in the warehouse.
func (d *Deltalake) CreateTable(ctx context.Context, tableName string, columns model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Deltalake) TrimErrorMessage(baseError error) error { _ = "STUB: not implemented"; return nil }

// columnsWithDataTypes returns the columns with their data types.
func columnsWithDataTypes(columns model.TableSchema, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}

// tableLocationQuery returns the location query for the table.
func (d *Deltalake) tableLocationQuery(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

// AddColumns adds columns to the table.
func (d *Deltalake) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// AlterColumn alters a column in the warehouse
func (*Deltalake) AlterColumn(context.Context, string, string, string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// LoadTable loads table for table name
func (d *Deltalake) LoadTable(
	ctx context.Context,
	tableName string,
) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Deltalake) loadTable(
	ctx context.Context,
	tableName string,
	tableSchemaInUpload model.TableSchema,
	tableSchemaAfterUpload model.TableSchema,
	skipTempTableDelete bool,
) (*types.LoadTableStats, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (d *Deltalake) copyIntoLoadTable(
	ctx context.Context,
	tableName string,
	stagingTableName string,
	tableSchemaInUpload model.TableSchema,
	tableSchemaAfterUpload model.TableSchema,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Deltalake) insertIntoLoadTable(
	ctx context.Context,
	tableName string,
	stagingTableName string,
	tableSchemaAfterUpload model.TableSchema,
) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Deltalake) mergeIntoLoadTable(
	ctx context.Context,
	tableName string,
	stagingTableName string,
	tableSchemaInUpload model.TableSchema,
) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tableSchemaDiff(tableSchemaInUpload, tableSchemaAfterUpload model.TableSchema) warehouseutils.TableSchemaDiff {
	_ = "STUB: not implemented"
	return *new(warehouseutils.TableSchemaDiff)
}

func columnNames(columns []string) string { _ = "STUB: not implemented"; return "" }

func stagingColumnNames(columns []string) string { _ = "STUB: not implemented"; return "" }

func columnsWithValues(columns []string) string { _ = "STUB: not implemented"; return "" }

func primaryKey(tableName string) string { _ = "STUB: not implemented"; return "" }

// sortedColumnNames returns the column names in the order of sortedColumnKeys
func (d *Deltalake) sortedColumnNames(tableSchemaInUpload model.TableSchema, sortedColumnKeys []string, diff warehouseutils.TableSchemaDiff) string {
	_ = "STUB: not implemented"
	return ""
}

// authQuery return authentication for AWS STS and SSE-C encryption
// STS authentication is only supported with S3A client.
func (d *Deltalake) authQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// canUseAuth returns true if the warehouse is configured to use RudderObjectStorage or STS tokens
func (d *Deltalake) canUseAuth() bool { _ = "STUB: not implemented"; return false }

// getLoadFolder returns the load folder for the warehouse load files
func (d *Deltalake) getLoadFolder(location string) string { _ = "STUB: not implemented"; return "" }

// hasAWSCredentials returns true if the warehouse is configured to use AWS credentials
func (d *Deltalake) hasAWSCredentials() bool { _ = "STUB: not implemented"; return false }

// LoadUserTables loads user tables
func (d *Deltalake) LoadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// getColumnProperties returns the column names and first value properties for the given table schema
func getColumnProperties(usersSchemaInWarehouse model.TableSchema) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadIdentityMergeRulesTable loads identifies merge rules tables
func (*Deltalake) LoadIdentityMergeRulesTable(context.Context) error {
	_ = "STUB: not implemented"

	// LoadIdentityMappingsTable loads identifies mappings table
	return nil
}

func (*Deltalake) LoadIdentityMappingsTable(context.Context) error {
	_ = "STUB: not implemented"

	// Cleanup cleans up the warehouse
	return nil
}

func (d *Deltalake) Cleanup(ctx context.Context) { _ = "STUB: not implemented"; return }

// IsEmpty checks if the warehouse is empty or not
func (*Deltalake) IsEmpty(context.Context, model.Warehouse) (bool, error) {
	_ = "STUB: not implemented"

	// TestConnection tests the connection to the warehouse
	return false, nil
}

func (d *Deltalake) TestConnection(ctx context.Context, _ model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

// DownloadIdentityRules downloadchecking if schema exists identity rules
func (*Deltalake) DownloadIdentityRules(context.Context, *misc.GZipWriter) error {
	_ = "STUB: not implemented"

	// Connect returns Client
	return nil
}

func (d *Deltalake) Connect(_ context.Context, warehouse model.Warehouse) (warehouseclient.Client, error) {
	_ = "STUB: not implemented"
	return *new(warehouseclient.Client), nil
}

// LoadTestTable loads the test table
func (d *Deltalake) TestLoadTable(ctx context.Context, location, tableName string, _ map[string]any, format string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Deltalake) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// SetConnectionTimeout sets the connection timeout
func (d *Deltalake) SetConnectionTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// ErrorMappings returns the error mappings
func (*Deltalake) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }

// DropTable drops a table in the warehouse
func (d *Deltalake) DropTable(ctx context.Context, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Deltalake) DeleteBy(context.Context, []string, warehouseutils.DeleteByParams) error {
	_ = "STUB: not implemented"
	return nil
}

// ShouldMerge returns true if:
// * the uploader says we cannot append
// * the user opted in to merging and we allow merging
func (d *Deltalake) ShouldMerge() bool { _ = "STUB: not implemented"; return false }
