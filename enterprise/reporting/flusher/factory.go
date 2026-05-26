package flusher

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

var supportedTables = []string{"tracked_users_reports"}

// This function has to be called once for each table
func CreateRunner(ctx context.Context, table string, log logger.Logger, stats stats.Stats, conf *config.Config, module string) (Runner, error) {
	_ = "STUB: not implemented"
	return *new(Runner), nil
}
