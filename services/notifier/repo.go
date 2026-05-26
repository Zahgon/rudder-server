package notifier

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"

	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
)

const (
	notifierTableName    = "pg_notifier_queue"
	notifierTableColumns = `
		id,
		batch_id,
		worker_id,
		workspace,
		attempt,
		status,
		job_type,
		priority,
		error,
		payload,
		created_at,
		updated_at,
		last_exec_time
`
)

type Opt func(*repo)

type scanFn func(dest ...any) error

func WithNow(now func() time.Time) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithStats(s stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }

type repo struct {
	db           *sqlmw.DB
	now          func() time.Time
	statsFactory stats.Stats
}

func newRepo(db *sqlmw.DB, opts ...Opt) *repo { _ = "STUB: not implemented"; return nil }

// ResetForWorkspace deletes all the jobs for a specified workspace.
func (n *repo) resetForWorkspace(
	ctx context.Context,
	workspaceIdentifier string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Insert inserts a jobs into the notifier queue.
func (n *repo) insert(
	ctx context.Context,
	publishRequest *PublishRequest,
	workspaceIdentifier string,
	batchID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PendingByBatchID returns the number of pending jobs for a batchID.
func (n *repo) pendingByBatchID(
	ctx context.Context,
	batchID string,
) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetByBatchID returns all the jobs for a batchID.
// TODO: ATM Hack to remove `UploadSchema` from the payload to have the similar implementation as the old notifier.
func (n *repo) getByBatchID(
	ctx context.Context,
	batchID string,
) ([]Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanJob(scan scanFn, job *Job) error { _ = "STUB: not implemented"; return nil }

// DeleteByBatchID deletes all the jobs for a batchID.
func (n *repo) deleteByBatchID(
	ctx context.Context,
	batchID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *repo) claim(
	ctx context.Context,
	workerID string,
) (*Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OnClaimFailed updates the status of a job to failed.
func (n *repo) onClaimFailed(
	ctx context.Context,
	job *Job,
	claimError error,
	maxAttempt int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnClaimSuccess updates the status of a job to succeed.
func (n *repo) onClaimSuccess(
	ctx context.Context,
	job *Job,
	payload json.RawMessage,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OrphanJobIDs returns the IDs of the jobs that are in executing state for more than the given interval.
func (n *repo) orphanJobIDs(
	ctx context.Context,
	intervalInSeconds int,
) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *repo) refreshClaim(ctx context.Context, jobId int64) error {
	_ = "STUB: not implemented"
	return nil
}

// timerStat returns a function that records the duration of a database action.
func (n *repo) timerStat(action string, extraTags stats.Tags) func() {
	_ = "STUB: not implemented"
	return nil
}
