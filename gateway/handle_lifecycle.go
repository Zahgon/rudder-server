package gateway

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-schemas/go/stream"

	"github.com/rudderlabs/rudder-server/app"
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/gateway/throttler"
	"github.com/rudderlabs/rudder-server/jobsdb"
	sourcedebugger "github.com/rudderlabs/rudder-server/services/debugger/source"
	"github.com/rudderlabs/rudder-server/services/rsources"
	"github.com/rudderlabs/rudder-server/services/transformer"
)

type msgToUpload struct {
	payload []byte
	fields  []logger.Field
}

/*
Setup initializes this module:
- Monitors backend config for changes.
- Starts web request batching goroutine, that batches incoming messages.
- Starts web request batch db writer goroutine, that writes incoming batches to JobsDB.
- Starts debugging goroutine that prints gateway stats.

This function will block until backend config is initially received.
*/
func (gw *Handle) Setup(
	ctx context.Context,
	config *config.Config, logger logger.Logger, stat stats.Stats,
	application app.App, backendConfig backendconfig.BackendConfig, jobsDB jobsdb.JobsDB,
	rateLimiter throttler.Throttler, versionHandler func(w http.ResponseWriter, r *http.Request),
	rsourcesService rsources.JobService, transformerFeaturesService transformer.FeaturesService,
	sourcehandle sourcedebugger.SourceDebugger, streamMsgValidator func(message *stream.Message) error,
	opts ...OptFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Port where GW is running

// Number of incoming requests that are batched before handing off to write workers

// Number of userWorkerBatchRequest that are batched before initiating write

// Multiple workers are used to batch user web requests

// Multiple DB writers are used to write data to DB

// Timeout after which batch is formed anyway with whatever requests are available

// Enables accepting requests without user id and anonymous id. This is added to prevent client 4xx retries.

// Maximum request size to gateway

// Enable rate limit on incoming events. false by default

// Enable suppress user feature. false by default

// Time period for diagnosis ticker

// if set to '0', it means disabled.

// enable webhook v2 handler. disabled by default

// Registering stats

// new bg ctx for leaky logger
// we don't want to cancel the main context.

func getLeakyUploaderFileManager(conf *config.Config, log logger.Logger) (filemanager.FileManager, error) {
	_ = "STUB: not implemented"
	return *new(filemanager.FileManager), nil
}

func leakyUploader(ctx context.Context, conf *config.Config, log logger.Logger, done chan struct{}, uploads <-chan msgToUpload, fm filemanager.FileManager) {
	_ = "STUB: not implemented"
	return
}

type OptFunc func(*Handle)

func WithInternalHttpHandlers(handlers map[string]http.Handler) OptFunc {
	_ = "STUB: not implemented"
	return *new(OptFunc)
}

func WithNow(now func() time.Time) OptFunc { _ = "STUB: not implemented"; return *new(OptFunc) }

// initUserWebRequestWorkers initiates `maxUserWebRequestWorkerProcess` number of `webRequestWorkers` that listen on their `webRequestQ` for new WebRequests.
func (gw *Handle) initUserWebRequestWorkers() { _ = "STUB: not implemented"; return }

// processBackendConfig processes backend config data and updates internal maps
func (gw *Handle) processBackendConfig(configData map[string]backendconfig.ConfigT) {
	_ = "STUB: not implemented"
	return
}

// backendConfigSubscriber gets the config from config backend and extracts source information from it.
func (gw *Handle) backendConfigSubscriber(ctx context.Context) { _ = "STUB: not implemented"; return }

// runUserWebRequestWorkers starts two goroutines for each worker:
//  1. `userWebRequestBatcher` batches the webRequests that a worker gets
//  2. `userWebRequestWorkerProcess` processes the requests in the batches and sends them as part of a `jobsList` to `dbWriterWorker`s.
func (gw *Handle) runUserWebRequestWorkers(ctx context.Context) { _ = "STUB: not implemented"; return }

// initDBWriterWorkers initiates `maxDBWriterProcess` number of dbWriterWorkers
func (gw *Handle) initDBWriterWorkers(ctx context.Context) { _ = "STUB: not implemented"; return }

//	userWorkerRequestBatcher batches together jobLists received on the `userWorkerBatchRequestQ` channel of the gateway
//	and queues the batch at the `batchUserWorkerBatchRequestQ` channel of the gateway.
//
// Initiated during the gateway Setup and keeps batching jobLists received from webRequestWorkers
func (gw *Handle) userWorkerRequestBatcher() { _ = "STUB: not implemented"; return }

// Append to request buffer

// dbWriterWorkerProcess goes over the batches of jobs-list, and stores each job in every jobList into gw_db
// sends a map of errors if any(errors mapped to the job.uuid) over the responseQ channel of the webRequestWorker.
// userWebRequestWorkerProcess method of the webRequestWorker is waiting for this errorMessageMap.
// This in turn sends the error over the done channel of each respective webRequest.
func (gw *Handle) dbWriterWorkerProcess() { _ = "STUB: not implemented"; return }

// rsources stats

/*
StartWebHandler starts all gateway web handlers, listening on gateway port.
Supports CORS from all origins. This function will block.
*/
func (gw *Handle) StartWebHandler(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// rudder-sources new APIs

// TODO: delete this handler once we are ready to remove support for the v1 api

// TODO: delete this handler once we are ready to remove support for the v1 api

// 15 mins

// Shutdown the gateway
func (gw *Handle) Shutdown() error { _ = "STUB: not implemented"; return nil }

// UserWebRequestWorkers
