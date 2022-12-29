package controllers

import (
	"net/http"
	"self_money_management_api_golang/helpers"
	"self_money_management_api_golang/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPlan(c echo.Context) error {

	result, err := models.FetchAllPlan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StorePlan(c echo.Context) error {
	id_user := helpers.ConvertStringToInt(c.FormValue("id_user"))
	name := c.FormValue("name")
	price := helpers.ConvertStringToInt(c.FormValue("price"))
	time := helpers.ConvertStringToInt(c.FormValue("time"))

	result, err := models.StorePlan(id_user, name, price, time)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}

func UpdatePlan(c echo.Context) error {

	id := helpers.ConvertStringToInt(c.FormValue("id"))
	// id_user := helpers.ConvertStringToInt(c.FormValue("id_user"))
	name := c.FormValue("name")
	price := helpers.ConvertStringToInt(c.FormValue("price"))
	time := helpers.ConvertStringToInt(c.FormValue("time"))

	result, err := models.UpdatePlan(id, name, price, time)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeletePlan(c echo.Context) error {

	id := helpers.ConvertStringToInt(c.FormValue("id"))

	result, err := models.DeletePlan(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// fetch plan by id
func FetchPlanById(c echo.Context) error {

	id := c.Param("id")

	result, err := models.FetchPlanById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
