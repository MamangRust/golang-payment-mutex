package requests

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateWithdrawRequest struct {
	CardNumber     string    `json:"card_number" validate:"required,min=1"`
	WithdrawAmount int       `json:"withdraw_amount" validate:"required,min=50000"`
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`
}

type UpdateWithdrawRequest struct {
	CardNumber     string    `json:"card_number" validate:"required,min=1"`
	WithdrawID     int       `json:"withdraw_id" validate:"required,min=1"`
	WithdrawAmount int       `json:"withdraw_amount" validate:"required,min=50000"`
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`
}

func (r *CreateWithdrawRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.WithdrawAmount < 50000 {
		return errors.New("withdraw amount must be at least 50,000")
	}

	if r.WithdrawTime.After(time.Now()) {
		return errors.New("withdraw time cannot be in the future")
	}

	return nil
}

func (r *UpdateWithdrawRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.WithdrawAmount < 50000 {
		return errors.New("withdraw amount must be at least 50,000")
	}

	if r.WithdrawTime.After(time.Now()) {
		return errors.New("withdraw time cannot be in the future")
	}

	return nil
}
