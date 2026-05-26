package pulsar

import (
	"context"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	pulsarLog "github.com/apache/pulsar-client-go/pulsar/log"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

type ClientConf struct {
	url               string
	operationTimeout  time.Duration
	connectionTimeout time.Duration
}

type Producer struct {
	pulsar.Producer
}

type ProducerAdapter interface {
	SendMessage(ctx context.Context, key, orderingKey string, msg []byte) error
	SendMessageAsync(ctx context.Context, key, orderingKey string, msg []byte, statusFunc func(id pulsar.MessageID, message *pulsar.ProducerMessage, err error))
	Close()
	Flush() error
}

type Client struct {
	pulsar.Client
}

// NewClient returns a new instance of Pulsar client
func NewClient(config *config.Config) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// NewProducer returns a new instance of Pulsar producer
func (c *Client) NewProducer(opts pulsar.ProducerOptions) (ProducerAdapter, error) {
	_ = "STUB: not implemented"
	return *new(ProducerAdapter), nil
}

func newPulsarClient(conf ClientConf, log logger.Logger) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// SendMessage sends a message to pulsar synchronously
func (p *Producer) SendMessage(ctx context.Context, key, orderingKey string, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SendMessageAsync sends a message to pulsar asynchronously
func (p *Producer) SendMessageAsync(ctx context.Context, key, orderingKey string, msg []byte, statusfunc func(id pulsar.MessageID, message *pulsar.ProducerMessage, err error)) {
	_ = "STUB: not implemented"
	return
}

func getClientConf(config *config.Config) ClientConf {
	_ = "STUB: not implemented"
	return *new(ClientConf)
}

type pulsarLogAdapter struct {
	logger.Logger
}

func (pl *pulsarLogAdapter) SubLogger(fields pulsarLog.Fields) pulsarLog.Logger {
	_ = "STUB: not implemented"
	return *new(pulsarLog.Logger)
}

// nolint:forbidigo

func (pl *pulsarLogAdapter) WithFields(fields pulsarLog.Fields) pulsarLog.Entry {
	_ = "STUB: not implemented"
	return *new(pulsarLog.Entry)
}

func (pl *pulsarLogAdapter) WithField(name string, value any) pulsarLog.Entry {
	_ = "STUB: not implemented"
	return *new(pulsarLog.Entry)
}

// nolint:forbidigo

func (pl *pulsarLogAdapter) WithError(err error) pulsarLog.Entry {
	_ = "STUB: not implemented"
	return *new(pulsarLog.Entry)
}
