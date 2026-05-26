package snowflake

import (
	"context"
	"errors"
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
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	provider       = whutils.SNOWFLAKE
	tableNameLimit = 127
)

var primaryKeyMap = map[string]string{
	usersTable:      "ID",
	identifiesTable: "ID",
	discardsTable:   "ROW_ID",
}

var partitionKeyMap = map[string]string{
	usersTable:      `"ID"`,
	identifiesTable: `"ID"`,
	discardsTable:   `"ROW_ID", "COLUMN_NAME", "TABLE_NAME"`,
}

var errNoGrants = errors.New("no grants found")

var (
	usersTable              = whutils.ToProviderCase(whutils.SNOWFLAKE, whutils.UsersTable)
	identifiesTable         = whutils.ToProviderCase(whutils.SNOWFLAKE, whutils.IdentifiesTable)
	discardsTable           = whutils.ToProviderCase(whutils.SNOWFLAKE, whutils.DiscardsTable)
	identityMergeRulesTable = whutils.ToProviderCase(whutils.SNOWFLAKE, whutils.IdentityMergeRulesTable)
	identityMappingsTable   = whutils.ToProviderCase(whutils.SNOWFLAKE, whutils.IdentityMappingsTable)
)

var errorsMappings = []model.JobError{
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`The requested warehouse does not exist or not authorized`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`The requested database does not exist or not authorized`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`failed to connect to db. verify account name is correct`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`Incorrect username or password was specified`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`Insufficient privileges to operate on table`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`IP .* is not allowed to access Snowflake. Contact your local security administrator or please create a case with Snowflake Support or reach us on our support line`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`User temporarily locked`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`Schema .* already exists, but current role has no privileges on it`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`The AWS Access Key Id you provided is not valid`),
	},
	{
		Type:   model.PermissionError,
		Format: regexp.MustCompile(`Location .* is not allowed by integration .*. Please use DESC INTEGRATION to check out allowed and blocked locations.`),
	},
	{
		Type:   model.InsufficientResourceError,
		Format: regexp.MustCompile(`Warehouse .* cannot be resumed because resource monitor .* has exceeded its quota`),
	},
	{
		Type:   model.InsufficientResourceError,
		Format: regexp.MustCompile(`Your free trial has ended and all of your virtual warehouses have been suspended. Add billing information in the Snowflake web UI to continue using the full set of Snowflake features.`),
	},
	{
		Type:   model.ResourceNotFoundError,
		Format: regexp.MustCompile(`Table .* does not exist`),
	},
	{
		Type:   model.ColumnCountError,
		Format: regexp.MustCompile(`Operation failed because soft limit on objects of type 'Column' per table was exceeded. Please reduce number of 'Column's or contact Snowflake support about raising the limit.`),
	},
}

type tableLoadResp struct {
	db           *sqlmw.DB
	stagingTable string
}

type optionalCreds struct {
	schemaName string
}

type duplicateMessage struct {
	id         string
	receivedAt time.Time
}
type privilegeGrant struct {
	name, privilege, grantedOn, granteeName string
}

func (m *duplicateMessage) String() string { _ = "STUB: not implemented"; return "" }

type Snowflake struct {
	DB             *sqlmw.DB
	Namespace      string
	CloudProvider  string
	ObjectStorage  string
	Warehouse      model.Warehouse
	Uploader       whutils.Uploader
	connectTimeout time.Duration
	conf           *config.Config
	logger         logger.Logger
	stats          stats.Stats
	tableManager   tableManager

	config struct {
		allowMerge         bool
		slowQueryThreshold time.Duration
		enableDeleteByJobs bool
		appendOnlyTables   []string

		debugDuplicateWorkspaceIDs   []string
		debugDuplicateTables         []string
		debugDuplicateIntervalInDays int
		debugDuplicateLimit          int
		privileges                   struct {
			fetchSchema struct {
				required []string
				enabled  bool
			}
		}
	}
}

func New(conf *config.Config, log logger.Logger, stat stats.Stats) *Snowflake {
	_ = "STUB: not implemented"
	return nil
}

// appendOnlyTables is a workaround introduced for Mattermost for now. It is only supported for snowflake.

// schemaIdentifier returns [DATABASE_NAME].[NAMESPACE] format to access the schema directly.
func (sf *Snowflake) schemaIdentifier() string { _ = "STUB: not implemented"; return "" }

func (sf *Snowflake) createTable(ctx context.Context, tableName string, columns model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) tableExists(ctx context.Context, tableName string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sf *Snowflake) columnExists(ctx context.Context, columnName, tableName string) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sf *Snowflake) schemaExists(ctx context.Context) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ignore err if no results for query

