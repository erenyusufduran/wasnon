package workers

import (
	"fmt"
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
func NewWorkerConfigs(repositories *repositories.Repositories) []WorkerConfig {
	return []WorkerConfig{
		{
			Name: "DisableExpiredProducts",
			Schedule: func() time.Duration {
				return schedule(15, time.Second)
			},
			OnTick: func() {
				err := disableExpiredProducts(repositories.ProductRepository)
				if err != nil {
					fmt.Println("There is an error at DisableExpiredProducts job")
				}
			},
		},
		{
			Name: "CleanupLogs",
			Schedule: func() time.Duration {
				return scheduleSpecificTime(12, 38)
			}, OnTick: func() {
				err := cleanupLogs()
				if err != nil {
					fmt.Println("There is an error at CleanupLogs job")
				}
			},
		},
	}
}

/* schedule every sec, minute, hours
 * schedule(30, time.Second)
 * schedule(2, time.Minute)
 * schedule(2, time.Hour)
 */
func schedule(hour uint, t time.Duration) time.Duration {
	return time.Duration(hour) * t
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
