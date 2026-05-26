package sqlquerywrapper

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/tx"
)

type Opt func(*DB)

type logger interface {
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
}

type DB struct {
	*sql.DB
	stats stats.Stats

	since              func(time.Time) time.Duration
	logger             logger
	keysAndValues      []any
	slowQueryThreshold time.Duration
	queryTimeout       time.Duration
	transactionTimeout time.Duration
	rollbackThreshold  time.Duration
	commitThreshold    time.Duration
	secretsRegex       map[string]string
}

type Rows struct {
	*sql.Rows
	context.CancelFunc
	logQ
}

func (r *Rows) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Rows) Next() bool { _ = "STUB: not implemented"; return false }

func (r *Rows) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }

func (r *Rows) Err() error { _ = "STUB: not implemented"; return nil }

type Row struct {
	*sql.Row
	context.CancelFunc
	once sync.Once
	logQ
}

func (r *Row) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }

// Err provides a way for wrapping packages to check for
// query errors without calling Scan.
func (r *Row) Err() error { _ = "STUB: not implemented"; return nil }

type Tx struct {
	*tx.Tx
	db *DB
	context.CancelFunc
}

func WithLogger(logger logger) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithStats(stats stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithKeyAndValues(keyAndValues ...any) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithSlowQueryThreshold(slowQueryThreshold time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithSecretsRegex(secretsRegex map[string]string) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithQueryTimeout imposes a timeout on each query
func WithQueryTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithTransactionTimeout imposes a timeout on the transaction
func WithTransactionTimeout(timeout time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func New(db *sql.DB, opts ...Opt) *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *DB) Query(query string, args ...any) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) QueryRow(query string, args ...any) *Row { _ = "STUB: not implemented"; return nil }

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) WithTx(ctx context.Context, fn func(*Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) logQuery(query string, since time.Time) logQ {
	_ = "STUB: not implemented"
	return *new(logQ)
}

type logQ func()

// Begin starts a transaction.
//
// Use BeginTx to pass context and options to the underlying driver.
func (db *DB) Begin() (*Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) Exec(query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (tx *Tx) Query(query string, args ...any) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) QueryRow(query string, args ...any) *Row { _ = "STUB: not implemented"; return nil }

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) Rollback() error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) Commit() error { _ = "STUB: not implemented"; return nil }

func queryContextWithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}
