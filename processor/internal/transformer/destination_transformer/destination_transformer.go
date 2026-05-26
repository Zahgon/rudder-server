package destination_transformer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	"github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/kafka"
	"github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/pubsub"
	"github.com/rudderlabs/rudder-server/processor/types"
	transformerfs "github.com/rudderlabs/rudder-server/services/transformer"
)

type warehouseClient interface {
	Transform(ctx context.Context, clientEvents []types.TransformerEvent) types.Response
	CompareResponsesAndUpload(ctx context.Context, events []types.TransformerEvent, legacyResponse types.Response)
}

type Opt func(*Client)

func WithClient(client transformerclient.Client) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithFeatureService is used to set the feature service for the transformer client.
// It is used to check if the destination transformer supports compacted payloads.
// If this option is omitted, the transformer client will not be able to use compacted payloads.
func WithFeatureService(featureService transformerfs.FeaturesService) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// Wait for the feature service to be ready

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

		maxLoggedEvents config.ValueLoader[int]

		warehouseTransformations struct {
			enable config.ValueLoader[bool]
			verify config.ValueLoader[bool]
		}
		compactionSupported bool
	}
	conf            *config.Config
	log             logger.Logger
	stat            stats.Stats
	client          transformerclient.Client
	warehouseClient warehouseClient

	stats struct {
		comparisonTime   stats.Timer
		matchedEvents    stats.Counter
		mismatchedEvents stats.Counter
	}

	loggedEvents        atomic.Int64
	samplingFileManager *filemanager.S3Manager
}

func (d *Client) transform(ctx context.Context, clientEvents []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// Transform is one to many mapping so returned
// response for each is an array. We flatten it out

func (d *Client) sendBatch(ctx context.Context, url string, labels types.TransformerMetricLabels, data []types.TransformerEvent) ([]types.TransformerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// consistent state for the entire request
// Call remote transformation

// This is returned by our JS engine so should  be parsable
// Panic the processor to avoid replays

func (d *Client) doPost(ctx context.Context, rawJSON []byte, url string, labels types.TransformerMetricLabels, extraHeaders map[string]string) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Header to let transformer know that the client understands event filter code

// Record metrics with labels

// We'll count response events after unmarshaling in the request method

// perform version compatibility check only on success

func (d *Client) destTransformURL(destType string) string { _ = "STUB: not implemented"; return "" }

type transformer func(ctx context.Context, clientEvents []types.TransformerEvent) types.Response

var embeddedTransformerImpls = map[string]transformer{
	"GOOGLEPUBSUB": pubsub.Transform,
	"KAFKA":        kafka.Transform,
}

func (c *Client) Transform(ctx context.Context, clientEvents []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

func deepCopy[T any](src T) T { _ = "STUB: not implemented"; return *new(T) }

func (c *Client) canRunWarehouseTransformations(destType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *Client) getRequestPayload(data []types.TransformerEvent, compactRequestPayloads bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSamplingUploader(conf *config.Config, log logger.Logger) (*filemanager.S3Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
