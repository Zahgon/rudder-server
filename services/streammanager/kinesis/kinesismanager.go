//go:generate mockgen -destination=../../../mocks/services/streammanager/kinesis/mock_kinesis.go -package mock_kinesis github.com/rudderlabs/rudder-server/services/streammanager/kinesis KinesisClient

package kinesis

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type KinesisProducer struct {
	client KinesisClient
}

type KinesisClient interface {
	PutRecord(ctx context.Context, input *kinesis.PutRecordInput, opts ...func(*kinesis.Options)) (*kinesis.PutRecordOutput, error)
}

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (common.Producer, error) {
	_ = "STUB: not implemented"
	return *new(common.Producer), nil
}

func parseKinesisError(err error) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

// Produce creates a producer and send data to Kinesis.
func (producer *KinesisProducer) Produce(jsonData json.RawMessage, destConfig any) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func (*KinesisProducer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
