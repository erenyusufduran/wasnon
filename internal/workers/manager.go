package workers

import (
	"log"
	"sync"

	"github.com/erenyusufduran/wasnon/internal/repositories"
	"gorm.io/gorm"
)

var (
	workers    = make(map[string]*Worker)
	workerLock sync.Mutex
)

func Workers() map[string]*Worker {
	return workers
}

// Initialize sets up the workers with necessary repositories
func Initialize(db *gorm.DB, productRepo repositories.ProductRepository) {
	configs := NewWorkerConfigs(productRepo)

	for _, config := range configs {
		workers[config.Name] = New(config.Name, config.Schedule, config.OnTick)
	}

	StartAll()
}

// StartWorker starts a worker by name
func Start(name string) {
	workerLock.Lock()
	defer workerLock.Unlock()

	worker, exists := workers[name]
	if !exists {
		log.Printf("Worker %s not found\n", name)
		return
	}

	worker.Start()
}

// StopWorker stops a worker by name
func Stop(name string) {
	workerLock.Lock()
	defer workerLock.Unlock()

	worker, exists := workers[name]
	if !exists {
		log.Printf("Worker %s not found\n", name)
		return
	}

	worker.Stop()
}

func StartAll() {
	for _, worker := range workers {
		worker.Start()
	}
}

func StopAll() {
	for _, worker := range workers {
		worker.Stop()
	}
}
