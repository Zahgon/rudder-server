package router

import (
	"context"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

func (job *UploadJob) exportData() error { _ = "STUB: not implemented"; return nil }

func (job *UploadJob) identifiesTableName() string { _ = "STUB: not implemented"; return "" }

func (job *UploadJob) usersTableName() string { _ = "STUB: not implemented"; return "" }

func (job *UploadJob) identityMergeRulesTableName() string { _ = "STUB: not implemented"; return "" }

func (job *UploadJob) identityMappingsTableName() string { _ = "STUB: not implemented"; return "" }

func (job *UploadJob) TablesToSkip() (map[string]model.PendingTableUpload, map[string]model.PendingTableUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Current upload and table upload succeeded

func (job *UploadJob) getLoadFilesTableMap() (loadFilesMap map[tableNameT]bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (job *UploadJob) exportUserTables(loadFilesTableMap map[tableNameT]bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) loadUserTables(loadFilesTableMap map[tableNameT]bool) ([]error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There is at least one table to load

// Load all user tables

// Skip loading user tables if identifies table schema is not present

func (job *UploadJob) updateSchema(tName string) error { _ = "STUB: not implemented"; return nil }

func (job *UploadJob) UpdateTableSchema(tName string, tableSchemaDiff whutils.TableSchemaDiff) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) alterColumnsToWarehouse(ctx context.Context, tName string, columnsMap model.TableSchema) error {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) addColumnsToWarehouse(ctx context.Context, tName string, columnsMap model.TableSchema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) processLoadTableResponse(errorMap map[string]error) (errors []error, tableUploadErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: set last_exec_time

// Since load is successful, we assume all events in load files are uploaded

func (job *UploadJob) exportIdentities() (err error) {
	_ = "STUB: not implemented"
	// Load Identities if enabled
	return nil
}

func (job *UploadJob) loadIdentityTables(populateHistoricIdentities bool) (loadErrors []error, tableUploadErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// var generated bool

func (job *UploadJob) areIdentityTablesLoadFilesGenerated(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (job *UploadJob) resolveIdentities(populateHistoricIdentities bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) exportRegularTables(specialTables []string, loadFilesTableMap map[tableNameT]bool) (err error) {
	_ = "STUB: not implemented"
	//[]string{job.identifiesTableName(), job.usersTableName(), job.identityMergeRulesTableName(), job.identityMappingsTableName()}
	// Export all other tables
	return nil
}

func (job *UploadJob) loadAllTablesExcept(skipLoadForTables []string, loadFilesTableMap map[tableNameT]bool) []error {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) loadTable(tName string) error { _ = "STUB: not implemented"; return nil }

// columnCountStat sent the column count for a table to statsd
// skip sending for S3_DATALAKE, GCS_DATALAKE, AZURE_DATALAKE
func (job *UploadJob) columnCountStat(tableName string) { _ = "STUB: not implemented"; return }

func (job *UploadJob) RefreshPartitions(loadFileStartID, loadFileEndID int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Refresh partitions if exists
