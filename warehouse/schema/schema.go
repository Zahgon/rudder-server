package schema

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

// deprecatedColumnsRegex
// This regex is used to identify deprecated columns in the warehouse
// Example: abc-deprecated-dba626a7-406a-4757-b3e0-3875559c5840
var deprecatedColumnsRegex = regexp.MustCompile(
	`.*-deprecated-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
)

type schemaRepo interface {
	GetForNamespace(ctx context.Context, destID, namespace string) (model.WHSchema, error)
	Insert(ctx context.Context, whSchema *model.WHSchema) error
}

type stagingFileRepo interface {
	GetSchemasByIDs(ctx context.Context, ids []int64) ([]model.Schema, error)
}

type fetchSchemaRepo interface {
	FetchSchema(ctx context.Context) (model.Schema, error)
}

type Handler interface {
	// Check if schema exists for the namespace
	IsSchemaEmpty(ctx context.Context) bool
	// Retrieves the schema for a specific table
	GetTableSchema(ctx context.Context, tableName string) model.TableSchema
	// Updates the schema with the provided schema definition
	UpdateSchema(ctx context.Context, updatedSchema model.Schema) error
	// Updates the schema for a specific table
	UpdateTableSchema(ctx context.Context, tableName string, tableSchema model.TableSchema) error
	// Returns the number of columns present in the schema for a given table
	GetColumnsCount(ctx context.Context, tableName string) (int, error)
	// Merges schemas from staging files with the schema to produce a consolidated schema.
	ConsolidateStagingFilesSchema(ctx context.Context, stagingFiles []*model.StagingFile) (model.Schema, error)
	// Computes the difference between the existing schema of a table and a newly provided schema.
	// Returns details of added and modified columns
	TableSchemaDiff(ctx context.Context, tableName string, tableSchema model.TableSchema) (whutils.TableSchemaDiff, error)
	// Checks if the cached schema is outdated compared to the warehouse
	IsSchemaOutdated(ctx context.Context) (bool, error)
}

type schema struct {
	stats struct {
		schemaSize stats.Histogram
	}
	warehouse                        model.Warehouse
	log                              logger.Logger
	ttlInMinutes                     time.Duration
	schemaRepo                       schemaRepo
	stagingFilesSchemaPaginationSize int
	stagingFileRepo                  stagingFileRepo
	enableIDResolution               bool
	fetchSchemaRepo                  fetchSchemaRepo
	now                              func() time.Time
	cachedSchema                     model.Schema
	cachedSchemaExpiresAt            time.Time // To prevent a DB lookup for getting the current value of expiresAt
	cachedSchemaMu                   sync.RWMutex
}

func New(
	ctx context.Context,
	warehouse model.Warehouse,
	conf *config.Config,
	slogger logger.Logger,
	statsFactory stats.Stats,
	fetchSchemaRepo fetchSchemaRepo,
	schemaRepo schemaRepo,
	stagingFileRepo stagingFileRepo,
) (Handler, error) {
	_ = "STUB: not implemented"
	return *new(Handler), nil
}

// cachedSchema can be computed in the constructor
// we need not worry about it getting expired in the middle of the job
// since we need the schema to be the same for the entireduration of the job

// No schema found in DB, try fetching from warehouse

func (sh *schema) fetchSchemaFromWarehouse(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *schema) IsSchemaEmpty(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (sh *schema) GetTableSchema(ctx context.Context, tableName string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}

func (sh *schema) UpdateSchema(ctx context.Context, updatedSchema model.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *schema) UpdateTableSchema(ctx context.Context, tableName string, tableSchema model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *schema) GetColumnsCount(ctx context.Context, tableName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sh *schema) ConsolidateStagingFilesSchema(ctx context.Context, stagingFiles []*model.StagingFile) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (sh *schema) IsSchemaOutdated(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sh *schema) isIDResolutionEnabled() bool { _ = "STUB: not implemented"; return false }

func (sh *schema) TableSchemaDiff(ctx context.Context, tableName string, tableSchema model.TableSchema) (whutils.TableSchemaDiff, error) {
	_ = "STUB: not implemented"
	return *new(whutils.TableSchemaDiff), nil
}

/*
Routine schema updates:

	Always pass cachedSchemaExpiresAt for expiresAt.
	This ensures the DB preserves the current expiresAt, so the system can eventually recover
	from a corrupted or stale schema by re-fetching from the warehouse when the expiry is reached.
	Note: If a new entry is being created with routine schema update, the expiresAt will have zero value.

Warehouse fetch:

	Only when saving a schema as a result of a successful warehouse fetch should an updated expiresAt be passed (i.e., to extend the expiry).
*/
func (sh *schema) saveSchema(ctx context.Context, updatedSchema model.Schema, expiresAt time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// consolidateStagingSchemas merges multiple schemas into one
// Prefer the type of the first schema, If the type is text, prefer text
func consolidateStagingSchemas(consolidatedSchema model.Schema, schemas []model.Schema) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

// consolidateWarehouseSchema overwrites the consolidatedSchema with the schemaInWarehouse
// Prefer the type of the schemaInWarehouse, If the type is text, prefer text
func consolidateWarehouseSchema(consolidatedSchema, warehouseSchema model.Schema) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

// overrideUsersWithIdentifiesSchema overrides the users table with the identifies table
// users(id) <-> identifies(user_id)
// Removes the user_id column from the users table
func overrideUsersWithIdentifiesSchema(consolidatedSchema model.Schema, warehouseType string, warehouseSchema model.Schema) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

// enhanceDiscardsSchema adds the discards table to the schema
// For bq, adds the loaded_at column to be segment compatible
func enhanceDiscardsSchema(consolidatedSchema model.Schema, warehouseType string) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

// enhanceSchemaWithIDResolution adds the merge rules and mappings table to the schema if IDResolution is enabled
func enhanceSchemaWithIDResolution(consolidatedSchema model.Schema, isIDResolutionEnabled bool, warehouseType string) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}

func removeDeprecatedColumns(schema model.Schema, warehouse model.Warehouse, log logger.Logger) {
	_ = "STUB: not implemented"
	return
}

func tableSchemaDiff(tableName string, schemaMap model.Schema, tableSchema model.TableSchema) whutils.TableSchemaDiff {
	_ = "STUB: not implemented"
	return *new(whutils.TableSchemaDiff)
}
