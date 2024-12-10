package response

import "time"

type SaldoResponse struct {
	ID             int       `json:"id"`
	CardNumber     string    `json:"card_number"`
	TotalBalance   int       `json:"total_balance"`
	WithdrawAmount int       `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}
