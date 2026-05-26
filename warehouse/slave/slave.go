package slave

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	"github.com/rudderlabs/rudder-server/warehouse/constraints"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
)

type slaveNotifier interface {
	Subscribe(ctx context.Context, workerId string, jobsBufferSize int) <-chan *notifier.ClaimJob
	RunMaintenance(ctx context.Context) error
	UpdateClaim(ctx context.Context, job *notifier.ClaimJob, response *notifier.ClaimJobResponse)
	RefreshClaim(ctx context.Context, jobId int64) error
}

type Slave struct {
	conf               *config.Config
	log                logger.Logger
	stats              stats.Stats
	notifier           slaveNotifier
	bcManager          *bcm.BackendConfigManager
	constraintsManager *constraints.Manager
	encodingFactory    *encoding.Factory

	config struct {
		noOfSlaveWorkerRoutines config.ValueLoader[int]
	}
}

func New(
	conf *config.Config,
	logger logger.Logger,
	stats stats.Stats,
	notifier slaveNotifier,
	bcManager *bcm.BackendConfigManager,
	constraintsManager *constraints.Manager,
	encodingFactory *encoding.Factory,
) *Slave {
	_ = "STUB: not implemented"
	return nil
}

func (s *Slave) SetupSlave(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
