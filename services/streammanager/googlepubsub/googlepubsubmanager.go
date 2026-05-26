package googlepubsub

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub/v2"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type PubSubConfig struct {
	Credentials     string              `json:"credentials"`
	ProjectId       string              `json:"projectId"`
	EventToTopicMap []map[string]string `json:"eventToTopicMap"`
	TestConfig      TestConfig          `json:"testConfig"`
}

type TestConfig struct {
	Endpoint string `json:"endpoint"`
}

type PubsubClient struct {
	pbs      *pubsub.Client
	topicMap map[string]*pubsub.Publisher
	opts     common.Opts
}

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("streammanager").Child("googlepubsub")
}

type GooglePubSubProducer struct {
	client *PubsubClient
	conf   *config.Config
	// Retry configuration
	enableRetry          *config.Reloadable[bool]
	retryInitialInterval *config.Reloadable[time.Duration]
	retryMaxInterval     *config.Reloadable[time.Duration]
	retryMaxElapsedTime  *config.Reloadable[time.Duration]
	retryMaxRetries      *config.Reloadable[int]
}

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (*GooglePubSubProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Normal configuration requires credentials

// Test configuration requires a custom endpoint

// Initialize retry configuration

func (producer *GooglePubSubProducer) publish(ctx context.Context, topic *pubsub.Publisher, message *pubsub.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// publishWithRetry publishes a message with retry logic for intermittent authentication errors
func (producer *GooglePubSubProducer) publishWithRetry(ctx context.Context, topic *pubsub.Publisher, message *pubsub.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Only retry on specific authentication error

// Return error to trigger retry

// For non-authentication errors, mark as permanent to avoid retry

// Use configurable exponential backoff

func (producer *GooglePubSubProducer) Produce(jsonData json.RawMessage, _ any) (statusCode int, respStatus, responseMessage string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

// Close closes a given producer
func (producer *GooglePubSubProducer) Close() error { _ = "STUB: not implemented"; return nil }

func getError(err error) (statusCode int) { _ = "STUB: not implemented"; return 0 }
