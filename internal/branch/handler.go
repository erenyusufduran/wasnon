package branch

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BranchHandler struct {
	repo BranchRepository
}

// NewBranchHandler creates a new BranchHandler with the given repository
func NewBranchHandler(repo BranchRepository) *BranchHandler {
	return &BranchHandler{repo: repo}
}

func (h *BranchHandler) CreateBranch(c echo.Context) error {
	var req CreateBranchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	branch := NewBranch(req.CompanyID, req.Name, req.Address, req.City, req.County, req.Email)

	if err := h.repo.Create(branch); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, branch)
}

// ListBranchs retrieves all Branchs and returns them as a JSON array
func (h *BranchHandler) ListBranches(c echo.Context) error {
	branchs, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, branchs)
}

// ListBranchs retrieves all Branchs and returns them as a JSON array
func (h *BranchHandler) GetBranch(c echo.Context) error {
	id := c.Param("id")

	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	branch, err := h.repo.GetOneById(uintId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, branch)
}
