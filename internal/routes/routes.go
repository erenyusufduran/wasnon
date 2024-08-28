package routes

import (
	"github.com/erenyusufduran/wasnon/internal/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(
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
