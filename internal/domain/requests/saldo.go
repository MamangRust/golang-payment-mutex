package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateSaldoRequest struct {
	UserID       int `json:"user_id" validate:"required"`
	TotalBalance int `json:"total_balance" validate:"required"`
}

type UpdateSaldoRequest struct {
	SaldoID        int       `json:"saldo_id" validate:"required"`
	UserID         int       `json:"user_id" validate:"required"`
	TotalBalance   int       `json:"total_balance" validate:"required"`
	WithdrawAmount int       `json:"withdraw_amount" validate:"required_without=WithdrawTime"`
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required_without=WithdrawAmount"`
}

func (r *CreateSaldoRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateSaldoRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil

}
