package transaction

import "time"

type Transaction struct {
	Transaction_id      int
	Transaction_type    string
	Transaction_date    time.Time
	Amount              int
	Sender_account_id   int
	Receiver_account_id int
}