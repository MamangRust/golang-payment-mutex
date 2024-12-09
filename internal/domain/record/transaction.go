package record

import "time"

type TransactionRecord struct {
	TransactionID   int       `json:"transaction_id"`
	UserID          int       `json:"user_id"`
	CardNumber      string    `json:"card_number"`
	Amount          int       `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	TransactionTime time.Time `json:"transaction_time"`
}
