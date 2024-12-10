package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type transactionService struct {
	merchantRepository    repository.MerchantRepository
	cardRepository        repository.CardRepository
	saldoRepository       repository.SaldoRepository
	transactionRepository repository.TransactionRepository
	logger                logger.Logger
	mapper                responseMapper.TransactionResponseMapper
}

func NewTransactionService(
	merchantRepository repository.MerchantRepository,
	cardRepository repository.CardRepository,
	saldoRepository repository.SaldoRepository,
	transactionRepository repository.TransactionRepository,
	logger logger.Logger,
	mapper responseMapper.TransactionResponseMapper,
) *transactionService {
	return &transactionService{
		merchantRepository:    merchantRepository,
		cardRepository:        cardRepository,
		saldoRepository:       saldoRepository,
		transactionRepository: transactionRepository,
		logger:                logger,
		mapper:                mapper,
	}
}

func (s *transactionService) FindAll() (*response.ApiResponse[[]*response.TransactionResponse], *response.ErrorResponse) {
	transactions, err := s.transactionRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}

	so := s.mapper.ToTransactionsResponse(transactions)

	return &response.ApiResponse[[]*response.TransactionResponse]{
		Status:  "success",
		Message: "Transaction found",
		Data:    so,
	}, nil
}

func (s *transactionService) FindById(transactionID int) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse) {
	transaction, err := s.transactionRepository.Read(transactionID)
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}

	so := s.mapper.ToTransactionResponse(*transaction)

	return &response.ApiResponse[*response.TransactionResponse]{
		Status:  "success",
		Message: "Transaction found",
		Data:    so,
	}, nil
}

func (s *transactionService) Create(apiKey string, request requests.CreateTransactionRequest) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse) {
	// Ambil informasi merchant berdasarkan API key
	merchant, err := s.merchantRepository.ReadByApiKey(apiKey)
	if err != nil {
		s.logger.Error("failed to find merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	// Ambil informasi kartu berdasarkan nomor kartu
	card, err := s.cardRepository.ReadByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	// Ambil informasi saldo kartu
	saldo, err := s.saldoRepository.ReadByCardNumber(card.CardNumber)
	if err != nil {
		s.logger.Error("failed to find saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found",
		}
	}

	// Periksa apakah saldo mencukupi
	if saldo.TotalBalance < request.Amount {
		s.logger.Error("insufficient balance", zap.Int("available_balance", saldo.TotalBalance), zap.Int("transaction_amount", request.Amount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance",
		}
	}

	// Proses pengurangan saldo
	saldo.TotalBalance -= request.Amount
	if _, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   card.CardNumber,
		TotalBalance: saldo.TotalBalance,
	}); err != nil {
		s.logger.Error("failed to update saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo",
		}
	}

	// Buat transaksi
	request.MerchantID = &merchant.MerchantID
	transaction, err := s.transactionRepository.Create(request)
	if err != nil {
		// Rollback saldo jika pembuatan transaksi gagal
		s.logger.Error("failed to create transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction record",
		}
	}

	// Ambil informasi kartu dan saldo merchant
	merchantCard, err := s.cardRepository.ReadByUserID(merchant.UserID)
	if err != nil {
		s.logger.Error("failed to find merchant card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant card not found",
		}
	}

	merchantSaldo, err := s.saldoRepository.ReadByCardNumber(merchantCard.CardNumber)
	if err != nil {
		s.logger.Error("failed to find merchant saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant saldo not found",
		}
	}

	// Tambahkan saldo ke merchant
	merchantSaldo.TotalBalance += request.Amount
	if _, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   merchantCard.CardNumber,
		TotalBalance: merchantSaldo.TotalBalance,
	}); err != nil {
		s.logger.Error("failed to update merchant saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant saldo",
		}
	}

	// Map hasil transaksi ke response
	so := s.mapper.ToTransactionResponse(*transaction)
	return &response.ApiResponse[*response.TransactionResponse]{
		Status:  "success",
		Message: "Transaction created successfully",
		Data:    so,
	}, nil
}

func (s *transactionService) Update(apiKey string, request requests.UpdateTransactionRequest) (*response.ApiResponse[*response.TransactionResponse], *response.ErrorResponse) {
	// Validate merchant based on API key
	merchant, err := s.merchantRepository.ReadByApiKey(apiKey)
	if err != nil {
		s.logger.Error("failed to find merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	// Find the existing transaction
	existingTransaction, err := s.transactionRepository.Read(request.TransactionID)
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}

	// Validate that the merchant owns the transaction
	if existingTransaction.MerchantID != merchant.MerchantID {
		s.logger.Error("merchant not authorized to update transaction",
			zap.Int("transaction_merchant_id", existingTransaction.MerchantID),
			zap.Int("current_merchant_id", merchant.MerchantID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Not authorized to update this transaction",
		}
	}

	// Prepare update fields
	updateRequest := requests.UpdateTransactionRequest{
		TransactionID:   request.TransactionID,
		CardNumber:      request.CardNumber,
		Amount:          request.Amount,
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      request.MerchantID,
		TransactionTime: request.TransactionTime,
	}

	// Perform balance adjustment if amount changes
	if request.Amount != existingTransaction.Amount {
		// Get card information
		card, err := s.cardRepository.ReadByCardNumber(request.CardNumber)
		if err != nil {
			s.logger.Error("failed to find card", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Card not found",
			}
		}

		// Get current card balance
		saldo, err := s.saldoRepository.ReadByCardNumber(card.CardNumber)
		if err != nil {
			s.logger.Error("failed to find saldo", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Saldo not found",
			}
		}

		// Calculate balance difference
		amountDifference := request.Amount - existingTransaction.Amount

		// Check if sufficient balance for increased amount
		if amountDifference > 0 && saldo.TotalBalance < amountDifference {
			s.logger.Error("insufficient balance",
				zap.Int("available_balance", saldo.TotalBalance),
				zap.Int("additional_amount", amountDifference))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Insufficient balance",
			}
		}

		// Update card balance
		saldo.TotalBalance -= amountDifference
		if _, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			CardNumber:   card.CardNumber,
			TotalBalance: saldo.TotalBalance,
		}); err != nil {
			s.logger.Error("failed to update saldo", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to update saldo",
			}
		}
	}

	// Update transaction
	updatedTransaction, err := s.transactionRepository.Update(updateRequest)
	if err != nil {
		s.logger.Error("failed to update transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction record",
		}
	}

	// Map and return updated transaction
	so := s.mapper.ToTransactionResponse(*updatedTransaction)
	return &response.ApiResponse[*response.TransactionResponse]{
		Status:  "success",
		Message: "Transaction updated successfully",
		Data:    so,
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
