package drain_config

import (
	"context"
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

const (
	defaultPollFrequency      = 10
	defaultPollFrequencyUnits = time.Second

	defaultCleanupFrequency      = 1
	defaultCleanupFrequencyUnits = time.Hour

	defaultMaxAge      = 24
	defaultMaxAgeUnits = time.Hour

	// drain configurations

	jobRunIDKey = "drain.jobRunIDs"
)

type drainConfigManager struct {
	log  logger.Logger
	conf *config.Config
	db   *sql.DB

	done *atomic.Bool
	wg   sync.WaitGroup
}

// NewDrainConfigManager returns a drainConfigManager
// If migration fails while setting up drain config, drainConfigManager object will be returned along with error
// Consumers must handle errors and non-nil drainConfigManager object according to their use case.
func NewDrainConfigManager(conf *config.Config, log logger.Logger, stats stats.Stats) (*drainConfigManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *drainConfigManager) CleanupRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *drainConfigManager) DrainConfigRoutine(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// map to hold the config values

// holds the config values fetched from the db

// compare config values, if different set the config

func (d *drainConfigManager) Stop() { _ = "STUB: not implemented"; return }

func migrate(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

// setupDBConn sets up the database connection
func setupDBConn(conf *config.Config, stats stats.Stats) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
