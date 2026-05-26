package trackingplan_validation

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	"github.com/rudderlabs/rudder-server/processor/types"
)

type Opt func(*Client)

func WithClient(client transformerclient.Client) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func New(conf *config.Config, log logger.Logger, stat stats.Stats, opts ...Opt) *Client {
	_ = "STUB: not implemented"
	return nil
}

type Client struct {
	config struct {
		destTransformationURL   string
		maxRetry                config.ValueLoader[int]
		failOnError             config.ValueLoader[bool]
		maxRetryBackoffInterval config.ValueLoader[time.Duration]
		timeoutDuration         time.Duration
		batchSize               config.ValueLoader[int]
	}
	conf   *config.Config
	log    logger.Logger
	stat   stats.Stats
	client transformerclient.Client
}

func (t *Client) Validate(ctx context.Context, clientEvents []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// Transform is one to many mapping so returned
// response for each is an array. We flatten it out

func (t *Client) sendBatch(ctx context.Context, url string, labels types.TransformerMetricLabels, clientEvents []types.TransformerEvent) []types.TransformerResponse {
	_ = "STUB: not implemented"
	return nil
}

// endless retry if transformer-control plane connection is down

// endless backoff loop, only nil error or panics inside

// disable max elapsed time --> endless

// control plane back up

// This is returned by our JS engine so should  be parsable
// Panic the processor to avoid replays

func (t *Client) doPost(ctx context.Context, rawJSON []byte, url string, labels types.TransformerMetricLabels) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Header to let transformer know that the client understands event filter code

// Record metrics with labels

// We'll count response events after unmarshaling in the request method

// perform version compatibility check only on success

func (t *Client) trackingPlanValidationURL() string { _ = "STUB: not implemented"; return "" }
