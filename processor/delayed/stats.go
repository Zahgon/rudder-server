package delayed

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

type eventStats struct {
	stats     stats.Stats
	threshold time.Duration
}

func NewEventStats(stats stats.Stats, config *config.Config) *eventStats {
	_ = "STUB: not implemented"
	return nil
}

func (s *eventStats) ObserveSourceEvents(source *backendconfig.SourceT, events []types.TransformerEvent) {
	_ = "STUB: not implemented"
	return
}
