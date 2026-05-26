package validators

import (
	"context"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

const (
	// This is integer representation of Postgres version.
	// For ex, integer representation of version 9.6.3 is 90603
	// Minimum postgres version needed for rudder server is 10
	minPostgresVersion = 100000
)

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("validators").Child("envValidator")
}

func createDBConnection() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func closeDBConnection(handle *sql.DB) error { _ = "STUB: not implemented"; return nil }

func killDanglingDBConnections(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

// IsPostgresCompatible checks the if the version of postgres is greater than minPostgresVersion
func IsPostgresCompatible(ctx context.Context, db *sql.DB) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ValidateEnv validates the current environment available for the server
func ValidateEnv() error { _ = "STUB: not implemented"; return nil }

// SQL statements in rudder-server are not executed with a timeout context, instead we are letting them take as much time as they need :)
// Due to the above, when a server shutdown is initiated in a cloud environment while long-running statements are being executed,
// the server process will not manage to shutdown gracefully, since it will be blocked by the SQL statements.
// The container orchestrator will eventually kill the server process, leaving one or more dangling connections in the database.
// This will ensure that before a new rudder-server instance starts working, all previous dangling connections belonging to this server are being killed.

// InitializeEnv initializes the environment for the server
func InitializeNodeMigrations() error { _ = "STUB: not implemented"; return nil }
