package workers

import (
	"fmt"
	"time"

	"github.com/erenyusufduran/wasnon/internal/repositories"
)

// WorkerConfig holds the configuration for a worker
type WorkerConfig struct {
	Name     string
	Schedule time.Duration
	OnTick   func()
}

// NewConfigs creates worker configurations for initialization
func NewWorkerConfigs(repositories *repositories.Repositories) []WorkerConfig {
	return []WorkerConfig{
		{
			Name:     "DisableExpiredProducts",
			Schedule: Schedule(15, time.Second),
			OnTick: func() {
				err := disableExpiredProducts(repositories.ProductRepository)
				if err != nil {
					fmt.Println("There is an error at DisableExpiredProducts job")
				}
			},
		},
		{
			Name:     "CleanupLogs",
			Schedule: ScheduleSpecificTime(12, 38),
			OnTick: func() {
				err := cleanupLogs()
				if err != nil {
					fmt.Println("There is an error at CleanupLogs job")
				}
			},
		},
	}
}
