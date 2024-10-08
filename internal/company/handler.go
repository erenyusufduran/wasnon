package company

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyHandler struct {
	repo CompanyRepository
}

func NewCompanyHandler(repo CompanyRepository) *CompanyHandler {
	return &CompanyHandler{repo: repo}
}

func (h *CompanyHandler) CreateCompany(c echo.Context) error {
	var req CreateCompanyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	company := NewCompanyWithName(req.Name)

	if err := h.repo.Create(company); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, company)
}

func (h *CompanyHandler) ListCompanies(c echo.Context) error {
	companies, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) ListCompaniesWithBranches(c echo.Context) error {
	companies, err := h.repo.GetManyWithBranches()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, companies)
}
