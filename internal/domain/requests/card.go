package requests

import (
	"fmt"
	methodtopup "payment-mutex/pkg/method_topup"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateCardRequest struct {
	UserID       int       `json:"user_id"`
	CardType     string    `json:"card_type"`
	ExpireDate   time.Time `json:"expire_date"`
	CVV          string    `json:"cvv"`
	CardProvider string    `json:"card_provider"`
}

func (r *CreateCardRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if r.CardType != "credit" && r.CardType != "debit" {
		return fmt.Errorf("card type must be credit or debit")
	}

	if !methodtopup.PaymentMethodValidator(r.CardProvider) {
		return fmt.Errorf("card provider not found")
	}

	if err != nil {
		return err
	}

	return nil
}

type UpdateCardRequest struct {
	CardID       int       `json:"card_id"`
	UserID       int       `json:"user_id"`
	CardType     string    `json:"card_type"`
	ExpireDate   time.Time `json:"expire_date"`
	CVV          string    `json:"cvv"`
	CardProvider string    `json:"card_provider"`
}

func (r *UpdateCardRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if r.CardType != "credit" && r.CardType != "debit" {
		return fmt.Errorf("card type must be credit or debit")
	}

	if !methodtopup.PaymentMethodValidator(r.CardProvider) {
		return fmt.Errorf("card provider not found")
	}

	if err != nil {
		return err
	}

	return nil
}
