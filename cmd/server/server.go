package main

import (
	"net/http"

	"github.com/erenyusufduran/wasnon/internal/branch"
	"github.com/erenyusufduran/wasnon/internal/company"
	"github.com/erenyusufduran/wasnon/internal/product"
	"github.com/erenyusufduran/wasnon/pkg/worker"
	"github.com/erenyusufduran/wasnon/shared/validator"
	"github.com/labstack/echo/v4"
)

// Initialize creates and configures an Echo instance with routes
func InitializeServer(repositories *Repositories) *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()

	// Initialize handlers
	productHandler := product.NewProductHandler(repositories.ProductRepository)
	companyHandler := company.NewCompanyHandler(repositories.CompanyRepository)
	branchHandler := branch.NewBranchHandler(repositories.BranchRepository)

	product.RegisterRoutes(e, productHandler)
	company.RegisterRoutes(e, companyHandler)
	branch.RegisterRoutes(e, branchHandler)
	registerWorkerRoutes(e)

	return e
}

func registerWorkerRoutes(e *echo.Echo) {
	e.POST("/workers/start/:name", func(c echo.Context) error {
		name := c.Param("name")
		err := worker.Start(name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"worker": name, "error": err.Error()})

		}
		return c.JSON(http.StatusOK, echo.Map{"worker": name, "message": "Worker started"})
	})
	e.POST("/workers/stop/:name", func(c echo.Context) error {
		name := c.Param("name")
		err := worker.Stop(name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"worker": name, "error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"worker": name, "message": "Worker will stop after task processing finish."})

	})
	e.GET("/workers/status", func(c echo.Context) error {
		statuses := make(map[string]worker.WorkerStatus)

		for name, worker := range worker.Workers {
			statuses[name] = worker.Status()
		}

		return c.JSON(http.StatusOK, echo.Map{"workers": statuses})
	})
}
