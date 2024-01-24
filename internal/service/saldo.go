package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type saldoService struct {
	repository repository.SaldoRepository
	logger     logger.Logger
}

func NewSaldoService(repository repository.SaldoRepository, logger logger.Logger) *saldoService {
	return &saldoService{
		repository: repository,
		logger:     logger,
	}
}

func (s *saldoService) FindAll() (*[]models.Saldo, error) {
	saldo, err := s.repository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all saldo: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) FindByUserID(userID int) (*models.Saldo, error) {
	saldo, err := s.repository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) FindById(saldoID int) (*models.Saldo, error) {
	saldo, err := s.repository.Read(saldoID)

	if err != nil {
		s.logger.Error("failed find saldo by id: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) Create(requests requests.CreateSaldoRequest) (*models.Saldo, error) {
	if requests.TotalBalance > 50000 {
		return nil, fmt.Errorf("total balance must be less than 50000")
	}

	saldo := models.Saldo{
		UserID:       requests.UserID,
		TotalBalance: requests.TotalBalance,
	}

	res, err := s.repository.Create(saldo)

	if err != nil {
		s.logger.Error("failed create saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}
func (s *saldoService) Update(requests requests.UpdateSaldoRequest) (*models.Saldo, error) {
	if requests.TotalBalance > 50000 {
		return nil, fmt.Errorf("total balance must be less than 50000")
	}

	saldo := models.Saldo{
		SaldoID:      requests.SaldoID,
		UserID:       requests.UserID,
		TotalBalance: requests.TotalBalance,
	}

	res, err := s.repository.Update(saldo)

	if err != nil {
		s.logger.Error("failed update saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *saldoService) Delete(saldoID int) error {
	err := s.repository.Delete(saldoID)

	if err != nil {
		s.logger.Error("failed delete saldo: ", zap.Error(err))
		return err
	}

	return nil
}
