package scenario

import (
	"context"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// NewSimple creates a jobsdb bench scenario which emulates basic behaviour:
// 1. It creates a new jobsdb instance.
// 2. It spawns w*s writer go-routines, where [w] is the number of write concurrency and [s] is the number of sources.
// 3. It spawns s reader go-routines and u*s status updating go-routines, where [s] is the number of sources and [u] is the number of update concurrency.
// 4. Readers read from jobsdb and mark jobs as succeeded.
func NewSimple(conf *config.Config, stats stats.Stats, log logger.Logger, db *sql.DB) *simple {
	_ = "STUB: not implemented"
	return nil
}

type simple struct {
	stats stats.Stats
	log   logger.Logger
	conf  *config.Config
	db    *sql.DB
}

func (p *simple) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// size of the event payload
// number of sources
// number of jobs writers go-routines for each source
// number of jobs writers write in one go
// number of jobs readers read in one go
// number of jobs status updater go-routines for each reader
// if 0, no limit will be applied on the size of the payload queried
// if 0, no limit will be applied on the number of data sets queried
// if 0, no backlog will be accumulated before processing starts

// channel used to signal that backlog has been reached

// total number of jobs written
// total number of jobs processed

// number of jobs written in the last second
// number of jobs read in the last second
// number of jobs updated in the last second

// nolint: nilerr

// we can only have one reader per source

// wait for backlog to be reached

// nolint: nilerr

// nolint: nilerr

// print stats every second
