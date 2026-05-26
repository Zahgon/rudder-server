package webhook

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/config"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

const (
	contentTypeJsonUTF8 = "application/json; charset=utf-8"
)

type sourceTransformAdapter interface {
	getTransformerEvent(authCtx *gwtypes.AuthRequestContext, eventRequest []byte) ([]byte, error)
	getTransformerURL(sourceType string) (string, error)
	getAdapterVersion() string
}

// ----- v1 adapter ---------

type v1Adapter struct {
	baseTransformerURL string
}

type V1TransformerEvent struct {
	EventRequest json.RawMessage       `json:"event"`
	Source       backendconfig.SourceT `json:"source"`
}

func (v1 *v1Adapter) getTransformerEvent(authCtx *gwtypes.AuthRequestContext, eventRequest []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v1 *v1Adapter) getTransformerURL(sourceType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v1 *v1Adapter) getAdapterVersion() string { _ = "STUB: not implemented"; return "" }

// ----- v2 adapter -----

type v2Adapter struct {
	baseTransformerURL string
}

type V2TransformerEvent struct {
	EventRequest json.RawMessage       `json:"request"`
	Source       backendconfig.SourceT `json:"source"`
}

func (v2 *v2Adapter) getTransformerEvent(authCtx *gwtypes.AuthRequestContext, eventRequest []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v2 *v2Adapter) getTransformerURL(sourceType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v2 *v2Adapter) getAdapterVersion() string { _ = "STUB: not implemented"; return "" }

// ------------------------------

func newSourceTransformAdapter(version string, conf *config.Config) sourceTransformAdapter {
	_ = "STUB: not implemented"
	// V0 Deprecation: this function returns v1 adapter by default, thereby deprecating v0
	return *new(sourceTransformAdapter)
}

// --- utilities -----

func getTransformerURL(version, sourceType, baseURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func prepareTransformerEventRequestV1(req *http.Request, sourceType string, sourceListForParsingParams []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If body is empty, set it to an empty JSON object

func prepareTransformerEventRequestV2(req *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type outputToSource struct {
	Body        []byte `json:"body"`
	ContentType string `json:"contentType"`
}

// transformerResponse will be populated using JSON unmarshall
// so we need to make fields public
type transformerResponse struct {
	Output         map[string]any  `json:"output"`
	Err            string          `json:"error"`
	StatusCode     int             `json:"statusCode"`
	OutputToSource *outputToSource `json:"outputToSource"`
}

type transformerBatchResponseT struct {
	batchError error
	responses  []transformerResponse
	statusCode int
}

func (bt *batchWebhookTransformerT) markResponseFail(reason string) transformerResponse {
	_ = "STUB: not implemented"
	return *new(transformerResponse)
}

func (bt *batchWebhookTransformerT) transform(events [][]byte, sourceTransformerURL string) transformerBatchResponseT {
	_ = "STUB: not implemented"
	return *new(transformerBatchResponseT)
}

/*
	expected response format
	[
		------Output to Gateway only---------
		{
			output: {
				batch: [
					{
						context: {...},
						properties: {...},
						userId: "U123"
					}
				]
			}
		}

		------Output to Source only---------
		{
			outputToSource: {
				"body": "eyJhIjoxfQ==", // base64 encode string
				"contentType": "application/json"
			}
		}

		------Output to Both Gateway and Source---------
		{
			output: {
				batch: [
					{
						context: {...},
						properties: {...},
						userId: "U123"
					}
				]
			},
			outputToSource: {
				"body": "eyJhIjoxfQ==", // base64 encode string
				"contentType": "application/json"
			}
		}

		------Error example---------
		{
			statusCode: 400,
			error: "event type is not supported"
		}

	]
*/

func (bt *batchWebhookTransformerT) doPost(transformerURL string, body io.Reader) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
