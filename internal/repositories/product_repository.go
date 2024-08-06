package repositories

import (
	"time"

	"github.com/erenyusufduran/wasnon/internal/models"
)

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	GetAll(limit int) ([]models.Product, error)
	GetOneById(id uint) (*models.Product, error) // Use a pointer to allow modifications
	GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*models.Product, error)
	UpdateProductsStatus(products []*models.Product, status string) error
}
