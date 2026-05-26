//go:generate mockgen --build_flags=--mod=mod -destination=../../../mocks/services/streammanager/wunderkind/mock_wunderkind.go -package mock_wunderkind github.com/rudderlabs/rudder-server/services/streammanager/wunderkind LambdaClient

package wunderkind

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/lambda"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type Producer struct {
	client LambdaClient
	logger logger.Logger
	conf   *config.Config
}

type LambdaClient interface {
	Invoke(ctx context.Context, input *lambda.InvokeInput, opts ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
}

// NewProducer creates a producer based on destination config
func NewProducer(conf *config.Config, destination *backendconfig.DestinationT, o common.Opts) (common.Producer, error) {
	_ = "STUB: not implemented"
	return *new(common.Producer), nil
}

// Produce creates a producer and send data to Lambda.
func (p *Producer) Produce(jsonData json.RawMessage, destConfig any) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

// handle a case where lambda invocation is successful, but there is an issue with the payload.

func (*Producer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
