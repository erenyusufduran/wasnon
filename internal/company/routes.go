package company

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, handler *CompanyHandler) {
	e.POST("/companies", handler.CreateCompany)
	e.GET("/companies", handler.ListCompanies)
}
