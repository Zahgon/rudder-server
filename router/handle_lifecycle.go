package router

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/throttler"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	"github.com/rudderlabs/rudder-server/services/rsources"
	transformerFeaturesService "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/services/transientsource"
)

// Setup initializes this module
func (rt *Handle) Setup(
	destinationDefinition backendconfig.DestinationDefinitionT,
	log logger.Logger,
	config *config.Config,
	backendConfig backendconfig.BackendConfig,
	jobsDB jobsdb.JobsDB,
	transientSources transientsource.Service,
	rsourcesService rsources.JobService,
	transformerFeaturesService transformerFeaturesService.FeaturesService,
	debugger destinationdebugger.DestinationDebugger,
	throttlerFactory throttler.Factory,
) {
	_ = "STUB: not implemented"
	return
}

// if noOfJobsPerChannel is more than max, set it as the new max

// Explicitly control destination types for which we want to support batching
// Avoiding stale configurations still having KAFKA batching enabled to cause issues with later versions of rudder-server

// periodically publish a zero counter for ensuring that stuck processing pipeline alert
// can always detect a stuck router

func (rt *Handle) setupReloadableVars() { _ = "STUB: not implemented"; return }

func (rt *Handle) Start() { _ = "STUB: not implemented"; return }

// always close the channel

// no-op, just wait

// waiting for transformer features

// proceed

// no-op, just wait

// start the ping loop

func (rt *Handle) Shutdown() { _ = "STUB: not implemented"; return }

// router is not started

// wait for all workers to stop first

// now it is safe to close the response channel

// statusInsertLoop will run in a separate goroutine
// Blocking method, returns when rt.responseQ channel is closed.
func (rt *Handle) statusInsertLoop() { _ = "STUB: not implemented"; return }

func (rt *Handle) backendConfigSubscriber() { _ = "STUB: not implemented"; return }

// Config key "throttlingCost" is expected to have the eventType as the first key and the call type
// as the second key (e.g. track, identify, etc...) or default to apply the cost to all call types:
// dDT["config"]["throttlingCost"] = `{"eventType":{"default":1,"track":2,"identify":3}}`
