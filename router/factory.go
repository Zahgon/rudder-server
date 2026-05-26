package router

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/throttler"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	"github.com/rudderlabs/rudder-server/services/rsources"
	transformerFeaturesService "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/services/transientsource"
	. "github.com/rudderlabs/rudder-server/utils/tx" //nolint:staticcheck
	utilTypes "github.com/rudderlabs/rudder-server/utils/types"
)

type Factory struct {
	Logger                     logger.Logger
	Reporting                  reporter
	BackendConfig              backendconfig.BackendConfig
	RouterDB                   jobsdb.JobsDB
	TransientSources           transientsource.Service
	RsourcesService            rsources.JobService
	TransformerFeaturesService transformerFeaturesService.FeaturesService
	ThrottlerFactory           throttler.Factory
	Debugger                   destinationdebugger.DestinationDebugger
	AdaptiveLimit              func(int64) int64
}

func (f *Factory) New(destination *backendconfig.DestinationT) *Handle {
	_ = "STUB: not implemented"
	return nil
}

type reporter interface {
	Report(ctx context.Context, metrics []*utilTypes.PUReportedMetric, txn *Tx) error
}
