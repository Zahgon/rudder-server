package transformer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

type featuresService struct {
	logger   logger.Logger
	waitChan chan struct{}
	options  FeaturesServiceOptions
	features json.RawMessage
	client   *http.Client
}

func (t *featuresService) isInitialized() bool { _ = "STUB: not implemented"; return false }

func (t *featuresService) SourceTransformerVersion() string { _ = "STUB: not implemented"; return "" }

// todo: should we panic as SourceTransformerVersion() is called before features are fetched?

// If transformer is upgraded to V2, enable V2 spec communication

// V0 Deprecation: This function will verify if `supportSourceTransformV1` is available and enabled
// if `supportSourceTransformV1` is not enabled, transformer is not compatible and server will panic with appropriate message.

func (t *featuresService) TransformerProxyVersion() string { _ = "STUB: not implemented"; return "" }

func (t *featuresService) RouterTransform(destType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *featuresService) Regulations() []string { _ = "STUB: not implemented"; return nil }

// SupportDestTransformCompactedPayloadV1 checks if the transformer supports compacted payload for destination transformation
func (t *featuresService) SupportDestTransformCompactedPayloadV1() bool {
	_ = "STUB: not implemented"
	return false
}

func (t *featuresService) Wait() chan struct{} { _ = "STUB: not implemented"; return nil }

func (t *featuresService) syncTransformerFeatureJson(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (t *featuresService) makeFeaturesFetchCall() bool { _ = "STUB: not implemented"; return false }

//  we are calling this to see if the transformer version is deprecated. if so, we panic.
