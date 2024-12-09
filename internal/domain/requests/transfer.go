package requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CreateTransferRequest struct {
	TransferFrom   string `json:"transfer_from" validate:"required"`
	TransferTo     string `json:"transfer_to" validate:"required,min=1"`
	TransferAmount int    `json:"transfer_amount" validate:"required,min=50000"`
}

type UpdateTransferRequest struct {
	TransferID     int    `json:"transfer_id" validate:"required,min=1"`
	TransferFrom   string `json:"transfer_from" validate:"required"`
	TransferTo     string `json:"transfer_to" validate:"required,min=1"`
	TransferAmount int    `json:"transfer_amount" validate:"required,min=50000"`
}

type UpdateTransferAmountRequest struct {
	TransferID     int `json:"transfer_id" validate:"required,min=1"`
	TransferAmount int `json:"transfer_amount" validate:"required,gt=0"`
}

func (r *CreateTransferRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TransferAmount < 50000 {
		return errors.New("transfer amount must be at least 50,000")
	}

	return nil
}

func (r *UpdateTransferRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TransferID <= 0 {
		return errors.New("transfer ID must be a positive integer")
	}

	if r.TransferAmount < 50000 {
		return errors.New("transfer amount must be at least 50,000")
	}

	return nil
}

func (r *UpdateTransferAmountRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TransferID <= 0 {
		return errors.New("transfer ID must be a positive integer")
	}

	if r.TransferAmount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	return nil
}
