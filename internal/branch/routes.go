package branch

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, handler *BranchHandler) {
	e.POST("/branches", handler.CreateBranch)
	e.GET("/branches", handler.ListBranches)
	e.GET("/branches/:id", handler.GetBranch)
}
