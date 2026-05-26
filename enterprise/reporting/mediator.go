package reporting

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/enterprise/reporting/flusher"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	"github.com/rudderlabs/rudder-server/utils/types"
)

const (
	TrackedUsersReportsTable = "tracked_users_reports"
)

type Mediator struct {
	log logger.Logger

	g         *errgroup.Group
	ctx       context.Context
	cancel    context.CancelFunc
	reporters []types.Reporting
	stats     stats.Stats

	cronRunners []flusher.Runner
}

func NewReportingMediator(ctx context.Context, conf *config.Config, log logger.Logger, enterpriseToken string, backendConfig backendconfig.BackendConfig) *Mediator {
	_ = "STUB: not implemented"
	return nil
}

// default reporting implementation

// error reporting implementation

// error index reporting implementation

func (rm *Mediator) Report(ctx context.Context, metrics []*types.PUReportedMetric, txn *Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (rm *Mediator) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	_ = "STUB: not implemented"
	return *new(types.ReportingSyncer)
}

//  TODO: Should we panic here?

func (rm *Mediator) Stop() { _ = "STUB: not implemented"; return }
