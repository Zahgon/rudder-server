package apphandlers

import (
	"context"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/app"
	"github.com/rudderlabs/rudder-server/app/cluster"
	"github.com/rudderlabs/rudder-server/internal/enricher"
	"github.com/rudderlabs/rudder-server/services/rsources"
	"github.com/rudderlabs/rudder-server/utils/types/deployment"
)

// AppHandler starts the app
type AppHandler interface {
	// Setup to be called only once before starting the app.
	Setup() error
	// StartRudderCore starts the app
	StartRudderCore(context.Context, func(), *app.Options) error
}

func GetAppHandler(application app.App, appType string, versionHandler func(w http.ResponseWriter, r *http.Request)) (AppHandler, error) {
	_ = "STUB: not implemented"
	return *new(AppHandler), nil
}

func rudderCoreDBValidator() error { _ = "STUB: not implemented"; return nil }

func rudderCoreNodeSetup() error { _ = "STUB: not implemented"; return nil }

// NewRsourcesService produces a rsources.JobService through environment configuration (env variables & config file)
func NewRsourcesService(ctx context.Context, deploymentType deployment.Type, shouldSetupSharedDB bool, stats stats.Stats) (rsources.JobService, error) {
	_ = "STUB: not implemented"
	return *new(rsources.JobService), nil
}

// For multitenant deployment type we shall require the existence of a SHARED_DB
// TODO: change default value of Rsources.FailOnMissingSharedDB to true, when shared DB is provisioned

func resolveModeProvider(log logger.Logger, deploymentType deployment.Type) (cluster.ChangeEventProvider, error) {
	_ = "STUB: not implemented"
	return *new(cluster.ChangeEventProvider), nil
}

// FIXME: hacky way to determine server mode

// terminalErrorFunction returns a function that cancels the errgroup g with an error when the returned function is called.
func terminalErrorFunction(ctx context.Context, g *errgroup.Group) func(error) {
	_ = "STUB: not implemented"
	return nil
}

func setupPipelineEnrichers(conf *config.Config, log logger.Logger, stats stats.Stats) ([]enricher.PipelineEnricher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
