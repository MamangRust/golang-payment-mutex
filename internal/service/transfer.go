package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type transferService struct {
	saldoRepository repository.SaldoRepository
	repository      repository.TransferRepository
	logger          logger.Logger
}

func NewTransferService(repository repository.TransferRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *transferService {
	return &transferService{
		repository:      repository,
		saldoRepository: saldoRepository,
		logger:          logger,
	}
}

func (s *transferService) FindAll() (*[]models.Transfer, error) {
	transfer, err := s.repository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all transfer: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil
}

func (s *transferService) FindByUserID(userID int) (*models.Transfer, error) {
	transfer, err := s.repository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find transfer by user id: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil
}

func (s *transferService) FindById(transferID int) (*models.Transfer, error) {
	transfer, err := s.repository.Read(transferID)

	if err != nil {
		s.logger.Error("failed find transfer by id: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil

}

func (s *transferService) Create(request requests.CreateTransferRequest) (*models.Transfer, error) {
	senderSaldo, err := s.saldoRepository.ReadByUserID(request.TransferFrom)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	receiverSaldo, err := s.saldoRepository.ReadByUserID(request.TransferTo)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	if senderSaldo.TotalBalance < request.TransferAmount {
		return nil, fmt.Errorf("sender saldo not enough")
	}

	senderSaldo.TotalBalance = senderSaldo.TotalBalance - request.TransferAmount
	receiverSaldo.TotalBalance = receiverSaldo.TotalBalance + request.TransferAmount

	_, err = s.saldoRepository.Update(*senderSaldo)

	if err != nil {
		s.logger.Error("failed update sender saldo: ", zap.Error(err))
		return nil, err
	}

	transferModel := models.Transfer{
		TransferFrom:   senderSaldo.UserID,
		TransferTo:     receiverSaldo.UserID,
		TransferAmount: request.TransferAmount,
	}

	res, err := s.repository.Create(transferModel)

	if err != nil {
		s.logger.Error("failed create transfer: ", zap.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *transferService) Update(requests requests.UpdateTransferRequest) (*models.Transfer, error) {
	senderSaldo, err := s.saldoRepository.ReadByUserID(requests.TransferFrom)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	receiverSaldo, err := s.saldoRepository.ReadByUserID(requests.TransferTo)

	if err != nil {
		s.logger.Error("failed find saldo by user id: ", zap.Error(err))
		return nil, err
	}

	if senderSaldo.TotalBalance < requests.TransferAmount {
		return nil, fmt.Errorf("sender saldo not enough")
	}

	senderSaldo.TotalBalance = senderSaldo.TotalBalance - requests.TransferAmount
	receiverSaldo.TotalBalance = receiverSaldo.TotalBalance + requests.TransferAmount

	_, err = s.saldoRepository.Update(*senderSaldo)

	if err != nil {
		s.logger.Error("failed update sender saldo: ", zap.Error(err))
		return nil, err
	}

	transferModel := models.Transfer{
		TransferID:     requests.TransferID,
		TransferFrom:   requests.TransferFrom,
		TransferTo:     requests.TransferTo,
		TransferAmount: requests.TransferAmount,
	}

	res, err := s.repository.Update(transferModel)

	if err != nil {
		s.logger.Error("failed update transfer: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *transferService) Delete(transferID int) error {
	err := s.repository.Delete(transferID)

	if err != nil {
		s.logger.Error("failed delete transfer: ", zap.Error(err))
		return err
	}

	return nil
}
