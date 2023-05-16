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

// func RegisterUser(db *sql.DB, user User_account) error {
// 	result, err := db.Exec("INSERT INTO User_account (Name, Phone_number, Password) VALUES (?,?,?)", user.Name, user.Phone_number, user.Password)
// 	if err != nil {
// 		return fmt.Errorf("failed to create user account: %v", err)
// 	}

// 	// Get the new album's generated ID for the client.
// 	_, errId := result.LastInsertId()
// 	if errId != nil {
// 		return fmt.Errorf("AddAlbum: %v", errId)
// 	}
// 	// Return the new album's ID.
// 	return nil
// }
