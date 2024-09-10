package product

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	repo ProductRepository
}

// NewProductHandler creates a new ProductHandler with the given repository
func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Parse the expiration date
	expiration, err := time.Parse("2006-01-02", req.Expiration)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid expiration date format"})
	}

	// Use the NewProduct constructor to create a new product
	product := NewProduct(req.Name, req.Description, req.Price, expiration, req.CompanyID)

	if err := h.repo.Create(product); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

// ListProducts retrieves all products and returns them as a JSON array
func (h *ProductHandler) ListProducts(c echo.Context) error {
	products, err := h.repo.GetAll(100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}

// ListProducts retrieves all products and returns them as a JSON array
func (h *ProductHandler) ApproveProduct(c echo.Context) error {
	var req ApproveProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	product, err := h.repo.GetOneById(req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if product.Status != Waiting {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Product is not at the waiting status"})
	}

	if product.Expiration.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Product's expiration date is past."})
	}

	product.Status = Active
	if err := h.repo.Update(product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}
