package apphandlers

import (
	"context"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/jobsdb"
)

type PartitionMigrator interface {
	Start() error
	Stop()
}

// ProcessorPartitionMigratorSetup holds the result of setting up the processor partition migrator.
type ProcessorPartitionMigratorSetup struct {
	PartitionMigrator PartitionMigrator
	GwDB              jobsdb.JobsDB
	RtDB              jobsdb.JobsDB
	BrtDB             jobsdb.JobsDB
	Finally           func()
}

// setupProcessorPartitionMigrator sets up the partition migrator for processor nodes (app running in processor or embedded mode)
func setupProcessorPartitionMigrator(ctx context.Context,
	shutdownFn func(), // called to initiate shutdown of the app
	dbPool *sql.DB, // database handle
	priorityPool *sql.DB, // priority database handle
	config *config.Config, stats stats.Stats,
	gwRODB, gwWODB, // gateway reader and writer jobsDB handles. if gwWODB is nil, gwRODB is used for reading and a new writer gw DB is created internally
	rtRWDB, brtRWDB jobsdb.JobsDB,
	etcdClientProvider func() (etcdclient.Client, error),
) (ProcessorPartitionMigratorSetup, error) {
	_ = "STUB: not implemented"
	return *new(ProcessorPartitionMigratorSetup), nil
}

// caller expects to get reader gw db back if writer is nil

// caller expects to get writer gw db back if writer is not nil

// setup partition buffer for gateway jobsDB

// we have separate reader and writer gw DBs, writer is externally managed
// and we are going to create a single buffer using both

// we have only a reader gw DB, so we create a buffer with reader and
// a new writer gw DB that we create here so that it can be used for flushing buffered jobs

// setup partition buffer for router jobsDB

// setup partition buffer for batchrouter jobsDB

// setup partition migrator

// this is for tests where we cannot start multiple grpc servers on the same port

func setupGatewayPartitionMigrator(ctx context.Context,
	dbPool *sql.DB, // database handle pool
	config *config.Config, stats stats.Stats,
	gwWODB jobsdb.JobsDB,
	etcdClientProvider func() (etcdclient.Client, error),
) (partitionMigrator PartitionMigrator, gwDB jobsdb.JobsDB, err error) {
	_ = "STUB: not implemented"
	return *new(PartitionMigrator), *new(jobsdb.JobsDB), nil
}

// setup partition migrator

type noOpPartitionMigrator struct{}

func (n *noOpPartitionMigrator) Start() error { _ = "STUB: not implemented"; return nil }

func (n *noOpPartitionMigrator) Stop() { _ = "STUB: not implemented"; return }
