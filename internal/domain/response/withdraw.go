package response

import "time"

type WithdrawResponse struct {
	ID             int       `json:"id"`
	CardNumber     string    `json:"card_number"`
	WithdrawAmount int       `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}
