package controllers

import (
	"net/http"
	"self_money_management_api_golang/helpers"
	"self_money_management_api_golang/models"

	// "github.com/go-playground/validator/translations/id"
	"github.com/labstack/echo/v4"
	// "time"
	// "github.com/dgrijalva/jwt-go"
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

// function checklogin and get user id from model checklogin without token
func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.CheckLogin(email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	//make return data to json
	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"user_id": result,
			"message": "login success",
		})

}

// function get user data from id
func FetchUserById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchUserById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}