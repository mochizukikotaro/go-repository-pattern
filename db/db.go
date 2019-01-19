package db

import (
	"database/sql"
	"fmt"
)

func Db() *sql.DB {
	db, err := sql.Open("mysql", "myuser:mypw@tcp(db:3306)/mydb?charset=utf8mb4,utf8")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
