package service

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type transactionService struct {
	cardRepository        repository.CardRepository
	saldoRepository       repository.SaldoRepository
	transactionRepository repository.TransactionRepository
	logger                logger.Logger
}

func NewTransactionService(
	cardRepository repository.CardRepository,
	saldoRepository repository.SaldoRepository,
	transactionRepository repository.TransactionRepository,
	logger logger.Logger,
) *transactionService {
	return &transactionService{
		cardRepository:        cardRepository,
		saldoRepository:       saldoRepository,
		transactionRepository: transactionRepository,
		logger:                logger,
	}
}

func (s *transactionService) FindAll() (*response.ApiResponse[[]*record.TransactionRecord], *response.ErrorResponse) {
	transactions, err := s.transactionRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}
	return &response.ApiResponse[[]*record.TransactionRecord]{
		Status:  "success",
		Message: "Transaction found",
		Data:    transactions,
	}, nil
}

func (s *transactionService) FindById(transactionID int) (*response.ApiResponse[*record.TransactionRecord], *response.ErrorResponse) {
	transaction, err := s.transactionRepository.Read(transactionID)
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}
	return &response.ApiResponse[*record.TransactionRecord]{
		Status:  "success",
		Message: "Transaction found",
		Data:    transaction,
	}, nil
}

func (s *transactionService) Create(request requests.CreateTransactionRequest) (*response.ApiResponse[*record.TransactionRecord], *response.ErrorResponse) {
	card, err := s.cardRepository.ReadByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	saldo, err := s.saldoRepository.ReadByCardNumber(card.CardNumber)
	if err != nil {
		s.logger.Error("failed to find saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found",
		}
	}

	// Periksa apakah saldo cukup
	if saldo.TotalBalance < request.Amount {
		s.logger.Error("insufficient balance", zap.Int("available_balance", saldo.TotalBalance), zap.Int("transaction_amount", request.Amount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance",
		}
	}

	// Mulai transaksi pengurangan saldo
	saldo.TotalBalance -= request.Amount

	// Perbarui saldo di repository
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   card.CardNumber,
		TotalBalance: saldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("failed to update saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo",
		}
	}

	// Buat transaksi
	transaction, err := s.transactionRepository.Create(request)
	if err != nil {
		// Rollback perubahan saldo jika transaksi gagal
		saldo.TotalBalance += request.Amount
		_, rollbackErr := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			CardNumber:   card.CardNumber,
			TotalBalance: saldo.TotalBalance,
		})
		if rollbackErr != nil {
			s.logger.Error("failed to rollback saldo", zap.Error(rollbackErr))
		}

		s.logger.Error("failed to create transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction record",
		}
	}

	return &response.ApiResponse[*record.TransactionRecord]{
		Status:  "success",
		Message: "Transaction created successfully",
		Data:    transaction,
	}, nil
}

func (s *transactionService) Update(request requests.UpdateTransactionRequest) (*response.ApiResponse[*record.TransactionRecord], *response.ErrorResponse) {
	card, err := s.cardRepository.ReadByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	saldo, err := s.saldoRepository.ReadByCardNumber(card.CardNumber)
	if err != nil {
		s.logger.Error("failed to find saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found",
		}
	}

	// Perbarui transaksi
	transaction, err := s.transactionRepository.Update(request)
	if err != nil {
		s.logger.Error("failed to update transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction record",
		}
	}

	// Jika transaksi melibatkan perubahan jumlah saldo, lakukan rollback jika perlu
	if request.Amount != 0 {
		saldo.TotalBalance += request.Amount // Sesuaikan saldo jika diperlukan

		_, rollbackErr := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			CardNumber:   card.CardNumber,
			TotalBalance: saldo.TotalBalance,
		})
		if rollbackErr != nil {
			s.logger.Error("failed to rollback saldo after transaction update", zap.Error(rollbackErr))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Transaction updated but failed to rollback saldo",
			}
		}
	}

	return &response.ApiResponse[*record.TransactionRecord]{
		Status:  "success",
		Message: "Transaction updated successfully",
		Data:    transaction,
	}, nil
}

func (s *transactionService) Delete(transactionID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.transactionRepository.Delete(transactionID)
	if err != nil {
		s.logger.Error("failed to delete transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction record",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Transaction deleted successfully",
		Data:    fmt.Sprintf("Transaction with ID %d has been deleted", transactionID),
	}, nil
}
