package workers

import (
	"time"

	"github.com/erenyusufduran/wasnon/internal/repositories"
)

// WorkerConfig holds the configuration for a worker
type WorkerConfig struct {
	Name     string
	Schedule func() time.Duration
	OnTick   func()
}

// NewConfigs creates worker configurations for initialization
func NewWorkerConfigs(repo repositories.ProductRepository) []WorkerConfig {
	return []WorkerConfig{
		{
			Name: "product",
			Schedule: func() time.Duration {
				return scheduleByMinute(1)
			},
			OnTick: func() {
				disableExpiredProducts(repo)
			},
		},
		{
			Name: "log",
			Schedule: func() time.Duration {
				return scheduleSpecificTime(1, 38)
			}, OnTick: func() {
				cleanupLogs()
			},
		},
	}
}

// scheduleEveryMinute returns a duration of 1 minute
func scheduleByMinute(minute uint) time.Duration {
	return time.Duration(minute) * time.Minute
}

// scheduleEveryHour returns a duration of 1 hour
func scheduleByHour(hour uint) time.Duration {
	return time.Duration(hour) * time.Hour
}

// scheduleSpecificTime returns the duration until the next occurrence of a specific time of day
func scheduleSpecificTime(hour, minute int) time.Duration {
	now := time.Now()
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

	// If the next scheduled time is before the current time, schedule it for the next day
	if nextRun.Before(now) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return time.Until(nextRun)
}
