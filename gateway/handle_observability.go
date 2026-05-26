package gateway

import (
	"github.com/rudderlabs/rudder-go-kit/stats"
	"github.com/rudderlabs/rudder-schemas/go/stream"

	gwstats "github.com/rudderlabs/rudder-server/gateway/internal/stats"
	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

// NewSourceStat creates a new source stat for a gateway request
func (gw *Handle) NewSourceStat(arctx *gwtypes.AuthRequestContext, reqType string) *gwstats.SourceStat {
	_ = "STUB: not implemented"
	return nil
}

func newSourceStatReporter(arctx *gwtypes.AuthRequestContext, reqType string) gwtypes.StatReporter {
	_ = "STUB: not implemented"
	return *new(gwtypes.StatReporter)
}

func (gw *Handle) newSourceStatTagsWithReason(properties stream.MessageProperties, reqType, reason, writeKey, name string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
