package notifier

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
)

const (
	queueName = "pg_notifier_queue"
	module    = "pgnotifier"
)

type JobType string

const (
	JobTypeUploadV2 JobType = "upload_v2"
	JobTypeAsync    JobType = "async_job"
)

type JobStatus string

const (
	Waiting   JobStatus = "waiting"
	Executing JobStatus = "executing"
	Succeeded JobStatus = "succeeded"
	Failed    JobStatus = "failed"
	Aborted   JobStatus = "aborted"
)

type Job struct {
	ID                  int64
	BatchID             string
	WorkerID            string
	WorkspaceIdentifier string

	Attempt  int
	Status   JobStatus
	Type     JobType
	Priority int
	Error    error

	Payload json.RawMessage

	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastExecTime time.Time
}

type PublishRequest struct {
	Payloads     []json.RawMessage
	UploadSchema json.RawMessage // ATM Hack to support merging schema with the payload at the postgres level
	JobType      JobType
	Priority     int
}

type PublishResponse struct {
	Jobs []Job
	Err  error
}

type ClaimJob struct {
	Job *Job
}

type ClaimJobResponse struct {
	Payload json.RawMessage
	Err     error
}

type notifierRepo interface {
	resetForWorkspace(context.Context, string) error
	insert(context.Context, *PublishRequest, string, string) error
	pendingByBatchID(context.Context, string) (int64, error)
	deleteByBatchID(context.Context, string) error
	orphanJobIDs(context.Context, int) ([]int64, error)
	getByBatchID(context.Context, string) ([]Job, error)
	claim(context.Context, string) (*Job, error)
	onClaimFailed(context.Context, *Job, error, int) error
	onClaimSuccess(context.Context, *Job, json.RawMessage) error
	refreshClaim(context.Context, int64) error
}

type Notifier struct {
	conf                *config.Config
	logger              logger.Logger
	statsFactory        stats.Stats
	db                  *sqlmw.DB
	repo                notifierRepo
	workspaceIdentifier string
	batchIDGenerator    func() uuid.UUID
	randGenerator       *rand.Rand
	now                 func() time.Time
	canRunMigrations    bool
	background          struct {
		group       *errgroup.Group
		groupCtx    context.Context
		groupCancel context.CancelFunc
		groupWait   func() error
	}

	config struct {
		host                       string
		port                       int
		user                       string
		password                   string
		database                   string
		sslMode                    string
		maxAttempt                 int
		maxOpenConnections         int
		shouldForceSetLowerVersion bool
		trackBatchInterval         time.Duration
		maxPollSleep               config.ValueLoader[time.Duration]
		jobOrphanTimeout           config.ValueLoader[time.Duration]
		queryTimeout               time.Duration
	}
	stats struct {
		insertRecords      stats.Counter
		publish            stats.Counter
		publishTime        stats.Timer
		claimSucceeded     stats.Counter
		claimSucceededTime stats.Timer
		claimFailed        stats.Counter
		claimFailedTime    stats.Timer
		claimUpdateFailed  stats.Counter
		abortedRecords     stats.Counter
	}
}

func New(
	conf *config.Config,
	log logger.Logger,
	statsFactory stats.Stats,
	workspaceIdentifier string,
	canRunMigrations bool,
) *Notifier {
	_ = "STUB: not implemented"
	return nil
}

func (n *Notifier) Setup(
	ctx context.Context,
	fallbackDSN string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Notifier) checkForNotifierEnvVars() bool { _ = "STUB: not implemented"; return false }

func (n *Notifier) connectionString() string { _ = "STUB: not implemented"; return "" }

func (n *Notifier) setupDatabase(
	ctx context.Context,
	dsn string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Notifier) setupTables() error { _ = "STUB: not implemented"; return nil }

func (n *Notifier) migrate() error { _ = "STUB: not implemented"; return nil }

func (n *Notifier) migrateAlways() error { _ = "STUB: not implemented"; return nil }

// ClearJobs deletes all jobs for the current workspace.
func (n *Notifier) ClearJobs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *Notifier) CheckHealth(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// Publish inserts the payloads into the database and returns a channel of type PublishResponse
func (n *Notifier) Publish(
	ctx context.Context,
	publishRequest *PublishRequest,
) (<-chan *PublishResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// trackBatch tracks the batch and returns a channel of type PublishResponse
func (n *Notifier) trackBatch(
	ctx context.Context,
	batchID string,
) <-chan *PublishResponse {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe returns a channel of type Job
func (n *Notifier) Subscribe(
	ctx context.Context,
	workerId string,
	bufferSize int,
) <-chan *ClaimJob {
	_ = "STUB: not implemented"
	return nil
}

// Claim claims a job from the notifier queue
func (n *Notifier) claim(
	ctx context.Context,
	workerID string,
) (*Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateClaim updates the notifier with the claimResponse
// In case if we are not able to update the claim, we are just logging it,
// maintenance workers can again mark the status as waiting, and it will be again claimed by somebody else.
// Although, there is a case that it is being picked up, but never getting updated. We can monitor it using claim lag.
// claim lag also helps us to make sure that even the maintenance workers are able to monitor the jobs correctly.
func (n *Notifier) UpdateClaim(
	ctx context.Context,
	claimedJob *ClaimJob,
	response *ClaimJobResponse,
) {
	_ = "STUB: not implemented"
	return
}

// RunMaintenance re-triggers zombie jobs which were left behind by dead workers in executing state
// Since it's a blocking call, it should be run in a separate goroutine
func (n *Notifier) RunMaintenance(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Shutdown waits for all the background jobs to be drained off.
func (n *Notifier) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (n *Notifier) RefreshClaim(ctx context.Context, jobId int64) error {
	_ = "STUB: not implemented"
	return nil
}
