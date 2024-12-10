package response

import "time"

type CardResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	CardNumber   string    `json:"card_number"`
	CardType     string    `json:"card_type"`
	ExpireDate   time.Time `json:"expire_date"`
	CVV          string    `json:"cvv"`
	CardProvider string    `json:"card_provider"`
}
