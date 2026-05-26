package rsources

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

const (
	defaultRetentionPeriodInHours             = 3 * 24
	defaultSharedRetentionBufferPeriodInHours = 24
	sharedCleanupLockID                       = 100020002
)

// ErrOperationNotSupported sentinel error indicating an unsupported operation
var ErrOperationNotSupported = errors.New("rsources: operation not supported")

// ErrInvalidPaginationToken sentinel error indicating an invalid pagination token
var ErrInvalidPaginationToken = errors.New("rsources: invalid pagination token")

// In postgres, the replication slot name can contain lower-case letters, underscore characters, and numbers.
var replSlotDisallowedChars *regexp.Regexp = regexp.MustCompile(`[^a-z0-9_]`)

type sourcesHandler struct {
	log                  logger.Logger
	config               JobServiceConfig
	localDB              *sql.DB
	sharedDB             *sql.DB
	cleanupTrigger       func() <-chan time.Time
	sharedCleanupTrigger func() <-chan time.Time
}

func (sh *sourcesHandler) GetStatus(ctx context.Context, jobRunId string, filter JobFilter) (JobStatus, error) {
	_ = "STUB: not implemented"
	return *new(JobStatus), nil
}

func (sh *sourcesHandler) getStatusInternal(ctx context.Context, db *sql.DB, jobRunId string, filter JobFilter) (JobStatus, error) {
	_ = "STUB: not implemented"
	return *new(JobStatus), nil
}

// IncrementStats checks for stats table and upserts the stats
func (*sourcesHandler) IncrementStats(ctx context.Context, tx *sql.Tx, jobRunId string, key JobTargetKey, stats Stats) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) AddFailedRecords(ctx context.Context, tx *sql.Tx, jobRunId string, key JobTargetKey, records []FailedRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) GetFailedRecords(ctx context.Context, jobRunId string, filter JobFilter, paging PagingInfo) (JobFailedRecordsV2, error) {
	_ = "STUB: not implemented"
	return *new(JobFailedRecordsV2), nil
}

// first find the list of ids (postgres query planner uses an inefficient plan if there is one id with millions of records and a few ids with a few records)

func (sh *sourcesHandler) GetFailedRecordsV1(ctx context.Context, jobRunId string, filter JobFilter, paging PagingInfo) (JobFailedRecordsV1, error) {
	_ = "STUB: not implemented"
	return *new(JobFailedRecordsV1), nil
}

// first find the list of ids (postgres query planner uses an inefficient plan if there is one id with millions of records and a few ids with a few records)

func (sh *sourcesHandler) Delete(ctx context.Context, jobRunId string, filter JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) DeleteFailedRecords(ctx context.Context, jobRunId string, filter JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) DeleteJobStatus(ctx context.Context, jobRunId string, filter JobFilter) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) CleanupLoop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) doCleanupTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// doCleanupSharedTables garbage collects rows orphaned in the shared db by nodes that left the cluster; gated by an advisory lock so only one node runs at a time.
func (sh *sourcesHandler) doCleanupSharedTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanupTablesBefore(ctx context.Context, tx *sql.Tx, before time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) readDB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (sh *sourcesHandler) init(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, config.GetDurationVar(60, time.Second, "Rsources.setupTimeout"))
	defer cancel()
	if sh.cleanupTrigger == nil {
		sh.cleanupTrigger = func() <-chan time.Time {
			return time.After(config.GetDurationVar(1, time.Hour, "Rsources.stats.cleanup.interval"))
		}
	}
	if sh.sharedCleanupTrigger == nil {
		sh.sharedCleanupTrigger = func() <-chan time.Time {
			return time.After(config.GetDurationVar(1, time.Hour, "Rsources.shared.cleanup.interval"))
		}
	}

	const lockID = 100020001

	if err := withAdvisoryLock(ctx, sh.localDB, lockID, func(tx *sql.Tx) error {
		sh.log.Debugn("setting up rsources tables", logger.NewStringField("hostname", sh.config.LocalHostname))
		if err := setupTables(ctx, sh.localDB, sh.config.LocalHostname, sh.log); err != nil {
			return err
		}
		sh.log.Debugn("rsources tables setup successfully", logger.NewStringField("hostname", sh.config.LocalHostname))
		return nil
	}); err != nil {
		return err
	}

	if sh.config.ShouldSetupSharedDB && sh.sharedDB != nil {
		if err := withAdvisoryLock(ctx, sh.sharedDB, lockID, func(_ *sql.Tx) error {
			sh.log.Debugn("setting up rsources tables for shared db", logger.NewStringField("sharedConn", sh.config.SharedConn))
			if err := setupTables(ctx, sh.sharedDB, "shared", sh.log); err != nil {
				return err
			}
			sh.log.Debugn("rsources tables for shared db setup successfully", logger.NewStringField("sharedConn", sh.config.SharedConn))

			sh.log.Debugn("setting up rsources logical replication", logger.NewStringField("hostname", sh.config.LocalHostname))
			if err := sh.setupLogicalReplication(ctx); err != nil {
				return fmt.Errorf("logical replication in %q: %w", sh.config.LocalHostname, err)
			}
			sh.log.Debugn("rsources logical replication setup successfully", logger.NewStringField("hostname", sh.config.LocalHostname))
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func setupTables(ctx context.Context, db *sql.DB, defaultDbName string, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func setupFailedKeysTable(ctx context.Context, db *sql.DB, defaultDbName string, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func setupStatsTable(ctx context.Context, db *sql.DB, defaultDbName string, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (sh *sourcesHandler) setupLogicalReplication(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Create subscription for the above publication (ignore already exists error)

// skipcq: GO-R4002

func sqlFilters(jobRunId string, filter JobFilter) (fragment string, params []any) {
	_ = "STUB: not implemented"
	return "", nil
}

func (sh *sourcesHandler) Monitor(ctx context.Context, lagGauge, replicationSlotGauge Gauger) {
	_ = "STUB: not implemented"
	return
}

// Indicates that shared db is unavailable

func withAdvisoryLock(ctx context.Context, db *sql.DB, lockId int64, f func(tx *sql.Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}
