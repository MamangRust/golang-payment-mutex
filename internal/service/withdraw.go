package service

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type withdrawService struct {
	userRepository     repository.UserRepository
	saldoRepository    repository.SaldoRepository
	withdrawRepository repository.WithdrawRepository
	logger             logger.Logger
}

func NewWithdrawService(
	userRepository repository.UserRepository,
	withdrawRepository repository.WithdrawRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *withdrawService {
	return &withdrawService{
		userRepository:     userRepository,
		saldoRepository:    saldoRepository,
		withdrawRepository: withdrawRepository,
		logger:             logger,
	}
}

func (s *withdrawService) FindAll() (*response.ApiResponse[[]*record.WithdrawRecord], *response.ErrorResponse) {
	withdraws, err := s.withdrawRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find all withdraws", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch all withdraw records.",
		}
	}

	return &response.ApiResponse[[]*record.WithdrawRecord]{
		Status:  "success",
		Message: "Successfully retrieved all withdraw records.",
		Data:    withdraws,
	}, nil
}

func (s *withdrawService) FindById(withdrawID int) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse) {
	withdraw, err := s.withdrawRepository.Read(withdrawID)
	if err != nil {
		s.logger.Error("failed to find withdraw by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw record by ID.",
		}
	}

	return &response.ApiResponse[*record.WithdrawRecord]{
		Status:  "success",
		Message: "Successfully retrieved withdraw record by ID.",
		Data:    withdraw,
	}, nil
}

func (s *withdrawService) FindByUserID(userID int) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse) {
	withdraw, err := s.withdrawRepository.ReadByUserID(userID)
	if err != nil {
		s.logger.Error("failed to find withdraw by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw record for user.",
		}
	}

	return &response.ApiResponse[*record.WithdrawRecord]{
		Status:  "success",
		Message: "Successfully retrieved withdraw record for user.",
		Data:    withdraw,
	}, nil
}

func (s *withdrawService) FindByUsersID(userID int) (*response.ApiResponse[[]*record.WithdrawRecord], *response.ErrorResponse) {
	withdraws, err := s.withdrawRepository.ReadByUsersID(userID)
	if err != nil {
		s.logger.Error("failed to find withdraws by users id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw records for users.",
		}
	}

	return &response.ApiResponse[[]*record.WithdrawRecord]{
		Status:  "success",
		Message: "Successfully retrieved withdraw records for users.",
		Data:    withdraws,
	}, nil
}

func (s *withdrawService) Create(request requests.CreateWithdrawRequest) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse) {
	// Cek saldo pengguna
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo == nil {
		s.logger.Error("Saldo not found for user", zap.Int("userID", request.UserID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found for the specified user ID.",
		}
	}

	// Periksa saldo mencukupi
	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.Int("userID", request.UserID), zap.Int("requested", request.WithdrawAmount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal.",
		}
	}

	// Update saldo setelah penarikan
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateData := requests.UpdateSaldoWithdraw{
		UserID:         request.UserID,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateData)
	if err != nil {
		s.logger.Error("Failed to update saldo after withdrawal", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo after withdrawal.",
		}
	}

	// Buat catatan withdraw
	withdrawRecord, err := s.withdrawRepository.Create(request)
	if err != nil {
		s.logger.Error("Failed to create withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create withdraw record.",
		}
	}

	return &response.ApiResponse[*record.WithdrawRecord]{
		Status:  "success",
		Message: "Withdrawal created successfully.",
		Data:    withdrawRecord,
	}, nil
}

func (s *withdrawService) Update(request requests.UpdateWithdrawRequest) (*response.ApiResponse[*record.WithdrawRecord], *response.ErrorResponse) {
	_, err := s.withdrawRepository.Read(request.WithdrawID)
	if err != nil {
		s.logger.Error("Failed to find withdraw record by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Withdraw record not found.",
		}
	}

	// Ambil saldo pengguna
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to fetch saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.Int("userID", request.UserID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal update.",
		}
	}

	// Update saldo baru
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateSaldoData := requests.UpdateSaldoWithdraw{
		UserID:         saldo.UserID,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateSaldoData)
	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo balance.",
		}
	}

	updatedWithdraw, err := s.withdrawRepository.Update(request)
	if err != nil {
		rollbackData := requests.UpdateSaldoBalance{
			UserID:       saldo.UserID,
			TotalBalance: saldo.TotalBalance,
		}
		_, rollbackErr := s.saldoRepository.UpdateBalance(rollbackData)
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback saldo after withdraw update failure", zap.Error(rollbackErr))
		}
		s.logger.Error("Failed to update withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update withdraw record.",
		}
	}

	return &response.ApiResponse[*record.WithdrawRecord]{
		Status:  "success",
		Message: "Withdraw record updated successfully.",
		Data:    updatedWithdraw,
	}, nil
}

func (s *withdrawService) Delete(withdrawID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.withdrawRepository.Delete(withdrawID)
	if err != nil {
		s.logger.Error("Failed to delete withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete withdraw record.",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Withdraw record deleted successfully.",
		Data:    "Record deleted",
	}, nil
}
