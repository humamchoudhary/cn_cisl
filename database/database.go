package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./cisl-database.db")
	if err != nil {
		fmt.Println(err)
		return err
	}
	DB = db
	return nil
}
