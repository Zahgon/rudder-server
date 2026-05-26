package router

import (
	"context"
	"fmt"
	"time"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

var (
	errUploadFrequencyExceeded          = fmt.Errorf("upload frequency exceeded")
	errCurrentTimeExistsInExcludeWindow = fmt.Errorf("current time exists in exclude window")
	errBeforeScheduledTime              = fmt.Errorf("before scheduled time")
	errManualSyncModeEnabled            = fmt.Errorf("manual sync mode is enabled, automatic uploads are blocked")
)

type createUploadAlwaysLoader interface {
	Load() bool
}

// canCreateUpload indicates if an upload can be started now for the warehouse based on its configured schedule
func (r *Router) canCreateUpload(ctx context.Context, warehouse model.Warehouse) error {
	_ = "STUB: not implemented"
	// can be set from rudder-cli to force uploads always
	return nil
}

// return true if the upload was triggered

// check if manual sync mode is enabled

// gets exclude window start time and end time

// start upload only if no upload has started in current window
// e.g. with prev scheduled time 14:00 and current time 15:00, start only if prev upload hasn't started after 14:00

func excludeWindowStartEndTimes(excludeWindow map[string]any) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func checkCurrentTimeExistsInExcludeWindow(currentTime time.Time, windowStartTime, windowEndTime string) bool {
	_ = "STUB: not implemented"
	return false
}

// startTime, currentTime, endTime: 05:09, 06:19, 09:07 - > window between this day 05:09 and 09:07

// startTime, currentTime, endTime: 22:09, 06:19, 09:07 -> window between this day 22:09 and tomorrow 09:07

// startTime, currentTime, endTime: 22:09, 23:19, 09:07 -> window between this day 22:09 and tomorrow 09:07

// prevScheduledTime returns the closest previous scheduled time
// e.g. Syncing every 3hrs starting at 13:00 (scheduled times: 13:00, 16:00, 19:00, 22:00, 01:00, 04:00, 07:00, 10:00)
// prev scheduled time for current time (e.g. 18:00 -> 16:00 same day, 00:30 -> 22:00 prev day)
func (r *Router) prevScheduledTime(syncFrequency, syncStartAt string, currTime time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// current time in minutes since start of day

// get position where current time can fit in the sorted list of allStartTimes

// case when currTime is greater than all the day's start time

// case when currTime is less than all the day's start time

// if current time is less than first start time in a day, take last start time in prev day

// scheduledTimes returns all possible start times (minutes from start of day) as per schedule
// e.g. Syncing every 3hrs starting at 13:00 (scheduled times: 13:00, 16:00, 19:00, 22:00, 01:00, 04:00, 07:00, 10:00)
func (r *Router) scheduledTimes(syncFrequency, syncStartAt string) []int {
	_ = "STUB: not implemented"
	return nil
}
