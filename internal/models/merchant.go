package models

type Merchant struct {
	MerchantID int    `json:"merchant_id"`
	Name       string `json:"name"`
	ApiKey     string `json:"api_key"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`
}
