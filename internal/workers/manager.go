package workers

import (
	"log"
)

var workers = make(map[string]*Worker)

func Workers() map[string]*Worker {
	return workers
}

// Initialize sets up the workers with necessary repositories
func Initialize(workerConfigs []WorkerConfig) {
	for _, config := range workerConfigs {
		workers[config.Name] = New(config.Name, config.Schedule, config.OnTick)
	}

	StartAll()
}

// StartWorker starts a worker by name
func Start(name string) {
	worker, exists := workers[name]
	if !exists {
		log.Printf("Worker %s not found\n", name)
		return
	}

	worker.Start()
}

// StopWorker stops a worker by name
func Stop(name string) {
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

func StopAll(wait bool) {
	for _, worker := range workers {
		if wait {
			worker.wg.Wait()
		}
		worker.Stop()
	}
}
