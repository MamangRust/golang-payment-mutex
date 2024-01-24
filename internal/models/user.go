package models

type User struct {
	UserID      int    `json:"user_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NocTransfer int    `json:"noc_transfer"`
}
