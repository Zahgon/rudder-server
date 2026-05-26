package bench

import (
	"context"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type Bench interface {
	Run(ctx context.Context) error
}

func New(conf *config.Config, stat stats.Stats, log logger.Logger, db *sql.DB) (Bench, error) {
	_ = "STUB: not implemented"
	return *new(Bench), nil
}
