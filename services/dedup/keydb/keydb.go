package keydb

import (
	"time"

	"github.com/rudderlabs/keydb/client"
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/dedup/types"
)

type keyDB struct {
	client *client.Client
	window config.ValueLoader[time.Duration]
	logger logger.Logger

	stats struct {
		getTimer stats.Timer
		setTimer stats.Timer
	}
}

func NewKeyDB(conf *config.Config, stat stats.Stats, log logger.Logger) (types.DB, error) {
	_ = "STUB: not implemented"
	return *new(types.DB), nil
}

// No MaxElapsedTime, the client will retry forever.
// To detect issues monitor the client metrics:
// https://github.com/rudderlabs/keydb/blob/v0.4.2-alpha/client/client.go#L160

// After a duration of this time if the client doesn't see any activity it
// pings the server to see if the transport is still alive.

// After having pinged for keepalive check, the client waits for a duration
// of Timeout and if no activity is seen even after that the connection is
// closed.

// If false, client sends keepalive pings even with no active RPCs. If true,
// when there are no active RPCs, KeepAliveTime and KeepAliveTimeout will be ignored and no
// keepalive pings will be sent.

// BackoffBaseDelay is the amount of time to backoff after the first failure.

// BackoffMultiplier is the factor with which to multiply backoffs after a
// failed retry. Should ideally be greater than 1.

// BackoffJitter is the factor with which backoffs are randomized.

// BackoffMaxDelay is the upper bound of backoff delay.

// MinConnectTimeout is the minimum amount of time we are willing to give a
// connection to complete.

func (d *keyDB) Get(keys []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *keyDB) Set(keys []string) error { _ = "STUB: not implemented"; return nil }

func (d *keyDB) Close() { _ = "STUB: not implemented"; return }
