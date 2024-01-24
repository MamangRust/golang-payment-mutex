package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateWithdrawRequest struct {
	UserID         int       `json:"user_id" validate:"required"`
	WithdrawAmount int       `json:"withdraw_amount" validate:"required"`
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`
}

func (r *CreateWithdrawRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateWithdrawRequest struct {
	UserID         int       `json:"user_id" validate:"required"`
	WithdrawID     int       `json:"withdraw_id" validate:"required"`
	WithdrawAmount int       `json:"withdraw_amount" validate:"required"`
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`
}

func (r *UpdateWithdrawRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
