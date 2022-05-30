package routes

import (
	"rest-api-e-comerce/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	// Address
	e.GET("/provinsi", controllers.FetchAllProvinsi)
	e.GET("/city/:id", controllers.FetchAllCity)
	e.GET("/subdistrict/:id", controllers.FetchAllSubdistrict)
	e.GET("/courier", controllers.FetchAllCourier)
	e.GET("/courierAktif", controllers.FetchAllCourierAktif)

	e.POST("/category", controllers.StoreCategory)
	e.GET("/category", controllers.FetchAllCategory)
	e.DELETE("/category/:id", controllers.DeleteCategory)
	e.Static("/category/image", "files/category")

	// Products
	e.POST("/products", controllers.StoreProduct)
	e.GET("/product/:id", controllers.FetchProductByID)
	e.GET("/products/Category/:id", controllers.FetchAllProductByCategory)
	e.GET("/products", controllers.FetchAllProduct)
	e.POST("/productDetails", controllers.InsertProductDetail)
	e.GET("/ListFavorite/:id", controllers.ListProductFavorite)

	// User
	e.POST("/register", controllers.RegisterCustomer)
	e.POST("/loginCustomer", controllers.CheckLoginCustomer)

	return e
}
