package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gularis")

	if err != nil {
		panic(err.Error())
	}

	return db
}
