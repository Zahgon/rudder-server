package forwarder

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/internal/pulsar"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/schema-forwarder/internal/transformer"
)

// JobsForwarder is a forwarder that transforms and forwards jobs to a pulsar topic
type JobsForwarder struct {
	BaseForwarder

	transformer    transformer.Transformer // transformer used to transform jobs to destination schema
	sampler        *sampler[string]        // sampler used to decide how often payloads should be sampled
	pulsarClient   *pulsar.Client          // pulsar client used to create producers
	pulsarProducer pulsar.ProducerAdapter  // pulsar producer used to forward jobs to pulsar

	topic                string                            // topic to which jobs are forwarded
	maxSampleSize        config.ValueLoader[int64]         // max payload size for the samples to include in the schema messages
	initialRetryInterval config.ValueLoader[time.Duration] // initial retry interval for the backoff mechanism
	maxRetryInterval     config.ValueLoader[time.Duration] // max retry interval for the backoff mechanism
	maxRetryElapsedTime  config.ValueLoader[time.Duration] // max retry elapsed time for the backoff mechanism
}

// NewJobsForwarder returns a new, properly initialized, JobsForwarder
func NewJobsForwarder(terminalErrFn func(error), schemaDB jobsdb.JobsDB, client *pulsar.Client, config *config.Config, backendConfig backendconfig.BackendConfig, log logger.Logger, stat stats.Stats) *JobsForwarder {
	_ = "STUB: not implemented"
	return nil
}

func (jf *JobsForwarder) setupReloadableVars() { _ = "STUB: not implemented"; return }

// Start starts the forwarder which will start forwarding jobs from database to the appropriate pulsar topics
func (jf *JobsForwarder) Start() error { _ = "STUB: not implemented"; return nil }

// we are shutting down
//nolint:nilerr

// we are signaling to shut down the app

// protects statuses and toRetry

// mark job as aborted

// if the sample is not an empty json object (redacted) and
// sample is too big or
// sample should be skipped
// by setting to sample to nil we are instructing the schema worker to keep the previous sample

// try to forward messageBatches to pulsar. Succeeded jobs are removed from messageBatches

// we are shutting down
//nolint:nilerr

// mark job as succeeded and remove from toRetry

// Retry to forward the batches to pulsar until there are no more messageBatches to retry or until maxRetryElapsedTime is reached

// mark all messageBatches left over as aborted

// we are shutting down
//nolint:nilerr

// we are signaling to shutdown the app

// Stop stops the forwarder
func (jf *JobsForwarder) Stop() { _ = "STUB: not implemented"; return }
