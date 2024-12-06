package service

import (
	"errors"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type topupService struct {
	userRepository  repository.UserRepository
	topupRepository repository.TopupRepository
	saldoRepository repository.SaldoRepository

	logger logger.Logger
}

func NewTopupService(
	userRepository repository.UserRepository,
	topupRepository repository.TopupRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *topupService {
	return &topupService{
		userRepository:  userRepository,
		topupRepository: topupRepository,
		saldoRepository: saldoRepository,
		logger:          logger,
	}
}

func (s *topupService) FindAll() (*[]models.Topup, error) {
	topup, err := s.topupRepository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all topup: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) FindById(topupID int) (*models.Topup, error) {
	topup, err := s.topupRepository.Read(topupID)

	if err != nil {
		s.logger.Error("failed find topup by id: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) FindByUserID(userID int) (*models.Topup, error) {
	_, err := s.userRepository.Read(userID)

	if err != nil {
		s.logger.Error("failed find user id: ", zap.Error(err))
	}

	topup, err := s.topupRepository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find topup by user id: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) FindByUsersID(userID int) (*[]models.Topup, error) {
	_, err := s.userRepository.Read(userID)

	if err != nil {
		s.logger.Error("failed find user id: ", zap.Error(err))
	}

	topup, err := s.topupRepository.ReadByUsersID(userID)

	if err != nil {
		s.logger.Error("failed find topup by user id: ", zap.Error(err))

		return nil, err
	}

	return topup, nil
}

func (s *topupService) Create(request requests.CreateTopupRequest) (*models.Topup, error) {
	// Find user
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err))
		return nil, errors.New("user not found")
	}

	// Create topup
	topup, err := s.topupRepository.Create(request)
	if err != nil {
		s.logger.Error("Failed to create topup", zap.Error(err))
		return nil, err
	}

	// Find current saldo
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		if rollbackErr := s.topupRepository.Delete(topup.TopupID); rollbackErr != nil {
			s.logger.Error("Failed to rollback topup creation", zap.Error(rollbackErr))
		}
		return nil, err
	}

	newBalance := saldo.TotalBalance + request.TopupAmount
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       request.UserID,
		TotalBalance: newBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))
		if rollbackErr := s.topupRepository.Delete(topup.TopupID); rollbackErr != nil {
			s.logger.Error("Failed to rollback topup creation", zap.Error(rollbackErr))
		}
		return nil, err
	}

	return topup, nil
}

func (s *topupService) Update(request requests.UpdateTopupRequest) (*models.Topup, error) {
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err))
		return nil, errors.New("user not found")
	}

	// Find the existing topup
	existingTopup, err := s.topupRepository.Read(request.TopupID)
	if err != nil || existingTopup == nil {
		s.logger.Error("Failed to find topup by ID", zap.Error(err))
		return nil, errors.New("topup not found")
	}

	topupDifference := request.TopupAmount - existingTopup.TopupAmount

	// Update the topup amount
	_, err = s.topupRepository.UpdateAmount(requests.UpdateTopupAmount{
		TopupID:     request.TopupID,
		TopupAmount: request.TopupAmount,
	})
	if err != nil {
		s.logger.Error("Failed to update topup amount", zap.Error(err))
		return nil, err
	}

	// Retrieve the current balance from saldo
	currentSaldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to retrieve current saldo", zap.Error(err))
		return nil, err
	}

	if currentSaldo == nil {
		s.logger.Error("No saldo found for user", zap.Int("userID", request.UserID))
		return nil, errors.New("saldo not found")
	}

	newBalance := currentSaldo.TotalBalance + topupDifference

	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       request.UserID,
		TotalBalance: newBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))

		// Rollback the topup update if saldo update fails
		_, rollbackErr := s.topupRepository.UpdateAmount(requests.UpdateTopupAmount{
			TopupID:     request.TopupID,
			TopupAmount: existingTopup.TopupAmount,
		})
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback topup update", zap.Error(rollbackErr))
		}

		return nil, err
	}

	updatedTopup, err := s.topupRepository.Read(request.TopupID)
	if err != nil || updatedTopup == nil {
		s.logger.Error("Failed to find updated topup by ID", zap.Error(err))
		return nil, errors.New("updated topup not found")
	}

	return updatedTopup, nil
}

func (s *topupService) Delete(topupID int) error {
	err := s.topupRepository.Delete(topupID)

	if err != nil {
		s.logger.Error("failed delete topup: ", zap.Error(err))
		return err
	}

	return nil
}
