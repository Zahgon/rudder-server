package stats

import (
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type SourceStat struct {
	Source string

	WriteKey      string
	ReqType       string
	SourceID      string
	WorkspaceID   string
	SourceType    string
	Version       string
	SourceDefName string

	reason string

	requests struct {
		total int

		succeeded  int
		failed     int
		dropped    int
		suppressed int
	}
	events struct {
		total int
		bot   int

		succeeded int
		failed    int
	}
}

// RequestSucceeded increments the requests total & succeeded counters by one
func (ss *SourceStat) RequestSucceeded() { _ = "STUB: not implemented"; return }

// RequestDropped increments the requests total & dropped counters by one
func (ss *SourceStat) RequestDropped() { _ = "STUB: not implemented"; return }

// RequestSuppressed increments the requests total & suppressed counters by one
func (ss *SourceStat) RequestSuppressed() { _ = "STUB: not implemented"; return }

// RequestFailed increments the requests total & failed counters by one
func (ss *SourceStat) RequestFailed(reason string) { _ = "STUB: not implemented"; return }

// RequestEventsSucceeded increments the requests total & succeeded counters by one, and the events total & succeeded counters by num
func (ss *SourceStat) RequestEventsSucceeded(num int) { _ = "STUB: not implemented"; return }

// RequestEventsFailed increments the requests total & failed counters by one, and the events total & failed counters by num
func (ss *SourceStat) RequestEventsFailed(num int, reason string) {
	_ = "STUB: not implemented"
	return
}

// EventsSuccess increments the events total & succeeded counters by num
func (ss *SourceStat) EventsSuccess(num int) { _ = "STUB: not implemented"; return }

// EventsFailed increments the events total & failed counters by num
func (ss *SourceStat) EventsFailed(num int, reason string) { _ = "STUB: not implemented"; return }

func (ss *SourceStat) RequestEventsBot(num int) { _ = "STUB: not implemented"; return }

// Report captured stats
func (ss *SourceStat) Report(s stats.Stats) { _ = "STUB: not implemented"; return }
