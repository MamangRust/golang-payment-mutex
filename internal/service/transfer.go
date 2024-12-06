package service

import (
	"errors"
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type transferService struct {
	userRepository     repository.UserRepository
	saldoRepository    repository.SaldoRepository
	transferRepository repository.TransferRepository
	logger             logger.Logger
}

func NewTransferService(
	userRepository repository.UserRepository,
	transferRepository repository.TransferRepository, saldoRepository repository.SaldoRepository, logger logger.Logger) *transferService {
	return &transferService{
		transferRepository: transferRepository,
		saldoRepository:    saldoRepository,
		logger:             logger,
	}
}

func (s *transferService) FindAll() (*[]models.Transfer, error) {
	transfer, err := s.transferRepository.ReadAll()

	if err != nil {
		s.logger.Error("failed find all transfer: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil
}

func (s *transferService) FindByUserID(userID int) (*models.Transfer, error) {
	_, err := s.userRepository.Read(userID)

	if err != nil {
		s.logger.Error("failed find user by user id: ", zap.Error(err))
	}

	transfer, err := s.transferRepository.ReadByUserID(userID)

	if err != nil {
		s.logger.Error("failed find transfer by user id: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil
}

func (s *transferService) FindByUsersID(userID int) (*[]models.Transfer, error) {
	_, err := s.userRepository.Read(userID)

	if err != nil {
		s.logger.Error("failed find user by user id: ", zap.Error(err))
	}

	transfer, err := s.transferRepository.ReadByUsersID(userID)

	if err != nil {
		s.logger.Error("failed find transers by user id: ", zap.Error(err))

		return nil, err
	}

	return transfer, nil

}

func (s *transferService) FindById(transferID int) (*models.Transfer, error) {
	transfer, err := s.transferRepository.Read(transferID)

	if err != nil {
		s.logger.Error("failed find transfer by id: ", zap.Error(err))
		return nil, err
	}

	return transfer, nil

}

func (s *transferService) Create(request requests.CreateTransferRequest) (*models.Transfer, error) {
	_, err := s.userRepository.Read(request.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender user by ID", zap.Error(err))
		return nil, errors.New("sender user not found")
	}

	_, err = s.userRepository.Read(request.TransferTo)
	if err != nil {
		s.logger.Error("Failed to find receiver user by ID", zap.Error(err))
		return nil, errors.New("receiver user not found")
	}

	// Retrieve sender and receiver saldo
	senderSaldo, err := s.saldoRepository.ReadByUserID(request.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender saldo by user ID", zap.Error(err))
		return nil, fmt.Errorf("failed to find sender saldo: %w", err)
	}

	receiverSaldo, err := s.saldoRepository.ReadByUserID(request.TransferTo)
	if err != nil {
		s.logger.Error("Failed to find receiver saldo by user ID", zap.Error(err))
		return nil, fmt.Errorf("failed to find receiver saldo: %w", err)
	}

	// Check if the sender has sufficient balance
	if senderSaldo.TotalBalance < request.TransferAmount {
		return nil, errors.New("insufficient balance for sender")
	}

	// Adjust sender's and receiver's balances
	senderSaldo.TotalBalance -= request.TransferAmount
	receiverSaldo.TotalBalance += request.TransferAmount

	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       senderSaldo.UserID,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update sender saldo", zap.Error(err))
		return nil, fmt.Errorf("failed to update sender saldo: %w", err)
	}

	// Update receiver's saldo
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       receiverSaldo.UserID,
		TotalBalance: receiverSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update receiver saldo", zap.Error(err))

		// Rollback sender's saldo if receiver update fails
		senderSaldo.TotalBalance += request.TransferAmount // Revert the sender's balance change
		rollbackErr, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			UserID:       senderSaldo.UserID,
			TotalBalance: senderSaldo.TotalBalance,
		})
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback sender saldo update", zap.Error(err))
		}

		return nil, fmt.Errorf("failed to update receiver saldo: %w", err)
	}

	res, err := s.transferRepository.Create(request)
	if err != nil {
		s.logger.Error("Failed to create transfer", zap.Error(err))

		// Rollback saldo changes if transfer creation fails
		senderSaldo.TotalBalance += request.TransferAmount
		receiverSaldo.TotalBalance -= request.TransferAmount

		rollbackErr, err := s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			UserID:       senderSaldo.UserID,
			TotalBalance: senderSaldo.TotalBalance,
		})
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback sender saldo update after transfer creation failure", zap.Error(err))
		}

		rollbackErr, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
			UserID:       receiverSaldo.UserID,
			TotalBalance: receiverSaldo.TotalBalance,
		})
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback receiver saldo update after transfer creation failure", zap.Error(err))
		}

		return nil, fmt.Errorf("failed to create transfer: %w", err)
	}

	return res, nil
}

