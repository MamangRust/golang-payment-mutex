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
	logger                logger.Logger
}

type OverviewData struct {
	TotalBalance   int            `json:"total_balance"`
	ActiveCards    int            `json:"active_cards"`
	TotalTransaksi int            `json:"total_transaksi"`
	TotalTopup     int            `json:"total_topup"`
	TopupAmount    int            `json:"topup_amount"`
	TotalWithdraw  int            `json:"total_withdraw"`
	WithdrawAmount int            `json:"withdraw_amount"`
	TotalTransfer  int            `json:"total_transfer"`
	ActivityTrends map[string]int `json:"activity_trends"`
}

func NewDashboardService(
	cardRepository repository.CardRepository,
	saldoRepository repository.SaldoRepository,
	transactionRepository repository.TransactionRepository,
	topupRepository repository.TopupRepository,
	withdrawRepository repository.WithdrawRepository,
	transferRepository repository.TransactionRepository,
	logger logger.Logger,
) *dashboardService {
	return &dashboardService{
		cardRepository:        cardRepository,
		saldoRepository:       saldoRepository,
		transactionRepository: transactionRepository,
		topupRepository:       topupRepository,
		withdrawRepository:    withdrawRepository,
		transferRepository:    transferRepository,
		logger:                logger,
	}
}

func (s *dashboardService) GetGlobalOverview() (*response.ApiResponse[*OverviewData], *response.ErrorResponse) {
	// Initialize the overview data
	overview := &OverviewData{
		ActivityTrends: make(map[string]int),
	}

	// 1. Calculate Total Balance and Active Cards
	cards, err := s.cardRepository.ReadAll()
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
			continue // Skip cards with no saldo data
		}
		totalBalance += saldo.TotalBalance
		activeCards++
	}
	overview.TotalBalance = totalBalance
	overview.ActiveCards = activeCards

	// 2. Count Total Transactions
	totalTransactions, err := s.transactionRepository.CountAll()
	if err != nil {
		s.logger.Error("failed to count transactions", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to count transactions",
		}
	}
	overview.TotalTransaksi = totalTransactions

	// 3. Count and Sum Topups
	topups, err := s.topupRepository.ReadAll()
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

	// 4. Count and Sum Withdrawals
	withdrawals, err := s.withdrawRepository.ReadAll()
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

	// 5. Count Total Transfers
	totalTransfers, err := s.transferRepository.CountAll()
	if err != nil {
		s.logger.Error("failed to count transfers", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to count transfer data",
		}
	}
	overview.TotalTransfer = totalTransfers

	// 6. Generate Activity Trends (Past 7 Days)
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

	// Return the overview response
	return &response.ApiResponse[*OverviewData]{
		Status:  "success",
		Message: "Global overview retrieved successfully",
		Data:    overview,
	}, nil
}
