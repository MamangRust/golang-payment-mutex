package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type transferService struct {
	userRepository     repository.UserRepository
	cardRepository     repository.CardRepository
	saldoRepository    repository.SaldoRepository
	transferRepository repository.TransferRepository
	logger             logger.Logger
	mapper             responseMapper.TransferResponseMapper
}

func NewTransferService(
	userRepository repository.UserRepository,
	cardRepository repository.CardRepository,
	transferRepository repository.TransferRepository,
	saldoRepository repository.SaldoRepository, logger logger.Logger, mapper responseMapper.TransferResponseMapper) *transferService {
	return &transferService{
		userRepository:     userRepository,
		transferRepository: transferRepository,
		saldoRepository:    saldoRepository,
		logger:             logger,
		mapper:             mapper,
	}
}

func (s *transferService) FindAll() (*response.ApiResponse[[]*response.TransferResponse], *response.ErrorResponse) {
	transfer, err := s.transferRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find all transfers", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfers",
		}
	}

	so := s.mapper.ToTransfersResponse(transfer)

	return &response.ApiResponse[[]*response.TransferResponse]{
		Status:  "success",
		Message: "Transfers retrieved successfully",
		Data:    so,
	}, nil
}

func (s *transferService) FindById(transferID int) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse) {
	transfer, err := s.transferRepository.Read(transferID)
	if err != nil {
		s.logger.Error("failed to find transfer by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transfer not found",
		}
	}

	so := s.mapper.ToTransferResponse(*transfer)

	return &response.ApiResponse[*response.TransferResponse]{
		Status:  "success",
		Message: "Transfer retrieved successfully",
		Data:    so,
	}, nil
}

func (s *transferService) Create(request requests.CreateTransferRequest) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse) {
	_, err := s.cardRepository.ReadByCardNumber(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender card by Number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Sender card not found",
		}
	}

	_, err = s.cardRepository.ReadByCardNumber(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver card by number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Receiver card not found",
		}
	}

	senderSaldo, err := s.saldoRepository.ReadByCardNumber(request.TransferFrom)
	if err != nil {
		s.logger.Error("failed to find sender saldo by card number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find sender saldo",
		}
	}

	receiverSaldo, err := s.saldoRepository.ReadByCardNumber(request.TransferTo)
	if err != nil {
		s.logger.Error("failed to find receiver saldo by card number", zap.Error(err))
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
		CardNumber:   senderSaldo.CardNumber,
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
		CardNumber:   receiverSaldo.CardNumber,
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

	so := s.mapper.ToTransferResponse(*transfer)

	return &response.ApiResponse[*response.TransferResponse]{
		Status:  "success",
		Message: "Transfer created successfully",
		Data:    so,
	}, nil
}

func (s *transferService) Update(request requests.UpdateTransferRequest) (*response.ApiResponse[*response.TransferResponse], *response.ErrorResponse) {
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
	senderSaldo, err := s.saldoRepository.ReadByCardNumber(transfer.TransferFrom)
	if err != nil {
		s.logger.Error("Failed to find sender's saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to find sender's saldo: %v", err),
		}
	}

	newSenderBalance := senderSaldo.TotalBalance - amountDifference
	if newSenderBalance < 0 {
		s.logger.Error("Insufficient balance for sender", zap.String("senderID", transfer.TransferFrom))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for sender",
		}
	}

	senderSaldo.TotalBalance = newSenderBalance
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		CardNumber:   senderSaldo.CardNumber,
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
	receiverSaldo, err := s.saldoRepository.ReadByCardNumber(transfer.TransferTo)
	if err != nil {
		s.logger.Error("Failed to find receiver's saldo by user ID", zap.Error(err))

		// Rollback the sender's saldo if the receiver's saldo update fails
		rollbackSenderBalance := requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferFrom,
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
		CardNumber:   receiverSaldo.CardNumber,
		TotalBalance: receiverSaldo.TotalBalance,
	})
	if err != nil {
		s.logger.Error("Failed to update receiver's saldo", zap.Error(err))

		// Rollback both sender's and receiver's saldos if the update fails
		rollbackSenderBalance := requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferTo,
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
			CardNumber:   transfer.TransferFrom,
			TotalBalance: senderSaldo.TotalBalance + amountDifference,
		}
		rollbackReceiverBalance := requests.UpdateSaldoBalance{
			CardNumber:   transfer.TransferTo,
			TotalBalance: receiverSaldo.TotalBalance - amountDifference,
		}

		s.saldoRepository.UpdateBalance(rollbackSenderBalance)
		s.saldoRepository.UpdateBalance(rollbackReceiverBalance)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update transfer: %v", err),
		}
	}

	so := s.mapper.ToTransferResponse(*updatedTransfer)

	// Return the updated transfer in a successful response
	return &response.ApiResponse[*response.TransferResponse]{
		Status:  "success",
		Message: "Transfer successfully updated.",
		Data:    so,
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
