package jobsdb

import (
	"github.com/rudderlabs/rudder-server/jobsdb/internal/lock"
)

// SchemaMigrationTable returns the table name used for storing current schema version.
func (jd *Handle) SchemaMigrationTable() string { _ = "STUB: not implemented"; return "" }

func (jd *Handle) RunAlwaysSchemaMigrationTable() string { _ = "STUB: not implemented"; return "" }

// setupDatabaseTables will initialize jobsdb tables using migration templates inside 'sql/migrations/jobsdb'.
// Dataset tables are not created via migration scripts, they can only be updated.
// The following data are passed to JobsDB migration templates:
// - Prefix: The table prefix used by this jobsdb instance.
// - Datasets: Array of existing dataset indices.
func (jd *Handle) setupDatabaseTables(templateData map[string]any) {
	_ = "STUB: not implemented"
	// setup migrator with appropriate schema migrations table
	return
}

// execute any necessary migrations

func (jd *Handle) runAlwaysChangesets(templateData map[string]any) {
	_ = "STUB: not implemented"
	// setup migrator with appropriate schema migrations table
	return
}

// execute any necessary migrations

func (jd *Handle) dropDatabaseTables(l lock.LockToken) { _ = "STUB: not implemented"; return }

func (jd *Handle) dropSchemaMigrationTables() { _ = "STUB: not implemented"; return }
