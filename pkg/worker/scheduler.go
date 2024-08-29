package worker

import "time"

/* Schedule every sec, minute, hours
 * Schedule(30, time.Second)
 * Schedule(2, time.Minute)
 * Schedule(2, time.Hour)
 */
func Schedule(hour uint, t time.Duration) time.Duration {
	return time.Duration(hour) * t
}

// ScheduleSpecificTime returns the duration until the next occurrence of a specific time of day
func ScheduleSpecificTime(hour, minute int) time.Duration {
	now := time.Now()
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

	// If the next scheduled time is before the current time, schedule it for the next day
	if nextRun.Before(now) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return time.Until(nextRun)
}
