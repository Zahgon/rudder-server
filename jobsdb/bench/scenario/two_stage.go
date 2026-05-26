package scenario

import (
	"context"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

// NewTwoStage creates a jobsdb bench scenario which emulates a two-stage pipeline:
// 1. It creates 2 new jobsdb instances, benchone and benchtwo.
// 2. It spawns w*s writer go-routines on benchone, where [w] is the number of write concurrency and [s] is the number of sources.
// 3. It spawns s reader go-routines from benchone and u*s status updating go-routines, where [s] is the number of sources and [u] is the number of update concurrency.
// 4. benchone readers read from benchone, forward to benchtow and mark jobs as succeeded.
// 5. It spawns s reader go-routines from benchtwo and u*s status updating go-routines, where [s] is the number of sources and [u] is the number of update concurrency.
// 4. benchtwo readers read from benchtwo and mark [100-f] jobs as succeeded and [f] jobs as failed, where [f] is the failure percentage.
func NewTwoStage(conf *config.Config, stats stats.Stats, log logger.Logger, db *sql.DB) *twoStage {
	_ = "STUB: not implemented"
	return nil
}

type twoStage struct {
	stats stats.Stats
	log   logger.Logger
	conf  *config.Config
	db    *sql.DB
}

func (p *twoStage) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// size of the event payload
// number of sources
// number of jobs writers go-routines for each source
// number of jobs status updater go-routines for each reader
// number of jobs writers write in one go
// number of jobs readers read in one go
// if 0, no limit will be applied on the size of the payload queried
// if 0, no limit will be applied on the number of data sets queried
// percentage of jobs that will fail, i.e. be marked as failed
// if 0, no backlog will be accumulated before processing starts

// channel used to signal that backlog has been reached

// total number of jobs sent to benchone
// total number of jobs processed by benchtwo successfully

// number of jobs marked as succeeded in the last second
// number of jobs marked as failed in the last second

// number of jobs written to benchone in the last second
// number of jobs read from benchone in the last second
// number of jobs updated in benchone in the last second

// write jobs in benchone

// nolint: nilerr

// read and move jobs from benchone to benchtwo
// we can only have one reader per source

// wait for backlog to be reached

// nolint: nilerr

// store jobs in benchtwo

// nolint: nilerr

// mark benchone jobs as complete

// nolint: nilerr

// read and process jobs from benchtwo
// we can only have one reader per source

// nolint: nilerr

// mark benchtwo jobs as complete/aborted/failed

// respect failure percentage

// nolint: nilerr

// print stats every second
