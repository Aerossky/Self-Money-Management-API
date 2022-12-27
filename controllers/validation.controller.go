package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"self_money_management_api_golang/models"
)

type User struct {
	Nama  string `json:"nama" validate:"required"`
	Umur  int    `validate:"required,gte=17,lte=40"`
	Email string `validate:"required,email"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	org := User{
		Nama:  "Rizky",
		Umur:  18,
		Email: "xxxx@gmail.com",
	}

	err := v.Struct(org)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{
				"message": err.Error()})
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Success"})
}


func TestVarValidation(c echo.Context) error {
	var res models.Response
	v := validator.New()
	nama := "Mychael"
	umur := 19
	email := "xxxx@jiohoi.bobo"

	var errordata = make(map[string]string)

	err1 := v.Var(email, "required,email")
	if err1 != nil {
		errordata["email"] = "Email not valid."
	}

	err2 := v.Var(nama, "required")
	if err2 != nil {
		errordata["name"] = "Name is required."
	}

	err3 := v.Var(umur, "required,gte=17,lte=40")
	if err3 != nil {
		errordata["umur"] = "Umur doesn't match the requirement."
	}

	if len(errordata) != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = errordata
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Success",
		})
}
