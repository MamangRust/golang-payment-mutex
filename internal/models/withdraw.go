package models

import "time"

type Withdraw struct {
	WithdrawID     int       `json:"withdraw_id"`
	UserID         int       `json:"user_id"`
	WithdrawAmount int       `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}
