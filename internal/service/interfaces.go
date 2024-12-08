package service

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
)

type AuthService interface {
	RegisterUser(request *requests.RegisterRequest) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse)
	Login(request *requests.AuthRequest) (*response.ApiResponse[string], error)
}

type UserService interface {
	FindAll() (*response.ApiResponse[[]*record.UserRecord], *response.ErrorResponse)
	FindByID(id int) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse)
	Create(request requests.CreateUserRequest) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse)
	Update(request requests.UpdateUserRequest) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse)
	Delete(userID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type SaldoService interface {
	FindAll() (*response.ApiResponse[[]*record.SaldoRecord], *response.ErrorResponse)
	FindByUserID(userID int) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse)
	FindByUsersID(userID int) (*response.ApiResponse[[]*record.SaldoRecord], *response.ErrorResponse)
	FindById(saldoID int) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse)
	Create(requests requests.CreateSaldoRequest) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse)
	Update(requests requests.UpdateSaldoRequest) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse)
	Delete(saldoID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type TopupService interface {
	FindAll() (*response.ApiResponse[[]*record.TopupRecord], *response.ErrorResponse)
	FindById(topupID int) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse)
	FindByUserID(userID int) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse)
	FindByUsersID(userID int) (*response.ApiResponse[[]*record.TopupRecord], *response.ErrorResponse)
	Create(requests requests.CreateTopupRequest) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse)
	Update(requests requests.UpdateTopupRequest) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse)
	Delete(topupID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type TransferService interface {
	FindAll() (*response.ApiResponse[[]*record.TransferRecord], *response.ErrorResponse)
	FindById(transferID int) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse)
	FindByUsersID(userID int) (*response.ApiResponse[[]*record.TransferRecord], *response.ErrorResponse)
	FindByUserID(userID int) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse)
	Create(requests requests.CreateTransferRequest) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse)
	Update(requests requests.UpdateTransferRequest) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse)
	Delete(transferID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type WithdrawService interface {
	FindAll() (*response.ApiResponse[[]*record.WithdrawRecord], *response.ErrorResponse)
	FindByUsersID(userID int) (*response.ApiResponse[[]*record.WithdrawRecord], *response.ErrorResponse)
	FindByUserID(userID int) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse)
	FindById(withdrawID int) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse)
	Create(requests requests.CreateWithdrawRequest) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse)
	Update(requests requests.UpdateWithdrawRequest) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse)
	Delete(withdrawID int) (*response.ApiResponse[string], *response.ErrorResponse)
}
