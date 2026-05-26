package dedup

import (
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/dedup/types"
)

type mode string

const (
	badgerOnlyMode mode = "badger"
	keyDBOnlyMode  mode = "keydb"
	mirrorToKeyDB  mode = "mirrorToKeyDB"
	mirrorToBadger mode = "mirrorToBadger"
)

// getMode determines which mode to use based on configuration
func getMode(conf *config.Config) mode { _ = "STUB: not implemented"; return *new(mode) }

// NewDB creates a new DB implementation based on configuration
func NewDB(conf *config.Config, s stats.Stats, log logger.Logger) (types.DB, error) {
	_ = "STUB: not implemented"
	return *new(types.DB), nil
}

// primary is badger, then we mirror to keydb

// primary is keydb, then we mirror to badger

// Default to badger only
