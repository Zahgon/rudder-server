/*
Implementation of JobsDB for keeping track of jobs (type JobT) and job status
(type JobStatusT). Jobs are stored in jobs_%d table while job status is stored
in job_status_%d table. Each such table pair (e.g. jobs_1, job_status_1) is called
a dataset (type dataSetT). After a dataset grows beyond a size, a new dataset is
created and jobs are written to a new dataset. When most of the jobs from a dataset
have been processed, we migrate the remaining jobs to a new intermediate
dataset and delete the old dataset. The range of job ids in a dataset are tracked
via the dataSetRangeT struct

The key reason for choosing this structure is to avoid costly DELETE and UPDATE
operations in DB. Instead, we just use WRITE (append) and DELETE TABLE (deleting a file)
operations which are fast.
Also, keeping each dataset small (enough to cache in memory) ensures that reads are
mostly serviced from memory cache.
*/

package jobsdb

//go:generate mockgen -destination=../mocks/jobsdb/mock_jobsdb.go -package=mocks_jobsdb github.com/rudderlabs/rudder-server/jobsdb JobsDB

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-go-kit/stats/collectors"

	"github.com/rudderlabs/rudder-server/jobsdb/internal/cache"
	"github.com/rudderlabs/rudder-server/jobsdb/internal/dsindex"
	"github.com/rudderlabs/rudder-server/jobsdb/internal/lock"
	"github.com/rudderlabs/rudder-server/services/rmetrics"
	"github.com/rudderlabs/rudder-server/utils/misc"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
)

var ErrStaleDsList = errors.New("stale dataset list")

const (
	pgReadonlyTableExceptionFuncName = "readonly_table_exception()"
	pgErrorCodeTableReadonly         = "RS001"
)

type payloadColumnType string

const (
	JSONB payloadColumnType = "jsonb"
	BYTEA payloadColumnType = "bytea"
	TEXT  payloadColumnType = "text"
)

// GetQueryParams is a struct to hold jobsdb query params.
type GetQueryParams struct {
	// query conditions

	WorkspaceID      string
	CustomValFilters []string
	ParameterFilters []ParameterFilterT
	PartitionFilters []string

	stateFilters                   []string
	afterJobID                     *int64
	ignoreReadPartitionsExclusions bool // if true, includes results from all partitions, ignoring any preconfigured excluded read partitions

	// query limits

	// Limit the total number of jobs.
	// A value less than or equal to zero will return no results
	JobsLimit int
	// Limit the total number of events, 1 job contains 1+ event(s).
	// A value less than or equal to zero will disable this limit (no limit),
	// only values greater than zero are considered as valid limits.
	EventsLimit int
	// Limit the total job payload size
	// A value less than or equal to zero will disable this limit (no limit),
	// only values greater than zero are considered as valid limits.
	PayloadSizeLimit int64
}

// StoreSafeTx sealed interface
type StoreSafeTx interface {
	Tx() *Tx
	SqlTx() *sql.Tx
	storeSafeTxIdentifier() string
	getLastDS() dataSetT
}

type storeSafeTx struct {
	tx       *Tx
	identity string
	lastDS   dataSetT
}

func (r *storeSafeTx) storeSafeTxIdentifier() string { _ = "STUB: not implemented"; return "" }

func (r *storeSafeTx) Tx() *Tx { _ = "STUB: not implemented"; return nil }

func (r *storeSafeTx) SqlTx() *sql.Tx { _ = "STUB: not implemented"; return nil }

func (r *storeSafeTx) getLastDS() dataSetT {
	_ = "STUB: not implemented"

	// EmptyStoreSafeTx returns an empty interface usable only for tests
	return *new(dataSetT)
}

func EmptyStoreSafeTx() StoreSafeTx { _ = "STUB: not implemented"; return *new(StoreSafeTx) }

// UpdateSafeTx sealed interface
type UpdateSafeTx interface {
	Tx() *Tx
	SqlTx() *sql.Tx
	getDSList() []dataSetT
	getDSRangeList() []dataSetRangeT
	setDSList([]dataSetT, []dataSetRangeT)
	updateSafeTxSealIdentifier() string
}
type updateSafeTx struct {
	tx          *Tx
	identity    string
	dsList      []dataSetT
	dsRangeList []dataSetRangeT
}

func (r *updateSafeTx) updateSafeTxSealIdentifier() string { _ = "STUB: not implemented"; return "" }

func (r *updateSafeTx) getDSList() []dataSetT { _ = "STUB: not implemented"; return nil }

func (r *updateSafeTx) getDSRangeList() []dataSetRangeT { _ = "STUB: not implemented"; return nil }

func (r *updateSafeTx) setDSList(dsList []dataSetT, dsRangeList []dataSetRangeT) {
	_ = "STUB: not implemented"
	return
}

func (r *updateSafeTx) Tx() *Tx { _ = "STUB: not implemented"; return nil }

func (r *updateSafeTx) SqlTx() *sql.Tx {
	_ = "STUB: not implemented"

	// EmptyUpdateSafeTx returns an empty interface usable only for tests
	return nil
}

func EmptyUpdateSafeTx() UpdateSafeTx { _ = "STUB: not implemented"; return *new(UpdateSafeTx) }

// HandleInspector is only intended to be used by tests for verifying the handle's internal state
type HandleInspector struct {
	*Handle
}

// MoreToken is a token that can be used to fetch more jobs
type MoreToken any

// MoreJobsResult is a JobsResult with a MoreToken
type MoreJobsResult struct {
	JobsResult
	More MoreToken
}

