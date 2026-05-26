package bigquery

import (
	"context"
	"regexp"
	"time"

	"cloud.google.com/go/bigquery"
	bqservice "google.golang.org/api/bigquery/v2"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/client"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/bigquery/middleware"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type BigQuery struct {
	db        *middleware.Client
	namespace string
	warehouse model.Warehouse
	projectID string
	uploader  warehouseutils.Uploader
	conf      *config.Config
	logger    logger.Logger
	now       func() time.Time

	config struct {
		setUsersLoadPartitionFirstEventFilter bool
		customPartitionsEnabled               bool
		enableDeleteByJobs                    bool
		customPartitionsEnabledWorkspaceIDs   []string
		slowQueryThreshold                    time.Duration
		loadByFolderPath                      bool
	}
}

type loadTableResponse struct {
	partitionDate string
}

const (
	provider       = warehouseutils.BQ
	tableNameLimit = 127
)

// dataTypesMap maps datatype stored in rudder to datatype in bigquery
var dataTypesMap = map[string]bigquery.FieldType{
	"boolean":  bigquery.BooleanFieldType,
	"int":      bigquery.IntegerFieldType,
	"float":    bigquery.FloatFieldType,
	"string":   bigquery.StringFieldType,
	"datetime": bigquery.TimestampFieldType,
}

// dataTypesMapToRudder maps datatype in bigquery to datatype stored in rudder
var dataTypesMapToRudder = map[bigquery.FieldType]string{
	"BOOLEAN":   "boolean",
	"BOOL":      "boolean",
	"INTEGER":   "int",
	"INT64":     "int",
	"NUMERIC":   "float",
	"FLOAT":     "float",
	"FLOAT64":   "float",
	"STRING":    "string",
	"BYTES":     "string",
	"DATE":      "datetime",
	"DATETIME":  "datetime",
	"TIME":      "datetime",
	"TIMESTAMP": "datetime",
}

var partitionKeyMap = map[string]string{
	warehouseutils.UsersTable:              "id",
	warehouseutils.IdentifiesTable:         "id",
	warehouseutils.DiscardsTable:           "row_id, column_name, table_name",
	warehouseutils.IdentityMappingsTable:   "merge_property_type, merge_property_value",
	warehouseutils.IdentityMergeRulesTable: "merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value",
}

var errorsMappings = []model.JobError{
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`googleapi: Error 403: Access Denied`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`googleapi: Error 404: Not found: Dataset .*, notFound`),
	},
	{
		Type:   model.ConcurrentQueriesError,
		Format: regexp.MustCompile(`googleapi: Error 400: Job exceeded rate limits: Your project_and_region exceeded quota for concurrent queries.`),
	},
	{
		Type:   model.ConcurrentQueriesError,
		Format: regexp.MustCompile(`googleapi: Error 400: Exceeded rate limits: too many concurrent queries for this project_and_region.`),
	},
	{
		Type:   model.ColumnCountError,
		Format: regexp.MustCompile(`googleapi: Error 400: Too many total leaf fields: .*, max allowed field count: 10000`),
	},
}

func New(conf *config.Config, log logger.Logger) *BigQuery { _ = "STUB: not implemented"; return nil }

