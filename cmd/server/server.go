package main

import (
	"github.com/erenyusufduran/wasnon/internal/handlers"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// Initialize creates and configures an Echo instance with routes
func InitializeServer(db *gorm.DB, repositories *repositories.Repositories) *echo.Echo {
	e := echo.New()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(repositories.ProductRepository)
	companyHandler := handlers.NewCompanyHandler(repositories.CompanyRepository)
	workerHandler := handlers.NewWorkerHandler()

	registerRoutes(e, productHandler, companyHandler, workerHandler)

	return e
}

// registerRoutes sets up the routes for the application
func registerRoutes(
	e *echo.Echo,
	productHandler *handlers.ProductHandler,
	companyHandler *handlers.CompanyHandler,
	workerHandler *handlers.WorkerHandler) {
	registerCompanyRoutes(e, companyHandler)
	registerProductRoutes(e, productHandler)
	registerWorkerRoutes(e, workerHandler)
}

func registerCompanyRoutes(e *echo.Echo, companyHandler *handlers.CompanyHandler) {
	e.POST("/companies", companyHandler.CreateCompany)
	e.GET("/companies", companyHandler.ListCompanies)
}

func registerProductRoutes(e *echo.Echo, productHandler *handlers.ProductHandler) {
	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.ListProducts)
}

func registerWorkerRoutes(e *echo.Echo, workerHandler *handlers.WorkerHandler) {
	e.POST("/workers/start/:name", workerHandler.StartWorker)
	e.POST("/workers/stop/:name", workerHandler.StopWorker)
	e.GET("/workers/status", workerHandler.CheckStatus)
}
