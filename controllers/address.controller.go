package controllers

import (
	"net/http"
	"rest-api-e-comerce/models"

	"github.com/labstack/echo/v4"
)

func FetchAllProvinsi(c echo.Context) error {
	result, err := models.FetchAllProvinsi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
func FetchAllCity(c echo.Context) error {
	Id := c.Param("id")
	result, err := models.FetchAllCity(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
func FetchAllSubdistrict(c echo.Context) error {
	Id := c.Param("id")
	result, err := models.FetchAllSubdistrict(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllCourier(c echo.Context) error {
	result, err := models.FetchAllCourier()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllCourierAktif(c echo.Context) error {
	result, err := models.FetchAllCourierAktif()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func Ongkir(c echo.Context) error {
	result, err := models.Ongkir("318", "city", "345", "city", "1000")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func AddAddress(c echo.Context) error {
	IdUser := c.FormValue("id_user")
	IdProvince := c.FormValue("id_province")
	Province := c.FormValue("province")
	IdCity := c.FormValue("id_city")
	City := c.FormValue("city")
	IdSubdistrict := c.FormValue("id_subdistrict")
	Subdistrict := c.FormValue("subdistrict")
	PostalCode := c.FormValue("postal_code")
	DetailAddress := c.FormValue("detail_address")

	result, err := models.AddAddress(&models.Address{
		IdUser:        IdUser,
		IdProvince:    IdProvince,
		Province:      Province,
		IdCity:        IdCity,
		City:          City,
		IdSubdistrict: IdSubdistrict,
		Subdistrict:   Subdistrict,
		PostalCode:    PostalCode,
		DetailAddress: DetailAddress,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func EditAddress(c echo.Context) error {
	Id := c.FormValue("id")
	IdUser := c.FormValue("id_user")
	IdProvince := c.FormValue("id_province")
	Province := c.FormValue("province")
	IdCity := c.FormValue("id_city")
	City := c.FormValue("city")
	IdSubdistrict := c.FormValue("id_subdistrict")
	Subdistrict := c.FormValue("subdistrict")
	PostalCode := c.FormValue("postal_code")
	DetailAddress := c.FormValue("detail_address")

	result, err := models.EditAddress(&models.Address{
		Id:            Id,
		IdUser:        IdUser,
		IdProvince:    IdProvince,
		Province:      Province,
		IdCity:        IdCity,
		City:          City,
		IdSubdistrict: IdSubdistrict,
		Subdistrict:   Subdistrict,
		PostalCode:    PostalCode,
		DetailAddress: DetailAddress,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAddressByIdUSer(c echo.Context) error {
	idUser := c.Param("id")

	result, err := models.FetchAddressByIdUSer(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
