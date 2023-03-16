package controllers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conn adalah variabel untuk menyimpan koneksi ke database

func Connect() *sql.DB {
	// Membuka koneksi ke database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/eksplorasi?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		panic(err.Error())
	}

	return db
}
