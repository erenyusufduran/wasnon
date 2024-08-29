package worker

import (
	"fmt"
	"time"
)

var Workers = make(map[string]*Worker)

// WorkerConfig holds the configuration for a worker
type WorkerConfig struct {
	Name     string
	Schedule time.Duration
	OnTick   func()
}

// Initialize sets up the workers with necessary repositories
func Initialize(workerConfigs []WorkerConfig) error {
	for _, config := range workerConfigs {
		Workers[config.Name] = New(config.Name, config.Schedule, config.OnTick)
	}

	return StartAll()
}

// StartWorker starts a worker by name
func Start(name string) error {
	worker, exists := Workers[name]
	if !exists {
		return fmt.Errorf("worker %s not found", name)
	}

	return worker.Start()
}

// StopWorker stops a worker by name
func Stop(name string) error {
	worker, exists := Workers[name]
	if !exists {
		return fmt.Errorf("worker %s not found", name)
	}

	return worker.Stop()
}

func StartAll() error {
	for _, worker := range Workers {
		err := worker.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func StopAll(wait bool) error {
	for _, worker := range Workers {
		if wait {
			worker.wg.Wait()
		}
		err := worker.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
