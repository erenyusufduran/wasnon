package main

import (
	"fmt"
	"time"

	"github.com/erenyusufduran/wasnon/internal/product"
	"github.com/erenyusufduran/wasnon/pkg/worker"
)

// NewConfigs creates worker configurations for initialization
func NewWorkerConfigs(repo *Repositories) []worker.WorkerConfig {
	return []worker.WorkerConfig{
		{
			Name:     "DisableExpiredProducts",
			Schedule: worker.Schedule(15, time.Second),
			OnTick: func() {
				err := product.DisableExpiredProducts(repo.ProductRepository)
				if err != nil {
					fmt.Println("There is an error at DisableExpiredProducts job")
				}
			},
		},
		// {
		// 	Name:     "CleanupLogs",
		// 	Schedule: workers.ScheduleSpecificTime(12, 38),
		// 	OnTick: func() {
		// 		err := cleanupLogs()
		// 		if err != nil {
		// 			fmt.Println("There is an error at CleanupLogs job")
		// 		}
		// 	},
		// },
	}
}
