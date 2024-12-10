package response

import "time"

type TransactionResponse struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CardNumber      string    `json:"card_number"`
	Amount          int       `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	TransactionTime time.Time `json:"transaction_time"`
}
