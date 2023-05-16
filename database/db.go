package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() (*sql.DB, error) {
	var db_connection = os.Getenv("dbconnection")
	db, err := sql.Open("mysql", db_connection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}