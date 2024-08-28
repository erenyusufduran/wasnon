package main

import (
	"github.com/erenyusufduran/wasnon/internal/handlers"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/erenyusufduran/wasnon/internal/routes"
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

	// Register routes
	routes.RegisterRoutes(e, productHandler, companyHandler, workerHandler)

	return e
}
