package models

import (
	// "database/sql"
	// "fmt"
	"net/http"
	"self_money_management_api_golang/db"

	// "self_money_management_api_golang/helpers"
	"github.com/go-playground/validator/v10"
)

type Plan struct {
	Id     int    `json:"id"`
	IdUser int    `json:"id_user"`
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Time   int    `json:"time" validate:"required"`
}

//! CRUD START

func FetchAllPlan() (Response, error) {
	var obj Plan
	// digunakan untuk menampung data Plan
	var arrObj []Plan
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM plans"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdUser, &obj.Name, &obj.Price, &obj.Time)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

// insert data plan
func StorePlan(id_user int, name string, price int, time int) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	pln := Plan{
		IdUser: id_user,
		Name:   name,
		Price:  price,
		Time:   time,
	}

	err := v.Struct(pln)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `plans`(`id_user`, `name`, `price`, `time`) VALUES (?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(id_user, name, price, time)

	if err != nil {
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}

	return res, nil

}

// func update user

func UpdatePlan(id int, name string, price int, time int) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE plans SET name=?,price=?,time=? WHERE id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price, time, id)

	if err != nil {
		return res, err
	}

	rowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": rowAffectedID,
	}

	return res, nil

}

// func delete user
func DeletePlan(id int) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM plans WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": rowAffectedID,
	}

	return res, nil

}

func FetchPlanById(id string) (Response, error) {
	var obj Plan
	var arrObj []Plan
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM plans WHERE id_user = ?"

	rows, err := con.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdUser, &obj.Name, &obj.Price, &obj.Time)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}
