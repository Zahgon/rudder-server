//go:generate mockgen --build_flags=--mod=mod -destination=../../../mocks/services/streammanager/personalize/mock_personalize.go -package mock_personalize github.com/rudderlabs/rudder-server/services/streammanager/personalize PersonalizeClient

package personalize

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type PersonalizeProducer struct {
	client PersonalizeClient
}

type PersonalizeClient interface {
	PutEvents(ctx context.Context, input *personalizeevents.PutEventsInput, opts ...func(*personalizeevents.Options)) (*personalizeevents.PutEventsOutput, error)
	PutUsers(ctx context.Context, input *personalizeevents.PutUsersInput, opts ...func(*personalizeevents.Options)) (*personalizeevents.PutUsersOutput, error)
	PutItems(ctx context.Context, input *personalizeevents.PutItemsInput, opts ...func(*personalizeevents.Options)) (*personalizeevents.PutItemsOutput, error)
}

func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (common.Producer, error) {
	_ = "STUB: not implemented"
	return *new(common.Producer), nil
}

func (producer *PersonalizeProducer) Produce(jsonData json.RawMessage, _ any) (statusCode int, respStatus, responseMessag string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func (*PersonalizeProducer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
