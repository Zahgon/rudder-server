package repo

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const whSchemaTableName = warehouseutils.WarehouseSchemasTable

const whSchemaTableColumns = `
	id,
   	source_id,
	namespace,
   	destination_id,
	destination_type,
	schema,
   	created_at,
   	updated_at,
	expires_at
`

type WHSchema struct {
	*repo

	log    logger.Logger
	config struct {
		enableTableLevelSchema config.ValueLoader[bool]
	}
}

func NewWHSchemas(db *sqlmiddleware.DB, conf *config.Config, log logger.Logger, opts ...Opt) *WHSchema {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a schema row in wh_schemas with the given schema.
// If Warehouse.enableTableLevelSchema is true in config, it also inserts/updates table-level schemas for each table in the schema.
func (sh *WHSchema) Insert(ctx context.Context, whSchema *model.WHSchema) error {
	_ = "STUB: not implemented"
	return nil
}

// update all schemas with the same destination_id and namespace but different source_id
// this is to ensure all the connections for a destination have the same schema copy

// Then, insert/update the new schema using the unique constraint

// If table-level schema is enabled, insert/update for each table

// Delete orphaned table-level schemas for the current source_id + destination_id + namespace
// This ensures consistency between raw schema and table-level schemas

// GetForNamespace fetches the schema for a namespace, supporting both legacy and table-level modes.
func (sh *WHSchema) GetForNamespace(ctx context.Context, destID, namespace string) (model.WHSchema, error) {
	_ = "STUB: not implemented"
	return *new(model.WHSchema), nil
}

func (sh *WHSchema) getForNamespace(ctx context.Context, destID, namespace string) (model.WHSchema, error) {
	_ = "STUB: not implemented"
	return *new(model.WHSchema), nil
}

// populateTableLevelSchemasWithTx inserts table-level schemas for each table in the parent schema
// that don't already exist as separate table-level schemas, using the provided transaction.
func (sh *WHSchema) populateTableLevelSchemasWithTx(ctx context.Context, tx *sqlmiddleware.Tx, destID, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// getTableLevelSchemasForNamespaceWithTx fetches the latest schema (by id) for each table in the given destID and namespace, regardless of source.
func (sh *WHSchema) getTableLevelSchemasForNamespaceWithTx(ctx context.Context, tx *sqlmiddleware.Tx, destID, namespace string) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func parseWHSchemas(rows *sqlmiddleware.Rows) ([]*model.WHSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sh *WHSchema) GetNamespace(ctx context.Context, sourceID, destID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (sh *WHSchema) GetTablesForConnection(ctx context.Context, connections []warehouseutils.SourceIDDestinationID) ([]warehouseutils.FetchTableInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// select all rows with max id for each source id and destination id pair

func (sh *WHSchema) SetExpiryForDestination(ctx context.Context, destinationID string, expiresAt time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDestinationNamespaces returns the most recent namespace for each source for a given destination ID.
func (sh *WHSchema) GetDestinationNamespaces(ctx context.Context, destinationID string) ([]model.NamespaceMapping, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
