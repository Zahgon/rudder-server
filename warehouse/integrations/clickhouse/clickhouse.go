package clickhouse

import (
	"context"
	"regexp"
	"time"

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
	partitionField = "received_at"
)

var clickhouseDefaultDateTime, _ = time.Parse(time.RFC3339, "1970-01-01 00:00:00")

// clickhouse doesn't support bool, they recommend to use Uint8 and set 1,0
var rudderDataTypesMapToClickHouse = map[string]string{
	"int":             "Int64",
	"array(int)":      "Array(Int64)",
	"float":           "Float64",
	"array(float)":    "Array(Float64)",
	"string":          "String",
	"array(string)":   "Array(String)",
	"datetime":        "DateTime",
	"array(datetime)": "Array(DateTime)",
	"boolean":         "UInt8",
	"array(boolean)":  "Array(UInt8)",
}

var clickhouseSpecificColumnNameMappings = map[string]string{
	"event":      "LowCardinality(String)",
	"event_text": "LowCardinality(String)",
}

var datatypeDefaultValuesMap = map[string]any{
	"int":      0,
	"float":    0.0,
	"boolean":  0,
	"datetime": clickhouseDefaultDateTime,
}

var clickhouseDataTypesMapToRudder = map[string]string{
	"Int8":                             "int",
	"Int16":                            "int",
	"Int32":                            "int",
	"Int64":                            "int",
	"Array(Int64)":                     "array(int)",
	"Array(Nullable(Int64))":           "array(int)",
	"Float32":                          "float",
	"Float64":                          "float",
	"Array(Float64)":                   "array(float)",
	"Array(Nullable(Float64))":         "array(float)",
	"String":                           "string",
	"Array(String)":                    "array(string)",
	"Array(Nullable(String))":          "array(string)",
	"DateTime":                         "datetime",
	"Array(DateTime)":                  "array(datetime)",
	"Array(Nullable(DateTime))":        "array(datetime)",
	"UInt8":                            "boolean",
	"Array(UInt8)":                     "array(boolean)",
	"Array(Nullable(UInt8))":           "array(boolean)",
	"LowCardinality(String)":           "string",
	"LowCardinality(Nullable(String))": "string",
	"Nullable(Int8)":                   "int",
	"Nullable(Int16)":                  "int",
	"Nullable(Int32)":                  "int",
	"Nullable(Int64)":                  "int",
	"Nullable(Float32)":                "float",
	"Nullable(Float64)":                "float",
	"Nullable(String)":                 "string",
	"Nullable(DateTime)":               "datetime",
	"Nullable(UInt8)":                  "boolean",
	"SimpleAggregateFunction(anyLast, Nullable(Int8))":     "int",
	"SimpleAggregateFunction(anyLast, Nullable(Int16))":    "int",
	"SimpleAggregateFunction(anyLast, Nullable(Int32))":    "int",
	"SimpleAggregateFunction(anyLast, Nullable(Int64))":    "int",
	"SimpleAggregateFunction(anyLast, Nullable(Float32))":  "float",
	"SimpleAggregateFunction(anyLast, Nullable(Float64))":  "float",
	"SimpleAggregateFunction(anyLast, Nullable(String))":   "string",
	"SimpleAggregateFunction(anyLast, Nullable(DateTime))": "datetime",
	"SimpleAggregateFunction(anyLast, Nullable(UInt8))":    "boolean",
}

var errorsMappings = []model.JobError{
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`code: 516, message: .*: Authentication failed: password is incorrect, or there is no user with such name`),
	},
	{
		Type:   model.InsufficientResourceError,
		Format: regexp.MustCompile(`code: 241, message: Memory limit .* exceeded: would use .*, maximum: .*`),
	},
}

type Clickhouse struct {
	DB                 *sqlmw.DB
	Namespace          string
	ObjectStorage      string
	Warehouse          model.Warehouse
	Uploader           warehouseutils.Uploader
	connectTimeout     time.Duration
	LoadFileDownloader downloader.Downloader

	conf   *config.Config
	logger logger.Logger
	stats  stats.Stats

	config struct {
		queryDebugLogs              string
		blockSize                   string
		poolSize                    string
		readTimeout                 string
		writeTimeout                string
		compress                    bool
		disableNullable             bool
		execTimeout                 time.Duration
		commitTimeout               time.Duration
		loadTableFailureRetries     int
		numWorkersDownloadLoadFiles int
		s3EngineEnabledWorkspaceIDs []string
		slowQueryThreshold          time.Duration
		randomLoadDelay             func(string) time.Duration
		disableLoadTableStats       func(string) bool
	}
}

type credentials struct {
	host       string
	database   string
	user       string
	password   string
	port       string
	secure     string
	skipVerify string
	tlsConfig  string
	timeout    time.Duration
}

type clickHouseStat struct {
	numRowsLoadFile       stats.Counter
	downloadLoadFilesTime stats.Timer
	syncLoadFileTime      stats.Timer
	commitTime            stats.Timer
	failRetries           stats.Counter
	execTimeouts          stats.Counter
	commitTimeouts        stats.Counter
}

// newClickHouseStat Creates a new clickHouseStat instance
func (ch *Clickhouse) newClickHouseStat(tableName string) *clickHouseStat {
	_ = "STUB: not implemented"
	return nil
}

