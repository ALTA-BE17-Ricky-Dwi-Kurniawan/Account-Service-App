package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/account_service_app")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if db.Ping() != nil {
		fmt.Println(db.Ping().Error())
		return nil
	}

	return db
}
