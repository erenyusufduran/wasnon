package product

import (
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *Product) error
	Update(product *Product) error
	GetAll(limit int) ([]Product, error)
	GetOneById(id uint) (*Product, error) // Use a pointer to allow modifications
	GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*Product, error)
	UpdateProductsStatus(products []*Product, status Status) error
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
func (r *ProductRepositoryImpl) Create(product *Product) error {
	return r.db.Create(product).Error
}

// Update modifies an existing product in the database
func (r *ProductRepositoryImpl) Update(product *Product) error {
	return r.db.Save(product).Error
}

// GetAll retrieves all products from the database
func (r *ProductRepositoryImpl) GetAll(limit int) ([]Product, error) {
	var products []Product
	err := r.db.Limit(limit).Find(&products).Error
	return products, err
}

// GetOneById retrieves a single product by ID
func (r *ProductRepositoryImpl) GetOneById(id uint) (*Product, error) {
	var product *Product
	err := r.db.First(product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*Product, error) {
	var expiredProducts []*Product
	err := r.db.Where("expiration < ? AND status = ?", currentTime, "active").Limit(limit).Find(&expiredProducts).Error
	return expiredProducts, err
}

func (r *ProductRepositoryImpl) UpdateProductsStatus(products []*Product, status Status) error {
	productIds := make([]uint, len(products))
	for i, product := range products {
		productIds[i] = product.ID
	}

	if err := r.db.Model(&Product{}).
		Where("id IN ?", productIds).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
