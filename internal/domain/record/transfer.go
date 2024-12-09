package record

import "time"

type TransferRecord struct {
	TransferID     int       `json:"transfer_id"`
	TransferFrom   string    `json:"transfer_from"`
	TransferTo     string    `json:"transfer_to"`
	TransferAmount int       `json:"transfer_amount"`
	TransferTime   time.Time `json:"transfer_time"`
}
