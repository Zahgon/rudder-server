//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/services/transformer/mock_features.go -package mock_features github.com/rudderlabs/rudder-server/services/transformer FeaturesService

package transformer

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
)

const (
	V0 = "v0"
	V1 = "v1"
	V2 = "v2"
)

type FeaturesServiceOptions struct {
	PollInterval             time.Duration
	TransformerURL           string
	FeaturesRetryMaxAttempts int
}

type FeaturesService interface {
	Regulations() []string
	SourceTransformerVersion() string
	RouterTransform(destType string) bool
	TransformerProxyVersion() string
	SupportDestTransformCompactedPayloadV1() bool
	Wait() chan struct{}
}

var defaultTransformerFeatures = `{
	"routerTransform": {
	  "MARKETO": true,
	  "HS": true
	},
	"regulations": ["AM"],
	"supportSourceTransformV1": true,
	"upgradedToSourceTransformV2": true,
  }`

func NewFeaturesService(ctx context.Context, config *config.Config, featConfig FeaturesServiceOptions) FeaturesService {
	_ = "STUB: not implemented"
	return *new(FeaturesService)
}

func NewNoOpService() FeaturesService { _ = "STUB: not implemented"; return *new(FeaturesService) }

type noopService struct{}

func (*noopService) Regulations() []string { _ = "STUB: not implemented"; return nil }

func (*noopService) SourceTransformerVersion() string {
	_ = "STUB: not implemented"
	// v0 is deprecated and upgrading to v2
	return ""
}

func (*noopService) TransformerProxyVersion() string { _ = "STUB: not implemented"; return "" }

func (*noopService) Wait() chan struct{} { _ = "STUB: not implemented"; return nil }

func (*noopService) RouterTransform(_ string) bool { _ = "STUB: not implemented"; return false }

func (*noopService) SupportDestTransformCompactedPayloadV1() bool {
	_ = "STUB: not implemented"
	return false
}
