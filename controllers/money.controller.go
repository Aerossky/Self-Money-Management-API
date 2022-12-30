package controllers

import (
	"net/http"
	"self_money_management_api_golang/helpers"
	"self_money_management_api_golang/models"

	"io/ioutil"

	"github.com/labstack/echo/v4"
	// "time"
	// "github.com/dgrijalva/jwt-go"
)

// !read
func FetchMoneyById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchMoneyById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// !total
func FetchTotalPemasukanById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchTotalPemasukanById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func FetchTotalPengeluaranById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchTotalPengeluaranById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateMoney(c echo.Context) error {

	id := helpers.ConvertStringToInt(c.FormValue("id"))
	total_money := helpers.ConvertStringToInt(c.FormValue("total_money"))
	note := c.FormValue("note")
	status := c.FormValue("status")

	result, err := models.UpdateMoney(id, total_money, note, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteMoney(c echo.Context) error {

	id := helpers.ConvertStringToInt(c.FormValue("id"))

	result, err := models.DeleteMoney(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// ! data by id
func FetchDataPemasukanByUserId(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchDataPemasukanByUserId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func FetchDataPengeluaranByUserId(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchDataPengeluaranByUserId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func FetchMoneyAPI(c echo.Context) error {
	// Make the request
	response, err := http.Post("https://api.frankfurter.app/latest?amount=10000&base=IDR", "application/json", c.Request().Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Read the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(data))
}
