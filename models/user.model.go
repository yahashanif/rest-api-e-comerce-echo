package models

import (
	"database/sql"
	"net/http"
	"rest-api-e-comerce/db"
	"rest-api-e-comerce/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Level    string `json:"level"`
	FullName string `json:"full_name"`
	Email    string `json:"Email"`
	UrlPhoto string `json:"url_photo"`
}

func RegisterCustomer(user *User, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO `users` ( `username`, `level`, `full_name`, `email`, `password`) VALUES ( ?, 'customer', ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	hash, _ := helpers.HashPassword(password)

	result, err := stmt.Exec(user.Username, user.FullName, user.Email, hash)
	if err != nil {
		return res, err
	}

	LastIDInsert, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "SUKSES Register Customer"
	res.Data = map[string]int{
		"LASTIDINSERT": int(LastIDInsert),
	}
	return res, nil
}

func CheckLoginCustomer(username, password string) (Response, error) {
	var res Response
	var user User

	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ? and level = 'customer'"

	err := con.QueryRow(sqlStatement, username).Scan(&user.Id, &user.Username, &user.Level, &user.FullName, &user.Email, &user.UrlPhoto, &pwd)

	if err == sql.ErrNoRows {
		res.Message = "USERNAME NOT FOUND"
		return res, err
	}

	if err != nil {
		return res, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		res.Message = "HASH and PASSWORD doesn't Match"
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES LOGIN"
	res.Data = user

	return res, nil

}

func CheckUser(id string) (Response, error) {
	var res Response
	var user User
	var pwd string

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM users WHERE id = ? and level = 'customer'"

	err := con.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Username, &user.Level, &user.FullName, &user.Email, &user.UrlPhoto, &pwd)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES CEK USER"
	res.Data = user

	return res, nil
}
