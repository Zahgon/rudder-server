package kafka

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/linkedin/goavro/v2"

	client "github.com/rudderlabs/rudder-go-kit/kafkaclient"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

// schema is the AVRO schema required to convert the data to AVRO
type avroSchema struct {
	SchemaId string
	Schema   string
}

// configuration is the config that is required to send data to Kafka
type configuration struct {
	Topic    string
	HostName string
	Port     string

	SslEnabled    bool
	CACertificate string
	UseSASL       bool
	SaslType      string
	Username      string
	Password      string

	ConvertToAvro     bool
	EmbedAvroSchemaID bool
	AvroSchemas       []avroSchema

	UseSSH  bool
	SSHHost string
	SSHPort string
	SSHUser string
}

func (c *configuration) validate() error { _ = "STUB: not implemented"; return nil }

// azureEventHubConfig is the config that is required to send data to Azure Event Hub.
// Make sure to select at least the Standard tier since the Basic tier does not support Kafka.
type azureEventHubConfig struct {
	// Topic is the name of the Event Hub on Azure (not the Event Hubs Namespace)
	Topic string
	// BootstrapServer should be in the form of "host:port" (the port is usually 9093 on Azure Event Hubs)
	BootstrapServer string
	// EventHubsConnectionString starts with "Endpoint=sb://" and contains the SharedAccessKey
	EventHubsConnectionString string
}

func (c *azureEventHubConfig) validate() error { _ = "STUB: not implemented"; return nil }

// confluentCloudConfig is the config that is required to send data to Confluent Cloud
type confluentCloudConfig struct {
	Topic           string
	BootstrapServer string
	APIKey          string
	APISecret       string
}

func (c *confluentCloudConfig) validate() error { _ = "STUB: not implemented"; return nil }

type publisher interface {
	Publish(context.Context, ...client.Message) error
}

type producerManager interface {
	io.Closer
	publisher
	getTimeout() time.Duration
	getEmbedAvroSchemaID() bool
	getCodecs() map[string]*goavro.Codec
}

type internalProducer interface {
	publisher
	Close(context.Context) error
}

type ProducerManager struct {
	p       internalProducer
	timeout time.Duration
	topic   string

	codecs            map[string]*goavro.Codec
	embedAvroSchemaID bool
}

func (p *ProducerManager) getTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (p *ProducerManager) getCodecs() map[string]*goavro.Codec {
	_ = "STUB: not implemented"
	return nil
}
func (p *ProducerManager) getEmbedAvroSchemaID() bool { _ = "STUB: not implemented"; return false }

type logger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
}

type managerStats struct {
	creationTime               stats.Measurement
	creationTimeConfluentCloud stats.Measurement
	creationTimeAzureEventHubs stats.Measurement
	missingUserID              stats.Measurement
	missingMessage             stats.Measurement
	publishTime                stats.Measurement
	produceTime                stats.Measurement
	prepareBatchTime           stats.Measurement
	closeProducerTime          stats.Measurement
	jsonSerializationMsgErr    stats.Measurement
	avroSerializationErr       stats.Measurement
}

const (
	defaultPublishTimeout = 10 * time.Second
)

var (
	_ producerManager = &ProducerManager{}

	clientCert, clientKey []byte
	kafkaStats            managerStats
	pkgLogger             logger

	now   = func() time.Time { return time.Now() }                   // skipcq: CRT-A0018
	since = func(t time.Time) time.Duration { return time.Since(t) } // skipcq: CRT-A0018
)

func Init() { _ = "STUB: not implemented"; return }

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (*ProducerManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SASL is enabled only with SSL

// NewProducerForAzureEventHubs creates a producer for Azure event hub based on destination config
func NewProducerForAzureEventHubs(destination *backendconfig.DestinationT, o common.Opts) (*ProducerManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewProducerForConfluentCloud creates a producer for Confluent cloud based on destination config
func NewProducerForConfluentCloud(destination *backendconfig.DestinationT, o common.Opts) (*ProducerManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareMessage(topic, key string, message []byte, timestamp time.Time) client.Message {
	_ = "STUB: not implemented"
	return *new(client.Message)
}

// This function is used to serialize the binary data according to the avroSchema.
// It iterates over the schemas provided by the customer and tries to serialize the data.
// If it's able to serialize the data then it returns the converted data otherwise it returns an error.
// We are using the LinkedIn goavro library for data serialization. Ref: https://github.com/linkedin/goavro
func serializeAvroMessage(schemaID string, embedSchemaID bool, value []byte, codec goavro.Codec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addAvroSchemaIDHeader(schemaID string, msgBytes []byte) (header []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes a given producer
func (p *ProducerManager) Close() error { _ = "STUB: not implemented"; return nil }

// Publish publishes a given message to Kafka
func (p *ProducerManager) Publish(ctx context.Context, msgs ...client.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Produce sends data to Kafka.
func (p *ProducerManager) Produce(jsonData json.RawMessage, _ any) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func sendMessage(ctx context.Context, jsonData json.RawMessage, p producerManager, defaultTopic string) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func publish(ctx context.Context, p producerManager, msgs ...client.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func makeErrorResponse(err error) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

// getStatusCodeFromError parses the error and returns the status so that event gets retried or failed.
func getStatusCodeFromError(err error) int { _ = "STUB: not implemented"; return 0 }

func newProducerConfig(destType string) client.ProducerConfig {
	_ = "STUB: not implemented"
	return *new(client.ProducerConfig)
}

func getSSHPrivateKey(ctx context.Context, destinationID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isValidPort(p string) error { _ = "STUB: not implemented"; return nil }
