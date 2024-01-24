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
	saldoRepository repository.SaldoRepository
	repository      repository.WithdrawRepository
	logger          logger.Logger
}

func NewWithdrawService(repository repository.WithdrawRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *withdrawService {
	return &withdrawService{
		saldoRepository: saldoRepository,
		repository:      repository,
		logger:          logger,
	}
}

func (s *withdrawService) FindAll() (*[]models.Withdraw, error) {
	withdraw, err := s.repository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all withdraw: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) FindByUserID(userID int) (*models.Withdraw, error) {
	withdraw, err := s.repository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find withdraw by user id: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) FindById(withdrawID int) (*models.Withdraw, error) {
	withdraw, err := s.repository.Read(withdrawID)

	if err != nil {
		s.logger.Error("failed find withdraw by id: ", zap.Error(err))
		return nil, err
	}

	return withdraw, nil
}

func (s *withdrawService) Create(withdraw requests.CreateWithdrawRequest) (*models.Withdraw, error) {
	saldo, err := s.saldoRepository.ReadByUserID(withdraw.UserID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	if saldo.TotalBalance < withdraw.WithdrawAmount {
		return nil, fmt.Errorf("sender saldo not enough")
	}

	saldo.TotalBalance = saldo.TotalBalance - withdraw.WithdrawAmount

	_, err = s.saldoRepository.Update(*saldo)

	if err != nil {
		s.logger.Error("failed update sender saldo: ", zap.Error(err))
		return nil, err
	}

	withdrawModel := models.Withdraw{
		WithdrawTime:   withdraw.WithdrawTime,
		WithdrawAmount: withdraw.WithdrawAmount,
		UserID:         withdraw.UserID,
	}

	res, err := s.repository.Create(withdrawModel)

	if err != nil {
		s.logger.Error("failed create withdraw: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *withdrawService) Update(requests requests.UpdateWithdrawRequest) (*models.Withdraw, error) {
	saldo, err := s.saldoRepository.ReadByUserID(requests.UserID)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	if saldo.TotalBalance < requests.WithdrawAmount {
		return nil, fmt.Errorf("sender saldo not enough")
	}

	saldo.TotalBalance = saldo.TotalBalance - requests.WithdrawAmount

	_, err = s.saldoRepository.Update(*saldo)

	if err != nil {
		s.logger.Error("failed update sender saldo: ", zap.Error(err))
		return nil, err
	}

	withdrawModel := models.Withdraw{
		WithdrawID:     requests.WithdrawID,
		WithdrawTime:   requests.WithdrawTime,
		WithdrawAmount: requests.WithdrawAmount,
		UserID:         requests.UserID,
	}

	res, err := s.repository.Update(withdrawModel)

	if err != nil {
		s.logger.Error("failed update withdraw: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *withdrawService) Delete(withdrawID int) error {
	err := s.repository.Delete(withdrawID)

	if err != nil {
		s.logger.Error("failed delete withdraw: ", zap.Error(err))
		return err
	}

	return nil
}
