//go:generate mockgen -destination=../../../mocks/services/streammanager/bqstream/mock_bqstream.go -package mock_bqstream github.com/rudderlabs/rudder-server/services/streammanager/bqstream BQClient

package bqstream

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/bigquery"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type Config struct {
	Credentials string `json:"credentials"`
	ProjectId   string `json:"projectId"`
	DatasetId   string `json:"datasetId"`
	TableId     string `json:"tableId"`
}

// https://stackoverflow.com/questions/55951812/insert-into-bigquery-without-a-well-defined-struct
type GenericRecord map[string]bigquery.Value

type BQClient interface {
	Put(ctx context.Context, datasetID, tableID string, records []*GenericRecord) error
	Close() error
}

type BQStreamProducer struct {
	Opts   common.Opts
	Client BQClient
}

type Client struct {
	bqClient *bigquery.Client
}

func (c *Client) Put(ctx context.Context, datasetID, tableID string, records []*GenericRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (rec GenericRecord) Save() (map[string]bigquery.Value, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

var pkgLogger logger.Logger

func Init() { _ = "STUB: not implemented"; return }

func init() {
	Init()
}

func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (*BQStreamProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (producer *BQStreamProducer) Produce(jsonData json.RawMessage, _ any) (statusCode int, respStatus, responseMessage string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

func (producer *BQStreamProducer) Close() error { _ = "STUB: not implemented"; return nil }

func createErr(err error, msg string) error { _ = "STUB: not implemented"; return nil }
