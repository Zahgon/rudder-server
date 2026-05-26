package client

import (
	"database/sql"

	"cloud.google.com/go/bigquery"

	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	SQLClient = "SQLClient"
	BQClient  = "BigQueryClient"
)

type Client struct {
	SQL  *sql.DB
	BQ   *bigquery.Client
	Type string
}

func (cl *Client) sqlQuery(statement string) (result warehouseutils.QueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(warehouseutils.QueryResult), nil
}

func (cl *Client) bqQuery(statement string) (result warehouseutils.QueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(warehouseutils.QueryResult), nil
}

func (cl *Client) Query(statement string) (result warehouseutils.QueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(warehouseutils.QueryResult), nil
}

func (cl *Client) Close() { _ = "STUB: not implemented"; return }
