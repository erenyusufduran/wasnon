package workers

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	Name      string
	Interval  time.Duration
	onTick    func()
	quit      chan struct{}
	wg        *sync.WaitGroup
	mu        sync.Mutex
	isRunning bool
}

// New creates a new worker with specified behavior
func New(name string, interval time.Duration, onTick func()) *Worker {
	return &Worker{
		Name:      name,
		Interval:  interval,
		onTick:    onTick,
		quit:      make(chan struct{}),
		wg:        &sync.WaitGroup{},
		isRunning: false,
	}
}

// Start begins the worker's task loop
func (w *Worker) Start() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.isRunning {
		fmt.Printf("Worker %s is already running\n", w.Name)
		return
	}

	w.isRunning = true
	w.quit = make(chan struct{}) // reset the quit channel for restarting

	ticker := time.NewTicker(w.Interval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				w.wg.Add(1)
				go func() {
					defer w.wg.Done()
					w.onTick()
				}()
			case <-w.quit:
				fmt.Printf("Worker %s received shutdown signal\n", w.Name)
				w.isRunning = false
				return
			}
		}
	}()
}

// Stop signals the worker to stop and waits for all tasks to complete
func (w *Worker) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.isRunning {
		fmt.Printf("Worker %s is not running\n", w.Name)
		return
	}

	close(w.quit)
	w.wg.Wait()
	w.isRunning = false
	fmt.Printf("Worker %s stopped\n", w.Name)
}

func (w *Worker) Status() string {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.isRunning {
		return "running"
	}
	return "stopped"
}
