package main

import (
	"self_money_management_api_golang/db"
	"self_money_management_api_golang/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":7070"))
}
