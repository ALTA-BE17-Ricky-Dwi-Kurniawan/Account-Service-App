package user_account

import (
	"database/sql"
	"fmt"
)

type User_account struct {
	Id           int
	Name         string
	Phone_number int
	Gender       bool
	Password     string
}
func RegisterUser(db *sql.DB, user User_account) error {
	insertquery := "insert into user_account (Name, Phone_number, Gender, Password) values (?, ?, ?, ?)"
	_, err := db.Exec(insertquery, user.Name, user.Phone_number, user.Gender, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user account: %v", err)
	}
	return nil
}