/*
JobsDB interface contains public methods to access JobsDB data
*/
type JobsDB interface {
	// Identifier returns the jobsdb's identifier, a.k.a. table prefix
	Identifier() string

	/* Commands */

	// WithTx begins a new transaction that can be used by the provided function.
	// If the function returns an error, the transaction will be rollbacked and return the error,
	// otherwise the transaction will be committed and a nil error will be returned.
	WithTx(context.Context, func(tx *Tx) error) error

	// WithStoreSafeTx prepares a store-safe environment and then starts a transaction
	// that can be used by the provided function.
	WithStoreSafeTx(context.Context, func(tx StoreSafeTx) error) error

	// WithStoreSafeTxFromTx prepares a store-safe environment for an existing transaction.
	WithStoreSafeTxFromTx(context.Context, *Tx, func(tx StoreSafeTx) error) error

	// Store stores the provided jobs to the database
	Store(ctx context.Context, jobList []*JobT) error

	// StoreInTx stores the provided jobs to the database using an existing transaction.
	// Please ensure that you are using an StoreSafeTx, e.g.
	//    jobsdb.WithStoreSafeTx(ctx, func(tx StoreSafeTx) error {
	//	      jobsdb.StoreInTx(ctx, tx, jobList)
	//    })
	StoreInTx(ctx context.Context, tx StoreSafeTx, jobList []*JobT) error

	// StoreEachBatchRetry tries to store all the provided job batches to the database
	//
	// returns the uuids of first job of each failed batch
	// Deprecated: use Store instead
	StoreEachBatchRetry(ctx context.Context, jobBatches [][]*JobT) map[uuid.UUID]string

	// StoreEachBatchRetryInTx tries to store all the provided job batches to the database, using an existing transaction.
	//
	// returns the uuids of first job of each failed batch
	//
	// Please ensure that you are using an StoreSafeTx, e.g.
	//    jobsdb.WithStoreSafeTx(func(tx StoreSafeTx) error {
	//	      jobsdb.StoreEachBatchRetryInTx(ctx, tx, jobBatches)
	//    })
	// Deprecated: use StoreInTx instead
	StoreEachBatchRetryInTx(ctx context.Context, tx StoreSafeTx, jobBatches [][]*JobT) (map[uuid.UUID]string, error)

	// WithUpdateSafeTx prepares an update-safe environment and then starts a transaction
	// that can be used by the provided function. An update-safe transaction shall be used if the provided function
	// needs to call UpdateJobStatusInTx.
	WithUpdateSafeTx(context.Context, func(tx UpdateSafeTx) error) error

	// WithUpdateSafeTxFromTx prepares an update-safe environment for an existing transaction.
	WithUpdateSafeTxFromTx(ctx context.Context, tx *Tx, f func(tx UpdateSafeTx) error) error

	// UpdateJobStatus updates the provided job statuses
	UpdateJobStatus(ctx context.Context, statusList []*JobStatusT) error

	// UpdateJobStatusInTx updates the provided job statuses in an existing transaction.
	// Please ensure that you are using an UpdateSafeTx, e.g.
	//    jobsdb.WithUpdateSafeTx(ctx, func(tx UpdateSafeTx) error {
	//	      jobsdb.UpdateJobStatusInTx(ctx, tx, statusList)
	//    })
	UpdateJobStatusInTx(ctx context.Context, tx UpdateSafeTx, statusList []*JobStatusT) error

	/* Queries */

	// GetJobs finds jobs in any of the provided state(s)
	GetJobs(ctx context.Context, states []string, params GetQueryParams) (JobsResult, error)

	// GetUnprocessed finds unprocessed jobs, i.e. new jobs whose state hasn't been marked in the database yet
	GetUnprocessed(ctx context.Context, params GetQueryParams) (JobsResult, error)

	// GetImporting finds jobs in importing state
	GetImporting(ctx context.Context, params GetQueryParams, more MoreToken) (*MoreJobsResult, error)

	// GetAborted finds jobs in aborted state
	GetAborted(ctx context.Context, params GetQueryParams) (JobsResult, error)

	// GetWaiting finds jobs in waiting state
	GetWaiting(ctx context.Context, params GetQueryParams) (JobsResult, error)

	// GetSucceeded finds jobs in succeeded state
	GetSucceeded(ctx context.Context, params GetQueryParams) (JobsResult, error)

	// GetFailed finds jobs in failed state
	GetFailed(ctx context.Context, params GetQueryParams) (JobsResult, error)

	// GetToProcess finds jobs in any of the following states: failed, waiting, unprocessed.
	// It also returns a MoreToken that can be used to fetch more jobs, if available, with a subsequent call.
	GetToProcess(ctx context.Context, params GetQueryParams, more MoreToken) (*MoreJobsResult, error)

	// GetPileUpCounts returns statistics (counters) of incomplete jobs
	// grouped by workspaceId and destination type
	GetPileUpCounts(ctx context.Context, cutoffTime time.Time, increaseFunc rmetrics.IncreasePendingEventsFunc) (err error)

	// GetDistinctParameterValues returns the list of distinct parameter("source_id", "destination_id", "workspace_id") values inside the jobs tables filtering for the passed customVal
	GetDistinctParameterValues(ctx context.Context, parameter ParameterName, customValFilter string) (values []string, err error)

	/* Admin */

	Ping() error
	DeleteExecuting()
	FailExecuting()
	RefreshDSList(ctx context.Context) error

	/* Journal */

	GetJournalEntries(opType string) (entries []JournalEntryT)
	JournalDeleteEntry(opID int64)
	JournalMarkStart(opType string, opPayload json.RawMessage) (int64, error)
	JournalMarkDone(opID int64) error

	IsMasterBackupEnabled() bool

	ReadExcludedPartitionsManager

	// lifecycle management
	Start() error
	Stop()
}

type ParameterName interface {
	string() string
}

type parameterName string

func (p parameterName) string() string { _ = "STUB: not implemented"; return "" }

const (
	SourceID      parameterName = "parameters->>'source_id'"
	DestinationID parameterName = "parameters->>'destination_id'"
	WorkspaceID   parameterName = "workspace_id"
)

/*
assertInterface contains public assert methods
*/
type assertInterface interface {
	assert(cond bool, errorString string)
	assertError(err error)
}

