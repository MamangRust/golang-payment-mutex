package service

import (
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"time"

	"go.uber.org/zap"
)

type dashboardService struct {
	cardRepository        repository.CardRepository
	saldoRepository       repository.SaldoRepository
	transactionRepository repository.TransactionRepository
	topupRepository       repository.TopupRepository
	withdrawRepository    repository.WithdrawRepository
	transferRepository    repository.TransactionRepository
	merchantRepository    repository.MerchantRepository
	logger                logger.Logger
}

func NewDashboardService(
	cardRepository repository.CardRepository,
	saldoRepository repository.SaldoRepository,
	transactionRepository repository.TransactionRepository,
	topupRepository repository.TopupRepository,
	withdrawRepository repository.WithdrawRepository,
	transferRepository repository.TransactionRepository,
	merchantRepository repository.MerchantRepository,
	logger logger.Logger,
) *dashboardService {
	return &dashboardService{
		cardRepository:        cardRepository,
		saldoRepository:       saldoRepository,
		transactionRepository: transactionRepository,
		topupRepository:       topupRepository,
		withdrawRepository:    withdrawRepository,
		transferRepository:    transferRepository,
		merchantRepository:    merchantRepository,
		logger:                logger,
	}
}

func (s *dashboardService) GetGlobalOverview() (*response.ApiResponse[*response.OverviewData], *response.ErrorResponse) {
	overview := &response.OverviewData{
		ActivityTrends: make(map[string]int),
	}

	cards, _, err := s.cardRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch cards", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card data",
		}
	}

	totalBalance := 0
	activeCards := 0
	for _, card := range cards {
		saldo, err := s.saldoRepository.ReadByCardNumber(card.CardNumber)
		if err != nil {
			continue
		}
		totalBalance += saldo.TotalBalance
		activeCards++
	}
	overview.TotalBalance = totalBalance
	overview.ActiveCards = activeCards

	totalTransactions, _, err := s.transactionRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction data",
		}
	}
	overview.TotalTransaksi = len(totalTransactions)

	topups, _, err := s.topupRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch topups", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch topup data",
		}
	}

	totalTopup := 0
	topupAmount := 0
	for _, topup := range topups {
		totalTopup++
		topupAmount += topup.TopupAmount
	}
	overview.TotalTopup = totalTopup
	overview.TopupAmount = topupAmount

	withdrawals, _, err := s.withdrawRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch withdrawals", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdrawal data",
		}
	}

	totalWithdraw := 0
	withdrawAmount := 0
	for _, withdraw := range withdrawals {
		totalWithdraw++
		withdrawAmount += withdraw.WithdrawAmount
	}
	overview.TotalWithdraw = totalWithdraw
	overview.WithdrawAmount = withdrawAmount

	transfers, _, err := s.transferRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch transfers", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transfer data",
		}
	}
	overview.TotalTransfer = len(transfers)

	now := time.Now()
	startDate := now.AddDate(0, 0, -7)
	for d := startDate; d.Before(now); d = d.AddDate(0, 0, 1) {
		date := d.Format("2006-01-02")
		transactions, _ := s.transactionRepository.CountByDate(date)
		topups, _ := s.topupRepository.CountByDate(date)
		withdrawals, _ := s.withdrawRepository.CountByDate(date)
		transfers, _ := s.transferRepository.CountByDate(date)

		overview.ActivityTrends[date] = transactions + topups + withdrawals + transfers
	}

	merchants, _, err := s.merchantRepository.ReadAll(1, 1000, "")
	if err != nil {
		s.logger.Error("failed to fetch merchants", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch merchant data",
		}
	}

	totalMerchants := len(merchants)
	activeMerchants := 0
	for _, merchant := range merchants {
		if merchant.Status == "active" {
			activeMerchants++
		}
	}
	overview.TotalMerchants = totalMerchants
	overview.ActiveMerchants = activeMerchants

	// Return the overview response
	return &response.ApiResponse[*response.OverviewData]{
		Status:  "success",
		Message: "Global overview retrieved successfully",
		Data:    overview,
	}, nil
}
