package worker

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	Name             string
	Schedule         time.Duration
	onTick           func()
	quit             chan struct{}
	wg               *sync.WaitGroup
	mu               sync.Mutex
	isRunning        bool
	isTaskProcessing bool
	nextRun          time.Time
}

type WorkerStatus struct {
	IsRunning        bool
	IsTaskProcessing bool
	NextRun          time.Time
}

// New creates a new worker with specified behavior
func New(name string, schedule time.Duration, onTick func()) *Worker {
	return &Worker{
		Name:             name,
		Schedule:         schedule,
		onTick:           onTick,
		quit:             make(chan struct{}),
		wg:               &sync.WaitGroup{},
		isRunning:        false,
		isTaskProcessing: false,
		nextRun:          time.Time{}, // Initialize to zero value
	}
}

// Start begins the worker's task loop
func (w *Worker) Start() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.isRunning {
		return fmt.Errorf("worker %s is already running", w.Name)
	}

	w.isRunning = true
	w.quit = make(chan struct{}) // reset the quit channel for restarting

	go func() {
		for {
			duration := w.Schedule
			w.mu.Lock()
			w.nextRun = time.Now().Add(duration) // Set the next run time
			w.mu.Unlock()

			select {
			case <-time.After(duration):
				w.mu.Lock()

				if w.isTaskProcessing {
					w.mu.Unlock()
					continue
				}

				w.isTaskProcessing = true
				w.mu.Unlock()

				w.wg.Add(1)
				go func() {
					defer func() {
						w.mu.Lock()
						w.isTaskProcessing = false
						w.mu.Unlock()
						w.wg.Done()
					}()

					w.onTick()
				}()
			case <-w.quit:
				fmt.Printf("Worker %s received shutdown signal\n", w.Name)
				w.wg.Wait() // Wait for all tasks to complete before returning
				w.mu.Lock()
				w.isRunning = false
				w.mu.Unlock()
				return
			}
		}
	}()

	return nil
}

// Stop signals the worker to stop and waits for all tasks to complete
func (w *Worker) Stop() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.isRunning {
		return fmt.Errorf("worker %s is not running", w.Name)
	}

	if w.quit != nil {
		close(w.quit)
	}

	go func() {
		w.wg.Wait() // This will still wait for all running tasks to complete
		w.isRunning = false
		w.nextRun = time.Time{}
		fmt.Printf("Worker %s stopped\n", w.Name)
	}()

	return nil
}

func (w *Worker) Status() WorkerStatus {
	w.mu.Lock()
	defer w.mu.Unlock()

	return WorkerStatus{
		IsRunning:        w.isRunning,
		IsTaskProcessing: w.isTaskProcessing,
		NextRun:          w.nextRun,
	}
}