func (s *transferService) Update(request requests.UpdateTransferRequest) (*models.Transfer, error) {
	// Retrieve the existing transfer
	transfer, err := s.transferRepository.Read(request.TransferID)
	if err != nil {
		s.logger.Error("Failed to find transfer by ID", zap.Error(err))
		return nil, fmt.Errorf("transfer with ID %d not found: %w", request.TransferID, err)
	}

	// Calculate the amount difference for the update
	amountDifference := request.TransferAmount - transfer.TransferAmount

	// Update sender's saldo
	senderSaldo, err := s.saldoRepository.ReadByUserID(transfer.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender's saldo by user ID", zap.Error(err))
		return nil, fmt.Errorf("failed to find sender's saldo: %w", err)
	}

	newSenderBalance := senderSaldo.TotalBalance - amountDifference
	if newSenderBalance < 0 {
		return nil, fmt.Errorf("insufficient balance for sender")
	}

	senderSaldo.TotalBalance = newSenderBalance
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       senderSaldo.UserID,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update sender's saldo", zap.Error(err))
		return nil, fmt.Errorf("failed to update sender's saldo: %w", err)
	}

	// Update receiver's saldo
	receiverSaldo, err := s.saldoRepository.ReadByUserID(transfer.TransferTo)
	if err != nil {
		s.logger.Error("Failed to find receiver's saldo by user ID", zap.Error(err))

		// Rollback the sender's saldo if the receiver's saldo update fails
		rollbackSenderBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance,
		}
		rollbackErr, err := s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		if rollbackErr != nil {

			return nil, err
		}

		return nil, fmt.Errorf("failed to find receiver's saldo: %w", err)
	}

	newReceiverBalance := receiverSaldo.TotalBalance + amountDifference
	receiverSaldo.TotalBalance = newReceiverBalance

	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       receiverSaldo.UserID,
		TotalBalance: receiverSaldo.TotalBalance,
	})

	if err != nil {
		s.logger.Error("Failed to update receiver's saldo", zap.Error(err))

		// Rollback both sender's and receiver's saldos if the update fails
		rollbackSenderBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance,
		}
		rollbackReceiverBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		s.saldoRepository.UpdateBalance(rollbackReceiverBalance)

		return nil, fmt.Errorf("failed to update receiver's saldo: %w", err)
	}

	updatedTransfer, err := s.transferRepository.Update(request)
	if err != nil {
		s.logger.Error("Failed to update transfer", zap.Error(err))

		// Rollback saldos if the transfer update fails
		rollbackSenderBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		s.saldoRepository.UpdateBalance(rollbackReceiverBalance)

		return nil, fmt.Errorf("failed to update transfer: %w", err)
	}

	return updatedTransfer, nil
}

func (s *transferService) Delete(transferID int) error {
	err := s.transferRepository.Delete(transferID)

	if err != nil {
		s.logger.Error("failed delete transfer: ", zap.Error(err))
		return err
	}

	return nil
}
