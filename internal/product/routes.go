package product

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, handler *ProductHandler) {
	e.POST("/products", handler.CreateProduct)
	e.GET("/products", handler.ListProducts)
	e.PATCH("/products/approve", handler.ApproveProduct)
}
