package user_account

import (
	"database/sql"
	"fmt"
)

type user_account struct {
	id           int
	name         string
	phone_number int
	gender       string
	password     string
}

func RegisterUser(db *sql.DB, user user_account) error {
	insertQuery := "insert into users (name, phone_number, gender, password) values (?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, user.name, user.phone_number, user.gender, user.password)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}
