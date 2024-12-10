package response

import "time"

type TopupResponse struct {
	ID          int       `json:"id"`
	CardNumber  string    `json:"card_number"`
	TopupNo     string    `json:"topup_no"`
	TopupAmount int       `json:"topup_amount"`
	TopupMethod string    `json:"topup_method"`
	TopupTime   time.Time `json:"topup_time"`
}
