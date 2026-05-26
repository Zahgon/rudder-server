// Package repo provides database repository functionality for warehouse operations.
package repo

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"

	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
)

// repo provides base repository functionality with database access and statistics tracking.
type repo struct {
	db           *sqlmiddleware.DB
	now          func() time.Time
	statsFactory stats.Stats
	repoType     string
}

// WithTx executes a function within a database transaction.
// Handles begin, commit, and rollback automatically.
func (r *repo) WithTx(ctx context.Context, f func(tx *sqlmiddleware.Tx) error) error {
	_ = "STUB: not implemented"
	return nil
}

// TimerStat returns a function that records the duration of a database action.
func (r *repo) TimerStat(action string, extraTags stats.Tags) func() {
	_ = "STUB: not implemented"
	return nil
}

func (r *repo) getRepoType() string { _ = "STUB: not implemented"; return "" }
