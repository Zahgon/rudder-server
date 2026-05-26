package switcher

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/router/throttler/internal/types"
)

// NewThrottlerSwitcher constructs a new throttler that can switch between two throttlers based on a configuration value.
func NewThrottlerSwitcher(
	useAlternative config.ValueLoader[bool],
	main, alternative types.PickupThrottler,
) types.PickupThrottler {
	_ = "STUB: not implemented"
	return *new(types.PickupThrottler)
}

type throttlerSwitcher struct {
	useAlternative config.ValueLoader[bool]
	main           types.PickupThrottler
	alternative    types.PickupThrottler
}

// CheckLimitReached checks the limit using the currently active throttler.
func (t *throttlerSwitcher) CheckLimitReached(ctx context.Context, cost int64) (limited bool, retErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ResponseCodeReceived forwards the response code to both main and alternative throttlers.
func (t *throttlerSwitcher) ResponseCodeReceived(code int) { _ = "STUB: not implemented"; return }

// Shutdown stops both main and alternative throttlers.
func (t *throttlerSwitcher) Shutdown() { _ = "STUB: not implemented"; return }

// GetLimitPerSecond returns the limit of the currently active throttler.
func (t *throttlerSwitcher) GetLimitPerSecond() int64 { _ = "STUB: not implemented"; return 0 }

// GetEventType returns the event type of the currently active throttler.
func (t *throttlerSwitcher) GetEventType() string { _ = "STUB: not implemented"; return "" }

// GetLastUsed returns the last used time of the currently active throttler.
func (t *throttlerSwitcher) GetLastUsed() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// throttler returns the currently active throttler based on the useAlternative config.
func (t *throttlerSwitcher) throttler() types.PickupThrottler {
	_ = "STUB: not implemented"
	return *new(types.PickupThrottler)
}
