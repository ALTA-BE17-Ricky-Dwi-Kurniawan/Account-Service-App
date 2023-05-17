package main

import (
	"account_service_app_project/user_account"
	"account_service_app_project/transaction"
	"account_service_app_project/database"
	"fmt"
	"log"
	// "time"
	

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.InitSQL()
	if err != nil {
		log.Fatal("failed to initialize database:", err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to the database!")

	var User_account user_account.User_account
	var menu int



	for menu != 99 {
		fmt.Println("")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Read Account")
		fmt.Println("4. Update Account")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Top-Up")
		fmt.Println("7. Transfer")
		fmt.Println("8. Top-Up History")
		fmt.Println("9. Transfer History")
		fmt.Println("10. View Other User Profile")
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

			user_account.RegisterUser(db, User_account)

			if User_account.Name == "" || User_account.Password == "" {
				log.Fatal("Gagal mendaftar. Mohon isi semua data yang dibutuhkan.")
			}
			fmt.Println("Registrasi berhasil!")
		case 2:
			fmt.Println("Login")
			fmt.Print("Enter phone\t: ")
			fmt.Scanln(&User_account.Phone_number)
			fmt.Print("Enter password\t: ")
			fmt.Scanln(&User_account.Password)

			err := user_account.LoginUser(db, &User_account)
			if err != nil {
				log.Fatal(err)
			}

			// Verifying the result of login
			if User_account.Name == "" {
				log.Fatal("Gagal login")
			}

	fmt.Println("Login berhasil!")

		case 3:
			fmt.Println("Masukkan Nomor Telepon Pencarian")
			fmt.Print("Enter phone\t: ")
			var phoneNumber string
			fmt.Scanln(&phoneNumber)

			user, err := user_account.ReadAccount(db, phoneNumber)
			if err != nil {
			log.Fatal(err)
			}

			fmt.Println("Name:", user.Name)
			fmt.Println("Phone Number:", user.Phone_number)
			fmt.Println("Password:", user.Password)
			
		case 4:
			fmt.Println("Update Account")
			fmt.Print("Enter phone\t: ")
			fmt.Scanln(&User_account.Phone_number)
			fmt.Print("Enter new name\t: ")
			fmt.Scanln(&User_account.Name)
			fmt.Print("Enter new password\t: ")
			fmt.Scanln(&User_account.Password)
		
			err := user_account.UpdateAccount(db, User_account)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Account updated successfully!")	
		case 5:
			fmt.Println("Delete Account")
			fmt.Print("Enter phone\t: ")
			fmt.Scanln(&User_account.Phone_number)
		
			err := user_account.DeleteAccount(db, User_account.Phone_number)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Account deleted successfully!")
		case 6:
			// Cek User sudah login atau belum 
			if User_account.Phone_number == "" { 
				log.Fatal("No logged-in account found") 
			   } 
			
			fmt.Printf("Using logged-in account: %s\n", User_account.Name) 
			
			// Top-Up
			fmt.Println("Top-Up")
			fmt.Print("Enter amount: ")
			var amount int
			fmt.Scanln(&amount)
			
			
			err := transaction.TopUp(db, User_account.Id, amount)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Top-up successful!")

		// case 7:
		// 	if !loggedIn {
		// 		log.Fatal("Anda harus login terlebih dahulu.")
		// 	}

		// 	fmt.Println("Transfer")
		// 	fmt.Print("Masukkan jumlah transfer: ")
		// 	var amount int
		// 	fmt.Scanln(&amount)
		// 	fmt.Print("Masukkan nomor akun tujuan: ")
		// 	var receiverID int
		// 	fmt.Scanln(&receiverID)

		// 	err := transaction.Transfer(db, User_account.Id, receiverID, amount)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	// Update saldo pengguna
		// 	User_account.Balance -= amount

		// 	fmt.Println("Transfer berhasil!")


		}}}