package migrator

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

// Migrator is responsible for migrating postgres tables
type Migrator struct {
	// MigrationsTable is name of the table that holds current migration version.
	// Each migration set requires a separate MigrationsTable.
	MigrationsTable string

	// Handle is the sql.DB handle used to execute migration statements
	Handle *sql.DB

	// Indicates if migration version should be force reset to latest on file in case of revert to lower version
	// Eg. DB has v3 set in MigrationsTable but latest version in MigrationsDir is v2
	ShouldForceSetLowerVersion bool

	// Indicates if all migrations should be run ignoring the current version in MigrationsTable
	RunAlways bool

	// Migration target version override. 0 means latest.
	Version uint
}

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("sql-migrator").Child("migrator")
}

// Migrate migrates database schema using migration SQL scripts.
func (m *Migrator) Migrate(migrationsDir string) error { _ = "STUB: not implemented"; return nil }

// migrate library reports that no change was required, using ErrNoChange

// MigrateFromTemplates migrates database with migration scripts provided by golang templates.
// Migration templates are read from all files in templatesDir and converted using provided context as template data.
// Directories inside templates directory are ignored.
func (m *Migrator) MigrateFromTemplates(templatesDir string, context any) error {
	_ = "STUB: not implemented"
	// look in templatesDir for migration template files
	return nil
}

// read files and create bindata source

// read template file

// parse template

// execute template with given context

// create destination driver from db.handle

// run the migration scripts

// migrate library reports that no change was required, using ErrNoChange

func (m *Migrator) getDestinationDriver() (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func latestSourceVersion(sourceDriver source.Driver) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Migrator) forceSetLowerVersion(migration *migrate.Migrate, sourceDriver source.Driver, destinationDriver database.Driver) error {
	_ = "STUB: not implemented"
	// get current version in database migrations table
	return nil
}

// check latest version on file

// force set version in DB to latestSourceVersion
// to handle cases where we are reverting back to old version
// this assumes applied changes on database are also compatible with older versions
