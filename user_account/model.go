package user_account

import (
	"database/sql"
	"fmt"
)

type UserModel struct {
	conn *sql.DB
}

func (R *UserModel) SetSQLConnection(db *sql.DB) {
	R.conn = db
}

func RegisterUser(db *sql.DB, user_account User_account) error {
	insertquery := "insert into user_account (name, phone_number, password) values (?, ?, ?)"
	_, err := db.Exec(insertquery, user_account.Name, user_account.Phone_number, user_account.Password)
	if err != nil {
		return fmt.Errorf("failed to create user account: %v", err)
	}
	return nil
}

func LoginUser(db *sql.DB, user_account User_account) error {
	selectquery := "SELECT * FROM user_account WHERE No_HP = ?"
	row := db.QueryRow(selectquery, user_account.Phone_number)
	err := row.Scan(&user_account.Name, &user_account.Phone_number, &user_account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to login: %v", err)
	}

	// Memeriksa kesesuaian kata sandi
	if user_account.Password != user_account.Password {
		return fmt.Errorf("incorrect password")
	}

	return nil
}

func ReadAccount(db *sql.DB, phoneNumber string) (User_account, error) {
	selectQuery := "SELECT * FROM user_account WHERE phone_number = ?"
	row := db.QueryRow(selectQuery, phoneNumber)

	var user User_account
	err := row.Scan(&user.Name, &user.Phone_number, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User_account{}, fmt.Errorf("user not found")
		}
		return User_account{}, fmt.Errorf("failed to read account: %v", err)
	}

	return user, nil
}

func UpdateAccount(db *sql.DB, user User_account) error {
	updateQuery := "UPDATE user_account SET name = ?, phone_number = ?, password = ? WHERE phone_number = ?"
	_, err := db.Exec(updateQuery, user.Name, user.Phone_number, user.Password, user.Phone_number)
	if err != nil {
		return fmt.Errorf("failed to update account: %v", err)
	}
	return nil
}

func DeleteAccount(db *sql.DB, phoneNumber string) error {
	deleteQuery := "DELETE FROM user_account WHERE phone_number = ?"
	_, err := db.Exec(deleteQuery, phoneNumber)
	if err != nil {
		return fmt.Errorf("failed to delete account: %v", err)
	}
	return nil
}

