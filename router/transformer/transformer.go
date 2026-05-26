package transformer

//go:generate mockgen -destination=../../mocks/router/transformer/mock_transformer.go -package=mocks_transformer github.com/rudderlabs/rudder-server/router/transformer Transformer

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-go-kit/sync"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	"github.com/rudderlabs/rudder-server/processor/integrations"
	"github.com/rudderlabs/rudder-server/router/types"
	oauthv2 "github.com/rudderlabs/rudder-server/services/oauth/v2"
	transformerfs "github.com/rudderlabs/rudder-server/services/transformer"
	"github.com/rudderlabs/rudder-server/utils/sysUtils"
)

const (
	BATCH            = "BATCH"
	ROUTER_TRANSFORM = "ROUTER_TRANSFORM"
	apiVersionHeader = "apiVersion"
)

// handle is the handle for this class
type handle struct {
	tr *http.Transport
	// http client for router transformation request
	client sysUtils.HTTPClientI
	// Mockable http.client for transformer proxy request
	proxyClient sysUtils.HTTPClientI
	// http client timeout for transformer proxy request
	destinationTimeout time.Duration
	// http client timeout for server-transformer request
	transformTimeout          time.Duration
	transformRequestTimerStat stats.Measurement
	logger                    logger.Logger

	stats stats.Stats

	// clientOAuthV2 is the HTTP client for router transformation requests using OAuth V2.
	clientOAuthV2 *http.Client
	// proxyClientOAuthV2 is the mockable HTTP client for transformer proxy requests using OAuth V2.
	proxyClientOAuthV2 sysUtils.HTTPClientI
	// expirationTimeDiff holds the configured time difference for token expiration.
	expirationTimeDiff config.ValueLoader[time.Duration]

	compactionSupported bool
}

type ProxyRequestMetadata struct {
	JobID         int64           `json:"jobId"`
	AttemptNum    int             `json:"attemptNum"`
	UserID        string          `json:"userId"`
	SourceID      string          `json:"sourceId"`
	DestinationID string          `json:"destinationId"`
	WorkspaceID   string          `json:"workspaceId"`
	Secret        json.RawMessage `json:"secret"`             // Receives OAuth destination secrets in the transformer and holds the current token during refresh token flows
	DestInfo      json.RawMessage `json:"destInfo,omitempty"` // Used by the transformer; potentially removable
	DontBatch     bool            `json:"dontBatch"`
}

type ProxyRequestPayload struct {
	integrations.PostParametersT
	Metadata          []ProxyRequestMetadata `json:"metadata"`
	DestinationConfig map[string]any         `json:"destinationConfig"`
}

type ProxyRequestParams struct {
	ResponseData ProxyRequestPayload
	DestName     string
	Adapter      transformerProxyAdapter
	Destination  *backendconfig.DestinationT
	Connection   backendconfig.Connection `json:"connection"`
}

type ProxyRequestResponse struct {
	ProxyRequestStatusCode   int
	ProxyRequestResponseBody string
	RespContentType          string
	RespStatusCodes          map[int64]int
	RespBodys                map[int64]string
	DontBatchDirectives      map[int64]bool
	OAuthErrorCategory       string
}

// Transformer provides methods to transform events
type Transformer interface {
	Transform(transformType string, transformMessage *types.TransformMessageT) []types.DestinationJobT
	ProxyRequest(ctx context.Context, proxyReqParams *ProxyRequestParams) ProxyRequestResponse
}

// NewTransformer creates a new transformer.
// If a nil [featuresService] is provided, the transformer will not use message compaction, even though transformation service might support it.
func NewTransformer(
	destType string,
	destinationTimeout, transformTimeout time.Duration,
	backendConfig backendconfig.BackendConfig,
	expirationTimeDiff config.ValueLoader[time.Duration],
	featuresService transformerfs.FeaturesService,
	conf *config.Config,
) Transformer {
	_ = "STUB: not implemented"
	return *new(Transformer)
}

var loggerOverride logger.Logger

// Add transformerMetricLabels struct and methods
type transformerMetricLabels struct {
	Endpoint        string // hostname of the service
	DestinationType string // BQ, etc.
	SourceType      string // webhook
	Stage           string // processor, router, gateway
	WorkspaceID     string // workspace identifier
	SourceID        string // source identifier
	DestinationID   string // destination identifier
}

// ToStatsTag converts transformerMetricLabels to stats.Tags
func (t transformerMetricLabels) ToStatsTag() stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

// Legacy tags: to be removed

// Transform transforms router jobs to destination jobs
func (trans *handle) Transform(transformType string, transformMessage *types.TransformMessageT) []types.DestinationJobT {
	_ = "STUB: not implemented"
	return nil
}

// consistent state for the entire request

// Create metric labels

// Record request metrics

// We should rarely have error communicating with our JS

// No point in retrying if we can't even create a request. Panicking as per convention.

// Header to let transformer know that the client understands event filter code

// If no err returned by client.Post, reading body.
// If reading body fails, retrying.

// Refresh the connection

// We don't need to handle it, as we can receive a string response even before executing OAuth operations like Refresh Token.
// It's acceptable if the structure of respData doesn't match the oauthv2.TransportResponse struct.

// re-assign originalResponse

// Validate the response received from the transformer

// invalid jobIDs are the ones that are in the response but were not included in the request

// Retrying. Go and fix transformer.

func (trans *handle) ProxyRequest(ctx context.Context, proxyReqParams *ProxyRequestParams) ProxyRequestResponse {
	_ = "STUB: not implemented"
	return *new(ProxyRequestResponse)
}

// Create metric labels

// Record request metrics

/*
	respData will be in ProxyResponseV0 or ProxyResponseV1
*/
// response that we get from oauth-interceptor in postRoundTrip

// unmarshal unsuccessful scenarios
// if respData is not a valid json

/**

	Structure of TransformerProxy Response:
	{
		output: {
			status: [destination status compatible with server]
			message: [ generic message for jobs_db payload]
			destinationResponse: [actual response payload from destination] <-- v0
			response: [actual response payload from destination] <-- v1
		}
	}
**/

// Conditions for which InterceptorResponse.StatusCode/Response will not be empty
// 1. authErrorCategory == CategoryRefreshToken
// 2. authErrorCategory == CategoryAuthStatusInactive
// 3. Any error occurred while performing authStatusInactive / RefreshToken
// Under these conditions, we will have to propagate the response from interceptor to JobsDB

func (trans *handle) setup(destType string, destinationTimeout, transformTimeout time.Duration, cache *oauthv2.OauthTokenCache, locker *sync.PartitionRWLocker, backendConfig backendconfig.BackendConfig, featuresService transformerfs.FeaturesService, conf *config.Config) {
	_ = "STUB: not implemented"
	return
}

// The timeout between server and transformer
// Basically this timeout is more for communication between transformer and server

// Destination API timeout
// Basically this timeout we will configure when we make final call to destination to send event

// This client is used for Router Transformation

// This client is used for Router Transformation using oauthV2

// This client is used for Transformer Proxy(delivered from transformer to destination)

// This client is used for Transformer Proxy(delivered from transformer to destination) using oauthV2

func (trans *handle) transformerClientConfig() *transformerclient.ClientConfig {
	_ = "STUB: not implemented"
	return nil
}

type httpProxyResponse struct {
	respData   []byte
	statusCode int
	err        error
}

func (trans *handle) doProxyRequest(ctx context.Context, proxyUrl string, proxyReqParams *ProxyRequestParams, payload []byte) httpProxyResponse {
	_ = "STUB: not implemented"
	return *new(httpProxyResponse)
}

// Make use of this header to set timeout in the transfomer's http client
// The header name may be worked out ?

// This stat will be useful in understanding the round trip time taken for the http req
// between server and transformer

// A timeout error occurred

// This was an error, but not a timeout

// Actually Router wouldn't send any destination to proxy unless it already exists
// But if accidentally such a request is sent, failing instead of aborting

// error handling if body is missing

// error handling while reading from resp.Body

// sending this as it is not getting sent at all

func getBatchURL() string { _ = "STUB: not implemented"; return "" }

func getRouterTransformURL() string { _ = "STUB: not implemented"; return "" }

type transformerResponse struct {
	AuthErrorCategory string `json:"authErrorCategory"`
}

// GetAuthErrorCategoryFromTransformResponse parses the response data from a transformerResponse
// to extract the authentication error category.
// {input: [{}]}
// {input: [{}, {}, {}, {}]}
// {input: [{}, {}, {}, {}]} -> {output: [{200}, {200}, {401,authErr}, {401,authErr}]}
func GetAuthErrorCategoryFromTransformResponse(respData []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// can be a valid scenario

func GetAuthErrorCategoryFromTransformProxyResponse(respData []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Helper function to get endpoint from URL
func getEndpointFromURL(urlStr string) string { _ = "STUB: not implemented"; return "" }

func (trans *handle) getRequestPayload(data *types.TransformMessageT, compactRequestPayloads bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
