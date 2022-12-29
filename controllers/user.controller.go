package controllers

import (
	"net/http"
	"self_money_management_api_golang/helpers"
	"self_money_management_api_golang/models"

	"github.com/labstack/echo/v4"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func FetchAllUser(c echo.Context) error {

	result, err := models.FetchAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StoreUser(c echo.Context) error {
	id := helpers.ConvertStringToInt(c.FormValue("id"))
	email := c.FormValue("email")
	username := c.FormValue("username")
	image := c.FormValue("image")
	password, _ := helpers.HashPassword(c.FormValue("password"))

	result, err := models.StoreUser(id, email, username, image, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateUser(c echo.Context) error {

	id := helpers.ConvertStringToInt(c.FormValue("id"))
	email := c.FormValue("email")
	username := c.FormValue("username")
	image := c.FormValue("image")
	password, _ := helpers.HashPassword(c.FormValue("password"))

	result, err := models.UpdateUser(id, email, username, image, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteUser(c echo.Context) error {

	id := c.FormValue("id")

	result, err := models.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// func CheckLogin(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	res, err := models.CheckLogin(email, password)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 			map[string]string{"message": err.Error()})
// 	}

// 	if !res {
// 		return echo.ErrUnauthorized
// 	}

// 	return c.JSON(http.StatusOK,
// 		map[string]string{
// 			"message": "login success",
// 			// "token":   mytoken,
// 		})
// }

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	mytoken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "login success",
			"token":   mytoken,
		})
}
