package product

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	db *pgxpool.Pool
}

// NewProductRepositoryImpl creates a new instance of ProductRepositoryImpl
func NewProductRepositoryImpl(db *pgxpool.Pool) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

// Create adds a new product to the database
func (r *ProductRepositoryImpl) Create(product *Product) error {
	query := `
		INSERT INTO 
		products (name, description, price, expiration, company_id, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	_, err := r.db.Exec(context.Background(), query, product.Name, product.Description, product.Price, product.Expiration, product.CompanyID, product.Status)
	if err != nil {
		return err
	}

	return nil
}

// Update modifies an existing product in the database
func (r *ProductRepositoryImpl) Update(product *Product) error {
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, expiration = $4, status = $5
		WHERE id = $6
	`
	_, err := r.db.Exec(context.Background(), query,
		product.Name,
		product.Description,
		product.Price,
		product.Expiration,
		product.Status,
		product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetAll retrieves all products from the database
func (r *ProductRepositoryImpl) GetAll(limit int) ([]Product, error) {
	query := `SELECT * FROM products LIMIT $1`
	rows, err := r.db.Query(context.Background(), query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Product])
}

// GetOneById retrieves a single product by ID
func (r *ProductRepositoryImpl) GetOneById(id uint) (*Product, error) {
	query := `SELECT id, name, description, price, expiration, company_id, status FROM products WHERE id = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return &Product{}, err
	}
	defer rows.Close()

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Product])
	if err != nil {
		if err == pgx.ErrNoRows {
			return &Product{}, fmt.Errorf("product with id %d not found", id)
		}
		return &Product{}, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) GetActiveExpiredProducts(currentTime time.Time, limit int) ([]*Product, error) {
	query := `SELECT * FROM products WHERE expiration < $1 AND status = $2 LIMIT $3`
	rows, err := r.db.Query(context.Background(), query, currentTime, Active, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	productPtrs := make([]*Product, len(products))
	for i, product := range products {
		productPtrs[i] = &product
	}

	return productPtrs, nil
}

func (r *ProductRepositoryImpl) UpdateProductsStatus(products []*Product, status Status) error {
	productIds := make([]uint, len(products))
	for i, product := range products {
		productIds[i] = product.ID
	}

	query := `
		UPDATE products 
		SET status = $1
		WHERE id = ANY($2)
	`

	_, err := r.db.Exec(context.Background(), query, status, productIds)
	if err != nil {
		return err
	}

	return nil
}
