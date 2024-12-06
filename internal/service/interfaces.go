package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
)

type AuthService interface {
	RegisterUser(requests *requests.RegisterRequest) (*models.User, error)
	Login(request *requests.AuthRequest) (string, error)
}

type UserService interface {
	FindAll() (*[]models.User, error)
	FindByID(id int) (*models.User, error)
	Create(request requests.CreateUserRequest) (*models.User, error)
	Update(request requests.UpdateUserRequest) (*models.User, error)
	Delete(userID int) error
}

type SaldoService interface {
	FindAll() (*[]models.Saldo, error)
	FindByUserID(userID int) (*models.Saldo, error)
	FindByUsersID(userID int) (*[]models.Saldo, error)
	FindById(saldoID int) (*models.Saldo, error)
	Create(requests requests.CreateSaldoRequest) (*models.Saldo, error)
	Update(requests requests.UpdateSaldoRequest) (*models.Saldo, error)
	Delete(saldoID int) error
}

type TopupService interface {
	FindAll() (*[]models.Topup, error)
	FindById(topupID int) (*models.Topup, error)
	FindByUserID(userID int) (*models.Topup, error)
	FindByUsersID(userID int) (*[]models.Topup, error)
	Create(requests requests.CreateTopupRequest) (*models.Topup, error)
	Update(requests requests.UpdateTopupRequest) (*models.Topup, error)
	Delete(topupID int) error
}

type TransferService interface {
	FindAll() (*[]models.Transfer, error)
	FindById(transferID int) (*models.Transfer, error)
	FindByUsersID(userID int) (*[]models.Transfer, error)
	FindByUserID(userID int) (*models.Transfer, error)
	Create(requests requests.CreateTransferRequest) (*models.Transfer, error)
	Update(requests requests.UpdateTransferRequest) (*models.Transfer, error)
	Delete(transferID int) error
}

type WithdrawService interface {
	FindAll() (*[]models.Withdraw, error)
	FindByUsersID(userID int) (*[]models.Withdraw, error)
	FindByUserID(userID int) (*models.Withdraw, error)
	FindById(withdrawID int) (*models.Withdraw, error)
	Create(requests requests.CreateWithdrawRequest) (*models.Withdraw, error)
	Update(requests requests.UpdateWithdrawRequest) (*models.Withdraw, error)
	Delete(withdrawID int) error
}
