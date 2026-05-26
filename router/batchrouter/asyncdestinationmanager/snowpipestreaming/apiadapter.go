package snowpipestreaming

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/snowpipestreaming/internal/model"
)

const (
	createChannelAPI = "create_channel"
	deleteChannelAPI = "delete_channel"
	insertAPI        = "insert"
	bulkStatusAPI    = "bulk_status"
)

func newApiAdapter(
	logger logger.Logger,
	statsFactory stats.Stats,
	api api,
	destination *backendconfig.DestinationT,
) api {
	_ = "STUB: not implemented"
	return *new(api)
}

func (a *apiAdapter) defaultTags(apiName string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

func (a *apiAdapter) CreateChannel(ctx context.Context, req *model.CreateChannelRequest) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *apiAdapter) DeleteChannel(ctx context.Context, channelID string, sync bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *apiAdapter) Insert(ctx context.Context, channelID string, insertRequest *model.InsertRequest) (*model.InsertResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *apiAdapter) GetBulkStatus(ctx context.Context, channelIDs []string) (*model.BulkStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *apiAdapter) recordDuration(tags stats.Tags) func() { _ = "STUB: not implemented"; return nil }
