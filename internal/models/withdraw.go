package models

import "time"

type Withdraw struct {
	WithdrawID     int       `json:"withdraw_id"`
	CardNumber     string    `json:"card_number"`
	WithdrawAmount int       `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}
