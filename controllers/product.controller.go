package controllers

import (
	"io"
	"net/http"
	"os"
	"rest-api-e-comerce/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func StoreProduct(c echo.Context) error {
	name := c.FormValue("name")
	idCategory := c.FormValue("id_category")
	idCategoryInt, err := strconv.Atoi(string(idCategory))
	if err != nil {
		return err
	}
	merk := c.FormValue("merk")
	harga := c.FormValue("harga")
	hargaInt, err := strconv.Atoi(string(harga))
	if err != nil {
		return err
	}
	description := c.FormValue("description")

	form, err := c.MultipartForm()
	var nameImage []string
	if err != nil {
		return err
	}
	files := form.File["images"]

	t := time.Now().Nanosecond()

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}

		defer src.Close()

		dst, err := os.Create("files/products/" + strconv.Itoa(t) + file.Filename)

		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		nameImage = append(nameImage, strconv.Itoa(t)+file.Filename)

	}
	result, err := models.StoreProduct(&models.Product{
		Name:        name,
		Category:    models.Category{Id: idCategoryInt},
		Merk:        merk,
		Harga:       hargaInt,
		Description: description,
	}, nameImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllProduct(c echo.Context) error {
	result, err := models.FetchAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)

}
func FetchAllProductByCategory(c echo.Context) error {
	IdCategory := c.Param("id")
	result, err := models.FetchAllProductByCategory(IdCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)

}

func FetchProductByID(c echo.Context) error {
	Id := c.Param("id")
	result, err := models.FetchProductByID(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)

}

func InsertProductDetail(c echo.Context) error {
	idproduct, _ := strconv.Atoi(c.FormValue("id_product"))
	size, _ := strconv.Atoi(c.FormValue("size"))
	quantity, _ := strconv.Atoi(c.FormValue("quantity"))

	result, err := models.InsertProductDetail(&models.ProductDetail{
		Size:     size,
		Quantity: quantity,
	}, idproduct)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)

}
func StoreCategory(c echo.Context) error {
	category := c.FormValue("category")

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	var filename string

	t := time.Now().Nanosecond()

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}

		defer src.Close()

		dst, err := os.Create("files/category/" + strconv.Itoa(t) + file.Filename)

		filename = strconv.Itoa(t) + file.Filename
		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	result, err := models.StoreCategory(&models.Category{
		Category: category,
		UrlPhoto: filename,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllCategory(c echo.Context) error {
	result, err := models.FetchAllCategory()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

// Function UPDATE CATEGORY BELOM

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := models.DeleteCategory(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func ListProductFavorite(c echo.Context) error {
	idUser := c.Param("id")
	result, err := models.ListProductFavorite(&models.Favorite{
		IdUser: idUser,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func IsFavorite(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProduct := c.FormValue("id_product")
	result, err := models.IsFavorite(&models.Favorite{
		IdUser:    IdUser,
		IdProduct: IdProduct,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteFavorite(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProduct := c.FormValue("id_product")
	result, err := models.DeleteFavorite(&models.Favorite{
		IdUser:    IdUser,
		IdProduct: IdProduct,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func AddCart(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProductDetail := c.FormValue("id_product_detail")
	quantity := c.FormValue("quantity")
	result, err := models.AddCart(IdUser, IdProductDetail, quantity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
func DeleteCart(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProductDetail := c.FormValue("id_product_detail")
	result, err := models.DeleteCart(IdUser, IdProductDetail)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
func AddQuantityCart(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProductDetail := c.FormValue("id_product_detail")
	result, err := models.AddQuantityCart(IdUser, IdProductDetail)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
func MinQuantityCart(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProductDetail := c.FormValue("id_product_detail")
	result, err := models.MinQuantityCart(IdUser, IdProductDetail)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func ListCart(c echo.Context) error {
	IdUser := c.Param("id_user")

	result, err := models.ListCart(IdUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
