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
