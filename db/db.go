package db

import (
	"database/sql"
	"self_money_management_api_golang/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	conn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", conn)

	if err != nil {
		panic("Connection failed/error!")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid!")
		// panic(err.Error())
	}

	// panic("Connected to database!")
}

func Createcon() *sql.DB {
	return db
}
