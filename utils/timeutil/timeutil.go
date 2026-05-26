package timeutil

import (
	"time"
)

// MinsOfDay returns minutes since start of day for a timestamp in format `15:30`
// eg. MinsOfDay("02:30") -> 150
// returns 0 for a timestamp not in format HH:MM
func MinsOfDay(str string) int { _ = "STUB: not implemented"; return 0 }

// StartOfDay returns start of the day
func StartOfDay(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Now returns the current time in UTC.
func Now() time.Time {
	_ = "STUB: not implemented"
	return *

	// GetElapsedMinsInThisDay() returns no of minutes elapsed in this day for the time provided
	new(time.Time)
}

func GetElapsedMinsInThisDay(currentTime time.Time) int { _ = "STUB: not implemented"; return 0 }
