package main

import (
	"account_service_app_project/database"
	"account_service_app_project/transaction"
	"account_service_app_project/user_account"
	"fmt"
	"log"
	"time"

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
		fmt.Println("10. Melihat Profile User Lain")
		fmt.Println("0. ExitSystem")
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
			fmt.Print("Masukkan Nomer Hanpdone\t: ")
			var phoneNumber string
			fmt.Scanln(&phoneNumber)

			user, err := user_account.ReadAccount(db, phoneNumber)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Nama:", user.Name)
			fmt.Println("Nomer Handphone:", user.Phone_number)
			fmt.Println("Kata sandi:", user.Password)

		case 4:
			fmt.Println("Update Account")
			fmt.Print("Masukkan nomer handphone\t: ")
			fmt.Scanln(&User_account.Phone_number)
			fmt.Print("Masukkan nama baru\t: ")
			fmt.Scanln(&User_account.Name)
			fmt.Print("Masukkan kata sandi baru\t: ")
			fmt.Scanln(&User_account.Password)

			err := user_account.UpdateAccount(db, User_account)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Update akun berhasil!")
		case 5:
			fmt.Println("Menghapus akun")
			fmt.Print("Masukkan nomor hp\t: ")
			fmt.Scanln(&User_account.Phone_number)

			err := user_account.DeleteAccount(db, User_account.Phone_number)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Akun berhasil dihapus!")
		case 6:
			// Cek User sudah login atau belum
			if User_account.Phone_number == "" {
				log.Fatal("Tidak login akun tidak ditemukan")
			}

			fmt.Printf("Akun yang sedang login: %s\n", User_account.Name)

			// Top-Up
			fmt.Println("Top-Up")
			fmt.Print("Masukkan Jumlah: ")
			var amount int
			fmt.Scanln(&amount)

			err := transaction.TopUp(db, User_account.Id, amount)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Top-up sukses!")

		case 7:
			if User_account.Phone_number == "" {
				log.Fatal("Anda harus login terlebih dahulu.")
			}
			fmt.Printf("Akun yang sedang login: %s\n", User_account.Name)

			// transfer
			fmt.Println("Transfer")
			fmt.Print("Masukkan nomor handphone\t: ")
			fmt.Scanln(&User_account.Phone_number)
			fmt.Print("Masukkan nomor akun tujuan\t: ")
			var receiverPhone string
			fmt.Scanln(&receiverPhone)
			fmt.Print("Masukkan jumlah transfer\t: ")
			var amount int
			fmt.Scanln(&amount)

			err := transaction.Transfer(db, User_account.Id, receiverPhone, amount)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Transfer berhasil!")

		case 8:
			if User_account.Phone_number == "" {
				log.Fatal("Anda harus login terlebih dahulu.")
			}
			fmt.Printf("Akun yang sedang login: %s\n", User_account.Name)

			// Melihat history top-up
			fmt.Println("History Top-up")

			// Mengambil data history top-up dari database berdasarkan akun pengguna
			query := "SELECT transaction_type, transaction_date, amount FROM transaction WHERE sender_account_id = ? AND transaction_type = 'Top-up'"
			rows, err := db.Query(query, User_account.Id)
			if err != nil {
				log.Fatal("Failed to retrieve history Top-up: ", err)
			}
			defer rows.Close()

			// Menampilkan data history top-up
			for rows.Next() {
				var transactionType string
				var transactionDateStr string
				var amount int
				err := rows.Scan(&transactionType, &transactionDateStr, &amount)
				if err != nil {
					log.Fatal("Failed to scan history Top-up rows: ", err)
				}

				// Konversi string tanggal menjadi time.Time
				transactionDate, err := time.Parse("2006-01-02 15:04:05", transactionDateStr)
				if err != nil {
					log.Fatal("Failed to parse transaction date: ", err)
				}

				fmt.Printf("Transaction Type: %s\n", transactionType)
				fmt.Printf("Transaction Date: %s\n", transactionDate.Format("2023-05-18 15:04:05"))
				fmt.Printf("Amount: %d\n", amount)
				fmt.Println("------------------------")
			}

			if err := rows.Err(); err != nil {
				log.Fatal("Failed to iterate over history Top-up rows: ", err)
			}

		case 9:
			if User_account.Phone_number == "" {
				log.Fatal("Anda harus login terlebih dahulu.")
			}
			fmt.Printf("Akun yang sedang login: %s\n", User_account.Name)

			// Melihat history transfer
			fmt.Println("History Transfer")

			// Mengambil data history Transfer dari database berdasarkan akun pengguna
			query := "SELECT transaction_type, transaction_date, amount FROM transaction WHERE sender_account_id = ? AND transaction_type = 'Transfer'"
			rows, err := db.Query(query, User_account.Id)
			if err != nil {
				log.Fatal("Failed to retrieve history Transfer: ", err)
			}
			defer rows.Close()

			// Menampilkan data history transfer
			for rows.Next() {
				var transactionType string
				var transactionDateStr string
				var amount int
				err := rows.Scan(&transactionType, &transactionDateStr, &amount)
				if err != nil {
					log.Fatal("Failed to scan history Transfer rows: ", err)
				}

				// Konversi string tanggal menjadi time.Time
				transactionDate, err := time.Parse("2006-01-02 15:04:05", transactionDateStr)
				if err != nil {
					log.Fatal("Failed to parse transaction date: ", err)
				}

				fmt.Printf("Transaction Type: %s\n", transactionType)
				fmt.Printf("Transaction Date: %s\n", transactionDate.Format("2023-05-18 15:04:05"))
				fmt.Printf("Amount: %d\n", amount)
				fmt.Println("------------------------")
			}

			if err := rows.Err(); err != nil {
				log.Fatal("Failed to iterate over history Top-up rows: ", err)
			}

		case 10:
			if User_account.Phone_number == "" {
				log.Fatal("Anda harus login terlebih dahulu.")
			}
			fmt.Println("Masukkan Nomor Telepon Pencarian")
			fmt.Print("Masukkan Nomer Handphone\t: ")
			var phoneNumber string
			fmt.Scanln(&phoneNumber)

			user, err := user_account.ViewUserProfile(db, phoneNumber)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Profile user: ")
			fmt.Println("Nama:", user.Name)
			fmt.Println("Nomer Handphone:", user.Phone_number)

		case 0:
			// Cek User sudah login atau belum
			if User_account.Phone_number == "" {
				log.Fatal("Tidak login akun tidak ditemukan")
			}

			user_account.ExitSystem()

		}
	}
}
