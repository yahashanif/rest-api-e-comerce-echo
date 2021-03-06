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
type Favorite struct {
	IdUser    string `json:"id_user"`
	IdProduct string `json:"id_product"`
}

type Cart struct {
	IdUser        string        `json:"id_user"`
	ProductDetail ProductDetail `json:"product_detail"`
	Product       Product       `json:"product"`
	Quantity      int           `json:"quantity"`
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
		idProduct = strconv.Itoa(product.Id)

		sqlStatementCategory := "Select * from category where id = ?"
		con.QueryRow(sqlStatementCategory, idCategory).Scan(&product.Category.Id, &product.Category.Category, &product.Category.UrlPhoto)

		sqlStatementProductImage := "Select url_image from product_image where id_product = " + idProduct
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
func FetchAllProductByCategory(IdCategory string) (Response, error) {
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
	sqlStatement := "Select * from products where id_category = " + IdCategory

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &idCategory, &product.Merk, &product.Harga, &product.Description)
		if err != nil {
			return res, err
		}
		idProduct = strconv.Itoa(product.Id)

		sqlStatementCategory := "Select * from category where id = ?"
		con.QueryRow(sqlStatementCategory, idCategory).Scan(&product.Category.Id, &product.Category.Category, &product.Category.UrlPhoto)

		sqlStatementProductImage := "Select url_image from product_image where id_product = " + idProduct
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

func FetchProductByID(Id string) (Response, error) {
	var product Product

	var productImage ProductImage
	var arrProductImage []ProductImage

	var productDetails ProductDetail
	var arrProductDetails []ProductDetail

	var idCategory string
	var idProduct string

	var res Response

	con := db.CreateCon()
	sqlStatement := "Select * from products where id =?"

	con.QueryRow(sqlStatement, Id).Scan(&product.Id, &product.Name, &idCategory, &product.Merk, &product.Harga, &product.Description)

	idProduct = strconv.Itoa(product.Id)

	sqlStatementCategory := "Select * from category where id = ?"
	con.QueryRow(sqlStatementCategory, idCategory).Scan(&product.Category.Id, &product.Category.Category, &product.Category.UrlPhoto)

	sqlStatementProductImage := "Select url_image from product_image where id_product = " + idProduct
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

	res.Status = http.StatusOK
	res.Message = "SUKSES GET DATA PRODUCT"
	res.Data = product

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

func IsFavorite(f *Favorite) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `favorite` (`id_user`, `id_product`) VALUES (?, ?);"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(f.IdUser, f.IdProduct)
	if err != nil {
		return res, nil
	}
	LastInsertID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES CREATE FAVORITE"
	res.Data = map[string]interface{}{
		"LastInsertID": LastInsertID,
	}

	return res, nil

}

func DeleteFavorite(f *Favorite) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE from `favorite` where id_user= ? and id_product= ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(f.IdUser, f.IdProduct)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES DELETE Favorite"
	res.Data = result

	return res, nil
}

func ListProductFavorite(f *Favorite) (Response, error) {
	var res Response

	var fav Favorite
	var arrFav []Favorite

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `favorite` where id_user = " + f.IdUser

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&fav.IdUser, &fav.IdProduct)
		if err != nil {
			return res, err
		}

		arrFav = append(arrFav, fav)
	}
	var arrRes []interface{}
	for _, data := range arrFav {
		ress, _ := FetchProductByID("2")
		arrRes = append(arrRes, ress.Data)
		fmt.Println(data)
	}
	res.Data = arrRes
	return res, nil
}

func AddCart(IdUser, idProductDetail, quantity string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatementCek := "Select id_product_details from cart where id_product_details = ?"
	var idcek string

	con.QueryRow(sqlStatementCek, idProductDetail).Scan(&idcek)
	if len(idcek) != 0 {
		result, err := AddQuantityCart(IdUser, idProductDetail)
		if err != nil {
			return res, err
		}

		res.Data = result
	} else {

		sqlStatement := "INSERT INTO `cart` (`id_user`, `id_product_details`, `quantity`) VALUES (?, ?, ?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(IdUser, idProductDetail, quantity)
		if err != nil {
			return res, err
		}

		lastInserId, err := result.LastInsertId()
		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "SUKSES ADD CART"
		res.Data = map[string]interface{}{
			"LastInsertID": lastInserId,
		}
	}
	return res, nil

}

func DeleteCart(Iduser, IdproductDetail string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE from `cart` where id_user= ? and id_product_details = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Iduser, IdproductDetail)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES DELETE Cart"
	res.Data = result

	return res, nil
}

func AddQuantityCart(Iduser, IdproductDetail string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "Select quantity from cart where id_user=? and id_product_details=?"
	var quantity int

	con.QueryRow(sqlStatement, Iduser, IdproductDetail).Scan(&quantity)

	quantity++

	sqlStatementUpdate := "UPDATE `cart` SET `quantity` = ? WHERE `cart`.`id_user` = ? AND `cart`.`id_product_details` = ?"

	stmt, err := con.Prepare(sqlStatementUpdate)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(quantity, Iduser, IdproductDetail)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES Tambah Quantity CART"
	res.Data = result

	return res, nil
}
func MinQuantityCart(Iduser, IdproductDetail string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "Select quantity from cart where id_user=? and id_product_details=?"
	var quantity int

	con.QueryRow(sqlStatement, Iduser, IdproductDetail).Scan(&quantity)
	if quantity == 1 {
		DeleteCart(Iduser, IdproductDetail)
		res.Status = http.StatusOK
		res.Message = "SUKSES DELETE Cart"
	} else {
		quantity--

		sqlStatementUpdate := "UPDATE `cart` SET `quantity` = ? WHERE `cart`.`id_user` = ? AND `cart`.`id_product_details` = ?"

		stmt, err := con.Prepare(sqlStatementUpdate)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(quantity, Iduser, IdproductDetail)

		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "SUKSES Kurang Quantity CART"
		res.Data = result
	}

	return res, nil
}

func ListCart(IdUser string) (Response, error) {
	var res Response

	var cart Cart
	var arrCart []Cart

	var productImage ProductImage
	var arrProductImage []ProductImage

	con := db.CreateCon()

	sqlStatement := "SELECT	cart.id_user,cart.quantity,product_details.id,product_details.size,products.id,products.name,products.merk,products.harga,products.description,category.category FROM `cart`INNER JOIN product_details ON cart.id_product_details=product_details.id INNER JOIN products ON product_details.id_product= products.id INNER JOIN category ON products.id_category=category.id where id_user = " + IdUser

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err := rows.Scan(&cart.IdUser, &cart.Quantity, &cart.ProductDetail.Id, &cart.ProductDetail.Size, &cart.Product.Id, &cart.Product.Name, &cart.Product.Merk, &cart.Product.Harga, &cart.Product.Description, &cart.Product.Category.Category)

		if err != nil {
			return res, err
		}
		idpString := strconv.Itoa(cart.Product.Id)

		sqlStatementProductImage := "SELECT url_image FROM `product_image` where id_product = " + idpString

		rowsProductImage, err := con.Query(sqlStatementProductImage)
		if err != nil {
			return res, err
		}
		arrProductImage = nil
		for rowsProductImage.Next() {
			err := rowsProductImage.Scan(&productImage.UrlImage)
			if err != nil {
				return res, err
			}

			arrProductImage = append(arrProductImage, productImage)

		}
		cart.Product.ProductImage = arrProductImage

		arrCart = append(arrCart, cart)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST CART"
	res.Data = arrCart

	return res, nil
}
