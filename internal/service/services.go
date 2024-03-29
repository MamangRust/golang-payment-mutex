package service

import (
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/auth"
	"payment-mutex/pkg/hash"
	"payment-mutex/pkg/logger"
)

type Services struct {
	Auth     AuthService
	Saldo    SaldoService
	Topup    TopupService
	Transfer TransferService
	Withdraw WithdrawService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     logger.Logger
	Hash       *hash.Hashing
	Token      auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth:  NewAuthService(*deps.Hash, deps.Repository.User, deps.Token, deps.Logger),
		Saldo: NewSaldoService(deps.Repository.Saldo, deps.Logger),
		Topup: NewTopupService(deps.Repository.Topup, deps.Repository.Saldo, deps.Logger),
		Transfer: NewTransferService(
			deps.Repository.Transfer,
			deps.Repository.Saldo,
			deps.Logger,
		),
		Withdraw: NewWithdrawService(deps.Repository.Withdraw, deps.Repository.Saldo, deps.Logger),
	}
}
