package response

type MerchantResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
	Status string `json:"status"`
}
