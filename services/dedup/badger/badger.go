package badger

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	obskit "github.com/rudderlabs/rudder-observability-kit/go/labels"

	"github.com/rudderlabs/rudder-server/rruntime"
	"github.com/rudderlabs/rudder-server/services/dedup/types"
)

type badgerDB struct {
	logger           loggerForBadger
	badgerDB         *badger.DB
	window           config.ValueLoader[time.Duration]
	path             string
	opts             badger.Options
	cleanupOnStartup bool

	wg     sync.WaitGroup
	bgCtx  context.Context
	cancel context.CancelFunc
	stats  struct {
		getTimer stats.Timer
		setTimer stats.Timer
		lsmSize  stats.Gauge
		vlogSize stats.Gauge
		totSize  stats.Gauge
	}
}

// DefaultPath returns the default path for the deduplication service's badger DB
func DefaultPath() string { _ = "STUB: not implemented"; return "" }

func NewBadgerDB(conf *config.Config, stat stats.Stats, path string) (types.DB, error) {
	_ = "STUB: not implemented"
	return *new(types.DB), nil
}

func (d *badgerDB) Get(keys []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *badgerDB) Set(keys []string) error { _ = "STUB: not implemented"; return nil }

func (d *badgerDB) Close() { _ = "STUB: not implemented"; return }

func (d *badgerDB) init() error {
	var err error
	if d.cleanupOnStartup {
		if err = os.RemoveAll(d.path); err != nil {
			err = fmt.Errorf("removing badger db directory: %w", err)
			return err
		}
	}
	openDB := func() (dbase *badger.DB, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic during badgerdb open: %v", r)
			}
		}()
		return badger.Open(d.opts)
	}
	d.badgerDB, err = openDB()
	if err != nil {
		// corrupted or incompatible db, clean up the directory and retry
		d.logger.Errorn("Error while opening dedup badger db, cleaning up the directory",
			obskit.Error(err),
		)
		if err = os.RemoveAll(d.opts.Dir); err != nil {
			err = fmt.Errorf("removing badger db directory: %w", err)
			return err
		}
		d.badgerDB, err = openDB()
		if err != nil {
			err = fmt.Errorf("opening badger db: %w", err)
			return err
		}
	}
	d.wg.Add(1)
	rruntime.Go(func() {
		defer d.wg.Done()
		d.gcLoop()
	})
	return err
}

func (d *badgerDB) gcLoop() { _ = "STUB: not implemented"; return }

// One call would only result in removal of at max one log file.
// As an optimization, you could also immediately re-run it whenever it returns nil error
// (this is why `goto again` is used).

type loggerForBadger struct {
	logger.Logger
}

func (l loggerForBadger) Warningf(fmt string, args ...any) { _ = "STUB: not implemented"; return }
