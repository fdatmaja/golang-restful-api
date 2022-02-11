package app

import (
	"belajar_golang_restful_api/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
