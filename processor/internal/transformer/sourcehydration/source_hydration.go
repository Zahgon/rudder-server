package sourcehydration

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	"github.com/rudderlabs/rudder-server/processor/types"
)

const srcHydrationStage = "source_hydration"

type Opt func(*Client)

func WithClient(client transformerclient.Client) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Client handles source hydration transformations
type Client struct {
	config struct {
		sourceHydrationURL      string
		maxRetry                config.ValueLoader[int]
		failOnError             config.ValueLoader[bool]
		maxRetryBackoffInterval config.ValueLoader[time.Duration]
		// Maximum time to wait before stopping retries
		maxEscapedTimeIncludingRetries config.ValueLoader[time.Duration]
		logLongRunningTransformAfter   time.Duration
		batchSize                      config.ValueLoader[int]
	}
	conf   *config.Config
	log    logger.Logger
	stat   stats.Stats
	client transformerclient.Client
}

// New creates a new source hydration client
func New(conf *config.Config, log logger.Logger, stat stats.Stats, opts ...Opt) *Client {
	_ = "STUB: not implemented"
	return nil
}

// Hydrate sends a batch of source events to the hydration endpoint and returns the hydrated events.
// It splits the input events into smaller batches (based on configured batch size), sends them
// concurrently for hydration, and aggregates the results.
//
// Error conditions:
//  1. Returns an error if the context is canceled.
//  2. Returns an error if all retry attempts fail and the client is configured to fail on errors.
//  3. Otherwise, if failOnError is disabled, a panic is raised.
func (c *Client) Hydrate(ctx context.Context, hydrationReq types.SrcHydrationRequest) (types.SrcHydrationResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.SrcHydrationResponse), nil
}

// Add all events from the batch to the result

func (c *Client) sendBatch(ctx context.Context, url string, source types.SrcHydrationSource, labels types.TransformerMetricLabels, hydrationEvents []types.SrcHydrationEvent) (types.SrcHydrationResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.SrcHydrationResponse), nil
}

// Create request in the format expected by the source hydration API

// Marshal request

func (c *Client) doPost(ctx context.Context, rawJSON []byte, url string, labels types.TransformerMetricLabels) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Record metrics with labels

func (c *Client) sourceHydrationURL(sourceType string) string {
	_ = "STUB: not implemented"
	// Based on the OpenAPI spec: /{version}/sources/{source}/hydrate
	return ""
}
