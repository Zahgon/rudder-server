package transformertest

import (
	"net/http"
	"net/http/httptest"

	routerTypes "github.com/rudderlabs/rudder-server/router/types"
)

// NewBuilder returns a new test transformer Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// Builder is a builder for a test transformer server
type Builder struct {
	routerTransforms            map[string]struct{}
	userTransformHandler        http.HandlerFunc
	destTransformHandlers       map[string]http.HandlerFunc
	trackingPlanHandler         http.HandlerFunc
	routerTransformHandler      http.HandlerFunc
	routerBatchTransformHandler http.HandlerFunc
	featuresHandler             http.HandlerFunc
	srcHydrationHandlers        map[string]http.HandlerFunc
}

func (b *Builder) WithFeaturesHandler(h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRouterTransformHandlerFunc sets the router transformation http handler function for the server
func (b *Builder) WithRouterTransformHandlerFunc(h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRouterTransformHandler sets the router transformation handler for the server
func (b *Builder) WithRouterTransformHandler(h RouterTransformerHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRouterBatchTransformHandlerFunc sets the router batch transformation http handler function for the server
func (b *Builder) WithRouterBatchTransformHandlerFunc(h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRouterBatchTransformHandler sets the router batch transformation handler for the server
func (b *Builder) WithRouterBatchTransformHandler(h RouterTransformerHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithUserTransformHandlerFunc sets the user transformation http handler function for the server
func (b *Builder) WithUserTransformHandlerFunc(h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithUserTransformHandler sets the user transformation handler for the server
func (b *Builder) WithUserTransformHandler(h TransformerHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithDesTransformHandlerFunc sets a destination specific transformation http handler function for the server
func (b *Builder) WithDesTransformHandlerFunc(destType string, h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) WithSrcHydrationHandlerFunc(srcName string, h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) WithSrcHydrationHandler(srcName string, h SrcHydrationHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithDestTransformHandler sets a destination specific transformation handler for the server
func (b *Builder) WithDestTransformHandler(destType string, h TransformerHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithTrackingPlanHandlerFunc sets the tracking plan validation http handler function for the server
func (b *Builder) WithTrackingPlanHandlerFunc(h http.HandlerFunc) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithTrackingPlanHandler sets the tracking plan validation handler for the server
func (b *Builder) WithTrackingPlanHandler(h TransformerHandler) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRouterTransform enables router transformation for a specific destination type
func (b *Builder) WithRouterTransform(destType string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build builds the test tranformer server
func (b *Builder) Build() *httptest.Server {
	_ = "STUB: not implemented"
	// user/custom transformation
	return nil
}

// tracking plan validtion

// destination transformation

// router transformation

// features

func transformerFunc(h TransformerHandler) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func srcTransformerFunc(h SrcHydrationHandler) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func routerTransformerFunc(h RouterTransformerHandler) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func routerBatchTransformerFunc(h RouterTransformerHandler) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func apiVersionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func routerCompactedTransformMessageToTransformMessage(data []byte) (routerTypes.TransformMessageT, error) {
	_ = "STUB: not implemented"
	return *new(routerTypes.TransformMessageT), nil
}
