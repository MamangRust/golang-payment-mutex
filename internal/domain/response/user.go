package response

type UserResponse struct {
	ID          int    `json:"user_id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	NocTransfer int    `json:"noc_transfer"`
}
