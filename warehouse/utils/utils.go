package warehouseutils

import (
	"database/sql"
	"encoding/json"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/awsutil"
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

const (
	RS                = "RS"
	BQ                = "BQ"
	SNOWFLAKE         = "SNOWFLAKE"
	SnowpipeStreaming = "SNOWPIPE_STREAMING"
	POSTGRES          = "POSTGRES"
	CLICKHOUSE        = "CLICKHOUSE"
	MSSQL             = "MSSQL"
	AzureSynapse      = "AZURE_SYNAPSE"
	DELTALAKE         = "DELTALAKE"
	S3Datalake        = "S3_DATALAKE"
	GCSDatalake       = "GCS_DATALAKE"
	AzureDatalake     = "AZURE_DATALAKE"
)

const (
	StagingFileSucceededState = "succeeded"
	StagingFileFailedState    = "failed"
	StagingFileExecutingState = "executing"
	StagingFileAbortedState   = "aborted"
	StagingFileWaitingState   = "waiting"
)

// warehouse table names
const (
	WarehouseStagingFilesTable              = "wh_staging_files"
	WarehouseStagingFileSchemaSnapshotTable = "wh_staging_file_schema_snapshots"
	WarehouseLoadFilesTable                 = "wh_load_files"
	WarehouseUploadsTable                   = "wh_uploads"
	WarehouseTableUploadsTable              = "wh_table_uploads"
	WarehouseSchemasTable                   = "wh_schemas"
	WarehouseAsyncJobTable                  = "wh_async_jobs"
)

const (
	DiscardsTable           = "rudder_discards"
	IdentityMergeRulesTable = "rudder_identity_merge_rules"
	IdentityMappingsTable   = "rudder_identity_mappings"
	ExcludeWindowStartTime  = "excludeWindowStartTime"
	ExcludeWindowEndTime    = "excludeWindowEndTime"
)

const (
	UsersTable      = "users"
	UsersView       = "users_view"
	IdentifiesTable = "identifies"
)

const (
	DatalakeTimeWindowFormat = "2006/01/02/15"
)

const (
	CTStagingTablePrefix = "setup_test_staging"
)

const (
	WAREHOUSE             = "warehouse"
	RudderMissingDatatype = "warehouse_rudder_missing_datatype"
)

const (
	stagingTablePrefix = "rudder_staging_"
)

// Object storages
const (
	S3                 = "S3"
	AzureBlob          = "AZURE_BLOB"
	GCS                = "GCS"
	MINIO              = "MINIO"
	DigitalOceanSpaces = "DIGITAL_OCEAN_SPACES"
)

// Cloud providers
const (
	AWS   = "AWS"
	GCP   = "GCP"
	AZURE = "AZURE"
)

var (
	pkgLogger          logger.Logger
	enableIDResolution bool

	TimeWindowDestinations = []string{S3Datalake, GCSDatalake, AzureDatalake}
	awsCredsExpiryInS      config.ValueLoader[int64]

	WarehouseDestinations     = []string{RS, BQ, SNOWFLAKE, POSTGRES, CLICKHOUSE, MSSQL, AzureSynapse, S3Datalake, GCSDatalake, AzureDatalake, DELTALAKE}
	IdentityEnabledWarehouses = []string{SNOWFLAKE, BQ}
	S3PathStyleRegex          = regexp.MustCompile(`https?://s3([.-](?P<region>[^.]+))?.amazonaws\.com/(?P<bucket>[^/]+)/(?P<keyname>.*)`)
	S3VirtualHostedRegex      = regexp.MustCompile(`https?://(?P<bucket>[^/]+).s3([.-](?P<region>[^.]+))?.amazonaws\.com/(?P<keyname>.*)`)

	PseudoWarehouseDestinationMap = pseudoWarehouseDestinations()
)

func pseudoWarehouseDestinations() map[string]struct{} { _ = "STUB: not implemented"; return nil }

var WHDestNameMap = map[string]string{
	BQ:            "bigquery",
	RS:            "redshift",
	MSSQL:         "mssql",
	POSTGRES:      "postgres",
	SNOWFLAKE:     "snowflake",
	CLICKHOUSE:    "clickhouse",
	DELTALAKE:     "deltalake",
	S3Datalake:    "s3_datalake",
	GCSDatalake:   "gcs_datalake",
	AzureDatalake: "azure_datalake",
	AzureSynapse:  "azure_synapse",
}

var ObjectStorageMap = map[string]string{
	RS:            S3,
	S3Datalake:    S3,
	BQ:            GCS,
	GCSDatalake:   GCS,
	AzureDatalake: AzureBlob,
}

var SnowflakeStorageMap = map[string]string{
	AWS:   S3,
	GCP:   GCS,
	AZURE: AzureBlob,
}

var DiscardsSchema = map[string]string{
	"table_name":   "string",
	"row_id":       "string",
	"column_name":  "string",
	"column_value": "string",
	"received_at":  "datetime",
	"uuid_ts":      "datetime",
	"reason":       "string",
}

const (
	LoadFileTypeCsv     = "csv"
	LoadFileTypeJson    = "json"
	LoadFileTypeParquet = "parquet"
)

func Init() { _ = "STUB: not implemented"; return }

func loadConfig() { _ = "STUB: not implemented"; return }

type DeleteByMetaData struct {
	JobRunId  string    `json:"job_run_id"`
	TaskRunId string    `json:"task_run_id"`
	StartTime time.Time `json:"start_time"`
}

type DeleteByParams struct {
	SourceId  string
	JobRunId  string
	TaskRunId string
	StartTime time.Time
}

func (d DeleteByParams) String() string { _ = "STUB: not implemented"; return "" }

type ColumnInfo struct {
	Name  string
	Value any
	Type  string
}

type GetLoadFilesOptions struct {
	Table   string
	StartID int64
	EndID   int64
	Limit   int64
}

type LoadFile struct {
	Location string
	Metadata json.RawMessage
}

type (
	ModelWarehouse   = model.Warehouse
	ModelTableSchema = model.TableSchema
)

func IDResolutionEnabled() bool { _ = "STUB: not implemented"; return false }

type TableSchemaDiff struct {
	Exists           bool
	TableToBeCreated bool
	ColumnMap        model.TableSchema
	UpdatedSchema    model.TableSchema
	AlteredColumnMap model.TableSchema
}

type QueryResult struct {
	Columns []string
	Values  [][]string
}

type SourceIDDestinationID struct {
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
}

type FetchTableInfo struct {
	SourceID      string   `json:"source_id"`
	DestinationID string   `json:"destination_id"`
	Namespace     string   `json:"namespace"`
	Tables        []string `json:"tables"`
}

func TimingFromJSONString(str sql.NullString) (status string, recordedTime time.Time) {
	_ = "STUB: not implemented"
	return "", *new(time.Time)
}

// zero values

// GetObjectFolder returns the folder path for the storage object based on the storage provider
// eg. For provider as S3: https://test-bucket.s3.amazonaws.com/test-object.csv --> s3://test-bucket/test-object.csv
func GetObjectFolder(provider, location string) (folder string) {
	_ = "STUB: not implemented"
	return ""
}

// GetObjectFolderForDeltalake returns the folder path for the storage object based on the storage provider for delta lake
// eg. For provider as S3: https://<bucket-name>.s3.amazonaws.com/<directory-name> --> s3://<bucket-name>/<directory-name>
// eg. For provider as GCS: https://storage.cloud.google.com/<bucket-name>/<directory-name> --> gs://<bucket-name>/<directory-name>
// eg. For provider as AZURE_BLOB: https://<storage-account-name>.blob.core.windows.net/<container-name>/<directory-name> --> wasbs://<container-name>@<storage-account-name>.blob.core.windows.net/<directory-name>
func GetObjectFolderForDeltalake(provider, location string) (folder string) {
	_ = "STUB: not implemented"
	return ""
}

func GetColumnsFromTableSchema(schema model.TableSchema) []string {
	_ = "STUB: not implemented"
	return nil
}

// GetObjectLocation returns the folder path for the storage object based on the storage provider
// e.g. For provider as S3: https://test-bucket.s3.amazonaws.com/test-object.csv --> s3://test-bucket/test-object.csv
func GetObjectLocation(provider, location string) (objectLocation string) {
	_ = "STUB: not implemented"
	return ""
}

// GetObjectName extracts object/key objectName from different buckets locations
// ex: https://bucket-endpoint/bucket-name/object -> object
func GetObjectName(location string, providerConfig any, objectProvider string) (objectName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CaptureRegexGroup returns capture as per the regex provided
func CaptureRegexGroup(r *regexp.Regexp, pattern string) (groups map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetS3Location parses path-style location http url to return in s3:// format
// [Path-style access] https://s3.amazonaws.com/test-bucket/test-object.csv --> s3://test-bucket/test-object.csv
// [Virtual-hosted–style access] https://test-bucket.s3.amazonaws.com/test-object.csv --> s3://test-bucket/test-object.csv
// TODO: Handle non regex matches.
func GetS3Location(location string) (s3Location, region string) {
	_ = "STUB: not implemented"
	return "", ""
}

// GetS3LocationFolder returns the folder path for a s3 object
// https://test-bucket.s3.amazonaws.com/myfolder/test-object.csv --> s3://test-bucket/myfolder
func GetS3LocationFolder(location string) string { _ = "STUB: not implemented"; return "" }

type GCSLocationOptions struct {
	TLDFormat string
}

// GetGCSLocation parses path-style location http url to return in gcs:// format
// https://storage.googleapis.com/test-bucket/test-object.csv --> gcs://test-bucket/test-object.csv
// tldFormat is used to set return format "<tldFormat>://..."
func GetGCSLocation(location string, options GCSLocationOptions) string {
	_ = "STUB: not implemented"
	return ""
}

// GetGCSLocationFolder returns the folder path for a gcs object
// https://storage.googleapis.com/test-bucket/myfolder/test-object.csv --> gcs://test-bucket/myfolder
func GetGCSLocationFolder(location string, options GCSLocationOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func GetGCSLocations(loadFiles []LoadFile, options GCSLocationOptions) (gcsLocations []string) {
	_ = "STUB: not implemented"
	return nil
}

func GetLocationFolder(location string) string { _ = "STUB: not implemented"; return "" }

// GetAzureBlobLocation parses path-style location http url to return in azure:// format
// https://myproject.blob.core.windows.net/test-bucket/test-object.csv  --> azure://myproject.blob.core.windows.net/test-bucket/test-object.csv
func GetAzureBlobLocation(location string) string { _ = "STUB: not implemented"; return "" }

// GetAzureBlobLocationFolder returns the folder path for an azure storage object
// https://myproject.blob.core.windows.net/test-bucket/myfolder/test-object.csv  --> azure://myproject.blob.core.windows.net/myfolder
func GetAzureBlobLocationFolder(location string) string { _ = "STUB: not implemented"; return "" }

func GetS3Locations(loadFiles []LoadFile) []LoadFile { _ = "STUB: not implemented"; return nil }

func JSONSchemaToMap(rawMsg json.RawMessage) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

func DestStat(statType, statName, id string) stats.Measurement {
	_ = "STUB: not implemented"
	return *new(stats.Measurement)
}

/*
ToSafeNamespace convert name of the namespace to one acceptable by warehouse
1. Remove symbols and joins continuous letters and numbers with single underscore and if first char is a number will append an underscore before the first number
2. adds an underscore if the name is a reserved keyword in the warehouse
3. truncate the length of namespace to 127 characters
4. return "stringempty" if name is empty after conversion
examples:
omega     to omega
omega v2  to omega_v2
9mega     to _9mega
mega&     to mega
ome$ga    to ome_ga
omega$    to omega
ome_ ga   to ome_ga
9mega________-________90 to _9mega_90
Cízǔ to C_z
*/
func ToSafeNamespace(provider, name string) string { _ = "STUB: not implemented"; return "" }

/*
ToProviderCase converts string provided to case generally accepted in the warehouse for table, column, schema names etc.
e.g. columns are uppercase in SNOWFLAKE and lowercase etc. in REDSHIFT, BIGQUERY etc
*/
func ToProviderCase(provider, str string) string { _ = "STUB: not implemented"; return "" }

func SnowflakeCloudProvider(config any) string { _ = "STUB: not implemented"; return "" }

func ObjectStorageType(destType string, config any, useRudderStorage bool) string {
	_ = "STUB: not implemented"
	return ""
}

func SortColumnKeysFromColumnMap(columnMap model.TableSchema) []string {
	_ = "STUB: not implemented"
	return nil
}

func IdentityMergeRulesTableName(warehouse model.Warehouse) string {
	_ = "STUB: not implemented"
	return ""
}

func IdentityMergeRulesWarehouseTableName(provider string) string {
	_ = "STUB: not implemented"
	return ""
}

func IdentityMappingsWarehouseTableName(provider string) string {
	_ = "STUB: not implemented"
	return ""
}

func IdentityMappingsTableName(warehouse model.Warehouse) string {
	_ = "STUB: not implemented"
	return ""
}

func IdentityMappingsUniqueMappingConstraintName(warehouse model.Warehouse) string {
	_ = "STUB: not implemented"
	return ""
}

func GetWarehouseIdentifier(destType, sourceID, destinationID string) string {
	_ = "STUB: not implemented"
	return ""
}

func DoubleQuoteAndJoinByComma(elems []string) string { _ = "STUB: not implemented"; return "" }

func GetTempFileExtension(destType string) string { _ = "STUB: not implemented"; return "" }

func GetTimeWindow(ts time.Time) time.Time {
	_ = "STUB: not implemented"

	// create and return time struct for window
	return *new(time.Time)
}

// GetTablePathInObjectStorage returns the path of the table relative to the object storage bucket
// <$WAREHOUSE_DATALAKE_FOLDER_NAME>/<namespace>/tableName
func GetTablePathInObjectStorage(namespace, tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

// JoinWithFormatting returns joined string for keys with the provided formatting function.
func JoinWithFormatting(keys []string, format func(idx int, str string) string, separator string) string {
	_ = "STUB: not implemented"
	return ""
}

func CreateAWSSessionConfig(destination *backendconfig.DestinationT, serviceName string) (*awsutil.SessionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetTemporaryS3Cred(destination *backendconfig.DestinationT) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

type Tag struct {
	Name  string
	Value string
}

func WHCounterStat(s stats.Stats, name string, warehouse *model.Warehouse, extraTags ...Tag) stats.Counter {
	_ = "STUB: not implemented"
	return *new(stats.Counter)
}

// FormatPemContent formats the content of certificates and keys by adding necessary newlines around specific markers.
func FormatPemContent(content string) string {
	_ = "STUB: not implemented"
	// Remove all existing newline characters
	return ""
}

// Add a newline after specific BEGIN markers

// Add a newline before and after specific END markers

type WriteSSLKeyError struct {
	errorText string
	errorTag  string
}

func (err *WriteSSLKeyError) IsError() bool { _ = "STUB: not implemented"; return false }

func (err *WriteSSLKeyError) Error() string { _ = "STUB: not implemented"; return "" }

func (err *WriteSSLKeyError) GetErrTag() string { _ = "STUB: not implemented"; return "" }

// WriteSSLKeys writes the ssl key(s) present in the destination config
// to the file system, this function checks whether a given config is
// already written to the file system, writes to the file system if the
// content is not already written
func WriteSSLKeys(destination backendconfig.DestinationT) WriteSSLKeyError {
	_ = "STUB: not implemented"
	return *new(WriteSSLKeyError)
}

// Permission files already written to FS

func GetSSLKeyDirPath(destinationID string) (whSSLRootDir string) {
	_ = "STUB: not implemented"
	return ""
}

func GetLoadFileType(destType string) string { _ = "STUB: not implemented"; return "" }

func GetLoadFileFormat(loadFileType string) string { _ = "STUB: not implemented"; return "" }

func StagingTablePrefix(provider string) string { _ = "STUB: not implemented"; return "" }

func StagingTableName(provider, tableName string, tableNameLimit int) string {
	_ = "STUB: not implemented"
	return ""
}

// RandHex returns a random hex string of length 32
func RandHex() string { _ = "STUB: not implemented"; return "" }

func ReadAsBool(key string, config map[string]any) bool { _ = "STUB: not implemented"; return false }

func GetConnectionTimeout(destType, destID string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func IsDatalakeDestination(destType string) bool { _ = "STUB: not implemented"; return false }