/*
UpdateJobStatusInTx updates the status of a batch of jobs in the past transaction
customValFilters[] is passed, so we can efficiently mark empty cache
Later we can move this to query
IMP NOTE: AcquireUpdateJobStatusLocks Should be called before calling this function
*/
func (jd *Handle) UpdateJobStatusInTx(ctx context.Context, tx UpdateSafeTx, statusList []*JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

/*
JobStatusT is used for storing status of the job. It is
the responsibility of the user of this module to set appropriate
job status. State can be one of
ENUM waiting, executing, succeeded, waiting_retry,  failed, aborted
*/
type JobStatusT struct {
	JobID         int64           `json:"JobID"`
	JobState      string          `json:"JobState"` // ENUM waiting, executing, succeeded, waiting_retry, filtered, failed, aborted, migrating, migrated, wont_migrate
	AttemptNum    int             `json:"AttemptNum"`
	ExecTime      time.Time       `json:"ExecTime"`
	RetryTime     time.Time       `json:"RetryTime"`
	ErrorCode     string          `json:"ErrorCode"`
	ErrorResponse json.RawMessage `json:"ErrorResponse"`
	Parameters    json.RawMessage `json:"Parameters"`
	JobParameters json.RawMessage `json:"-"`           // not stored in DB
	WorkspaceId   string          `json:"WorkspaceId"` // TODO: do we really need this field stored in DB?
	PartitionID   string          `json:"-"`           // not stored in DB
	CustomVal     string          `json:"-"`           // not stored in DB
}

type ConnectionDetails struct {
	SourceID      string
	DestinationID string
}

func (r *JobStatusT) sanitizeJson() error { _ = "STUB: not implemented"; return nil }

/*
JobT is the basic type for creating jobs. The JobID is generated
by the system and LastJobStatus is populated when reading a processed
job  while rest should be set by the user.
*/
type JobT struct {
	UUID          uuid.UUID       `json:"UUID"`
	JobID         int64           `json:"JobID"`
	UserID        string          `json:"UserID"`
	CreatedAt     time.Time       `json:"CreatedAt"`
	ExpireAt      time.Time       `json:"ExpireAt"`
	CustomVal     string          `json:"CustomVal"`
	EventCount    int             `json:"EventCount"`
	EventPayload  json.RawMessage `json:"EventPayload"`
	LastJobStatus JobStatusT      `json:"LastJobStatus"`
	Parameters    json.RawMessage `json:"Parameters"`
	WorkspaceId   string          `json:"WorkspaceId"`
	PartitionID   string          `json:"PartitionId"`
}

func (job *JobT) String() string { _ = "STUB: not implemented"; return "" }

func (job *JobT) sanitizeJSON() error { _ = "STUB: not implemented"; return nil }

// The struct fields need to be exposed to JSON package
type dataSetT struct {
	JobTable       string `json:"job"`
	JobStatusTable string `json:"status"`
	Index          string `json:"index"`
}

func (ds dataSetT) String() string { _ = "STUB: not implemented"; return "" }

type dataSetTList []dataSetT

func (l dataSetTList) String() string { _ = "STUB: not implemented"; return "" }

// dropDSEntry tracks a dataset to drop along with the dslist version that needs to be drained from readers before actually being able to drop the dataset
type dropDSEntry struct {
	ds      dataSetT // dataset to drop
	version uint64   // the dslist version that needs to be drained from readers before dropping the dataset
}

type dataSetRangeT struct {
	minJobID int64
	maxJobID int64
	ds       dataSetT
}

func (ds dataSetRangeT) String() string { _ = "STUB: not implemented"; return "" }

type dataSetRangeTList []dataSetRangeT

func (l dataSetRangeTList) String() string { _ = "STUB: not implemented"; return "" }

type dsRangeMinMax struct {
	minJobID sql.NullInt64
	maxJobID sql.NullInt64
}

/*
Handle is the main type implementing the database for implementing
jobs. The caller must call the SetUp function on a Handle object
*/
type Handle struct {
	dbHandle             *sql.DB
	priorityPool         *sql.DB // dedicated connection pool for high-priority operations (e.g., partition migration)
	sharedConnectionPool bool
	ownerType            OwnerType
	tablePrefix          string
	logger               logger.Logger
	stats                stats.Stats

	dsList              *versionedDSList
	dropDSListLock      sync.RWMutex  // protects dropDSList
	dropDSList          []dropDSEntry // list of datasets to be dropped asynchronously in a non-blocking fashion
	dropNotify          chan struct{} // used to notify the dropDSLoop of new entries in dropDSList
	dsRangeFuncMap      map[string]func() (dsRangeMinMax, error)
	distinctValuesCache *distinctValuesCache
	dsListLock          *lock.Locker
	dsMigrationLock     *lock.Locker
	// lastMigrateProbeIndex stores the dsindex of the last dataset probed by
	// getMigrationList when no eligible datasets were found and scanning was
	// cut short by maxMigrateDSProbe. The next invocation resumes from here
	// instead of re-scanning from the beginning.
	// Only accessed from the single migrateDSLoop goroutine.
	lastMigrateProbeIndex *dsindex.Index
	noResultsCache        *cache.NoResultsCache[ParameterFilterT]

	excludedReadPartitionsLock sync.RWMutex
	excludedReadPartitions     map[string]struct{}

	// table count stats
	statTableCount        stats.Measurement
	statPreDropTableCount stats.Measurement

	statReadExcludedPartitionsCount stats.Gauge

	// ds creation and drop period stats
	statNewDSPeriod               stats.Measurement
	newDSCreationTime             time.Time
	statDropDSPeriod              stats.Measurement
	dsDropTime                    time.Time
	isStatNewDSPeriodInitialized  bool
	isStatDropDSPeriodInitialized bool

	backgroundCancel context.CancelFunc
	backgroundGroup  *errgroup.Group

	// skipSetupDBSetup is useful for testing as we mock the database client
	// TriggerAddNewDS, TriggerMigrateDS is useful for triggering addNewDS to run from tests.
	// TODO: Ideally we should refactor the code to not use this override.
	TriggerAddNewDS  func() <-chan time.Time
	migrateDSPaused  atomic.Bool
	TriggerMigrateDS func() <-chan time.Time
	TriggerRefreshDS func() <-chan time.Time

	lifecycle struct {
		mu      sync.Mutex
		started bool
	}

	config *config.Config
	conf   struct {
		payloadColumnType               payloadColumnType
		maxTableSize                    config.ValueLoader[int64]
		cacheExpiration                 config.ValueLoader[time.Duration]
		addNewDSLoopSleepDuration       config.ValueLoader[time.Duration]
		addNewDSTimeout                 config.ValueLoader[time.Duration]
		refreshDSListLoopSleepDuration  config.ValueLoader[time.Duration]
		refreshDSTimeout                config.ValueLoader[time.Duration]
		minDSRetentionPeriod            config.ValueLoader[time.Duration]
		maxDSRetentionPeriod            config.ValueLoader[time.Duration]
		jobMaxAge                       config.ValueLoader[time.Duration]
		writeCapacity                   chan struct{}
		readCapacity                    chan struct{}
		enableWriterQueue               bool
		enableReaderQueue               bool
		clearAll                        bool
		skipMaintenanceError            bool
		dsLimit                         config.ValueLoader[int]
		maxReaders                      int
		maxWriters                      int
		maxOpenConnections              int
		analyzeThreshold                config.ValueLoader[int]
		MaxDSSize                       config.ValueLoader[int]
		numPartitions                   int // if zero or negative, no partitioning
		partitionFunction               func(job *JobT) string
		warnOnStatusMissingPartitionID  config.ValueLoader[bool]
		holdDSListLockDuringStore       config.ValueLoader[bool] // escape hatch: hold the dsList read lock for the entire store callback
		noResultsCacheStateOptimization config.ValueLoader[bool]
		dbTablesVersion                 int // version of the database tables schema (0 means latest)

		migration struct {
			maxMigrateOnce, maxMigrateDSProbe config.ValueLoader[int]
			vacuumFullStatusTableThreshold    config.ValueLoader[int64]
			vacuumAnalyzeStatusTableThreshold config.ValueLoader[int64]
			jobStatusMigrateThres             config.ValueLoader[float64]
			jobMinRowsLeftMigrateThreshold    config.ValueLoader[float64]
			migrateDSLoopSleepDuration        config.ValueLoader[time.Duration]
			migrateDSTimeout                  config.ValueLoader[time.Duration]
			// nonBlockingCompletedDSDrop routes datasets with zero pending jobs
			// through the async dropDSLoop instead of the in-TX postMigrateHandleDS
			// path, so the drop is migration-lock-free and does not block concurrent readers.
			nonBlockingCompletedDSDrop config.ValueLoader[bool]
		}
		backup struct {
			masterBackupEnabled config.ValueLoader[bool]
		}
	}
}

func (jd *Handle) IsMasterBackupEnabled() bool { _ = "STUB: not implemented"; return false }

// The struct which is written to the journal
type journalOpPayloadT struct {
	From []dataSetT `json:"from"`
	To   dataSetT   `json:"to"`
}

type ParameterFilterT struct {
	Name  string
	Value string
}

func (p ParameterFilterT) String() string { _ = "STUB: not implemented"; return "" }

func (p ParameterFilterT) GetName() string { _ = "STUB: not implemented"; return "" }

func (p ParameterFilterT) GetValue() string { _ = "STUB: not implemented"; return "" }

type ParameterFilterList []ParameterFilterT

func (l ParameterFilterList) String() string { _ = "STUB: not implemented"; return "" }

var dbInvalidJsonErrors = map[string]struct{}{
	"22P02": {},
	"22P05": {},
	"22025": {},
	"22019": {},
	"22021": {}, // invalid byte sequence for encoding "UTF8"
}

// Some helper functions
func (jd *Handle) assertError(err error) { _ = "STUB: not implemented"; return }

func (jd *Handle) assert(cond bool, errorString string) { _ = "STUB: not implemented"; return }

type jobStateT struct {
	isValid    bool
	isTerminal bool
	State      string
}

// State definitions
var (
	// Not valid, Not terminal
	Unprocessed = jobStateT{isValid: false, isTerminal: false, State: "not_picked_yet"}

	// Valid, Not terminal
	Failed    = jobStateT{isValid: true, isTerminal: false, State: "failed"}
	Executing = jobStateT{isValid: true, isTerminal: false, State: "executing"}
	Waiting   = jobStateT{isValid: true, isTerminal: false, State: "waiting"}
	Importing = jobStateT{isValid: true, isTerminal: false, State: "importing"}

	// Valid, Terminal
	Succeeded = jobStateT{isValid: true, isTerminal: true, State: "succeeded"}
	Aborted   = jobStateT{isValid: true, isTerminal: true, State: "aborted"}
	Migrated  = jobStateT{isValid: true, isTerminal: true, State: "migrated"}
	Filtered  = jobStateT{isValid: true, isTerminal: true, State: "filtered"}

	terminalStates         map[string]struct{}
	validTerminalStates    []string
	validNonTerminalStates []string
)

// Adding a new state to this list, will require an enum change in postgres db.
var jobStates = []jobStateT{
	Unprocessed,
	Failed,
	Executing,
	Waiting,
	Succeeded,
	Aborted,
	Migrated,
	Importing,
	Filtered,
}

// OwnerType for this jobsdb instance
type OwnerType string

const (
	// Read : Only Reader of this jobsdb instance
	Read OwnerType = "READ"
	// Write : Only Writer of this jobsdb instance
	Write OwnerType = "WRITE"
	// ReadWrite : Reader and Writer of this jobsdb instance
	ReadWrite OwnerType = ""
)

func (ot OwnerType) Identifier() string { _ = "STUB: not implemented"; return "" }

func init() {
	terminalStates = make(map[string]struct{})
	for _, js := range jobStates {
		if !js.isValid {
			continue
		}
		if js.isTerminal {
			terminalStates[js.State] = struct{}{}
			validTerminalStates = append(validTerminalStates, js.State)
		} else {
			validNonTerminalStates = append(validNonTerminalStates, js.State)
		}
	}
}

type OptsFunc func(jd *Handle)

// WithClearDB if set to true it will remove all existing tables
func WithClearDB(clearDB bool) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

func WithDSLimit(limit config.ValueLoader[int]) OptsFunc {
	_ = "STUB: not implemented"
	return *new(OptsFunc)
}

func WithDBHandle(dbHandle *sql.DB) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

func WithConfig(c *config.Config) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

func WithStats(s stats.Stats) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

func WithSkipMaintenanceErr(ignore bool) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

func WithJobMaxAge(jobMaxAge config.ValueLoader[time.Duration]) OptsFunc {
	_ = "STUB: not implemented"
	return *new(OptsFunc)
}

func WithNumPartitions(numPartitions int) OptsFunc {
	_ = "STUB: not implemented"
	return *new(OptsFunc)
}

// numPartitions must be a power-of-two number

// default partition function using a 32-bit key space and Murmur3 hash

// WithPriorityPoolDB sets a dedicated connection pool for high-priority operations.
// Operations that use WithPriorityPool(ctx) context will use this pool and bypass
// the regular reader/writer queues.
func WithPriorityPoolDB(pool *sql.DB) OptsFunc { _ = "STUB: not implemented"; return *new(OptsFunc) }

// withDatabaseTablesVersion sets the database tables version to use (internal use only for verifying database table migrations)
func withDatabaseTablesVersion(dbVersion int) OptsFunc {
	_ = "STUB: not implemented"
	return *new(OptsFunc)
}

func NewForRead(tablePrefix string, opts ...OptsFunc) *Handle {
	_ = "STUB: not implemented"
	return nil
}

func NewForWrite(tablePrefix string, opts ...OptsFunc) *Handle {
	_ = "STUB: not implemented"
	return nil
}

func NewForReadWrite(tablePrefix string, opts ...OptsFunc) *Handle {
	_ = "STUB: not implemented"
	return nil
}

func newOwnerType(ownerType OwnerType, tablePrefix string, opts ...OptsFunc) *Handle {
	_ = "STUB: not implemented"
	return nil
}

/*
Setup is used to initialize the HandleT structure.
clearAll = True means it will remove all existing tables
tablePrefix must be unique and is used to separate
multiple users of JobsDB
*/
func (jd *Handle) Setup(
	ownerType OwnerType, clearAll bool, tablePrefix string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) init() {
	jd.dsListLock = lock.NewLocker()
	jd.dsMigrationLock = lock.NewLocker()
	jd.dsList = newVersionedDSList(nil, nil)
	jd.dropNotify = make(chan struct{}, 1)
	if jd.logger == nil {
		jd.logger = logger.NewLogger().Child("jobsdb").Child(jd.tablePrefix)
	}
	jd.dsRangeFuncMap = make(map[string]func() (dsRangeMinMax, error))
	jd.distinctValuesCache = NewDistinctValuesCache()

	if jd.config == nil {
		jd.config = config.Default
	}

	if string(jd.conf.payloadColumnType) == "" {
		jd.conf.payloadColumnType = TEXT
	}

	if jd.stats == nil {
		jd.stats = stats.Default
	}

	jd.loadConfig()

	// Initialize dbHandle if not already set
	if jd.dbHandle != nil {
		jd.sharedConnectionPool = true
	} else {
		var err error
		psqlInfo := misc.GetConnectionString(jd.config, "jobsdb_"+jd.tablePrefix)
		jd.dbHandle, err = sql.Open("postgres", psqlInfo)
		jd.assertError(err)

		jd.assertError(
			jd.stats.RegisterCollector(
				collectors.NewDatabaseSQLStats(
					"jobsdb_"+jd.tablePrefix+"_"+jd.ownerType.Identifier(),
					jd.dbHandle,
				),
			),
		)

		var maxConns int
		if !jd.conf.enableReaderQueue || !jd.conf.enableWriterQueue {
			maxConns = jd.conf.maxOpenConnections
		} else {
			maxConns = 2 // buffer
			maxConns += jd.conf.maxReaders + jd.conf.maxWriters
			switch jd.ownerType {
			case Read:
				maxConns += 3 // migrate, refreshDsList, dropDS
			case Write:
				maxConns += 1 // addNewDS
			case ReadWrite:
				maxConns += 3 // migrate, addNewDS, dropDS
			}
			if maxConns >= jd.conf.maxOpenConnections {
				maxConns = jd.conf.maxOpenConnections
			}
		}
		jd.dbHandle.SetMaxOpenConns(maxConns)

		jd.assertError(jd.dbHandle.Ping())
	}

	jd.workersAndAuxSetup()

	err := jd.WithTx(context.Background(), func(tx *Tx) error {
		// only one migration should run at a time and block all other processes from adding or removing tables
		return jd.withDistributedLock(context.Background(), tx, "schema_migrate", func() error {
			// Database schema migration should happen early, even before jobsdb is started,
			// so that we can be sure that all the necessary tables are created and considered to be in
			// the latest schema version, before rudder-migrator starts introducing new tables.
			jd.dsListLock.WithLock(func(l lock.LockToken) {
				writer := jd.ownerType == Write || jd.ownerType == ReadWrite
				if writer && jd.conf.clearAll {
					jd.dropDatabaseTables(l)
				}
				templateData := func() map[string]any {
					// Important: if jobsdb type is acting as a writer then refreshDSList
					// doesn't return the full list of datasets, only the rightmost two.
					// But we need to run the schema migration against all datasets, no matter
					// whether jobsdb is a writer or not.
					datasets, err := getDSList(jd, jd.dbHandle, jd.tablePrefix)
					jd.assertError(err)

					datasetIndices := make([]string, 0)
					for _, dataset := range datasets {
						datasetIndices = append(datasetIndices, dataset.Index)
					}

					return map[string]any{
						"Prefix":              jd.tablePrefix,
						"Datasets":            datasetIndices,
						"PartitioningEnabled": jd.conf.numPartitions > 0,
					}
				}()

				if writer {
					jd.setupDatabaseTables(templateData)
				}

				// Run changesets that should always run for both writer and reader jobsdbs.
				//
				// When running separate gw and processor instances we cannot control the order of execution
				// and we cannot guarantee that after a gw migration completes, processor
				// will not create new tables using the old schema.
				//
				// Changesets that run always can help in such cases, by bringing non-migrated tables into a usable state.
				jd.runAlwaysChangesets(templateData)

				// finally refresh the dataset list to make sure [datasetList] field is populated
				err := jd.doRefreshDSRangeList(l)
				jd.assertError(err)
			})
			return nil
		})
	})
	if err != nil {
		panic(fmt.Errorf("failed to run schema migration for %s: %w", jd.tablePrefix, err))
	}
}

func (jd *Handle) workersAndAuxSetup() { _ = "STUB: not implemented"; return }

func (jd *Handle) loadConfig() {
	_ = "STUB: not implemented"
	// maxTableSizeInMB: Maximum Table size in MB
	return
}

// addNewDSLoopSleepDuration: How often is the loop (which checks for adding new DS) run

// refreshDSListLoopSleepDuration: How often is the loop (which refreshes DSList) run

// migrationConfig

// migrateDSLoopSleepDuration: How often is the loop (which checks for migrating DS) run

// jobStatusMigrateThres: A DS is migrated if the job_status exceeds this (* no_of_jobs)

// jobMinRowsLeftMigrateThreshold: A DS with a low number of pending rows should be eligible for migration if the number of pending rows are
// less than jobMinRowsLeftMigrateThreshold percent of maxDSSize (e.g. if jobMinRowsLeftMigrateThreshold is 0.5
// then DSs that have less than 50% of maxDSSize pending rows are eligible for migration)

// maxMigrateOnce: Maximum number of DSs that are migrated together into one destination

// maxMigrateDSProbe: Maximum number of DSs that are checked from left to right if they are eligible for migration

// masterBackupEnabled = true => all the jobsdb are eligible for backup

// maxDSSize: Maximum size of a DS. The process which adds new DS runs in the background
// (every few seconds) so a DS may go beyond this size
// passing `maxDSSize` by reference, so it can be hot reloaded

// starting with false as default since initial set of migrated jobs will not have partitionID set

// Default false: snapshot lastDS and release the dsList read lock before running the store callback,
// so long-running stores don't block dsList writers. Flip to true to revert to holding the lock for the whole callback.

// when true, the per-state noResultsCache optimization is enabled: stateFilters are narrowed
// against the cache before querying, and (!ok && !limitsReached) is used as a commit predicate.

func (jd *Handle) configKeys(key string, additionalKeys ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the jobsdb worker and housekeeping (migration, archive) threads.
// Start should be called before any other jobsdb methods are called.
func (jd *Handle) Start() error { _ = "STUB: not implemented"; return nil }

func (jd *Handle) setUpForOwnerType(ctx context.Context, ownerType OwnerType) {
	_ = "STUB: not implemented"
	return
}

func (jd *Handle) readerSetup(ctx context.Context, l lock.LockToken) {
	_ = "STUB: not implemented"
	return
}

// This is a thread-safe operation.
// Even if two different services (gateway and processor) perform this operation, there should not be any problem.

func (jd *Handle) writerSetup(ctx context.Context, l lock.LockToken) {
	_ = "STUB: not implemented"
	return
}

// This is a thread-safe operation.
// Even if two different services (gateway and processor) perform this operation, there should not be any problem.

// If no DS present, add one

func (jd *Handle) readerWriterSetup(ctx context.Context, l lock.LockToken) {
	_ = "STUB: not implemented"
	return
}

// Stop stops the background goroutines and waits until they finish.
// Stop should be called once only after Start.
// Only Start and Close can be called after Stop.
func (jd *Handle) Stop() { _ = "STUB: not implemented"; return }

// TearDown stops the background goroutines,
//
//	waits until they finish and closes the database.
func (jd *Handle) TearDown() { _ = "STUB: not implemented"; return }

// Close closes the database connection.
//
//	Stop should be called before Close.
//
//	Noop if the connection pool is shared with the handle.
func (jd *Handle) Close() { _ = "STUB: not implemented"; return }

/*
Utility function to return an ordered list of datasets (for tests)
*/
func (jd *Handle) getDSListSnapshot() dataSetTList {
	_ = "STUB: not implemented"
	return *new(dataSetTList)
}

// getLastDS returns the last dataset in the list. Caller must have the dsListLock readlocked
func (jd *Handle) getLastDS() dataSetT { _ = "STUB: not implemented"; return *new(dataSetT) }

// doRefreshDSList refreshes the ds list from the database
func (jd *Handle) doRefreshDSList(l lock.LockToken, db sqlDbOrTx) (dataSetTList, error) {
	_ = "STUB: not implemented"
	return *new(dataSetTList), nil
}

// report table count metrics before shrinking the datasetList

// If the owner of this jobsdb is a writer, then shrinking datasetList to have only last dataset
// which is being written to.
// Writers only write to the last dataset and if this dataset is full, then create a new dataset.

// addCompletedDSToDropList adds the given datasets to the dropDSList and removes them from the dsList. Caller must have the dsListLock write-locked.
// It returns the freshly published dsList snapshot (with the queued datasets filtered out) so callers can avoid an extra read.
func (jd *Handle) addCompletedDSToDropList(ctx context.Context, l lock.LockToken, dsList ...dataSetT) (dataSetTList, error) {
	_ = "STUB: not implemented"
	return *new(dataSetTList), nil
}

func (jd *Handle) dropNotifyPing() { _ = "STUB: not implemented"; return }

func (jd *Handle) doRefreshDSRangeList(l lock.LockToken) error {
	_ = "STUB: not implemented"
	return nil
}

// doRefreshDSRangeList first refreshes the DS list and then calculate the DS range list
func (jd *Handle) doRefreshDSRangeListWithDB(l lock.LockToken, db sqlDbOrTx) error {
	_ = "STUB: not implemented"

	// At this point we must have write-locked dsListLock
	return nil
}

// We store ranges EXCEPT for
// 1. the last element (which is being actively written to)
// 2. Migration target ds

// Skipping asserts and updating prevMax if a ds is found to be empty
// Happens if this function is called between addNewDS and populating data in two scenarios
// Scenario-1: During internal migrations
// Scenario-2: During scaleup scaledown

// acquireDSListForRead returns the dsList and dsRangeList. Caller should call the release function after done reading from the lists.
func (jd *Handle) acquireDSListForRead(ctx context.Context) (
	list dataSetTList, ranges dataSetRangeTList, release func(), err error,
) {
	_ = "STUB: not implemented"
	return *new(dataSetTList), *new(dataSetRangeTList), nil, nil
}

func (jd *Handle) checkIfFullDSInTx(tx *Tx, ds dataSetT) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

/*
Function to add a new dataset. DataSet can be added to the end (e.g when last
becomes full OR in between during migration. DataSets are assigned numbers
monotonically when added  to end. So, with just add to end, numbers would be
like 1,2,3,4, and so on. These are called level0 datasets. And the Index is
called level0 Index
During internal migration, we add datasets in between. In the example above, if we migrate
1 & 2, we would need to create a new DS between 2 & 3. This is assigned the number 2_1.
This is called a level1 dataset and the Index (2_1) is called level1
Index. We may migrate 2_1 into 2_2 and so on so there may be multiple level 1 datasets.

Immediately after creating a level_1 dataset (2_1 above), everything prior to it is
deleted.
Hence, there should NEVER be any requirement for having more than two levels.

There is an exception to this. In case of cross node migration during a scale up/down,
we continue to accept new events in level0 datasets. To maintain the ordering guarantee,
we write the imported jobs to the previous level1 datasets. Now if an internal migration
is to happen on one of the level1 dataset, we have to migrate them to level2 dataset

Eg. When the node has 1, 2, 3, 4 data sets and an import is triggered, new events start
going to 5, 6, 7... so on. And the imported data start going to 4_1, 4_2, 4_3... so on
Now if an internal migration is to happen and we migrate 1, 2, 3, 4, 4_1, we need to
create a newDS between 4_1 and 4_2. This is assigned to 4_1_1, 4_1_2 and so on.
*/

func mapDSToLevel(ds dataSetT) (levelInt int, levelVals []int, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Currently we don't have a scenario where we need more than 3 levels.

func newDataSet(tablePrefix, dsIdx string) dataSetT {
	_ = "STUB: not implemented"
	return *new(dataSetT)
}

func (jd *Handle) addNewDS(ctx context.Context, l lock.LockToken, ds dataSetT) {
	_ = "STUB: not implemented"
	return
}

// NOTE: If addNewDSInTx is directly called, make sure to explicitly call refreshDSRangeList(l) to update the DS list in cache, once transaction has completed.
func (jd *Handle) addNewDSInTx(ctx context.Context, tx *Tx, l lock.LockToken, dsList []dataSetT, ds dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

// Tracking time interval between new ds creations. Hence calling end before start

func (jd *Handle) computeNewIdxForAppend(l lock.LockToken) string {
	_ = "STUB: not implemented"
	return ""
}

func (jd *Handle) doComputeNewIdxForAppend(dList []dataSetT) string {
	_ = "STUB: not implemented"
	return ""
}

// Last one can only be Level0

type transactionHandler interface {
	Exec(string, ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	// If required, add other definitions that are common between *sql.DB and *sql.Tx
	// Never include Commit and Rollback in this interface
	// That ensures that whoever is acting on a transactionHandler can't commit or rollback
	// Only the function that passes *sql.Tx should do the commit or rollback based on the error it receives
}

func (jd *Handle) createDSInTx(ctx context.Context, tx *Tx, newDS dataSetT) error {
	_ = "STUB: not implemented"
	// Mark the start of operation. If we crash somewhere here, we delete the
	// DS being added
	return nil
}

// Create the jobs and job_status tables

func (jd *Handle) createDSTablesInTx(ctx context.Context, tx *Tx, newDS dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) createDSIndicesInTx(ctx context.Context, tx *Tx, newDS dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

// index used for maxDSRetention during migration

func (jd *Handle) setSequenceNumberInTx(tx *Tx, l lock.LockToken, dsList []dataSetT, newDSIdx string) error {
	_ = "STUB: not implemented"
	return nil
}

// Now set the min JobID for the new DS just added to be 1 more than previous max

// GetMaxDSIndex returns max dataset index in the DB
func (jd *Handle) GetMaxDSIndex() (maxDSIndex int64) { _ = "STUB: not implemented"; return 0 }

func (jd *Handle) prepareAndExecStmtInTx(tx *sql.Tx, sqlStatement string) {
	_ = "STUB: not implemented"
	return
}

func (jd *Handle) prepareAndExecStmtInTxAllowMissing(tx *sql.Tx, sqlStatement string) {
	_ = "STUB: not implemented"
	return
}

func (jd *Handle) dropDS(ds dataSetT) error { _ = "STUB: not implemented"; return nil }

func (jd *Handle) dropDSWithCtx(ctx context.Context, ds dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) markPreDropDS(ctx context.Context, dsList ...dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) markPreDropDSInTx(ctx context.Context, tx *Tx, ds dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) cleanupPreDropTables(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// dropDS drops a dataset
func (jd *Handle) dropDSInTx(tx *Tx, ds dataSetT) error { _ = "STUB: not implemented"; return nil }

func (jd *Handle) startDropDSLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

func (jd *Handle) dropDSLoop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait until all operations using this dataset are done.
// This ensures that we don't drop a dataset which is currently being read.

// drop the dataset

// update the lists and cache

// Remove the entry from dropDSList

// delete the entry from dsRangeFuncMap

// Invalidate the distinctValuesCache for the dropped dataset

// If there are more datasets to drop, notify the dropDSLoop to check the next one

// Drop a dataset and ignore if a table is missing
func (jd *Handle) dropDSForRecovery(ds dataSetT) { _ = "STUB: not implemented"; return }

func (jd *Handle) postDropDs(ds dataSetT) { _ = "STUB: not implemented"; return }

// Tracking time interval between drop ds operations. Hence calling end before start

func (jd *Handle) dropAllDS(l lock.LockToken) error { _ = "STUB: not implemented"; return nil }

// Update the lists

func (jd *Handle) internalStoreJobsInTx(ctx context.Context, tx *Tx, ds dataSetT, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) WithStoreSafeTx(ctx context.Context, f func(tx StoreSafeTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) WithStoreSafeTxFromTx(ctx context.Context, tx *Tx, f func(tx StoreSafeTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) inStoreSafeCtx(ctx context.Context, f func(lastDS dataSetT) error) error {
	_ = "STUB: not implemented"
	return nil
}

// The last dataset has already been refreshed, so we can just retry

func (jd *Handle) WithUpdateSafeTx(ctx context.Context, f func(tx UpdateSafeTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) WithUpdateSafeTxFromTx(ctx context.Context, tx *Tx, f func(tx UpdateSafeTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) inUpdateSafeCtx(ctx context.Context, f func(dsList []dataSetT, dsRangeList []dataSetRangeT) error) error {
	_ = "STUB: not implemented"
	// The order of lock is very important. The migrateDSLoop
	// takes lock in this order so reversing this will cause
	// deadlocks
	return nil
}

func (jd *Handle) WithTx(ctx context.Context, f func(tx *Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) invalidateCacheForJobs(ds dataSetT, jobList []*JobT) {
	_ = "STUB: not implemented"
	return
}

// If there is no partition id, use "none" so that cache invalidation doesn't invalidate the whole tree
// No need to log a warning, because partitioning is optional, it is not enabled for all jobsdbs

type moreToken struct {
	afterJobID *int64
}

func (jd *Handle) GetToProcess(ctx context.Context, params GetQueryParams, more MoreToken) (*MoreJobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return nil, nil
}

var cacheParameterFilters = []string{"source_id", "destination_id"}

func (jd *Handle) GetPileUpCounts(ctx context.Context, cutoffTime time.Time, increaseFunc rmetrics.IncreasePendingEventsFunc) error {
	_ = "STUB: not implemented"
	// pause migration to avoid any read locks being blocked during pileup count
	return nil
}

func (jd *Handle) getDistinctValuesPerDataset(
	ctx context.Context,
	dsList []string,
	param ParameterName,
	customVal string,
) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jd *Handle) GetDistinctParameterValues(ctx context.Context, parameter ParameterName, customValFilter string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jd *Handle) doStoreJobsInTx(ctx context.Context, tx *Tx, ds dataSetT, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// Assign partition ID if not already assigned

type JobsResult struct {
	Jobs            []*JobT
	LimitsReached   bool
	DSLimitsReached bool
	EventsCount     int
	PayloadSize     int64
}

/*
stateFilters and customValFilters do a OR query on values passed in array
parameterFilters do a AND query on values included in the map.
A JobsLimit less than or equal to zero indicates no limit.
*/
func (jd *Handle) getJobsDS(ctx context.Context, ds dataSetT, lastDS bool, params GetQueryParams) (JobsResult, bool, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), false, nil
}

// exclude states for which we already know that there are no jobs

// avoid setting result as noJobs if
//  (1) state is unprocessed and
//  (2) jobsdb owner is a reader and
//  (3) ds is the right most one

// excludedReadPartitions are mutually exclusive with partitionFilters
// Use NOT EXISTS against the exclusions table to avoid generating large NOT IN lists and
// improve performance by taking advantage of anti-join optimizations.

// If we are not querying for unprocessed jobs, we can use an inner join

// If we are querying only for unprocessed jobs, we should join with the status table instead of the view (performance reasons)

// If there is a single job in the dataset containing more events than the EventsLimit, we should return it,
// otherwise processing will halt.
// Therefore, we always retrieve one more job from the database than our limit dictates.
// This job will only be returned to the result in case of the aforementioned scenario, otherwise it gets filtered out
// later, during row scanning

// we don't need the payload_size but still need to scan it because it is part of the resultset
// The query uses it for limits checking. The variable is declared before the for loop to avoid extra allocations,
// but if we were to actually use it in the future for returning in the result, we would need to move its declaration
// inside the loop

// events limit overflow is triggered as long as we have read at least one job

// payload size limit overflow is triggered as long as we have read at least one job

// we are adding the job only after testing for limitsReached
// so that we don't always overflow

// we reached the jobs limit
// we reached the events limit
// we reached the payload limit

// we are committing the cache Tx only if
// (a) no jobs are returned by the query or
// (b) the state is not present in the resultset and limits have not been reached
//     (skipped when the noResultsCache state-filter optimization is disabled)

// updateJobStatusStats is a map containing statistics of job status updates grouped by: partition -> workspace -> state -> set of params (stringified) -> stats
type updateJobStatusStats map[partitionIDKey]map[workspaceIDKey]map[customValKey]map[jobStateKey]map[parameterFiltersKey]*UpdateJobStatusStats

// partitionIDKey represents partition id as key
type partitionIDKey string

// workspaceIDKey represents workspace id as key
type workspaceIDKey string

// customValKey represents custom value as key
type customValKey string

// jobStateKey represents job state as key (failed, succeeded, etc)
type jobStateKey string

// parameterFiltersKey represents a list of job parameter filters (stringified) as key
type parameterFiltersKey string

// Merges metrics from two updateJobStatusStats together
func (ujss updateJobStatusStats) Merge(other updateJobStatusStats) {
	_ = "STUB: not implemented"
	return
}

// Aggregates metrics by state across all workspaces
func (ujss updateJobStatusStats) StatsByCustomValAndState() map[customValKey]map[jobStateKey]map[parameterFiltersKey]*UpdateJobStatusStats {
	_ = "STUB: not implemented"
	return nil
}

// Stats for jobs grouped by status and parameters
type UpdateJobStatusStats struct {
	// job parameters
	parameters ParameterFilterList
	// number of jobs
	count int
	// total size of error responses in bytes
	bytes int
}

func (jd *Handle) updateJobStatusDSInTx(ctx context.Context, tx *Tx, ds dataSetT, statusList []*JobStatusT) (updatedStates updateJobStatusStats, err error) {
	_ = "STUB: not implemented"
	return *new(updateJobStatusStats), nil
}

// reset in case of retry

// log a warning if partition id is not set but partitioning is enabled

//  Handle the case when google analytics returns gif in response

/*
The next set of functions are the user visible functions to get/set job status.
For reading jobs, it scans from the oldest DS to the latest till it has found
enough jobs. For updating status, it finds the DS to which the job belongs
(using the in-memory range list) and adds the status to the appropriate DS.
These functions can race with the internal function to add new DS and create
new DS. Synchronization is handled by locks as described below.

In theory, we can keep just one lock. All operations which
change the DS structure (e.g. adding new dataset or moving records
from one DS to another thearby updating the DS range) can take a write lock
while functions which don't update the DS structure (as in list of DS or
ranges within DS can take the read lock) as they can run in paralle.

The drawback with this approach is that migrating a DS can take a long
time and can potentially block the jobs/job-batch store call. Blocking jobs store
is bad since user ACK won't be sent unless jobs store returns.

To handle this, we separate out the locks into dsListLock and dsMigrationLock.
Store() only needs to access the last element of dsList and is not
impacted by movement of data across ds so it only takes the dsListLock.
Other functions are impacted by movement of data across DS in background
so take both the list and data lock
*/
func (jd *Handle) addNewDSLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// Adding a new DS only creates a new DS & updates the cache. It doesn't move any data so we only take the list lock.
// start a transaction

// cannot run while schema migration is running
// only one add_ds can run at a time

// refresh ds list

// make sure we are operating on the latest version of the list

// checkIfFullDS is true for last DS in the list

// We acquire the list lock only after we have acquired the advisory lock.
// We will release the list lock after the transaction ends, that's why we need to use an async lock

// previous DS should become read only

// maybe another node added a new DS that we need to make visible to us

// to get the updated DS list in the cache after createDS transaction has been committed.

func (jd *Handle) getAdvisoryLockForOperation(operation string) int64 {
	_ = "STUB: not implemented"
	return 0
}

func setReadonlyDsInTx(ctx context.Context, tx *Tx, latestDS dataSetT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) refreshDSListLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// RefreshDSList refreshes the list of datasets in memory if the database view of the list has changed.
func (jd *Handle) RefreshDSList(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Identifier returns the identifier of the jobsdb. Here it is tablePrefix.
func (jd *Handle) Identifier() string { _ = "STUB: not implemented"; return "" }

// getDB returns the appropriate database handle based on context.
// If the context requests priority pool usage and a priority pool is configured,
// it returns the priority pool. Otherwise, it returns the regular dbHandle.
func (jd *Handle) getDB(ctx context.Context) *sql.DB { _ = "STUB: not implemented"; return nil }

/*
We keep a journal of all the operations. The journal helps
*/
const (
	addDSOperation             = "ADD_DS"
	migrateCopyOperation       = "MIGRATE_COPY"
	postMigrateDSOperation     = "POST_MIGRATE_DS_OP"
	dropDSOperation            = "DROP_DS"
	RawDataDestUploadOperation = "S3_DEST_UPLOAD"
)

type JournalEntryT struct {
	OpID      int64
	OpType    string
	OpDone    bool
	OpPayload json.RawMessage
}

func (jd *Handle) dropJournal() { _ = "STUB: not implemented"; return }

func (jd *Handle) JournalMarkStart(opType string, opPayload json.RawMessage) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (jd *Handle) JournalMarkStartInTx(tx *Tx, opType string, opPayload json.RawMessage) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// JournalMarkDone marks the end of a journal action
func (jd *Handle) JournalMarkDone(opID int64) error { _ = "STUB: not implemented"; return nil }

// JournalMarkDoneInTx marks the end of a journal action in a transaction
func (jd *Handle) journalMarkDoneInTx(tx *Tx, opID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) JournalDeleteEntry(opID int64) { _ = "STUB: not implemented"; return }

func (jd *Handle) GetJournalEntries(opType string) (entries []JournalEntryT) {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) recoverFromCrash(owner OwnerType, goRoutineType string) {
	_ = "STUB: not implemented"
	return
}

// Nothing to recover

// Need to recover the last failed operation
// Get the payload and undo

// Drop the table we were trying to create

// Delete the destination of the interrupted
// migration. After we start, code should
// redo the migration

const (
	addDSGoRoutine = "addDS"
	mainGoRoutine  = "main"
)

func (jd *Handle) recoverFromJournal(owner OwnerType) { _ = "STUB: not implemented"; return }

func (jd *Handle) UpdateJobStatus(ctx context.Context, statusList []*JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

/*
internalUpdateJobStatusInTx updates the status of a batch of jobs
customValFilters[] is passed, so we can efficiently mark empty cache
Later we can move this to query
*/
func (jd *Handle) internalUpdateJobStatusInTx(ctx context.Context, tx *Tx, dsList []dataSetT, dsRangeList []dataSetRangeT, statusList []*JobStatusT) error {
	_ = "STUB: not implemented"
	// capture stats
	return nil
}

// do update

// clear cache

// if no keys, we need to invalidate all keys

// if no keys, we need to invalidate all keys

// if no keys, we need to invalidate all keys

// if no keys, we need to invalidate all keys

// gather unique parameter filters

// from all JobStatusMetrics

// uniqueness by string representation

// invalidate cache for this combination

// use the aggregated stats from updateJobStatusInTx

/*
doUpdateJobStatusInTx updates the status of a batch of jobs
customValFilters[] is passed, so we can efficiently mark empty cache
Later we can move this to query
*/
func (jd *Handle) doUpdateJobStatusInTx(ctx context.Context, tx *Tx, dsList []dataSetT, dsRangeList []dataSetRangeT, statusList []*JobStatusT) (updatedStatesByDS map[dataSetT]updateJobStatusStats, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First we sort by JobID

// We scan through the list of jobs and map them to DS

// We have processed upto (but excluding) lastPos on statusList.
// Hence, that element must lie in this or subsequent dataset's
// range

// The JobID is outside this DS's range

// do not set for ds without any new state written as it would clear emptyCache

// Reached the end. Need to process this range

// do not set for ds without any new state written as it would clear emptyCache

// The last (most active DS) might not have range element as it is being written to

// Make sure range is missing for the last ds and migration ds (if at all present)

// Update status in the last element

// do not set for ds without any new state written as it would clear emptyCache

// Store stores new jobs to the jobsdb.
// If enableWriterQueue is true, this goes through writer worker pool.
func (jd *Handle) Store(ctx context.Context, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// StoreInTx stores new jobs to the jobsdb.
// If enableWriterQueue is true, this goes through writer worker pool.
func (jd *Handle) StoreInTx(ctx context.Context, tx StoreSafeTx, jobList []*JobT) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) StoreEachBatchRetry(
	ctx context.Context,
	jobBatches [][]*JobT,
) map[uuid.UUID]string {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) StoreEachBatchRetryInTx(
	ctx context.Context,
	tx StoreSafeTx,
	jobBatches [][]*JobT,
) (map[uuid.UUID]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jd *Handle) internalStoreEachBatchRetryInTx(ctx context.Context, tx *Tx, ds dataSetT, jobBatches [][]*JobT) (errorMessagesMap map[uuid.UUID]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retry storing each batch separately

// stop trying treat all remaining as failed

// savepoint

// rollback to savepoint

/*
printLists is a debugging function used to print
the current in-memory copy of jobs and job ranges
*/
func (jd *Handle) printLists(console bool) { _ = "STUB: not implemented"; return }

// This being an internal function, we don't lock

// GetUnprocessed finds unprocessed jobs, i.e. new jobs whose state hasn't been marked in the database yet
func (jd *Handle) GetUnprocessed(ctx context.Context, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

// GetImporting finds jobs in importing state
func (jd *Handle) GetImporting(ctx context.Context, params GetQueryParams, more MoreToken) (*MoreJobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	// Importing jobs are not to be migrated, they are a special case of [executing] jobs, which get queried by
	// batchrouter's async handler periodically to check for the status of imports. Eventually, these jobs will
	// transition to [succeeded]/[failed]/[aborted] state once the import is complete.
	// Hence, we ignore read partition exclusions when fetching importing jobs, so that we can always
	// find all of them regardless of any partition exclusions.
	return nil, nil
}

// GetAborted finds jobs in aborted state
func (jd *Handle) GetAborted(ctx context.Context, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

// GetWaiting finds jobs in waiting state
func (jd *Handle) GetWaiting(ctx context.Context, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

// GetSucceeded finds jobs in succeeded state
func (jd *Handle) GetSucceeded(ctx context.Context, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

// GetFailed finds jobs in failed state
func (jd *Handle) GetFailed(ctx context.Context, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

/*
getJobs returns events of a given state. This does not update any state itself and
realises on the caller to update it. That means that successive calls to getJobs("failed")
can return the same set of events. It is the responsibility of the caller to call it from
one thread, update the state (to "waiting") in the same thread and pass on the processors
*/
func (jd *Handle) getJobs(ctx context.Context, params GetQueryParams, more MoreToken) (*MoreJobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return nil, nil
}

// The order of lock is very important. The migrateDSLoop
// takes lock in this order so reversing this will cause
// deadlocks

// ranges are not stored for the last ds
// so the following condition cannot be applied the last ds

// skip this ds as afterJobID is beyond this ds's range

// we have reached the ds where afterJobID would be present
// so clear it for next ds

// decrement our limits for the next query

// number of actual ds tables that we queried
// number of ds tables that we skipped querying due to noResultsCache

// number of times that we queried and got no jobs

// number of jobs that we queried
// number of bytes that we queried

/*
GetJobs returns events of a given state. This does not update any state itself and
realises on the caller to update it. That means that successive calls to GetJobs("failed")
can return the same set of events. It is the responsibility of the caller to call it from
one thread, update the state (to "waiting") in the same thread and pass on the processors
*/
func (jd *Handle) GetJobs(ctx context.Context, states []string, params GetQueryParams) (JobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return *new(JobsResult), nil
}

func (jd *Handle) getMoreJobs(ctx context.Context, states []string, params GetQueryParams, more MoreToken) (*MoreJobsResult, error) {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return nil, nil
}

type queryResult struct {
	JobsResult
	err error
}

func queryResultWrapper(res *MoreJobsResult, err error) queryResult {
	_ = "STUB: not implemented"
	return *new(queryResult)
}

type moreQueryResult struct {
	*MoreJobsResult
	err error
}

func moreQueryResultWrapper(res *MoreJobsResult, err error) moreQueryResult {
	_ = "STUB: not implemented"
	return *new(moreQueryResult)
}

func (jd *Handle) withDistributedLock(ctx context.Context, tx *Tx, operation string, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (jd *Handle) withDistributedSharedLock(ctx context.Context, tx *Tx, operation string, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

// DefaultParititionFunction is the default function to compute partition key for a job
func DefaultParititionFunction(job *JobT, numPartitions int) string {
	_ = "STUB: not implemented"
	return ""
}
