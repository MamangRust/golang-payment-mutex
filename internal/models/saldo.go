package models

import "time"

type Saldo struct {
	SaldoID        int       `json:"saldo_id"`
	CardNumber     string    `json:"card_number"`
	TotalBalance   int       `json:"total_balance"`
	WithdrawAmount int       `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}
