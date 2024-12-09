package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateCardRequest struct {
	UserID       int       `json:"user_id"`
	CardNumber   string    `json:"card_number"`
	CardType     string    `json:"card_type"`
	ExpireDate   time.Time `json:"expire_date"`
	CVV          string    `json:"cvv"`
	CardProvider string    `json:"card_provider"`
}

func (r *CreateCardRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateCardRequest struct {
	CardID       int       `json:"card_id"`
	UserID       int       `json:"user_id"`
	CardNumber   string    `json:"card_number"`
	CardType     string    `json:"card_type"`
	ExpireDate   time.Time `json:"expire_date"`
	CVV          string    `json:"cvv"`
	CardProvider string    `json:"card_provider"`
}

func (r *UpdateCardRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
