package repository

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
)

type UserRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.UserRecord, int, error)
	Read(userID int) (*record.UserRecord, error)
	ReadByEmail(email string) (*record.UserRecord, error)
	Create(request requests.CreateUserRequest) (*record.UserRecord, error)
	Update(request requests.UpdateUserRequest) (*record.UserRecord, error)
	Delete(userID int) error
}

type SaldoRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.SaldoRecord, int, error)
	Read(saldoID int) (*record.SaldoRecord, error)
	ReadByCardNumber(cardNumber string) (*record.SaldoRecord, error)
	Create(request requests.CreateSaldoRequest) (*record.SaldoRecord, error)
	Update(request requests.UpdateSaldoRequest) (*record.SaldoRecord, error)
	UpdateBalance(request requests.UpdateSaldoBalance) (*record.SaldoRecord, error)
	UpdateSaldoWithdraw(request requests.UpdateSaldoWithdraw) (*record.SaldoRecord, error)
	Delete(saldoID int) error
}

type TopupRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.TopupRecord, int, error)
	Read(topupID int) (*record.TopupRecord, error)
	CountByDate(date string) (int, error)
	Create(request requests.CreateTopupRequest) (*record.TopupRecord, error)
	Update(request requests.UpdateTopupRequest) (*record.TopupRecord, error)
	UpdateAmount(request requests.UpdateTopupAmount) (*record.TopupRecord, error)
	Delete(topupID int) error
}

type TransferRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.TransferRecord, int, error)
	Read(transferID int) (*record.TransferRecord, error)
	CountByDate(date string) (int, error)
	CountAll() (int, error)
	Create(request requests.CreateTransferRequest) (*record.TransferRecord, error)
	Update(request requests.UpdateTransferRequest) (*record.TransferRecord, error)
	Delete(transferID int) error
}

type WithdrawRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.WithdrawRecord, int, error)
	Read(withdrawID int) (*record.WithdrawRecord, error)
	CountByDate(date string) (int, error)
	Create(request requests.CreateWithdrawRequest) (*record.WithdrawRecord, error)
	Update(request requests.UpdateWithdrawRequest) (*record.WithdrawRecord, error)
	Delete(transferID int) error
}

type CardRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.CardRecord, int, error)
	Read(cardID int) (*record.CardRecord, error)
	ReadByCardNumber(cardNumber string) (*record.CardRecord, error)
	ReadByUsersID(userID int) ([]*record.CardRecord, error)
	ReadByUserID(userID int) (*record.CardRecord, error)
	Create(request requests.CreateCardRequest) (*record.CardRecord, error)
	Update(request requests.UpdateCardRequest) (*record.CardRecord, error)
	Delete(cardID int) error
}

type TransactionRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.TransactionRecord, int, error)
	CountByDate(date string) (int, error)
	CountAll() (int, error)
	Read(transactionID int) (*record.TransactionRecord, error)
	Create(request requests.CreateTransactionRequest) (*record.TransactionRecord, error)
	Update(request requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	Delete(transactionID int) error
}

type MerchantRepository interface {
	ReadAll(page int, pageSize int, search string) ([]*record.MerchantRecord, int, error)
	Read(merchantID int) (*record.MerchantRecord, error)
	ReadByName(name string) (*record.MerchantRecord, error)
	ReadByApiKey(apiKey string) (*record.MerchantRecord, error)
	Create(request requests.CreateMerchantRequest) (*record.MerchantRecord, error)
	Update(request requests.UpdateMerchantRequest) (*record.MerchantRecord, error)
	Delete(merchantID int) error
}
