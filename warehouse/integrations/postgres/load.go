package postgres

import (
	"context"
	"database/sql"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/types"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

type loadUsersTableResponse struct {
	identifiesError error
	usersError      error
}

func (pg *Postgres) LoadTable(ctx context.Context, tableName string) (*types.LoadTableStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pg *Postgres) loadTable(
	ctx context.Context,
	txn *sqlmiddleware.Tx,
	tableName string,
	tableSchemaInUpload model.TableSchema,
) (*types.LoadTableStats, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (pg *Postgres) loadDataIntoStagingTable(
	ctx context.Context,
	stmt *sql.Stmt,
	fileName string,
	sortedColumnKeys []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) deleteFromLoadTable(
	ctx context.Context,
	txn *sqlmiddleware.Tx,
	tableName string,
	stagingTableName string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pg *Postgres) insertIntoLoadTable(
	ctx context.Context,
	txn *sqlmiddleware.Tx,
	tableName string,
	stagingTableName string,
	sortedColumnKeys []string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pg *Postgres) LoadUserTables(ctx context.Context) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func (pg *Postgres) loadUsersTable(
	ctx context.Context,
	tx *sqlmiddleware.Tx,
	identifiesSchemaInUpload,
	usersSchemaInUpload,
	usersSchemaInWarehouse model.TableSchema,
) loadUsersTableResponse {
	_ = "STUB: not implemented"
	return *new(loadUsersTableResponse)
}

// Deduplication
// Delete from users table if the id is present in the staging table

// Insert rows from staging table to users table

func (pg *Postgres) shouldMerge(tableName string) bool { _ = "STUB: not implemented"; return false }

// If we are here it's because canSkipComputingLatestUserTraits is true.
// preferAppend doesn't apply to the users table, so we are just checking skipDedupDestinationIDs for
// backwards compatibility.
