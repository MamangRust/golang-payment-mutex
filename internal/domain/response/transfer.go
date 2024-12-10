package response

import "time"

type TransferResponse struct {
	ID             int       `json:"id"`
	TransferFrom   string    `json:"transfer_from"`
	TransferTo     string    `json:"transfer_to"`
	TransferAmount int       `json:"transfer_amount"`
	TransferTime   time.Time `json:"transfer_time"`
}
