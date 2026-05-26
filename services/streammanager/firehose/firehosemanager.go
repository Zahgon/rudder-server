//go:generate mockgen -destination=../../../mocks/services/streammanager/firehose/mock_firehose.go -package mock_firehose github.com/rudderlabs/rudder-server/services/streammanager/firehose FirehoseClient

package firehose

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/firehose"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type FirehoseProducer struct {
	client FirehoseClient
}

type FirehoseClient interface {
	PutRecord(ctx context.Context, input *firehose.PutRecordInput, opts ...func(*firehose.Options)) (*firehose.PutRecordOutput, error)
}

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (common.Producer, error) {
	_ = "STUB: not implemented"
	return *new(common.Producer), nil
}

// Produce creates a producer and send data to Firehose.
func (producer *FirehoseProducer) Produce(jsonData json.RawMessage, _ any) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func (*FirehoseProducer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
