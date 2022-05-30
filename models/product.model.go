package models

import (
	"fmt"
	"net/http"
	"rest-api-e-comerce/db"
	"strconv"
)

type Product struct {
	Id            int             `json:"id"`
	Name          string          `json:"name"`
	Category      Category        `json:"category"`
	Merk          string          `json:"merk"`
	Harga         int             `json:"harga"`
	Description   string          `json:"description"`
	ProductDetail []ProductDetail `json:"product_details"`
	ProductImage  []ProductImage  `json:"product_image"`
}

type ProductDetail struct {
	Id       int `json:"id" query:"id"`
	Size     int `json:"size" query:"size"`
	Quantity int `json:"quantity" query:"quantity"`
}

type ProductImage struct {
	UrlImage string `json:"url_image"`
}

type Category struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	UrlPhoto string `json:"url_photo"`
}

func StoreProduct(p *Product, image []string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO `products` (`id`, `name`, `id_category`, `merk`, `harga`, `description`) VALUES (NULL, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(p.Name, p.Category.Id, p.Merk, p.Harga, p.Description)
	if err != nil {
		return res, err
	}

	LastIDInsert, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	for _, url := range image {
		sqlStatementImage := "INSERT INTO `product_image` (`id_product`, `url_image`) VALUES (?, ?)"
		stmtImage, err := con.Prepare(sqlStatementImage)

		if err != nil {
			return res, err
		}

		resultImage, err := stmtImage.Exec(LastIDInsert, url)

		if err != nil {
			return res, err
		}

		fmt.Println(resultImage)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES INPUT Products"
	res.Data = map[string]int{
		"LASTIDINSERT": int(LastIDInsert),
	}
	return res, nil

}

func FetchAllProduct() (Response, error) {
	var product Product
	var arrProduct []Product

	var productImage ProductImage
	var arrProductImage []ProductImage

	var productDetails ProductDetail
	var arrProductDetails []ProductDetail

	var idCategory string
	var idProduct string

	var res Response

	con := db.CreateCon()
	sqlStatement := "Select * from products"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &idCategory, &product.Merk, &product.Harga, &product.Description)
		if err != nil {
			return res, err
		}
		fmt.Println(product.Id)
		idProduct = strconv.Itoa(product.Id)

		sqlStatementCategory := "Select * from category where id = " + idCategory
		rowsCategory, err := con.Query(sqlStatementCategory)
		if err != nil {
			return res, err
		}

		for rowsCategory.Next() {
			err = rowsCategory.Scan(&product.Category.Id, &product.Category.Category, &product.Category.UrlPhoto)
			if err != nil {
				return res, err
			}

		}
		sqlStatementProductImage := "Select url_image from product_image where id_product = " + idProduct
		fmt.Println(sqlStatementProductImage)
		rowsProductImage, err := con.Query(sqlStatementProductImage)
		if err != nil {
			return res, err
		}

		arrProductImage = nil
		for rowsProductImage.Next() {
			err = rowsProductImage.Scan(&productImage.UrlImage)
			if err != nil {
				return res, err
			}
			arrProductImage = append(arrProductImage, productImage)
			product.ProductImage = arrProductImage

		}
		sqlStatementProductDetail := "Select id,size,quantity from product_details where id_product = " + idProduct
		fmt.Println(sqlStatementProductDetail)
		rowsProductDetail, err := con.Query(sqlStatementProductDetail)
		if err != nil {
			return res, err
		}

		arrProductDetails = nil
		for rowsProductDetail.Next() {
			err = rowsProductDetail.Scan(&productDetails.Id, &productDetails.Size, &productDetails.Quantity)
			if err != nil {
				return res, err
			}
			arrProductDetails = append(arrProductDetails, productDetails)
			product.ProductDetail = arrProductDetails

		}

		arrProduct = append(arrProduct, product)
	}
	res.Status = http.StatusOK
	res.Message = "SUKSES GET DATA PRODUCT"
	res.Data = arrProduct

	return res, nil
}

func InsertProductDetail(pd *ProductDetail, idProduct int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatementImage := "INSERT INTO `product_details` (`id`, `id_product`, `size`, `quantity`) VALUES (NULL, ?, ?, ?);"
	stmtImage, err := con.Prepare(sqlStatementImage)

	if err != nil {
		return res, err
	}

	resultImage, err := stmtImage.Exec(idProduct, pd.Size, pd.Quantity)

	if err != nil {
		return res, err
	}

	fmt.Println(resultImage)

	res.Status = http.StatusOK
	res.Message = "SUKSES INPUT Products"
	res.Data = map[string]int{
		"LAST INSERT ID PRODUCT ": idProduct,
	}
	return res, nil
}

func StoreCategory(k *Category) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO `category` (`id`, `category`, `url_photo`) VALUES (NULL, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(k.Category, k.UrlPhoto)
	if err != nil {
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES UPLOAD CATEGORY"
	res.Data = map[string]interface{}{
		"Result": result,
	}

	return res, nil
}

func FetchAllCategory() (Response, error) {
	var cat Category
	var arrCat []Category

	var res Response

	con := db.CreateCon()

	sqlStatement := "Select * from category"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&cat.Id, &cat.Category, &cat.UrlPhoto)
		if err != nil {
			return res, err
		}

		arrCat = append(arrCat, cat)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSESS GET CATEGORY"
	res.Data = arrCat

	return res, nil
}

func UpdateCategory(k *Category) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE `category` SET `category` = ?, `url_photo` = ? WHERE `category`.`id` = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(k.Category, k.UrlPhoto, k.Id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES UPDATE CATEGORY"
	res.Data = result

	return res, nil
}

func DeleteCategory(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE from `category` where id=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES DELETE CATEGORI"
	res.Data = result

	return res, nil
}
