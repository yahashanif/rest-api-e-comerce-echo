package routes

import (
	"rest-api-e-comerce/controllers"
	"rest-api-e-comerce/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	auth := e.Group("/api/auth", middleware.IsAuth)
	{
		auth.GET("/provinsi", controllers.FetchAllProvinsi)
		auth.GET("/city/:id", controllers.FetchAllCity)
		auth.GET("/subdistrict/:id", controllers.FetchAllSubdistrict)
		auth.GET("/courier", controllers.FetchAllCourier)
		auth.GET("/courierAktif", controllers.FetchAllCourierAktif)

		auth.POST("/AddAddress", controllers.AddAddress)
		auth.POST("/EditAddress", controllers.EditAddress)
		auth.GET("/Address/:id", controllers.FetchAddressByIdUSer)

		auth.POST("/category", controllers.StoreCategory)
		auth.GET("/category", controllers.FetchAllCategory)
		auth.DELETE("/category/:id", controllers.DeleteCategory)

		// Products
		auth.POST("/products", controllers.StoreProduct)
		auth.GET("/product/:id", controllers.FetchProductByID)
		auth.GET("/products/Category/:id", controllers.FetchAllProductByCategory)
		auth.GET("/products", controllers.FetchAllProduct)
		auth.POST("/productDetails", controllers.InsertProductDetail)
		auth.GET("/favorite/:id", controllers.ListProductFavorite)
		auth.DELETE("/favorite", controllers.DeleteFavorite)
		auth.POST("/favorite", controllers.IsFavorite)
		// Cart
		auth.POST("/addCart", controllers.AddCart)
		auth.GET("/cart/:id_user", controllers.ListCart)
		auth.DELETE("/deleteCart", controllers.DeleteCart)
		auth.POST("/addQuantityCart", controllers.AddQuantityCart)
		auth.POST("/minQuantityCart", controllers.MinQuantityCart)

		// User

		auth.GET("/checkUser/:id", controllers.CheckUser)

	}
	// Address

	// User
	e.Static("/category/image", "files/category")
	e.Static("/product/image", "files/products")
	e.POST("/register", controllers.RegisterCustomer)
	e.POST("/loginCustomer", controllers.CheckLoginCustomer)

	e.GET("/ongkir", controllers.Ongkir)

	return e
}
