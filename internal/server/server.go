package server

import (
	"github.com/erenyusufduran/wasnon/internal/handlers"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/erenyusufduran/wasnon/internal/routes"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// Initialize creates and configures an Echo instance with routes
func Initialize(db *gorm.DB, productRepo repositories.ProductRepository, companyRepo repositories.CompanyRepository) *echo.Echo {
	e := echo.New()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productRepo)
	companyHandler := handlers.NewCompanyHandler(companyRepo)
	workerHandler := handlers.NewWorkerHandler()

	// Register routes
	routes.RegisterRoutes(e, productHandler, companyHandler, workerHandler)

	return e
}
