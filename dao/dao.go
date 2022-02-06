package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "winter:winter@tcp(localhost:3306)/winter_work")
	if err != nil {
		panic(err)
	}

	dB = db
}
