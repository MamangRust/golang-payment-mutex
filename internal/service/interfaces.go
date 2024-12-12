package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
)

type AuthService interface {
	RegisterUser(request *requests.RegisterRequest) (*response.ApiResponse[response.UserResponse], *response.ErrorResponse)
	Login(request *requests.AuthRequest) (*response.ApiResponse[string], error)
}

type UserService interface {
	FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.UserResponse], *response.ErrorResponse)
	FindByID(id int) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse)
	Create(request requests.CreateUserRequest) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse)
	Update(request requests.UpdateUserRequest) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse)
	Delete(userID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type SaldoService interface {
	FindAll() (*response.ApiResponse[[]*response.SaldoResponse], *response.ErrorResponse)
	FindById(saldoID int) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse)
	Create(requests requests.CreateSaldoRequest) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse)
	Update(requests requests.UpdateSaldoRequest) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse)
	Delete(saldoID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type TopupService interface {
	FindAll() (*response.ApiResponse[[]*response.TopupResponse], *response.ErrorResponse)
	FindById(topupID int) (*response.ApiResponse[*response.TopupResponse], *response.ErrorResponse)
	Create(requests requests.CreateTopupRequest) (*response.ApiResponse[*response.TopupResponse], *response.ErrorResponse)
	Update(requests requests.UpdateTopupRequest) (*response.ApiResponse[*response.TopupResponse], *response.ErrorResponse)
	Delete(topupID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type TransferService interface {
	FindAll() (*response.ApiResponse[[]*response.TransferResponse], *response.ErrorResponse)
	FindById(transferID int) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse)
	Create(requests requests.CreateTransferRequest) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse)
	Update(requests requests.UpdateTransferRequest) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse)
	Delete(transferID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type CardService interface {
	FindAll() (*response.ApiResponse[[]*response.CardResponse], *response.ErrorResponse)
	FindById(cardID int) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse)
	Create(requests requests.CreateCardRequest) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse)
	Update(requests requests.UpdateCardRequest) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse)
	Delete(cardID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type WithdrawService interface {
	FindAll() (*response.ApiResponse[[]*response.WithdrawResponse], *response.ErrorResponse)
	FindById(withdrawID int) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse)
	Create(requests requests.CreateWithdrawRequest) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse)
	Update(requests requests.UpdateWithdrawRequest) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse)
	Delete(withdrawID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type TransactionService interface {
	FindAll() (*response.ApiResponse[[]*response.TransactionResponse], *response.ErrorResponse)
	FindById(transactionID int) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse)
	Create(apikey string, requests requests.CreateTransactionRequest) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse)
	Update(apikey string, requests requests.UpdateTransactionRequest) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse)
	Delete(transactionID int) (*response.ApiResponse[string], *response.ErrorResponse)
}

type DashboardService interface {
	GetGlobalOverview() (*response.ApiResponse[*response.OverviewData], *response.ErrorResponse)
}

type MerchantService interface {
	FindAll() (*response.ApiResponse[[]*response.MerchantResponse], *response.ErrorResponse)
	FindByName(name string) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse)
	FindByApiKey(apiKey string) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse)
	FindByID(merchantID int) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse)
	Create(requests requests.CreateMerchantRequest) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse)
	Update(requests requests.UpdateMerchantRequest) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse)
	Delete(merchantID int) (*response.ApiResponse[string], *response.ErrorResponse)
}
