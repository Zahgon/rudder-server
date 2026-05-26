package batchrouter

import (
	"github.com/rudderlabs/rudder-go-kit/config"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	"github.com/rudderlabs/rudder-server/services/rsources"
	"github.com/rudderlabs/rudder-server/services/transientsource"
	"github.com/rudderlabs/rudder-server/utils/types"
)

// Setup initializes the batch router
func (brt *Handle) Setup(
	destType string,
	backendConfig backendconfig.BackendConfig,
	jobsDB jobsdb.JobsDB,
	reporting types.Reporting,
	transientSources transientsource.Service,
	rsourcesService rsources.JobService,
	debugger destinationdebugger.DestinationDebugger,
	conf *config.Config,
) {
	_ = "STUB: not implemented"
	return
}

// periodically publish a zero counter for ensuring that stuck processing pipeline alert
// can always detect a stuck batch router

func (brt *Handle) setupReloadableVars() { _ = "STUB: not implemented"; return }

func (brt *Handle) startAsyncDestinationManager() { _ = "STUB: not implemented"; return }

// Start starts the batch router's main loop
func (brt *Handle) Start() { _ = "STUB: not implemented"; return }

// Shutdown stops the batch router
func (brt *Handle) Shutdown() {
	_ = "STUB: not implemented"
	// Signal all goroutines to stop via context cancellation
	return
}

// Wait for all background goroutines to complete

func (brt *Handle) initAsyncDestinationStruct(destination *backendconfig.DestinationT) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) refreshDestination(destination backendconfig.DestinationT) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) crashRecover() { _ = "STUB: not implemented"; return }

// Backward compatibility. If old entries dont have config, just delete journal entry

func (brt *Handle) backendConfigSubscriber() { _ = "STUB: not implemented"; return }

// initialize map to track encountered anonymousIds for a warehouse destination