func (sf *Snowflake) createSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func checkAndIgnoreAlreadyExistError(err error) bool {
	_ = "STUB: not implemented"

	// TODO: throw error if column already exists but of different type
	return false
}

func (sf *Snowflake) authString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (sf *Snowflake) DeleteBy(ctx context.Context, tableNames []string, params whutils.DeleteByParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) loadTable(
	ctx context.Context,
	tableName string,
	tableSchemaInUpload model.TableSchema,
	skipClosingDBSession bool,
) (*types.LoadTableStats, *tableLoadResp, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Truncating the columns by default to avoid size limitation errors
// https://docs.snowflake.com/en/sql-reference/sql/copy-into-table.html#copy-options-copyoptions

func (sf *Snowflake) mergeIntoLoadTable(
	ctx context.Context,
	db *sqlmw.DB,
	schemaIdentifier,
	tableName string,
	stagingTableName string,
	sortedColumnNames string,
	strKeys []string,
) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is being added in order to get the updates count

func (sf *Snowflake) joinColumnsWithFormatting(columns []string, format string) string {
	_ = "STUB: not implemented"
	return ""
}

func (sf *Snowflake) sampleDuplicateMessages(
	ctx context.Context,
	db *sqlmw.DB,
	mainTableName,
	stagingTableName string,
) ([]duplicateMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *Snowflake) copyInto(
	ctx context.Context,
	db *sqlmw.DB,
	schemaIdentifier string,
	tableName string,
	sortedColumnNames string,
	copyTargetTable string,
) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *Snowflake) LoadIdentityMergeRulesTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) LoadIdentityMappingsTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ShouldMerge returns true if:
// * the uploader says we cannot append
// * the server configuration says we can merge
// * the user opted-in
func (sf *Snowflake) ShouldMerge(tableName string) bool { _ = "STUB: not implemented"; return false }

func (sf *Snowflake) LoadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// replace staging stable with temp table, because in APPEND mode the previous "loadTable" call
// did not leave us with the ability to determine which records were inserted

// This is to handle cases when column in users table not present in identities table

func (sf *Snowflake) connect(ctx context.Context, opts optionalCreds) (*sqlmw.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unlike Snowflake destination, we don't expose the useKeyPairAuth config for Snowpipe Streaming in the UI
// So we should set it explicitly here

func (sf *Snowflake) CreateSchema(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) CreateTable(ctx context.Context, tableName string, columnMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) DropTable(ctx context.Context, tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) AddColumns(ctx context.Context, tableName string, columnsInfo []whutils.ColumnInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*Snowflake) AlterColumn(context.Context, string, string, string) (model.AlterTableResponse, error) {
	_ = "STUB: not implemented"
	return *new(model.AlterTableResponse), nil
}

// DownloadIdentityRules gets distinct combinations of anonymous_id, user_id from tables in warehouse
func (sf *Snowflake) DownloadIdentityRules(ctx context.Context, gzWriter *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// check if table in warehouse has anonymous_id and user_id and construct accordingly

// TODO: Handle case for missing anonymous_id, user_id columns

// avoid setting null merge_property_1 to avoid not null constraint in local postgres

func (sf *Snowflake) IsEmpty(ctx context.Context, warehouse model.Warehouse) (empty bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sf *Snowflake) Setup(ctx context.Context, warehouse model.Warehouse, uploader whutils.Uploader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) TestConnection(ctx context.Context, _ model.Warehouse) error {
	_ = "STUB: not implemented"
	return nil
}

// FetchSchema queries the snowflake database and returns the schema
func (sf *Snowflake) FetchSchema(ctx context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (sf *Snowflake) Cleanup(context.Context) { _ = "STUB: not implemented"; return }

func (sf *Snowflake) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *Snowflake) Connect(ctx context.Context, warehouse model.Warehouse) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func (sf *Snowflake) TestLoadTable(
	ctx context.Context, location, tableName string, _ map[string]any, _ string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) TestFetchSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) getRoles(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *Snowflake) getGrantedRoles(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sf *Snowflake) getShowGrantsPrivileges(ctx context.Context, sqlStatement string) ([]privilegeGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanRows(rows *sqlmw.Rows, columns []string, processRow func([]any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *Snowflake) SetConnectionTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (*Snowflake) ErrorMappings() []model.JobError { _ = "STUB: not implemented"; return nil }

func getSortedColumnsFromTableSchema(tableSchemaInUpload model.TableSchema) []string {
	_ = "STUB: not implemented"
	return nil
}
