package repo

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

type Opt func(*repo)

func WithNow(now func() time.Time) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithStats(s stats.Stats) Opt { _ = "STUB: not implemented"; return *new(Opt) }
