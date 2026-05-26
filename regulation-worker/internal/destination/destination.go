package destination

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/utils/pubsub"
)

var pkgLogger = logger.NewLogger().Child("client")

//go:generate mockgen -source=destination.go -destination=mock_destination.go -package=destination github.com/rudderlabs/rudder-server/regulation-worker/internal/Destination/destination
type destMiddleware interface {
	Subscribe(ctx context.Context, topic backendconfig.Topic) pubsub.DataChannel
}

type DestinationConfig struct {
	mu           sync.RWMutex
	destinations map[string]*backendconfig.DestinationT
	Dest         destMiddleware
}

func (d *DestinationConfig) GetDestination(destID string) (*backendconfig.DestinationT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start starts listening for configuration updates and updates the destinations.
// The method blocks until the first update is received.
func (d *DestinationConfig) Start(ctx context.Context) { _ = "STUB: not implemented"; return }
