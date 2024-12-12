package requests

import "github.com/go-playground/validator/v10"

type CreateMerchantRequest struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type UpdateMerchantRequest struct {
	MerchantID int    `json:"merchant_id"`
	Name       string `json:"name"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`
}

func (r CreateMerchantRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r UpdateMerchantRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
