package product

import (
	"time"
)

// Status represents the status of a product
type Status string

// Enum values for Status
const (
	Active  Status = "active"
	Past    Status = "past"
	Deleted Status = "deleted"
	Waiting Status = "waiting"
)

// Product represents a product in the system
type Product struct {
	ID          uint      `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Expiration  time.Time `json:"expiration"`
	CompanyID   uint      `json:"company_id"`
	Status      Status    `json:"status"`
}

// NewProduct creates a new Product with default values
func NewProduct(name, description string, price float64, expiration time.Time, companyID uint) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Expiration:  expiration,
		CompanyID:   companyID,
		Status:      Waiting, // Default status
	}
}
