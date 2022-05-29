package routes

import (
	"rest-api-e-comerce/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.POST("/category", controllers.StoreCategory)
	e.GET("/category", controllers.FetchAllCategory)
	e.DELETE("/category/:id", controllers.DeleteCategory)
	e.Static("/category/image", "files/category")

	// Products
	e.POST("/products", controllers.StoreProduct)
	e.GET("/products", controllers.FetchAllProduct)

	return e
}