func New(conf *config.Config, log logger.Logger, stat stats.Stats) *Clickhouse {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) connectToClickhouse(includeDBInConn bool) (*sqlmw.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// connectionCredentials returns the credentials for connecting to clickhouse
// Each destination will have separate tls config, hence using destination id as tlsName
func (ch *Clickhouse) connectionCredentials() (*credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// registerTLSConfig will create a global map, use different names for the different tls config.
// clickhouse will access the config by mentioning the key in connection string
func registerTLSConfig(key, certificate string) error { _ = "STUB: not implemented"; return nil }

func (ch *Clickhouse) defaultLogFields() []any { _ = "STUB: not implemented"; return nil }

// ColumnsWithDataTypes creates columns and its datatype into sql format for creating table
func (ch *Clickhouse) ColumnsWithDataTypes(tableName string, columns model.TableSchema, notNullableColumns []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ch *Clickhouse) getClickHouseCodecForColumnType(columnType, tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func getClickhouseColumnTypeForSpecificColumn(columnName, columnType string, isNullable bool) string {
	_ = "STUB: not implemented"
	return ""
}

// getClickHouseColumnTypeForSpecificTable gets suitable columnType based on the tableName
func (ch *Clickhouse) getClickHouseColumnTypeForSpecificTable(tableName, columnName, columnType string, notNullableKey bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Nullable is not disabled for users and identity table

func (*Clickhouse) DeleteBy(context.Context, []string, warehouseutils.DeleteByParams) error {
	_ = "STUB: not implemented"
	return nil
}

func generateArgumentString(length int) string { _ = "STUB: not implemented"; return "" }

func (ch *Clickhouse) castStringToArray(data, dataType string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Since we are converting true/false to 1/0 in warehouse slave
// We need to unmarshal into []int32 first to load the data into the table
// If it is unsuccessful, we unmarshal for []bool

// typecastDataFromType typeCasts string data to the mentioned data type
func (ch *Clickhouse) typecastDataFromType(data, dataType string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// loadTable loads table to clickhouse from the load files
func (ch *Clickhouse) loadTable(ctx context.Context, tableName string, tableSchemaInUpload model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) UseS3CopyEngineForLoading() bool { _ = "STUB: not implemented"; return false }

func (ch *Clickhouse) loadByDownloadingLoadFiles(ctx context.Context, tableName string, tableSchemaInUpload model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Clickhouse stats

func (ch *Clickhouse) credentials() (accessKeyID, secretAccessKey string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (ch *Clickhouse) loadByCopyCommand(ctx context.Context, tableName string, tableSchemaInUpload model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

// 1
// 2
// 3
// 4
// 5
// 6
// 7

type tableError struct {
	enableRetry bool
	err         error
}

func (ch *Clickhouse) loadTablesFromFilesNamesWithRetry(ctx context.Context, tableName string, tableSchemaInUpload model.TableSchema, fileNames []string, chStats *clickHouseStat) (terr tableError) {
	_ = "STUB: not implemented"
	return *new(tableError)
}

// sort column names

func (ch *Clickhouse) schemaExists(ctx context.Context, schemaName string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ignore err if no results for query

/*
createUsersTable creates a user's table with engine AggregatingMergeTree,
this lets us choose aggregation logic before merging records with same user id.
current behaviour is to replace user  properties with the latest non-null values
*/
func (ch *Clickhouse) createUsersTable(ctx context.Context, name string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) partitionByClause() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ch *Clickhouse) partitionExpr() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getSortKeyTuple(sortKeyFields []string) string { _ = "STUB: not implemented"; return "" }

// CreateTable creates table with engine ReplacingMergeTree(), this is used for dedupe event data and replace it will the latest data if duplicate data found. This logic is handled by clickhouse
// The engine differs from MergeTree in that it removes duplicate entries with the same sorting key value.
func (ch *Clickhouse) CreateTable(ctx context.Context, tableName string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) DropTable(ctx context.Context, tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) CreateSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) clusterClause() string { _ = "STUB: not implemented"; return "" }

func (*Clickhouse) AlterColumn(_ context.Context, _, _, _ string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// TestConnection is used destination connection tester to test the clickhouse connection
func (ch *Clickhouse) TestConnection(ctx context.Context, _ model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) Setup(_ context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FetchSchema queries clickhouse and returns the schema associated with provided namespace
func (ch *Clickhouse) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (ch *Clickhouse) LoadUserTables(ctx context.Context) (errorMap map[string]error) {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ch *Clickhouse) Cleanup(_ context.Context) { _ = "STUB: not implemented"; return }

func (*Clickhouse) LoadIdentityMergeRulesTable(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Clickhouse) LoadIdentityMappingsTable(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Clickhouse) DownloadIdentityRules(context.Context, *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Clickhouse) IsEmpty(_ context.Context, _ model.Warehouse) (empty bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ch *Clickhouse) totalCountIntable(ctx context.Context, tableName string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ch *Clickhouse) Connect(_ context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (ch *Clickhouse) GetLogIdentifier(args ...string) string { _ = "STUB: not implemented"; return "" }

func (ch *Clickhouse) TestLoadTable(ctx context.Context, _, tableName string, payloadMap map[string]any, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *Clickhouse) SetConnectionTimeout(timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (*Clickhouse) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }
