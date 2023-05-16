package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() *sql.DB {
	var db_connection = os.Getenv("dbconnection")
	db, err := sql.Open("mysql", db_connection)
	fmt.Println("tidak berhasil", err.Error())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if db.Ping() != nil {
		fmt.Println("tidak berhasil")
		fmt.Println(db.Ping().Error())
		return nil
	}

	return db
}

