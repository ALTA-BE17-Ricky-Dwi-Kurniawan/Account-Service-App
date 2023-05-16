package main

import (

	// "fakhry/rawsql/controllers"
	// "fakhry/rawsql/entities"

	//"account_service_app_project/database"
	// "account_service_app_project/user_account"
	"account_service_app_project/user_account"
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"os"
)

func main() {
	// fmt.Println("test")
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
	// fmt.Println("berhasil")
	var User_account user_account.User_account
	var menu int

	for menu != 99 {
		fmt.Println("")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			fmt.Print("Enter name\t: ")
			fmt.Scanln(&User_account.Name)
			fmt.Print("Enter phone\t: ")
			fmt.Scanln(&User_account.Phone_number)
			fmt.Print("Enter password\t: ")
			fmt.Scanln(&User_account.Password)

			if User_account.Name == "" || User_account.Password == "" {
				log.Fatal("Gagal mendaftar. Mohon isi semua data yang dibutuhkan.")
			}

			query := "INSERT INTO user (nama, no_hp, password) VALUES (?, ?, ?)"
			_, err := db.Exec(query, User_account.Name, User_account.Phone_number, User_account.Password)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Registrasi berhasil!")
			// case 2:
			// 	fmt.Println("Login")
			// Implement your login logic here
		}
	}
}
