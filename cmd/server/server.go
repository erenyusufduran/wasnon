package main

import (
	"net/http"

	"github.com/erenyusufduran/wasnon/internal/company"
	"github.com/erenyusufduran/wasnon/internal/product"
	"github.com/erenyusufduran/wasnon/pkg/worker"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// Initialize creates and configures an Echo instance with routes
func InitializeServer(db *gorm.DB, repositories *Repositories) *echo.Echo {
	e := echo.New()

	// Initialize handlers
	productHandler := product.NewProductHandler(repositories.ProductRepository)
	companyHandler := company.NewCompanyHandler(repositories.CompanyRepository)

	registerRoutes(e, productHandler, companyHandler)

	return e
}

// registerRoutes sets up the routes for the application
func registerRoutes(
	e *echo.Echo,
	productHandler *product.ProductHandler,
	companyHandler *company.CompanyHandler) {
	registerCompanyRoutes(e, companyHandler)
	registerProductRoutes(e, productHandler)
	registerWorkerRoutes(e)
}

func registerCompanyRoutes(e *echo.Echo, companyHandler *company.CompanyHandler) {
	e.POST("/companies", companyHandler.CreateCompany)
	e.GET("/companies", companyHandler.ListCompanies)
}

func registerProductRoutes(e *echo.Echo, productHandler *product.ProductHandler) {
	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.ListProducts)
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
