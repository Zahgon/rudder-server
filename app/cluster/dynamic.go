package cluster

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/utils/types/servermode"
)

var (
	controller     = "ETCD"
	controllerType = "Dynamic"
)

type ChangeEventProvider interface {
	ServerMode(ctx context.Context) <-chan servermode.ChangeEvent
	// EtcdClient returns an etcd client if supported by the provider. Otherwise, it returns an error.
	EtcdClient() (etcdclient.Client, error)
}

type lifecycle interface {
	Start() error
	Stop()
}

type Dynamic struct {
	Provider ChangeEventProvider

	GatewayComponent bool

	GatewayDB     lifecycle
	RouterDB      lifecycle
	BatchRouterDB lifecycle
	EventSchemaDB lifecycle
	ArchivalDB    lifecycle

	PartitionMigrator lifecycle
	Processor         lifecycle
	Router            lifecycle

	SchemaForwarder lifecycle
	Archiver        lifecycle

	currentMode servermode.Mode

	serverStartTimeStat  stats.Measurement
	serverStopTimeStat   stats.Measurement
	serverStartCountStat stats.Measurement
	serverStopCountStat  stats.Measurement

	logger logger.Logger

	once sync.Once
}

func (d *Dynamic) init() {
	d.currentMode = servermode.DegradedMode
	d.logger = logger.NewLogger().Child("cluster")
	tag := stats.Tags{
		"controlled_by":   controller,
		"controller_type": controllerType,
	}
	d.serverStartTimeStat = stats.Default.NewTaggedStat("cluster.server_start_time", stats.TimerType, tag)
	d.serverStopTimeStat = stats.Default.NewTaggedStat("cluster.server_stop_time", stats.TimerType, tag)
	d.serverStartCountStat = stats.Default.NewTaggedStat("cluster.server_start_count", stats.CountType, tag)
	d.serverStopCountStat = stats.Default.NewTaggedStat("cluster.server_stop_count", stats.CountType, tag)
}

func (d *Dynamic) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *Dynamic) start() error { _ = "STUB: not implemented"; return nil }

func (d *Dynamic) stop() { _ = "STUB: not implemented"; return }

func (d *Dynamic) handleModeChange(newMode servermode.Mode) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dynamic) Mode() servermode.Mode { _ = "STUB: not implemented"; return *new(servermode.Mode) }
