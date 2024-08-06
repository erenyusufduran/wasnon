package workers

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	Name          string
	Schedule      func() time.Duration
	onTick        func()
	quit          chan struct{}
	wg            *sync.WaitGroup
	runningMu     sync.Mutex
	taskRunningMu sync.Mutex
	isRunning     bool
	isTaskRunning bool
}

// New creates a new worker with specified behavior
func New(name string, schedule func() time.Duration, onTick func()) *Worker {
	return &Worker{
		Name:          name,
		Schedule:      schedule,
		onTick:        onTick,
		quit:          make(chan struct{}),
		wg:            &sync.WaitGroup{},
		isRunning:     false,
		isTaskRunning: false,
	}
}

// Start begins the worker's task loop
func (w *Worker) Start() {
	w.runningMu.Lock()
	defer w.runningMu.Unlock()

	if w.isRunning {
		fmt.Printf("Worker %s is already running\n", w.Name)
		return
	}

	w.isRunning = true
	w.quit = make(chan struct{}) // reset the quit channel for restarting

	go func() {
		for {
			duration := w.Schedule()
			fmt.Printf("Worker %s will start in %v\n", w.Name, duration)

			select {
			case <-time.After(duration):
				w.taskRunningMu.Lock()
				defer w.taskRunningMu.Unlock()

				if w.isTaskRunning {
					fmt.Printf("Worker %s is still processing the previous task\n", w.Name)
					continue
				}

				w.isTaskRunning = true

				w.wg.Add(1)
				go func() {
					defer func() {
						w.taskRunningMu.Lock()
						w.isTaskRunning = false
						w.taskRunningMu.Unlock()
						w.wg.Done()
					}()

					w.onTick()
				}()
			case <-w.quit:
				fmt.Printf("Worker %s received shutdown signal\n", w.Name)
				w.wg.Wait() // Wait for all tasks to complete before returning
				w.isRunning = false
				return
			}
		}
	}()
}

// Stop signals the worker to stop and waits for all tasks to complete
func (w *Worker) Stop() {
	w.runningMu.Lock()
	defer w.runningMu.Unlock()

	if !w.isRunning {
		fmt.Printf("Worker %s is not running\n", w.Name)
		return
	}

	close(w.quit)
	w.wg.Wait() // Ensure all running tasks are completed
	w.isRunning = false
	fmt.Printf("Worker %s stopped\n", w.Name)
}

func (w *Worker) Status() string {
	w.runningMu.Lock()
	defer w.runningMu.Unlock()

	if w.isRunning {
		return "running"
	}
	return "stopped"
}
