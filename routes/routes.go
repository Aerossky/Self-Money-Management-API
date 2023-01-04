package routes

import (
	"net/http"
	"self_money_management_api_golang/controllers"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+"!")
}

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is user!")
	})

	// e.GET("/user/:name", getUser)

	// login
	e.POST("/login", controllers.CheckLogin)

	// user
	// e.GET("/user", controllers.FetchAllUser,middleware.IsAuthenticated)
	e.GET("/user", controllers.FetchAllUser)
	e.POST("/user", controllers.StoreUser)
	e.PATCH("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)
	e.GET("/user/:id", controllers.FetchUserById)

	// plan
	e.GET("/plan", controllers.FetchAllPlan)
	e.GET("/plan/:id", controllers.FetchPlanById)
	e.POST("/plan", controllers.StorePlan)
	e.PATCH("/plan", controllers.UpdatePlan)
	e.DELETE("/plan", controllers.DeletePlan)

	//money
	e.GET("/money/:id", controllers.FetchMoneyById)
	e.GET("/moneyTotalPemasukan/:id", controllers.FetchTotalPemasukanById)
	e.GET("/moneyTotalPengeluaran/:id", controllers.FetchTotalPengeluaranById)
	e.PATCH("/money", controllers.UpdateMoney)
	e.DELETE("/money", controllers.DeleteMoney)
	e.GET("/currency", controllers.FetchMoneyAPI)
	//1pemasukan
	e.GET("/moneyPemasukan/:id", controllers.FetchDataPemasukanByUserId)
	e.GET("/moneyPengeluaran/:id", controllers.FetchDataPengeluaranByUserId)
	e.POST("/pemasukan", controllers.StorePepe)

	// validation
	e.POST("/test-validation", controllers.TestStructValidation)
	e.POST("/test-validation-var", controllers.TestVarValidation)
	return e
}
