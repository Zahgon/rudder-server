package client

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff/v5"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

const (
	StatRequestTotalBytes     = "reporting_client_http_request_total_bytes"
	StatTotalDurationsSeconds = "reporting_client_http_total_duration_seconds"
	StatRequestLatency        = "reporting_client_http_request_latency"
	StatHttpRequest           = "reporting_client_http_request"
)

const (
	RouteMetrics      Route = "/metrics?version=v1"
	RouteRecordErrors Route = "/recordErrors"
	RouteTrackedUsers Route = "/trackedUser"
)

// Route contains the HTTP path and query string for the service.
type Route string

// URL returns the absolute URL for the route, given a base URL.
// * baseURL provides only the scheme, host, and port.
// * Route provides path and query parameters.
func (p Route) URL(baseURL string) (url.URL, error) {
	_ = "STUB: not implemented"
	return *new(url.URL), nil
}

// Client handles sending metrics to the reporting service
type Client struct {
	route               Route
	reportingServiceURL string
	userName            string
	password            string

	httpClient *http.Client
	stats      stats.Stats
	log        logger.Logger

	moduleName string
	instanceID string

	conf *config.Config
}

func backoffOptsFromConfig(conf *config.Config) (opts []backoff.RetryOption) {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new reporting client
func New(path Route, conf *config.Config, log logger.Logger, stats stats.Stats) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Send(ctx context.Context, payload any) error {
	_ = "STUB: not implemented"
	return nil
}

// Record total bytes sent

// Record request duration

// getTags returns the common tags for reporting metrics
func (c *Client) getTags() stats.Tags { _ = "STUB: not implemented"; return *new(stats.Tags) }

func (f *Client) isHTTPRequestSuccessful(status int) bool { _ = "STUB: not implemented"; return false }
