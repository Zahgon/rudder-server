//go:generate mockgen -destination=../../mocks/transformer-client/mock_transformer_client.go -package=mocks_transformer_client github.com/rudderlabs/rudder-server/internal/transformer-client Client

package transformerclient

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/httplb"
	"github.com/bufbuild/httplb/conn"
	"github.com/bufbuild/httplb/picker"

	"github.com/rudderlabs/rudder-go-kit/retryablehttp"
)

type perpetualRetriesStatsTagsKey struct{}

// WithPerpetualRetriesStatsTags returns a context carrying extra tags to be
// added to the transformer_client_perpetual_retry_count stat for any request
// made with this context. Callers are responsible for keeping tag cardinality
// low.
func WithPerpetualRetriesStatsTags(ctx context.Context, tags map[string]string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func perpetualRetriesStatsTagsFromContext(ctx context.Context) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

const (
	defaultDisableKeepAlives   = true
	defaultMaxConnsPerHost     = 100
	defaultMaxIdleConnsPerHost = 10
	defaultIdleConnTimeout     = 30 * time.Second
	defaultClientTimeout       = 600 * time.Second
	defaultClientTTL           = 10 * time.Second
	defaultRecycleTTL          = 60 * time.Second

	defaultRetryRudderErrorsMaxRetry        = -1
	defaultRetryRudderErrorsInitialInterval = 1 * time.Second
	defaultRetryRudderErrorsMaxInterval     = 30 * time.Second
	defaultRetryRudderErrorsMaxElapsedTime  = 0
	defaultRetryRudderErrorsMultiplier      = 2.0
)

type ClientConfig struct {
	TransportConfig struct {
		DisableKeepAlives   bool          //	true
		MaxConnsPerHost     int           //	100
		MaxIdleConnsPerHost int           //	10
		IdleConnTimeout     time.Duration //	30*time.Second
	}

	ClientTimeout time.Duration //	600*time.Second
	ClientTTL     time.Duration //	10*time.Second
	ClientType    string        // stdlib(default), httplb
	PickerType    string        // power_of_two(default), round_robin, least_loaded_random, least_loaded_round_robin, random
	Recycle       bool          // false
	RecycleTTL    time.Duration // 60s

	// Configuration for retryable HTTP client in case of [X-Rudder-Should-Retry: true] HTTP 503 responses
	RetryRudderErrors struct {
		Enabled         bool          // false
		MaxRetry        int           // -1 - no limit
		InitialInterval time.Duration // 1s
		MaxInterval     time.Duration // 30s
		MaxElapsedTime  time.Duration // 0s - no limit
		Multiplier      float64       // 2.0
	}
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient(name string, config *ClientConfig) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// buildStandardClient creates a standard HTTP client with configuration applied
func buildStandardClient(name string, config *ClientConfig) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// buildHTTPLBClient creates an HTTP load balancer client
func buildHTTPLBClient(name string, config *ClientConfig) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// buildConfiguredTransport creates a transport with configuration applied
func buildConfiguredTransport(config *ClientConfig) *http.Transport {
	_ = "STUB: not implemented"
	return nil
}

// buildRetryableConfig creates retryable configuration if enabled
func buildRetryableConfig(clientConfig *ClientConfig) *retryablehttp.Config {
	_ = "STUB: not implemented"
	return nil
}

// Use ClientConfig values directly

// Helper functions to get configuration values with defaults
func getClientTimeout(config *ClientConfig) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getClientTTL(config *ClientConfig) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getRecycleTTL(config *ClientConfig) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func newRetryableHTTPClient(name string, baseClient Client, retryableConfig *retryablehttp.Config) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func getPicker(pickerType string) func(prev picker.Picker, allConns conn.Conns) picker.Picker {
	_ = "STUB: not implemented"
	return nil
}

type httplbtransport struct {
	MaxConnsPerHost     int
	MaxIdleConnsPerHost int
	*http.Transport
}

func (s httplbtransport) NewRoundTripper(_, _ string, opts httplb.TransportConfig) httplb.RoundTripperResult {
	_ = "STUB: not implemented"
	return *new(httplb.RoundTripperResult)
}
