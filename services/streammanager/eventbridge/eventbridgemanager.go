//go:generate mockgen -destination=../../../mocks/services/streammanager/eventbridge/mock_eventbridge.go -package mock_eventbridge github.com/rudderlabs/rudder-server/services/streammanager/eventbridge EventBridgeClient

package eventbridge

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	common "github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type EventBridgeProducer struct {
	client EventBridgeClient
}

type EventBridgeClient interface {
	PutEvents(ctx context.Context, input *eventbridge.PutEventsInput, opts ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (common.Producer, error) {
	_ = "STUB: not implemented"
	return *new(common.Producer), nil
}

// Produce creates a producer and send data to EventBridge.
func (producer *EventBridgeProducer) Produce(jsonData json.RawMessage, _ any) (int, string, string) {
	_ = "STUB: not implemented"
	// get producer
	return 0, "", ""
}

// return 400 if producer is invalid

// create eventbridge event

// create eventbridge request

// send request to event bridge

// Since we are sending only one event, Entries should have only one entry

// Considering only the first entry as we sent only one event

// if one of the required fields(Detail, DetailType, Source) is missing, the error returned by PutEvents will be nil.
// In this case, outputEntry will contain the error code and message

// request has failed if errorCode and errorMessage are not nil

func (*EventBridgeProducer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
