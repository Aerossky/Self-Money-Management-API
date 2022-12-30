package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"self_money_management_api_golang/db"
	"self_money_management_api_golang/helpers"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       int    `json:"id"` // tag json digunakan untuk menentukan nama field yang akan di tampilkan di response
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Image    string `json:"image" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//! CRUD START

func FetchAllUser() (Response, error) {
	var obj User
	// digunakan untuk menampung data user
	var arrObj []User
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Email, &obj.Username,&obj.Image, &obj.Password)

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

// insert data user
func StoreUser(id int, email string, username string, image string, password string) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	usr := User{
		Id:       id,
		Email:    email,
		Username: username,
		Image:    image,
		Password: password,
	}

	err := v.Struct(usr)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `users`(`id`, `email`, `username`,`image`, `password`) VALUES (?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(id, email, username, image, password)

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

func UpdateUser(id int, email string, username string, image string, password string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE users SET email=?,username=?,image=?,password=? WHERE id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, username, image, password, id)

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
func DeleteUser(id string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM users WHERE id=?"
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

// !CRUD END

// !VALIDATION START
// check login
func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string
	con := db.Createcon()

	sqlStatement := "SELECT * FROM users WHERE email = ?"
	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id, &obj.Email, &obj.Username, &obj.Image, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Email not found!")
		return false, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}

// !VALIDATION END
