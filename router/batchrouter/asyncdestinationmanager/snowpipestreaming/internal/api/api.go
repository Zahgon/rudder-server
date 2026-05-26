package api

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/snowpipestreaming/internal/model"
)

type (
	API struct {
		clientURL   string
		requestDoer requestDoer
		config      struct {
			enableCompression config.ValueLoader[bool]
		}
		stats struct {
			insertRequestBodySize stats.Histogram
		}
	}

	requestDoer interface {
		Do(*http.Request) (*http.Response, error)
	}
)

var ErrChannelNotFound = errors.New("channel not found")

func New(conf *config.Config, statsFactory stats.Stats, clientURL string, requestDoer requestDoer) *API {
	_ = "STUB: not implemented"
	return nil
}

func mustRead(r io.Reader) []byte { _ = "STUB: not implemented"; return nil }

// CreateChannel creates a new channel with the given request.
func (a *API) CreateChannel(ctx context.Context, channelReq *model.CreateChannelRequest) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteChannel deletes the channel with the given ID.
// If sync is true, the server waits for the flushing of all records in the channel, then do the soft delete.
// If sync is false, the server do the soft delete immediately and we need to wait for the flushing of all records.
func (a *API) DeleteChannel(ctx context.Context, channelID string, sync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetChannel retrieves the channel with the given ID.
func (a *API) GetChannel(ctx context.Context, channelID string) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Insert inserts the given rows into the channel with the given ID.
func (a *API) Insert(ctx context.Context, channelID string, insertRequest *model.InsertRequest) (*model.InsertResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatus retrieves the status of the channel with the given ID.
func (a *API) GetStatus(ctx context.Context, channelID string) (*model.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBulkStatus retrieves statuses for the given channel IDs.
func (a *API) GetBulkStatus(ctx context.Context, channelIDs []string) (*model.BulkStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gzippedReader(reqJSON []byte) (io.Reader, int, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), 0, nil
}
