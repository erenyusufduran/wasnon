package repositories

import (
	"time"

	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	GetAll(limit int) ([]models.Product, error)
	GetOneById(id uint) (*models.Product, error) // Use a pointer to allow modifications
	GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*models.Product, error)
	UpdateProductsStatus(products []*models.Product, status models.Status) error
}

// ProductRepositoryImpl is a concrete implementation of ProductRepository using Gorm
type ProductRepositoryImpl struct {
	db *gorm.DB
}

// NewProductRepositoryImpl creates a new instance of ProductRepositoryImpl
func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

// Create adds a new product to the database
func (r *ProductRepositoryImpl) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// Update modifies an existing product in the database
func (r *ProductRepositoryImpl) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// GetAll retrieves all products from the database
func (r *ProductRepositoryImpl) GetAll(limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Limit(limit).Find(&products).Error
	return products, err
}

// GetOneById retrieves a single product by ID
func (r *ProductRepositoryImpl) GetOneById(id uint) (*models.Product, error) {
	var product *models.Product
	err := r.db.First(product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*models.Product, error) {
	var expiredProducts []*models.Product
	err := r.db.Where("expiration < ? AND status = ?", currentTime, "active").Limit(limit).Find(&expiredProducts).Error
	return expiredProducts, err
}

func (r *ProductRepositoryImpl) UpdateProductsStatus(products []*models.Product, status models.Status) error {
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
