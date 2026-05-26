package repo

import (
	"context"

	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	sourceJobTableName = whutils.WarehouseAsyncJobTable
	sourceJobColumns   = `
		id,
		source_id,
		destination_id,
		status,
		created_at,
		updated_at,
		tablename,
		error,
		async_job_type,
		metadata,
		attempt,
		workspace_id
	`
)

type Source repo

func NewSource(db *sqlmw.DB, opts ...Opt) *Source { _ = "STUB: not implemented"; return nil }

func (s *Source) Insert(ctx context.Context, sourceJobs []model.SourceJob) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Source) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Source) GetToProcess(ctx context.Context, limit int64) ([]model.SourceJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanSourceJobs(rows *sqlmw.Rows) ([]model.SourceJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanSourceJob(scan scanFn, sourceJob *model.SourceJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) GetByJobRunTaskRun(ctx context.Context, jobRunID, taskRunID string) (*model.SourceJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Source) OnUpdateSuccess(ctx context.Context, id int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) OnUpdateFailure(ctx context.Context, id int64, error error, maxAttempt int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) MarkExecuting(ctx context.Context, ids []int64) error {
	_ = "STUB: not implemented"
	return nil
}
