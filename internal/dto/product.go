package dto

// CreateProductRequest defines the input structure for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Expiration  string  `json:"expiration"`
	CompanyID   uint    `json:"company_id"`
}

type ApproveProductRequest struct {
	ID uint `json:"id"`
}
