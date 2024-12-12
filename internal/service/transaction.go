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

func (s *transactionService) FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.TransactionResponse], *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	transactions, totalRecords, err := s.transactionRepository.ReadAll(page, pageSize, search)

	if err != nil {
		s.logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transactions",
		}
	}

	if len(transactions) == 0 {
		s.logger.Error("no transactions found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No transactions found",
		}
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	so := s.mapper.ToTransactionsResponse(transactions)

	return &response.APIResponsePagination[[]*response.TransactionResponse]{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    so,
		Meta: response.PaginationMeta{
			TotalRecords: totalRecords,
			CurrentPage:  page,
			TotalPages:   totalPages,
			PageSize:     pageSize,
		},
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
	// s.logger.Info("Checking balance", zap.Int("AvailableBalance", saldo.TotalBalance), zap.Int("TransactionAmount", request.Amount))
	if saldo.TotalBalance < request.Amount {
		s.logger.Error("insufficient balance", zap.Int("AvailableBalance", saldo.TotalBalance), zap.Int("TransactionAmount", request.Amount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance",
		}
	}

	// Proses pengurangan saldo
	saldo.TotalBalance -= request.Amount
	// s.logger.Info("Balance after deduction", zap.Int("NewBalance", saldo.TotalBalance))
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
	// s.logger.Info("Creating transaction", zap.Any("Request", zapcore.Field.A))
	transaction, err := s.transactionRepository.Create(request)
	if err != nil {
		// Rollback saldo jika pembuatan transaksi gagal
		saldo.TotalBalance += request.Amount
		s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			CardNumber:   card.CardNumber,
			TotalBalance: saldo.TotalBalance,
		})
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
	// s.logger.Info("Updating merchant saldo", zap.Int("NewMerchantBalance", merchantSaldo.TotalBalance))
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
	transaction, err := s.transactionRepository.Read(request.TransactionID)
	if err != nil {
		s.logger.Error("failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}

	// Validasi apakah merchant API key sesuai dengan merchant ID dalam transaksi
	merchant, err := s.merchantRepository.ReadByApiKey(apiKey)

	if err != nil || transaction.MerchantID != merchant.MerchantID {
		s.logger.Error("unauthorized access to transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized access to transaction",
		}
	}

	// Ambil informasi kartu dan saldo berdasarkan kartu yang terkait dengan transaksi
	card, err := s.cardRepository.ReadByCardNumber(transaction.CardNumber)
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

	// Kembalikan saldo ke pengguna untuk jumlah transaksi lama
	saldo.TotalBalance += transaction.Amount
	s.logger.Debug("Restoring balance for old transaction amount", zap.Int("RestoredBalance", saldo.TotalBalance))
	if _, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   card.CardNumber,
		TotalBalance: saldo.TotalBalance,
	}); err != nil {
		s.logger.Error("failed to restore balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore balance",
		}
	}

	// Validasi saldo untuk jumlah transaksi baru
	if saldo.TotalBalance < request.Amount {
		s.logger.Error("insufficient balance for updated amount", zap.Int("AvailableBalance", saldo.TotalBalance), zap.Int("UpdatedAmount", request.Amount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for updated transaction",
		}
	}

	// Perbarui saldo berdasarkan jumlah transaksi baru
	saldo.TotalBalance -= request.Amount
	s.logger.Info("Updating balance for updated transaction amount")
	if _, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   card.CardNumber,
		TotalBalance: saldo.TotalBalance,
	}); err != nil {
		s.logger.Error("failed to update balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update balance",
		}
	}

	transaction.Amount = request.Amount
	transaction.PaymentMethod = request.PaymentMethod

	res, err := s.transactionRepository.Update(requests.UpdateTransactionRequest{
		TransactionID:   transaction.TransactionID,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      &transaction.MerchantID,
		TransactionTime: transaction.TransactionTime,
	})

	if err != nil {
		s.logger.Error("failed to update transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction",
		}
	}

	so := s.mapper.ToTransactionResponse(*res)
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
