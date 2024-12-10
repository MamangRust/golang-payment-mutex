package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateTransactionRequest struct {
	CardNumber    string `json:"card_number"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`

	MerchantID      *int      `json:"merchant_id"`
	TransactionTime time.Time `json:"transaction_time"`
}

func (r *CreateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateTransactionRequest struct {
	TransactionID   int       `json:"transaction_id"`
	CardNumber      string    `json:"card_number"`
	Amount          int       `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`

	MerchantID      *int      `json:"merchant_id"`
	TransactionTime time.Time `json:"transaction_time"`
}

func (r *UpdateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
