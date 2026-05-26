package dedup

import (
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/dedup/types"
)

const defaultMaxRoutines = 3000

type mirrorDB struct {
	primary       types.DB
	mirror        types.DB
	mode          mode
	group         *errgroup.Group
	groupLimit    int
	errs          chan error
	stopPrintErrs chan struct{}

	logger logger.Logger

	metrics struct {
		getErrorsCount   stats.Counter
		setErrorsCount   stats.Counter
		maxRoutinesCount stats.Counter
	}
}

func NewMirrorDB(primary, mirror types.DB, mode mode, conf *config.Config, s stats.Stats, log logger.Logger) *mirrorDB {
	_ = "STUB: not implemented"
	return nil
}

func (m *mirrorDB) Get(keys []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	// Execute on primary and return results
	return nil, nil
}

// Asynchronously execute on mirror

func (m *mirrorDB) Set(keys []string) error {
	_ = "STUB: not implemented"
	// Execute on primary and return results
	return nil
}

// Asynchronously execute on mirror

func (m *mirrorDB) Close() {
	_ = "STUB: not implemented"
	// First we need to stop all mirroring goroutines
	return
}

// Then close both primary and mirror DBs

func (m *mirrorDB) logLeakyErr(err error) { _ = "STUB: not implemented"; return }

// leaky bucket to avoid filling the logs if the system fails badly

func (m *mirrorDB) printErrs(interval time.Duration) { _ = "STUB: not implemented"; return }
