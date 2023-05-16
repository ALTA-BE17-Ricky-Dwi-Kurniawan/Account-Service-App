package main

import (

	// "fakhry/rawsql/controllers"
	// "fakhry/rawsql/entities"

	// "account_service_app_project/database"
	// "account_service_app_project/user_account"
	"account_service_app_project/user_account"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"os"
)



func main() {
	fmt.Println("test")
	var db_connection = os.Getenv("dbconnection")
	db, err := sql.Open("mysql", db_connection)
	// fmt.Println("tidak berhasil", err.Error())
	if err != nil {
		fmt.Println(err)
		
	}

	if db.Ping() != nil {
		fmt.Println("tidak berhasil")
		fmt.Println(db.Ping().Error())
	}

// 	conn := database.InitSQL()
	
// 	mdl := user_account.UserModel{}
	
// 	mdl.SetSQLConnection(conn)
	fmt.Println("berhasil")
	var name user_account.User_account

	
	errorRegisterUser:= user_account.RegisterUser(db, name)
	fmt.Scanln("Nama Anda:", name)
	if errorRegisterUser != nil {
		fmt.Println("Nama anda tidak valid")
		fmt.Println(errorRegisterUser.Error())
}
}




