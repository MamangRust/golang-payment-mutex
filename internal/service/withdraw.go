package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
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

func (s *withdrawService) FindAll() (*[]models.Withdraw, error) {
	withdraw, err := s.withdrawRepository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all withdraw: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) FindByUserID(userID int) (*models.Withdraw, error) {
	withdraw, err := s.withdrawRepository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find withdraw by user id: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) FindByUsersID(userID int) (*[]models.Withdraw, error) {
	withdraw, err := s.withdrawRepository.ReadByUsersID(userID)

	if err != nil {
		s.logger.Error("failed find withdraw by users id: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) FindById(withdrawID int) (*models.Withdraw, error) {
	withdraw, err := s.withdrawRepository.Read(withdrawID)

	if err != nil {
		s.logger.Error("failed find withdraw by id: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) Create(request requests.CreateWithdrawRequest) (*models.Withdraw, error) {
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		return nil, fmt.Errorf("saldo with user ID %d not found: %w", request.UserID, err)
	}

	if saldo == nil {
		s.logger.Error("Saldo not found for user ID", zap.Int("userID", request.UserID))
		return nil, fmt.Errorf("saldo not found")
	}

	// Check for sufficient balance
	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.Int("userID", request.UserID), zap.Int("requested", request.WithdrawAmount))
		return nil, fmt.Errorf("insufficient balance")
	}
	s.logger.Info("User has sufficient balance for withdrawal")

	// Update the saldo balance after withdrawal
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount

	updateData := requests.UpdateSaldoWithdraw{
		UserID:         request.UserID,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateData)
	if err != nil {
		s.logger.Error("Failed to update sender's saldo", zap.Error(err))
		return nil, fmt.Errorf("failed to update sender's saldo: %w", err)
	}

	// Create the withdraw record
	withdrawRecord, err := s.withdrawRepository.Create(request)
	if err != nil {
		s.logger.Error("Failed to create withdraw record", zap.Error(err))
		return nil, fmt.Errorf("failed to create withdraw record: %w", err)
	}

	return withdrawRecord, nil
}

func (s *withdrawService) Update(request requests.UpdateWithdrawRequest) (*models.Withdraw, error) {
	_, err := s.withdrawRepository.Read(request.WithdrawID)
	if err != nil {
		s.logger.Error("Failed to find withdraw record by ID", zap.Error(err))
		return nil, fmt.Errorf("withdraw record not found")
	}

	// Retrieve saldo
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		return nil, fmt.Errorf("saldo not found")
	}

	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.Int("userID", request.UserID))
		return nil, fmt.Errorf("insufficient balance")
	}

	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateData := requests.UpdateSaldoWithdraw{
		UserID:         saldo.UserID,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
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
		return nil, fmt.Errorf("failed to update withdraw record: %w", err)
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateData)
	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))
		return nil, fmt.Errorf("failed to update saldo balance: %w", err)
	}

	return updatedWithdraw, nil
}

func (s *withdrawService) Delete(withdrawID int) error {
	err := s.withdrawRepository.Delete(withdrawID)

	if err != nil {
		s.logger.Error("failed delete withdraw: ", zap.Error(err))
		return err
	}

	return nil
}
