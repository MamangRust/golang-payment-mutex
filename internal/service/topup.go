package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"time"

	"go.uber.org/zap"
)

type topupService struct {
	saldoRepository repository.SaldoRepository
	repository      repository.TopupRepository
	logger          logger.Logger
}

func NewTopupService(repository repository.TopupRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *topupService {
	return &topupService{
		repository:      repository,
		saldoRepository: saldoRepository,
		logger:          logger,
	}
}

func (s *topupService) FindAll() (*[]models.Topup, error) {
	topup, err := s.repository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all topup: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) FindById(topupID int) (*models.Topup, error) {
	topup, err := s.repository.Read(topupID)

	if err != nil {
		s.logger.Error("failed find topup by id: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) FindByUserID(userID int) (*models.Topup, error) {
	topup, err := s.repository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find topup by user id: ", zap.Error(err))
		return nil, err
	}

	return topup, nil
}

func (s *topupService) Create(requests requests.CreateTopupRequest) (*models.Topup, error) {
	if requests.TopupAmount > 50000 {
		return nil, fmt.Errorf("topup amount must be less than 50000")
	}

	saldo, err := s.saldoRepository.ReadByUserID(requests.UserID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	topup := models.Topup{
		UserID:      requests.UserID,
		TopupAmount: requests.TopupAmount,
		TopupMethod: requests.TopupMethod,
		TopupTime:   time.Now(),
	}

	res, err := s.repository.Create(topup)

	if err != nil {
		s.logger.Error("failed create topup: ", zap.Error(err))
		return nil, err
	}

	saldo.TotalBalance = saldo.TotalBalance - requests.TopupAmount

	_, err = s.saldoRepository.Update(*saldo)

	if err != nil {
		s.logger.Error("failed update saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *topupService) Update(requests requests.UpdateTopupRequest) (*models.Topup, error) {
	if requests.TopupAmount > 50000 {
		return nil, fmt.Errorf("topup amount must be less than 50000")
	}

	topup := models.Topup{
		TopupID:     requests.TopupID,
		TopupAmount: requests.TopupAmount,
		TopupMethod: requests.TopupMethod,
	}

	res, err := s.repository.Update(topup)

	if err != nil {
		s.logger.Error("failed update topup: ", zap.Error(err))
		return nil, err
	}

	saldo, err := s.saldoRepository.ReadByUserID(requests.UserID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	saldo.TotalBalance = saldo.TotalBalance + requests.TopupAmount

	_, err = s.saldoRepository.Update(*saldo)

	if err != nil {
		s.logger.Error("failed update saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *topupService) Delete(topupID int) error {
	err := s.repository.Delete(topupID)

	if err != nil {
		s.logger.Error("failed delete topup: ", zap.Error(err))
		return err
	}

	return nil
}
