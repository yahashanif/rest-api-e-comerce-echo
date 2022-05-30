package controllers

import (
	"net/http"
	"rest-api-e-comerce/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RegisterCustomer(c echo.Context) error {
	username := c.FormValue("username")
	fullName := c.FormValue("full_name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.RegisterCustomer(&models.User{
		Username: username,
		FullName: fullName,
		Email:    email,
	}, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func CheckLoginCustomer(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.CheckLoginCustomer(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["level"] = "application"

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messageww": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  result,
		"token": t,
	})
}
