package repository

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
)

type UserRepository interface {
	ReadAll() []models.User
	Read(userID int) (*models.User, error)
	ReadByEmail(email string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(userID int, newUser models.User) (*models.User, error)
	Delete(userID int) bool
}

type SaldoRepository interface {
	ReadAll() (*[]models.Saldo, error)
	Read(saldoID int) (*models.Saldo, error)
	ReadByUserID(userID int) (*models.Saldo, error)
	ReadByUsersID(userID int) (*[]models.Saldo, error)
	Create(request requests.CreateSaldoRequest) (*models.Saldo, error)
	Update(request requests.UpdateSaldoRequest) (*models.Saldo, error)
	UpdateBalance(request requests.UpdateSaldoBalance) (*models.Saldo, error)
	UpdateSaldoWithdraw(request requests.UpdateSaldoWithdraw) (*models.Saldo, error)
	Delete(saldoID int) error
}

type TopupRepository interface {
	ReadAll() (*[]models.Topup, error)
	Read(topupID int) (*models.Topup, error)
	ReadByUserID(userID int) (*models.Topup, error)
	ReadByUsersID(userID int) (*[]models.Topup, error)
	Create(request requests.CreateTopupRequest) (*models.Topup, error)
	Update(request requests.UpdateTopupRequest) (*models.Topup, error)
	UpdateAmount(request requests.UpdateTopupAmount) (*models.Topup, error)
	Delete(topupID int) error
}

type TransferRepository interface {
	ReadAll() (*[]models.Transfer, error)
	Read(transferID int) (*models.Transfer, error)
	ReadByUsersID(userID int) (*[]models.Transfer, error)
	ReadByUserID(userID int) (*models.Transfer, error)
	Create(request requests.CreateTransferRequest) (*models.Transfer, error)
	Update(request requests.UpdateTransferRequest) (*models.Transfer, error)
	Delete(transferID int) error
}

type WithdrawRepository interface {
	ReadAll() (*[]models.Withdraw, error)
	Read(withdrawID int) (*models.Withdraw, error)
	ReadByUsersID(userID int) (*[]models.Withdraw, error)
	ReadByUserID(userID int) (*models.Withdraw, error)
	Create(request requests.CreateWithdrawRequest) (*models.Withdraw, error)
	Update(request requests.UpdateWithdrawRequest) (*models.Withdraw, error)
	Delete(transferID int) error
}
