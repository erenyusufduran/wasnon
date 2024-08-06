package workers

import (
	"time"

	"github.com/erenyusufduran/wasnon/internal/repositories"
)

// WorkerConfig holds the configuration for a worker
type WorkerConfig struct {
	Name     string
	Interval time.Duration
	OnTick   func()
}

// NewConfigs creates worker configurations for initialization
func NewWorkerConfigs(repo repositories.ProductRepository) []WorkerConfig {
	return []WorkerConfig{
		{
			Name:     "product",
			Interval: 1 * time.Minute,
			OnTick: func() {
				disableExpiredProducts(repo)
			},
		},
		{
			Name:     "log",
			Interval: 1 * time.Hour,
			OnTick: func() {
				cleanupLogs()
			},
		},
	}
}
