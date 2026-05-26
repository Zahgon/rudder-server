package misc

import (
	"context"
	"database/sql"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// GetConnectionString Returns Jobs DB connection configuration
func GetConnectionString(c *config.Config, componentName string) string {
	_ = "STUB: not implemented"
	return ""
}

// same database with potentially different idle tx timeouts per component

// [application_name] can be any string of less than NAMEDATALEN characters (64 characters in a standard PostgreSQL build).
// Format: [component_prefix][hostname_suffix]
//   - component_prefix: first 2 characters of componentName + "-" (omitted if componentName is empty)
//   - hostname_suffix:  first 60 characters of hostname
// It is important not to rely on postgresql's own truncation mechanism, since we are using the [app_name] for terminating dangling connections
// and we need all connections from the same host to be terminated, regardless of the component name, thus the last part of the [app_name] needs to be the same
// across all components.

type DatabaseConnectionPoolConfig struct {
	MaxOpenConns    config.ValueLoader[int]           // MaxOpenConns controls the maximum number of open connections to the database. Setting it to zero means unlimited.
	MaxIdleConns    config.ValueLoader[int]           // MaxIdleConns controls the maximum number of connections in the idle connection pool. Setting it to zero means no idle connections are retained.
	ConnMaxIdleTime config.ValueLoader[time.Duration] // ConnMaxIdleTime sets the maximum amount of time a connection may be idle. Setting it to zero means no limit.
	ConnMaxLifetime config.ValueLoader[time.Duration] // ConnMaxLifetime sets the maximum amount of time a connection may be reused. Setting it to zero means no limit.
	UpdateInterval  time.Duration                     // UpdateInterval sets the interval at which the connection pool configuration is reloaded. Setting it to zero means no reloads.
}

func NewDatabaseConnectionPool(ctx context.Context, componentName string, conf DatabaseConnectionPoolConfig, c *config.Config, stat stats.Stats) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updatePoolConfig[T comparable](setter func(T), current *T, conf config.ValueLoader[T]) {
	_ = "STUB: not implemented"
	return
}

// SetAppNameInDBConnURL sets application name in db connection url
// if application name is already present in dns it will get override by the appName
func SetAppNameInDBConnURL(connectionUrl, appName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func QuoteLiteral(literal string) string { _ = "STUB: not implemented"; return "" }

// DBCopyIn generates a COPY ... FROM STDIN statement for the given table and columns.
func DBCopyIn(table string, columns ...string) string { _ = "STUB: not implemented"; return "" }
