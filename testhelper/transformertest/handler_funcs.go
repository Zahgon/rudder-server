package transformertest

import (
	"net/http"

	"github.com/rudderlabs/rudder-server/processor/integrations"
	proctypes "github.com/rudderlabs/rudder-server/processor/types"
	"github.com/rudderlabs/rudder-server/router/types"
)

// TransformerHandler is a function that takes a transformer request and returns a response
type TransformerHandler func(request []proctypes.TransformerEvent) []proctypes.TransformerResponse

// RouterTransformerHandler is a function that takes a router transformer request and returns a response
type RouterTransformerHandler func(request types.TransformMessageT) types.DestinationJobs

// SrcHydrationHandler is a function that takes a src hydration request and returns a response
type SrcHydrationHandler func(request proctypes.SrcHydrationRequest) (response proctypes.SrcHydrationResponse, err error)

// MirroringTransformerHandler mirrors the request payload in the response
var MirroringTransformerHandler TransformerHandler = func(request []proctypes.TransformerEvent) (response []proctypes.TransformerResponse) {
	for i := range request {
		req := request[i]
		response = append(response, proctypes.TransformerResponse{
			Metadata:   req.Metadata,
			Output:     req.Message,
			StatusCode: http.StatusOK,
		})
	}
	return response
}

// MirroringRouterTransformerHandler mirrors the router request payload in the response
var MirroringRouterTransformerHandler RouterTransformerHandler = func(request types.TransformMessageT) (response types.DestinationJobs) {
	response = make(types.DestinationJobs, len(request.Data))
	for j := range request.Data {
		req := request.Data[j]
		response[j] = types.DestinationJobT{
			Message:          req.Message,
			JobMetadataArray: []types.JobMetadataT{req.JobMetadata},
			Destination:      req.Destination,
			Connection:       req.Connection,
			StatusCode:       http.StatusOK,
		}
	}
	return response
}

// ErrorTransformerHandler mirrors the request payload in the response but uses an error status code
func ErrorTransformerHandler(code int, err string) TransformerHandler {
	_ = "STUB: not implemented"
	return *new(TransformerHandler)
}

// ViolationErrorTransformerHandler mirrors the request payload in the response but uses an error status code along with the provided validation errors
func ViolationErrorTransformerHandler(code int, err string, validationErrors []proctypes.ValidationError) TransformerHandler {
	_ = "STUB: not implemented"
	return *new(TransformerHandler)
}

// EmptyTransformerHandler returns an empty response
var EmptyTransformerHandler TransformerHandler = func(request []proctypes.TransformerEvent) []proctypes.TransformerResponse {
	return []proctypes.TransformerResponse{}
}

// DestTransformerHandler returns an empty response
func DestTransformerHandler(f func(event proctypes.TransformerEvent) integrations.PostParametersT) func(request []proctypes.TransformerEvent) []proctypes.TransformerResponse {
	_ = "STUB: not implemented"
	return nil
}

// RESTJSONDestTransformerHandler transforms the request payload into a REST JSON destination request using the original message as the payload
func RESTJSONDestTransformerHandler(method, url string) func(request []proctypes.TransformerEvent) []proctypes.TransformerResponse {
	_ = "STUB: not implemented"
	return nil
}

// WarehouseTransformerHandler mirrors the request payload in the response but uses an error, status code along with warehouse compatible output
func WarehouseTransformerHandler(tableName string, code int, err string) TransformerHandler {
	_ = "STUB: not implemented"
	return *new(TransformerHandler)
}
