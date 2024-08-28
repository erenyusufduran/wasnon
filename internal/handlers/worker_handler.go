package handlers

import (
	"net/http"

	"github.com/erenyusufduran/wasnon/internal/workers"
	"github.com/labstack/echo/v4"
)

// WorkerHandler provides HTTP handlers for worker control
type WorkerHandler struct{}

// NewWorkerHandler creates a new WorkerHandler
func NewWorkerHandler() *WorkerHandler {
	return &WorkerHandler{}
}

// StartWorker starts a worker by name
func (h *WorkerHandler) StartWorker(c echo.Context) error {
	name := c.Param("name")
	workers.Start(name)
	return c.JSON(http.StatusOK, echo.Map{"status": "Worker started", "worker": name})
}

// StopWorker stops a worker by name
func (h *WorkerHandler) StopWorker(c echo.Context) error {
	name := c.Param("name")
	workers.Stop(name)
	return c.JSON(http.StatusOK, echo.Map{"status": "Worker will stop after task processing finish.", "worker": name})
}

func (h *WorkerHandler) CheckStatus(c echo.Context) error {
	statuses := make(map[string]workers.WorkerStatus)

	for name, worker := range workers.Workers() {
		statuses[name] = worker.Status()
	}

	return c.JSON(http.StatusOK, echo.Map{"workers": statuses})
}
