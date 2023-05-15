package main

import (

	// "fakhry/rawsql/controllers"
	// "fakhry/rawsql/entities"

	"account_service_app_project/database"
	"account_service_app_project/user_account"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	conn := database.InitSQL()
	mdl := user_account.UserModel{}
	mdl.SetSQLConnection(conn)
}

// func registerUser(db *sql.DB, name string, password string) {

// 	query := "select username from user where username = ?"
// 	row := db.QueryRow(query, name)

// 	var existingUsername string

// 	err := row.Scan(&existingUsername)
// 	if err == nil {
// 		fmt.Println("name", existingUsername, "sudah terdaftar.")
// 		return
// 	}

// 	insertQuery := "insert into user_account (name, password), values (?, ?)"
// 	_, err = db.Exec(insertQuery, name, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Registerasi berhasil. Username:", name)

// }
