package models

import "time"

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	CardNumber      string    `json:"card_number"`
	Amount          int       `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	MerchantID      int       `json:"merchant_id"`
	TransactionTime time.Time `json:"transaction_time"`
}
