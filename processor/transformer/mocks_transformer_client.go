package transformer

import (
	"context"

	"github.com/rudderlabs/rudder-server/processor/types"
)

// SimpleMockDestinationClient is a minimal mock for DestinationClient
type SimpleMockDestinationClient struct {
	// Fixed responses to return
	TransformOutput types.Response
}

// Transform implements the DestinationClient interface
func (m *SimpleMockDestinationClient) Transform(_ context.Context, _ []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *

	// SimpleMockUserClient is a minimal mock for UserClient
	new(types.Response)
}

type SimpleMockUserClient struct {
	// Fixed response to return
	TransformOutput types.Response
}

// Transform implements the UserClient interface
func (m *SimpleMockUserClient) Transform(_ context.Context, _ []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *

	// SimpleMockTrackingPlanClient is a minimal mock for TrackingPlanClient
	new(types.Response)
}

type SimpleMockTrackingPlanClient struct {
	// Fixed response to return
	ValidateOutput types.Response
}

// Validate implements the TrackingPlanClient interface
func (m *SimpleMockTrackingPlanClient) Validate(_ context.Context, _ []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

type SimpleMockSrcHydrationClient struct {
	HydratedOutput types.SrcHydrationResponse
	err            error
}

func (m *SimpleMockSrcHydrationClient) Hydrate(_ context.Context, _ types.SrcHydrationRequest) (types.SrcHydrationResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.SrcHydrationResponse), nil
}

// SimpleClients is a minimal implementation of TransformerClients
type SimpleClients struct {
	userClient         UserClient
	userMirrorClient   UserClient
	destinationClient  DestinationClient
	trackingPlanClient TrackingPlanClient
	sycHydrationClient SrcHydrationClient
}

// NewSimpleClients creates a new instance of SimpleClients with empty responses
func NewSimpleClients() *SimpleClients { _ = "STUB: not implemented"; return nil }

// User returns the user client
func (s *SimpleClients) User() UserClient { _ = "STUB: not implemented"; return *new(UserClient) }

func (s *SimpleClients) UserMirror() UserClient {
	_ = "STUB: not implemented"
	return *

	// Destination returns the destination client
	new(UserClient)
}

func (s *SimpleClients) Destination() DestinationClient {
	_ = "STUB: not implemented"
	return *new(DestinationClient)
}

// TrackingPlan returns the tracking plan client
func (s *SimpleClients) TrackingPlan() TrackingPlanClient {
	_ = "STUB: not implemented"
	return *new(TrackingPlanClient)
}

func (s *SimpleClients) SrcHydration() SrcHydrationClient {
	_ = "STUB: not implemented"
	return *new(SrcHydrationClient)
}

// SetUserTransformOutput sets the response for the User transformer
func (s *SimpleClients) SetUserTransformOutput(response types.Response) {
	_ = "STUB: not implemented"
	return
}

// SetDestinationTransformOutput sets the response for the Destination transformer
func (s *SimpleClients) SetDestinationTransformOutput(response types.Response) {
	_ = "STUB: not implemented"
	return
}

// SetTrackingPlanValidateOutput sets the response for the TrackingPlan validator
func (s *SimpleClients) SetTrackingPlanValidateOutput(response types.Response) {
	_ = "STUB: not implemented"
	return
}

// SetSrcHydrationOutput sets the response for the Source Hydration client
func (s *SimpleClients) SetSrcHydrationOutput(response types.SrcHydrationResponse, err error) {
	_ = "STUB: not implemented"
	return
}

// WithDynamicUserTransform sets a custom function for User transformer
func (s *SimpleClients) WithDynamicUserTransform(transformFn func(context.Context, []types.TransformerEvent) types.Response) {
	_ = "STUB: not implemented"
	return
}

// WithDynamicUserMirrorTransform sets a custom function for UserMirror transformer
func (s *SimpleClients) WithDynamicUserMirrorTransform(transformFn func(context.Context, []types.TransformerEvent) types.Response) {
	_ = "STUB: not implemented"
	return
}

// WithDynamicDestinationTransform sets a custom function for Destination transformer
func (s *SimpleClients) WithDynamicDestinationTransform(transformFn func(context.Context, []types.TransformerEvent) types.Response) {
	_ = "STUB: not implemented"
	return
}

// WithDynamicTrackingPlanValidate sets a custom function for TrackingPlan validator
func (s *SimpleClients) WithDynamicTrackingPlanValidate(validateFn func(context.Context, []types.TransformerEvent) types.Response) {
	_ = "STUB: not implemented"
	return
}

// WithDynamicSrcHydration sets a custom function for Source Hydration
func (s *SimpleClients) WithDynamicSrcHydration(hydrateFn func(context.Context, types.SrcHydrationRequest) (types.SrcHydrationResponse, error)) {
	_ = "STUB: not implemented"
	return
}

// Helper types for dynamic behavior

type dynamicUserClient struct {
	transformFn func(context.Context, []types.TransformerEvent) types.Response
}

func (d *dynamicUserClient) Transform(ctx context.Context, events []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

type dynamicDestinationClient struct {
	transformFn func(context.Context, []types.TransformerEvent) types.Response
}

func (d *dynamicDestinationClient) Transform(ctx context.Context, events []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

type dynamicTrackingPlanClient struct {
	validateFn func(context.Context, []types.TransformerEvent) types.Response
}

func (d *dynamicTrackingPlanClient) Validate(ctx context.Context, events []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

type dynamicSrcHydrationClient struct {
	hydrateFn func(context.Context, types.SrcHydrationRequest) (types.SrcHydrationResponse, error)
}

func (d *dynamicSrcHydrationClient) Hydrate(ctx context.Context, req types.SrcHydrationRequest) (types.SrcHydrationResponse, error) {
	_ = "STUB: not implemented"
	return *new(types.SrcHydrationResponse), nil
}

// Helper functions to create common responses

// EmptySuccessResponse creates an empty successful response
func EmptySuccessResponse() types.Response { _ = "STUB: not implemented"; return *new(types.Response) }

// SuccessResponse creates a response with successfully processed events
func SuccessResponse(outputs []map[string]any, metadatas []types.Metadata) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// ErrorResponse creates a response with error information
func ErrorResponse(errorMsg string, statusCode int, metadata types.Metadata) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}
