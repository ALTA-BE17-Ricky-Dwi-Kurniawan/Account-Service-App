package user_account

import "database/sql"

type UserModel struct {
	conn *sql.DB
}

func (R *UserModel) SetSQLConnection(db *sql.DB) {
	R.conn = db
}


