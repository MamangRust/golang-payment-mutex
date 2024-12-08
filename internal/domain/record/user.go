package record

type UserRecord struct {
	UserID      int     `json:"user_id"`
	FirstName   string  `json:"firstname"`
	LastName    string  `json:"lastname"`
	Email       string  `json:"email"`
	Password    *string `json:"password"`
	NocTransfer int     `json:"noc_transfer"`
}
