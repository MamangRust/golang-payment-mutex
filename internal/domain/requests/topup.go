package requests

import "github.com/go-playground/validator/v10"

type CreateTopupRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	TopupNo     string `json:"topup_no" validate:"required"`
	TopupAmount int    `json:"topup_amount" validate:"required"`
	TopupMethod string `json:"topup_method" validate:"required"`
}

func (r *CreateTopupRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateTopupRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	TopupID     int    `json:"topup_id" validate:"required"`
	TopupAmount int    `json:"topup_amount" validate:"required"`
	TopupMethod string `json:"topup_method" validate:"required"`
}

func (r *UpdateTopupRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
