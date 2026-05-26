package flusher

import (
	"context"
	"time"

	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type Runner interface {
	Run()
	Stop()
}

type NOPCronRunner struct{}

func (c *NOPCronRunner) Run()  { _ = "STUB: not implemented"; return }
func (c *NOPCronRunner) Stop() { _ = "STUB: not implemented"; return }

type CronRunner struct {
	ctx    context.Context
	cancel context.CancelFunc
	g      *errgroup.Group

	stats stats.Stats
	log   logger.Logger

	instanceId    string
	table         string
	module        string
	flusher       *Flusher
	sleepInterval config.ValueLoader[time.Duration]

	flushTimer   stats.Measurement
	reportingLag stats.Measurement

	started atomic.Bool
}

func NewCronRunner(ctx context.Context, log logger.Logger, stats stats.Stats, conf *config.Config, flusher *Flusher, table, module string) *CronRunner {
	_ = "STUB: not implemented"
	return nil
}

func (c *CronRunner) initStats() { _ = "STUB: not implemented"; return }

func (c *CronRunner) Run() { _ = "STUB: not implemented"; return }

func (c *CronRunner) startFlushing(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CronRunner) Stop() { _ = "STUB: not implemented"; return }