func getTableSchema(tableSchema model.TableSchema) []*bigquery.FieldSchema {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) DeleteTable(ctx context.Context, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTable creates a table in BigQuery with the provided schema
// It also creates a view for the table to deduplicate the data
// If custom partitioning is enabled, it creates a table with custom partitioning based on the partition column and type
// only if the partition column exists in the schema.
// Otherwise, it creates a table with ingestion-time partitioning
func (bq *BigQuery) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

// If partition column is _PARTITIONTIME and partition type is not empty, then we only set the partition type

// Checking if the partition column exists in the schema, because in case of
// 1. rudder_discards: we only have timestamp column.
// 2. rudder_identity_merge_rules: we don't have any column.
// 3. rudder_identity_mappings: we don't have any column.

// createTableView creates a view for the table to deduplicate the data
// If custom partition is enabled, it creates a view with the partition column and type. Otherwise, it creates a view with ingestion-time partitioning
func (bq *BigQuery) createTableView(ctx context.Context, tableName string, columnMap model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) deduplicationQuery(tableName string, columnMap model.TableSchema) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// assuming it has field named id upon which dedup is done in view
// the following view takes the last two months into consideration i.e. 60 * 60 * 24 * 60 * 1000000

func (bq *BigQuery) DropTable(ctx context.Context, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) schemaExists(ctx context.Context, _, _ string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (bq *BigQuery) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func checkAndIgnoreAlreadyExistError(err error) bool { _ = "STUB: not implemented"; return false }

// 409 is returned when we try to create a table that already exists
// 400 is returned for all kinds of invalid input - so we need to check the error message too

func (bq *BigQuery) DeleteBy(ctx context.Context, tableNames []string, params warehouseutils.DeleteByParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) loadTable(ctx context.Context, tableName string) (
	*types.LoadTableStats, *loadTableResponse, error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we don't support merging in BigQuery due to its cost limitations

func (bq *BigQuery) gcsReferences(
	ctx context.Context,
	tableName string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadTableByAppend loads data into a table by appending to it
//
// In BigQuery, tables created by RudderStack are typically ingestion-time partitioned tables
// with a pseudo-column named _PARTITIONTIME. BigQuery automatically assigns rows to partitions
// based on the time when BigQuery ingests the data. To support custom field partitions, it is
// important to avoid loading data into partitioned tables with names like tableName$20191221.
// Instead, ensure that data is loaded into the appropriate ingestion-time partition, allowing
// BigQuery to manage partitioning based on the data's ingestion time.
//
// TODO: Support custom field partition on users & identifies tables
func (bq *BigQuery) loadTableByAppend(
	ctx context.Context,
	tableName string,
	gcsRef *bigquery.GCSReference,
	log logger.Logger,
) (*types.LoadTableStats, *loadTableResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func jobStatusError(status *bigquery.JobStatus) error { _ = "STUB: not implemented"; return nil }

// jobStatistics returns statistics for a job
// In case of rate limit error, it returns empty statistics
func (bq *BigQuery) jobStatistics(
	ctx context.Context,
	job *bigquery.Job,
) (*bqservice.JobStatistics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In case of rate limit error, return empty statistics

func (bq *BigQuery) LoadUserTables(ctx context.Context) (errorMap map[string]error) {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) createAndLoadStagingUsersTable(ctx context.Context, stagingTable string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) connect(ctx context.Context) (*middleware.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bq *BigQuery) dropDanglingStagingTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) IsEmpty(
	ctx context.Context,
	warehouse model.Warehouse,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (bq *BigQuery) Setup(ctx context.Context, warehouse model.Warehouse, uploader warehouseutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*BigQuery) TestConnection(_ context.Context, _ model.Warehouse) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bq *BigQuery) AddColumns(ctx context.Context, tableName string, columnsInfo []warehouseutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle error in case of single column

func (*BigQuery) AlterColumn(_ context.Context, _, _, _ string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// FetchSchema queries bigquery and returns the schema associated with provided namespace
func (bq *BigQuery) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

// if dataset resource is not found, return empty schema

// lower case all column names from bigquery

func (bq *BigQuery) Cleanup(ctx context.Context) { _ = "STUB: not implemented"; return }

func (bq *BigQuery) LoadIdentityMergeRulesTable(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) LoadIdentityMappingsTable(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) tableExists(ctx context.Context, tableName string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (bq *BigQuery) columnExists(ctx context.Context, columnName, tableName string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

type identityRules struct {
	MergeProperty1Type  string `json:"merge_property_1_type"`
	MergeProperty1Value string `json:"merge_property_1_value"`
	MergeProperty2Type  string `json:"merge_property_2_type"`
	MergeProperty2Value string `json:"merge_property_2_value"`
}

func (bq *BigQuery) DownloadIdentityRules(ctx context.Context, gzWriter *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// check if table in warehouse has anonymous_id and user_id and construct accordingly

func (bq *BigQuery) Connect(ctx context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (bq *BigQuery) TestLoadTable(ctx context.Context, location, tableName string, _ map[string]any, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func loadFolder(objectLocation string) string { _ = "STUB: not implemented"; return "" }

func (*BigQuery) SetConnectionTimeout(_ time.Duration) { _ = "STUB: not implemented"; return }

func (*BigQuery) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }
