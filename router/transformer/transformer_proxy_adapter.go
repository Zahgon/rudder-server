package transformer

//go:generate mockgen -destination=../../mocks/router/transformer/mock_transformer.go -package=mocks_transformer github.com/rudderlabs/rudder-server/router/transformer Transformer

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/processor/integrations"
)

type transformerProxyAdapter interface {
	getPayload(proxyReqParams *ProxyRequestParams) ([]byte, error)
	getProxyURL(destType string) (string, error)
	getResponse(response []byte, respCode int, metadata []ProxyRequestMetadata) (TransResponse, error)
}

type ProxyRequestPayloadV0 struct {
	integrations.PostParametersT
	Metadata          ProxyRequestMetadata `json:"metadata"`
	DestinationConfig map[string]any       `json:"destinationConfig"`
}

type ProxyResponseV0 struct {
	Message             string `json:"message"`
	DestinationResponse any    `json:"destinationResponse"`
	AuthErrorCategory   string `json:"authErrorCategory"`
}

type ProxyResponseV1 struct {
	Message           string           `json:"message"`
	Response          []TPDestResponse `json:"response"`
	AuthErrorCategory string           `json:"authErrorCategory"`
}

type TransResponse struct {
	routerJobResponseCodes       map[int64]int
	routerJobResponseBodys       map[int64]string
	routerJobDontBatchDirectives map[int64]bool
	authErrorCategory            string
}

type TPDestResponse struct {
	StatusCode int                  `json:"statusCode"`
	Metadata   ProxyRequestMetadata `json:"metadata"`
	Error      string               `json:"error"`
}

type (
	v0Adapter struct {
		logger logger.Logger
	}
	v1Adapter struct {
		logger logger.Logger
	}
)

func (v0 *v0Adapter) getPayload(proxyReqParams *ProxyRequestParams) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v0 *v0Adapter) getProxyURL(destType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v0 *v0Adapter) getResponse(respData []byte, respCode int, metadata []ProxyRequestMetadata) (TransResponse, error) {
	_ = "STUB: not implemented"
	return *new(TransResponse), nil
}

func (v1 *v1Adapter) getPayload(proxyReqParams *ProxyRequestParams) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v1 *v1Adapter) getProxyURL(destType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v1 *v1Adapter) getResponse(respData []byte, respCode int, metadata []ProxyRequestMetadata) (TransResponse, error) {
	_ = "STUB: not implemented"
	return *new(TransResponse), nil
}

// router/transformer/transformer_proxy_adapter.go
// getTransformerProxyURL constructs the transformer proxy URL, prioritizing DELIVERY_TRANSFORMER_URL for dedicated deployments.
// Prefer DELIVERY_TRANSFORMER_URL for dedicated deployments, fallback to DEST_TRANSFORM_URL for backward compatibility
func getTransformerProxyURL(version, destType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func NewTransformerProxyAdapter(version string, logger logger.Logger) transformerProxyAdapter {
	_ = "STUB: not implemented"
	return *new(transformerProxyAdapter)
}
