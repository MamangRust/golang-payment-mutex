package requests

import "github.com/go-playground/validator/v10"

type CreateTransferRequest struct {
	TransferFrom   int `json:"transfer_from" validate:"required"`
	TransferTo     int `json:"transfer_to" validate:"required"`
	TransferAmount int `json:"transfer_amount" validate:"required"`
}

func (r *CreateTransferRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateTransferRequest struct {
	TransferID     int `json:"transfer_id" validate:"required"`
	TransferFrom   int `json:"transfer_from" validate:"required"`
	TransferTo     int `json:"transfer_to" validate:"required"`
	TransferAmount int `json:"transfer_amount" validate:"required"`
}

func (r *UpdateTransferRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
