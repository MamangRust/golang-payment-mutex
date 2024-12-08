package record

import "time"

type TopupRecord struct {
	TopupID     int       `json:"topup_id"`
	UserID      int       `json:"user_id"`
	TopupNo     string    `json:"topup_no"`
	TopupAmount int       `json:"topup_amount"`
	TopupMethod string    `json:"topup_method"`
	TopupTime   time.Time `json:"topup_time"`
}
