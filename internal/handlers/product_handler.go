package handlers

import (
	"net/http"
	"time"

	"github.com/erenyusufduran/wasnon/internal/dto"
	"github.com/erenyusufduran/wasnon/internal/models"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	repo repositories.ProductRepository
}

// NewProductHandler creates a new ProductHandler with the given repository
func NewProductHandler(repo repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req dto.CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Parse the expiration date
	expiration, err := time.Parse("2006-01-02", req.Expiration)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid expiration date format"})
	}

	// Use the NewProduct constructor to create a new product
	product := models.NewProduct(req.Name, req.Description, req.Price, expiration, req.CompanyID)

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

	// Prepare response
	response := make([]echo.Map, len(products))
	for i, product := range products {
		response[i] = echo.Map{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"expiration":  product.Expiration.Format("2006-01-02"),
			"company_id":  product.CompanyID,
			"status":      product.Status,
		}
	}

	return c.JSON(http.StatusOK, response)
}

// ListProducts retrieves all products and returns them as a JSON array
func (h *ProductHandler) ApproveProduct(c echo.Context) error {
	var req dto.ApproveProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	product, err := h.repo.GetOneById(req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve product"})
	}

	if product.Status != models.Waiting {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Product is not at the waiting status"})
	}

	if product.Expiration.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Product's expiration date is past."})
	}

	product.Status = models.Active
	if err := h.repo.Update(product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error when updating"})
	}

	return c.JSON(http.StatusCreated, product)
}
