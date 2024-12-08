package service

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
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

func (s *transferService) FindAll() (*response.ApiResponse[[]*record.TransferRecord], *response.ErrorResponse) {
	transfer, err := s.transferRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find all transfers", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfers",
		}
	}

	return &response.ApiResponse[[]*record.TransferRecord]{
		Status:  "success",
		Message: "Transfers retrieved successfully",
		Data:    transfer,
	}, nil
}

func (s *transferService) FindById(transferID int) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse) {
	transfer, err := s.transferRepository.Read(transferID)
	if err != nil {
		s.logger.Error("failed to find transfer by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transfer not found",
		}
	}

	return &response.ApiResponse[*record.TransferRecord]{
		Status:  "success",
		Message: "Transfer retrieved successfully",
		Data:    transfer,
	}, nil
}

func (s *transferService) FindByUserID(userID int) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(userID)
	if err != nil {
		s.logger.Error("failed to find user by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	transfer, err := s.transferRepository.ReadByUserID(userID)
	if err != nil {
		s.logger.Error("failed to find transfer by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer",
		}
	}

	return &response.ApiResponse[*record.TransferRecord]{
		Status:  "success",
		Message: "Transfer retrieved successfully",
		Data:    transfer,
	}, nil
}

func (s *transferService) FindByUsersID(userID int) (*response.ApiResponse[[]*record.TransferRecord], *response.ErrorResponse) {
	transfers, err := s.transferRepository.ReadByUsersID(userID)
	if err != nil {
		s.logger.Error("Failed to find transfer records by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer records for the specified user.",
		}
	}

	if transfers == nil || len(transfers) == 0 {
		s.logger.Error("No transfer records found for user ID", zap.Int("userID", userID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No transfer records found for the specified user ID.",
		}
	}

	return &response.ApiResponse[[]*record.TransferRecord]{
		Status:  "success",
		Message: "Successfully retrieved transfer records for the user.",
		Data:    transfers,
	}, nil
}

func (s *transferService) Create(request requests.CreateTransferRequest) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Sender user not found",
		}
	}

	_, err = s.userRepository.Read(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Receiver user not found",
		}
	}

	senderSaldo, err := s.saldoRepository.ReadByUserID(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find sender saldo",
		}
	}

	receiverSaldo, err := s.saldoRepository.ReadByUserID(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find receiver saldo",
		}
	}

	if senderSaldo.TotalBalance < request.TransferAmount {
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for sender",
		}
	}

	senderSaldo.TotalBalance -= request.TransferAmount
	receiverSaldo.TotalBalance += request.TransferAmount

	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       senderSaldo.UserID,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("failed to update sender saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update sender saldo",
		}
	}

	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       receiverSaldo.UserID,
		TotalBalance: receiverSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("failed to update receiver saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update receiver saldo",
		}
	}

	transfer, err := s.transferRepository.Create(request)
	if err != nil {
		s.logger.Error("failed to create transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transfer",
		}
	}

	return &response.ApiResponse[*record.TransferRecord]{
		Status:  "success",
		Message: "Transfer created successfully",
		Data:    transfer,
	}, nil
}

func (s *transferService) Update(request requests.UpdateTransferRequest) (*response.ApiResponse[*record.TransferRecord], *response.ErrorResponse) {
	// Retrieve the existing transfer
	transfer, err := s.transferRepository.Read(request.TransferID)
	if err != nil {
		s.logger.Error("Failed to find transfer by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Transfer with ID %d not found: %v", request.TransferID, err),
		}
	}

	// Calculate the amount difference for the update
	amountDifference := request.TransferAmount - transfer.TransferAmount

	// Update sender's saldo
	senderSaldo, err := s.saldoRepository.ReadByUserID(transfer.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender's saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to find sender's saldo: %v", err),
		}
	}

	newSenderBalance := senderSaldo.TotalBalance - amountDifference
	if newSenderBalance < 0 {
		s.logger.Error("Insufficient balance for sender", zap.Int("senderID", transfer.TransferFrom))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for sender",
		}
	}

	senderSaldo.TotalBalance = newSenderBalance
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       senderSaldo.UserID,
		TotalBalance: senderSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update sender's saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update sender's saldo: %v", err),
		}
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
		_, rollbackErr := s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback sender's saldo after receiver lookup failure", zap.Error(rollbackErr))
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to find receiver's saldo: %v", err),
		}
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
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := requests.UpdateSaldoBalance{
			UserID:       transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		s.saldoRepository.UpdateBalance(rollbackReceiverBalance)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update receiver's saldo: %v", err),
		}
	}

	// Update transfer record
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update transfer: %v", err),
		}
	}

	// Return the updated transfer in a successful response
	return &response.ApiResponse[*record.TransferRecord]{
		Status:  "success",
		Message: "Transfer successfully updated.",
		Data:    updatedTransfer,
	}, nil
}

func (s *transferService) Delete(transferID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.transferRepository.Delete(transferID)
	if err != nil {
		s.logger.Error("failed to delete transfer", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transfer",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Transfer deleted successfully",
		Data:    fmt.Sprintf("Transfer with ID %d has been deleted", transferID),
	}, nil
}
