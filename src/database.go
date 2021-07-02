package main

import (
	"database/sql"
	"time"
)


func GetConnection() *sql.DB {
	db,err := sql.Open("mysql","golang:golang@(database:3306)/golang")
	if err != nil {
		panic(err)
	}


	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifeTime(60 * time.Minute)

	return db
}