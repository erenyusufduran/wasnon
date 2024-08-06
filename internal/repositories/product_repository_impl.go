package repositories

import (
	"time"

	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

// GormProductRepository is a concrete implementation of ProductRepository using Gorm
type GormProductRepository struct {
	db *gorm.DB
}

// NewGormProductRepository creates a new instance of GormProductRepository
func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

// Create adds a new product to the database
func (r *GormProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// Update modifies an existing product in the database
func (r *GormProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// GetAll retrieves all products from the database
func (r *GormProductRepository) GetAll(limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Limit(limit).Find(&products).Error
	return products, err
}

// GetOneById retrieves a single product by ID
func (r *GormProductRepository) GetOneById(id uint) (*models.Product, error) {
	var product *models.Product
	err := r.db.First(product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *GormProductRepository) GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*models.Product, error) {
	var expiredProducts []*models.Product
	err := r.db.Where("expiration < ? AND status = ?", currentTime, "active").Limit(limit).Find(&expiredProducts).Error
	return expiredProducts, err
}

func (r *GormProductRepository) UpdateProductsStatus(products []*models.Product, status string) error {
	productIds := make([]uint, len(products))
	for i, product := range products {
		productIds[i] = product.ID
	}

	if err := r.db.Model(&models.Product{}).
		Where("id IN ?", productIds).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
