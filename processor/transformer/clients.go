//go:generate mockgen -destination=../../mocks/processor/transformer/mock_transformer_clients.go -package=mocks_transformer_clients github.com/rudderlabs/rudder-server/processor/transformer TransformerClients

package transformer

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer"
	"github.com/rudderlabs/rudder-server/processor/types"
	transformerfs "github.com/rudderlabs/rudder-server/services/transformer"
)

type DestinationClient interface {
	Transform(ctx context.Context, events []types.TransformerEvent) types.Response
}

type UserClient interface {
	Transform(ctx context.Context, events []types.TransformerEvent) types.Response
}

type TrackingPlanClient interface {
	Validate(ctx context.Context, events []types.TransformerEvent) types.Response
}

type SrcHydrationClient interface {
	Hydrate(ctx context.Context, req types.SrcHydrationRequest) (types.SrcHydrationResponse, error)
}

type Clients struct {
	user         UserClient
	userMirror   UserClient
	destination  DestinationClient
	trackingplan TrackingPlanClient
	srcHydration SrcHydrationClient
}

type TransformerClients interface {
	User() UserClient
	UserMirror() UserClient
	Destination() DestinationClient
	TrackingPlan() TrackingPlanClient
	SrcHydration() SrcHydrationClient
}

// WithFeatureService is used to set the feature service for the destination transformer.
func WithFeatureService(featuresService transformerfs.FeaturesService) func(*opts) {
	_ = "STUB: not implemented"
	return nil
}

// NewClients creates a new instance of TransformerClients.
func NewClients(conf *config.Config, log logger.Logger, statsFactory stats.Stats, options ...func(*opts)) TransformerClients {
	_ = "STUB: not implemented"
	return *new(TransformerClients)
}

func (c *Clients) User() UserClient { _ = "STUB: not implemented"; return *new(UserClient) }

func (c *Clients) UserMirror() UserClient { _ = "STUB: not implemented"; return *new(UserClient) }

func (c *Clients) Destination() DestinationClient {
	_ = "STUB: not implemented"
	return *new(DestinationClient)
}

func (c *Clients) TrackingPlan() TrackingPlanClient {
	_ = "STUB: not implemented"
	return *new(TrackingPlanClient)
}

func (c *Clients) SrcHydration() SrcHydrationClient {
	_ = "STUB: not implemented"
	return *new(SrcHydrationClient)
}

type opts struct {
	destinationOpts []destination_transformer.Opt
}
