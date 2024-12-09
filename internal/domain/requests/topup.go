package requests

import (
	"errors"
	methodtopup "payment-mutex/pkg/method_topup"

	"github.com/go-playground/validator/v10"
)

type CreateTopupRequest struct {
	CardNumber  string `json:"card_number" validate:"required,min=1"`
	TopupNo     string `json:"topup_no" validate:"required"`
	TopupAmount int    `json:"topup_amount" validate:"required,min=50000"`
	TopupMethod string `json:"topup_method" validate:"required"`
}

type UpdateTopupRequest struct {
	CardNumber  string `json:"card_number" validate:"required,min=1"`
	TopupID     int    `json:"topup_id" validate:"required,min=1"`
	TopupAmount int    `json:"topup_amount" validate:"required,min=50000"`
	TopupMethod string `json:"topup_method" validate:"required"`
}

type UpdateTopupAmount struct {
	TopupID     int `json:"topup_id" validate:"required,min=1"`
	TopupAmount int `json:"topup_amount" validate:"required,min=50000"`
}

func (r *CreateTopupRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupNo == "" {
		return errors.New("top-up number is required")
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	if r.TopupMethod == "" {
		return errors.New("top-up method is required")
	}

	if !methodtopup.PaymentMethodValidator(r.TopupMethod) {
		return errors.New("topup method not found")
	}

	return nil
}

func (r *UpdateTopupRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupID <= 0 {
		return errors.New("top-up ID must be a positive integer")
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	if r.TopupMethod == "" {
		return errors.New("top-up method is required")
	}

	if !methodtopup.PaymentMethodValidator(r.TopupMethod) {
		return errors.New("topup method not found")
	}

	return nil
}

func (r *UpdateTopupAmount) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupID <= 0 {
		return errors.New("top-up ID must be a positive integer")
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	return nil
}
