package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type saldoService struct {
	userRepository  repository.UserRepository
	saldoRepository repository.SaldoRepository
	logger          logger.Logger
}

func NewSaldoService(
	userRepository repository.UserRepository,
	saldoRepository repository.SaldoRepository,
	logger logger.Logger) *saldoService {
	return &saldoService{
		userRepository:  userRepository,
		saldoRepository: saldoRepository,
		logger:          logger,
	}
}

func (s *saldoService) FindAll() (*[]models.Saldo, error) {
	saldo, err := s.saldoRepository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all saldo: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) FindById(saldoID int) (*models.Saldo, error) {
	saldo, err := s.saldoRepository.Read(saldoID)

	if err != nil {
		s.logger.Error("failed find saldo by id: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) FindByUsersID(userID int) (*[]models.Saldo, error) {
	saldo, err := s.saldoRepository.ReadByUsersID(userID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))

		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) FindByUserID(userID int) (*models.Saldo, error) {
	saldo, err := s.saldoRepository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	return saldo, nil
}

func (s *saldoService) Create(requests requests.CreateSaldoRequest) (*models.Saldo, error) {
	_, err := s.userRepository.Read(requests.UserID)

	if err != nil {
		s.logger.Error("failed find user not found")

		return nil, err
	}

	res, err := s.saldoRepository.Create(requests)

	if err != nil {
		s.logger.Error("failed create saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *saldoService) Update(requests requests.UpdateSaldoRequest) (*models.Saldo, error) {
	_, err := s.userRepository.Read(requests.UserID)

	if err != nil {
		s.logger.Error("failed find user not found")

		return nil, err
	}

	res, err := s.saldoRepository.Update(requests)

	if err != nil {
		s.logger.Error("failed update saldo: ", zap.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *saldoService) Delete(saldoID int) error {
	err := s.saldoRepository.Delete(saldoID)

	if err != nil {
		s.logger.Error("failed delete saldo: ", zap.Error(err))
		return err
	}

	return nil
}
