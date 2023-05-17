package transaction

import (
	"database/sql"
	"fmt"
	"time"

)

type UserModel struct {
	conn *sql.DB
}

func TopUp(db *sql.DB, Sender_account_id int, amount int) error {
	// Membuat objek Transaction untuk top-up
	transaction := Transaction{
	Transaction_type: "Top-up",
	Transaction_date: time.Now(),
	Amount: amount,
	Sender_account_id: Sender_account_id,
	Receiver_account_id: Sender_account_id, // Top-up merupakan transfer dari akun sendiri ke akun sendiri
	}
	// Menyimpan transaksi ke dalam database
	insertQuery := "INSERT INTO transaction (transaction_type, transaction_date, amount, sender_account_id, receiver_account_id) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, transaction.Transaction_type, transaction.Transaction_date, transaction.Amount, transaction.Sender_account_id, transaction.Receiver_account_id)
	if err != nil {
		return fmt.Errorf("failed to store top-up transaction: %v", err)
	}
	return nil
}

// func Transfer(db *sql.DB, Sender_account_id int, amount int) error {
// 	transaction := Transaction{
// 		Transaction_type: "Top-up",
// 		Transaction_date: time.Now(),
// 		Amount: amount,
// 		Sender_account_id: Sender_account_id,
// 		Receiver_account_id: Sender_account_id, // Top-up merupakan transfer dari akun sendiri ke akun sendiri
// 		}
// 		// Menyimpan transaksi ke dalam database
// 		insertQuery := "INSERT INTO transaction (transaction_type, transaction_date, amount, sender_account_id, receiver_account_id) VALUES (?, ?, ?, ?, ?)"
// 		_, err := db.Exec(insertQuery, transaction.Transaction_type, transaction.Transaction_date, transaction.Amount, transaction.Sender_account_id, transaction.Receiver_account_id)
// 		if err != nil {
// 			return fmt.Errorf("failed to store top-up transaction: %v", err)
// 		}
	
// 		return nil
// 	}
