package repository

import "payment-mutex/internal/models"

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
	Create(saldo models.Saldo) (*models.Saldo, error)
	Update(newSaldo models.Saldo) (*models.Saldo, error)
	Delete(saldoID int) error
}

type TopupRepository interface {
	ReadAll() (*[]models.Topup, error)
	Read(topupID int) (*models.Topup, error)
	ReadByUserID(userID int) (*models.Topup, error)
	Create(topup models.Topup) (*models.Topup, error)
	Update(newTopup models.Topup) (*models.Topup, error)
	Delete(topupID int) error
}

type TransferRepository interface {
	ReadAll() (*[]models.Transfer, error)
	Read(transferID int) (*models.Transfer, error)
	ReadByUserID(userID int) (*models.Transfer, error)
	Create(transfer models.Transfer) (*models.Transfer, error)
	Update(newTransfer models.Transfer) (*models.Transfer, error)
	Delete(transferID int) error
}

type WithdrawRepository interface {
	ReadAll() (*[]models.Withdraw, error)
	Read(withdrawID int) (*models.Withdraw, error)
	ReadByUserID(userID int) (*models.Withdraw, error)
	Create(withdraw models.Withdraw) (*models.Withdraw, error)
	Update(newWithdraw models.Withdraw) (*models.Withdraw, error)
	Delete(transferID int) error
}
