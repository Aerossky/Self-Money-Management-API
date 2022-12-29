package models

import (
	// "database/sql"
	// "fmt"
	"net/http"
	"self_money_management_api_golang/db"
	// "self_money_management_api_golang/helpers"
	// "github.com/go-playground/validator/v10"
)

type Money struct {
	Id         int    `json:"id"`
	IdUser     int    `json:"id_user"`
	TotalMoney int    `json:"total_money" validate:"required"`
	Note       string `json:"note" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

// fetch money by user id
func FetchMoneyById(id string) (Response, error) {
	var obj Money
	var arrObj []Money
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM moneys WHERE id_user = ?"

	rows, err := con.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdUser, &obj.TotalMoney, &obj.Note, &obj.Status)

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

// !total
func FetchTotalPemasukanById(id string) (Response, error) {
	var obj Money
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT SUM(total_money) FROM moneys WHERE  id_user = ? AND status = 'pemasukan' "

	err := con.QueryRow(sqlStatement, id).Scan(&obj.TotalMoney)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

// fetch money by id AND status = pemasukan
func FetchTotalPengeluaranById(id string) (Response, error) {
	var obj Money
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT SUM(total_money) FROM moneys WHERE  id_user = ? AND status = 'pengeluaran' "

	err := con.QueryRow(sqlStatement, id).Scan(&obj.TotalMoney)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

// func delete money
func DeleteMoney(id int) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM moneys WHERE id=?"
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
// func update user

func UpdateMoney(id int, total_money int, note string, status string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE moneys SET total_money=?,note=?,status=? WHERE id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(total_money, note, status, id)

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