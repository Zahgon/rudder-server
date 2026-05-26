package middleware

import (
	"context"
	"time"

	"cloud.google.com/go/bigquery"
)

type Opt func(*Client)

type loggerMW interface {
	Infow(msg string, keysAndValues ...any)
}

type Client struct {
	*bigquery.Client

	since              func(time.Time) time.Duration
	logger             loggerMW
	keysAndValues      []any
	slowQueryThreshold time.Duration
}

func WithLogger(logger loggerMW) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithKeyAndValues(keyAndValues ...any) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithSlowQueryThreshold(slowQueryThreshold time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithSince(since func(time.Time) time.Duration) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func New(client *bigquery.Client, opts ...Opt) *Client { _ = "STUB: not implemented"; return nil }

func (client *Client) Run(ctx context.Context, query *bigquery.Query) (*bigquery.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *Client) Read(ctx context.Context, query *bigquery.Query) (it *bigquery.RowIterator, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *Client) logQuery(query *bigquery.Query, elapsed time.Duration) {
	_ = "STUB: not implemented"
	return
}
