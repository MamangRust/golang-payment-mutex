package service

import (
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/auth"
	"payment-mutex/pkg/hash"
	"payment-mutex/pkg/logger"
)

type Services struct {
	Auth        AuthService
	Saldo       SaldoService
	Topup       TopupService
	Transfer    TransferService
	Withdraw    WithdrawService
	User        UserService
	Card        CardService
	Transaction TransactionService
	Dashboard   DashboardService
	Merchant    MerchantService
}

type Deps struct {
	Repository     *repository.Repositories
	Logger         logger.Logger
	Hash           *hash.Hashing
	Token          auth.TokenManager
	MapperResponse responseMapper.ResponseMapper
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth:  NewAuthService(*deps.Hash, deps.Repository.User, deps.Token, deps.Logger, deps.MapperResponse.UserResponseMapper),
		Saldo: NewSaldoService(deps.Repository.Card, deps.Repository.Saldo, deps.Logger, deps.MapperResponse.SaldoResponseMapper),
		Topup: NewTopupService(deps.Repository.Card, deps.Repository.Topup, deps.Repository.Saldo, deps.Logger, deps.MapperResponse.TopupResponseMapper),
		Transfer: NewTransferService(
			deps.Repository.User,
			deps.Repository.Card,
			deps.Repository.Transfer,
			deps.Repository.Saldo,
			deps.Logger,
			deps.MapperResponse.TransferResponseMapper,
		),
		Withdraw: NewWithdrawService(deps.Repository.User, deps.Repository.Withdraw, deps.Repository.Saldo, deps.Logger, deps.MapperResponse.WithdrawResponseMapper),
		User:     NewUserService(deps.Repository.User, deps.Logger, deps.MapperResponse.UserResponseMapper),
		Card:     NewCardService(deps.Repository.Card, deps.Repository.User, deps.Repository.Saldo, deps.Logger, deps.MapperResponse.CardResponseMapper),
		Transaction: NewTransactionService(
			deps.Repository.Merchant,
			deps.Repository.Card, deps.Repository.Saldo, deps.Repository.Transaction, deps.Logger, deps.MapperResponse.TransactionResponseMapper),
		Dashboard: NewDashboardService(
			deps.Repository.Card, deps.Repository.Saldo, deps.Repository.Transaction, deps.Repository.Topup, deps.Repository.Withdraw, deps.Repository.Transaction,
			deps.Logger,
		),
		Merchant: NewMerchantService(
			deps.Repository.Merchant,
			deps.Logger,
			deps.MapperResponse.MerchantResponseMapper,
		),
	}
}
