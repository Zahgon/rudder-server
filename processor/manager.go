package processor

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/enterprise/trackedusers"
	"github.com/rudderlabs/rudder-server/internal/enricher"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/processor/transformer"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	transformationdebugger "github.com/rudderlabs/rudder-server/services/debugger/transformation"
	"github.com/rudderlabs/rudder-server/services/fileuploader"
	"github.com/rudderlabs/rudder-server/services/rmetrics"
	"github.com/rudderlabs/rudder-server/services/rsources"
	transformerFeaturesService "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/services/transientsource"
	"github.com/rudderlabs/rudder-server/utils/types"
)

type LifecycleManager struct {
	Handle                     *Handle
	mainCtx                    context.Context
	currentCancel              context.CancelFunc
	waitGroup                  interface{ Wait() }
	gatewayDB                  jobsdb.JobsDB
	routerDB                   jobsdb.JobsDB
	batchRouterDB              jobsdb.JobsDB
	esDB                       jobsdb.JobsDB
	arcDB                      jobsdb.JobsDB
	clearDB                    *bool
	ReportingI                 types.Reporting // need not initialize again
	BackendConfig              backendconfig.BackendConfig
	TransformerClients         *transformer.Clients
	transientSources           transientsource.Service
	fileuploader               fileuploader.Provider
	rsourcesService            rsources.JobService
	transformerFeaturesService transformerFeaturesService.FeaturesService
	destDebugger               destinationdebugger.DestinationDebugger
	transDebugger              transformationdebugger.TransformationDebugger
	enrichers                  []enricher.PipelineEnricher
	trackedUsersReporter       trackedusers.UsersReporter
	pendingEventsRegistry      rmetrics.PendingEventsRegistry
}

// Start starts a processor, this is not a blocking call.
// If the processor is not completely started and the data started coming then also it will not be problematic as we
// are assuming that the DBs will be up.
func (proc *LifecycleManager) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stops the processor, this is a blocking call.
func (proc *LifecycleManager) Stop() { _ = "STUB: not implemented"; return }

// New creates a new Processor instance
func New(
	ctx context.Context,
	clearDb *bool,
	gwDb, rtDb, brtDb, esDB, arcDB jobsdb.JobsDB,
	reporting types.Reporting,
	transientSources transientsource.Service,
	fileuploader fileuploader.Provider,
	rsourcesService rsources.JobService,
	transformerFeaturesService transformerFeaturesService.FeaturesService,
	destDebugger destinationdebugger.DestinationDebugger,
	transDebugger transformationdebugger.TransformationDebugger,
	enrichers []enricher.PipelineEnricher,
	trackedUsersReporter trackedusers.UsersReporter,
	pendingEventsRegistry rmetrics.PendingEventsRegistry,
	opts ...Opts,
) *LifecycleManager {
	_ = "STUB: not implemented"
	return nil
}

type Opts func(l *LifecycleManager)

func WithAdaptiveLimit(adaptiveLimitFunction func(int64) int64) Opts {
	_ = "STUB: not implemented"
	return *new(Opts)
}

func WithStats(stats stats.Stats) Opts { _ = "STUB: not implemented"; return *new(Opts) }

func WithTransformerClients(transformerClients transformer.TransformerClients) Opts {
	_ = "STUB: not implemented"
	return *new(Opts)
}
