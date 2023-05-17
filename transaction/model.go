package transaction

import (
	// "account_service_app_project/user_account"
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
		Transaction_type:    "Top-up",
		Transaction_date:    time.Now(),
		Amount:              amount,
		Sender_account_id:   Sender_account_id,
		Receiver_account_id: Sender_account_id, // Top-up merupakan transfer dari akun sendiri ke akun sendiri
	}

	// Menyimpan transaksi ke dalam database
	insertQuery := "INSERT INTO transaction (transaction_type, transaction_date, amount, sender_account_id, receiver_account_id) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, transaction.Transaction_type, transaction.Transaction_date, transaction.Amount, transaction.Sender_account_id, transaction.Receiver_account_id)
	if err != nil {
		return fmt.Errorf("gagal top up transaction: %v", err)
	}

	// Memperbarui balance pada user_account
	updateQuery := "UPDATE user_account SET balance = balance + ? WHERE Id = ?"
	_, err = db.Exec(updateQuery, amount, Sender_account_id)
	if err != nil {
		return fmt.Errorf("gagal update balance: %v", err)
	}

	return nil
}

func Transfer(db *sql.DB, Sender_account_id int, Receiver_phone_number string, amount int) error {
	// Mengambil ID akun pengirim berdasarkan Sender_account_id
	query := "SELECT Id FROM user_account WHERE Id = ?"
	row := db.QueryRow(query, Sender_account_id)
	var SenderID int
	err := row.Scan(&SenderID)
	if err != nil {
		return fmt.Errorf("failed to retrieve sender's account: %v", err)
	}

	// Mengambil ID akun penerima berdasarkan phoneNumber penerima
	query = "SELECT Id FROM user_account WHERE phone_number = ?"
	row = db.QueryRow(query, Receiver_phone_number)
	var ReceiverID int
	err = row.Scan(&ReceiverID)
	if err != nil {
		return fmt.Errorf("failed to retrieve receiver's account: %v", err)
	}

	// Membuat objek Transaction untuk transfer
	transaction := Transaction{
		Transaction_type:    "Transfer",
		Transaction_date:    time.Now(),
		Amount:              amount,
		Sender_account_id:   Sender_account_id,
		Receiver_account_id: ReceiverID,
	}
	// Menyimpan transaksi ke dalam database
	insertQuery := "INSERT INTO transaction (transaction_type, transaction_date, amount, sender_account_id, receiver_account_id) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(insertQuery, transaction.Transaction_type, transaction.Transaction_date, transaction.Amount, transaction.Sender_account_id, transaction.Receiver_account_id)
	if err != nil {
		return fmt.Errorf("failed to store transfer transaction: %v", err)
	}

	// Memperbarui balance pada akun pengirim
	updateSenderQuery := "UPDATE user_account SET balance = balance - ? WHERE Id = ?"
	_, err = db.Exec(updateSenderQuery, amount, Sender_account_id)
	if err != nil {
		return fmt.Errorf("failed to update sender's balance: %v", err)
	}

	// Memperbarui balance pada akun penerima
	updateReceiverQuery := "UPDATE user_account SET balance = balance + ? WHERE Id = ?"
	_, err = db.Exec(updateReceiverQuery, amount, ReceiverID)
	if err != nil {
		return fmt.Errorf("failed to update receiver's balance: %v", err)
	}

	return nil

}

func HistoryTopUp(db *sql.DB, accountID int) error {
	// Mengambil riwayat top-up dari database berdasarkan ID akun
	query := "SELECT transaction_type, transaction_date, amount FROM transaction WHERE sender_account_id = ? AND transaction_type = 'Top-up'"
	rows, err := db.Query(query, accountID)
	if err != nil {
		return fmt.Errorf("failed to retrieve top-up history: %v", err)
	}
	defer rows.Close()

	// Menampilkan riwayat top-up
	fmt.Println("Top-up History:")
	for rows.Next() {
		var transactionType string
		var transactionDate time.Time
		var amount int
		err := rows.Scan(&transactionType, &transactionDate, &amount)
		if err != nil {
			return fmt.Errorf("failed to scan top-up history rows: %v", err)
		}

		fmt.Printf("Transaction Type: %s\n", transactionType)
		fmt.Printf("Transaction Date: %s\n", transactionDate.Format("2006-01-02 15:04:05"))
		fmt.Printf("Amount: %d\n", amount)
		fmt.Println("------------------------")
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("failed to iterate over top-up history rows: %v", err)
	}

	return nil
}

// func HistoryTransfer
