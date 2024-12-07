package service

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"strconv"

	"go.uber.org/zap"
)

type topupService struct {
	userRepository  repository.UserRepository
	topupRepository repository.TopupRepository
	saldoRepository repository.SaldoRepository
	logger          logger.Logger
}

func NewTopupService(
	userRepository repository.UserRepository,
	topupRepository repository.TopupRepository,
	saldoRepository repository.SaldoRepository,
	logger logger.Logger) *topupService {
	return &topupService{
		userRepository:  userRepository,
		topupRepository: topupRepository,
		saldoRepository: saldoRepository,
		logger:          logger,
	}
}

func (s *topupService) FindAll() (*response.ApiResponse[[]*record.TopupRecord], *response.ErrorResponse) {
	topup, err := s.topupRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed to find all topups", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch all topup records",
		}
	}

	return &response.ApiResponse[[]*record.TopupRecord]{
		Status:  "success",
		Message: "Successfully fetched all topup records",
		Data:    topup,
	}, nil
}

func (s *topupService) FindById(topupID int) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse) {
	topup, err := s.topupRepository.Read(topupID)
	if err != nil {
		s.logger.Error("failed to find topup by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Topup record not found",
		}
	}

	return &response.ApiResponse[*record.TopupRecord]{
		Status:  "success",
		Message: "Successfully fetched topup record",
		Data:    topup,
	}, nil
}

func (s *topupService) FindByUserID(userID int) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse) {
	res, err := s.topupRepository.ReadByUserID(userID)
	if err != nil {
		s.logger.Error("Failed to find top-up record by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve top-up record for the specified user.",
		}
	}

	return &response.ApiResponse[*record.TopupRecord]{
		Status:  "success",
		Message: "Successfully retrieved top-up record for the user.",
		Data:    res,
	}, nil
}

func (s *topupService) FindByUsersID(userID int) (*response.ApiResponse[[]*record.TopupRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(userID)
	if err != nil {
		s.logger.Error("failed to find user by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	topup, err := s.topupRepository.ReadByUsersID(userID)
	if err != nil {
		s.logger.Error("failed to find topup by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No topup records found for the given user",
		}
	}

	return &response.ApiResponse[[]*record.TopupRecord]{
		Status:  "success",
		Message: "Successfully fetched topup records for user",
		Data:    topup,
	}, nil
}

func (s *topupService) Create(request requests.CreateTopupRequest) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("failed to find user by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	// Create topup
	topup, err := s.topupRepository.Create(request)
	if err != nil {
		s.logger.Error("failed to create topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create topup record",
		}
	}

	// Find current saldo
	saldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("failed to find saldo by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch user's saldo",
		}
	}

	newBalance := saldo.TotalBalance + request.TopupAmount
	_, err = s.saldoRepository.UpdateBalance(requests.UpdateSaldoBalance{
		UserID:       request.UserID,
		TotalBalance: newBalance,
	})
	if err != nil {
		s.logger.Error("failed to update saldo balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo balance",
		}
	}

	return &response.ApiResponse[*record.TopupRecord]{
		Status:  "success",
		Message: "Topup record created successfully",
		Data:    topup,
	}, nil
}

func (s *topupService) Update(request requests.UpdateTopupRequest) (*response.ApiResponse[*record.TopupRecord], *response.ErrorResponse) {
	// Check if the user exists
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	// Find the existing topup
	existingTopup, err := s.topupRepository.Read(request.TopupID)
	if err != nil || existingTopup == nil {
		s.logger.Error("Failed to find topup by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Topup not found",
		}
	}

	topupDifference := request.TopupAmount - existingTopup.TopupAmount

	// Update the topup amount
	_, err = s.topupRepository.UpdateAmount(requests.UpdateTopupAmount{
		TopupID:     request.TopupID,
		TopupAmount: request.TopupAmount,
	})
	if err != nil {
		s.logger.Error("Failed to update topup amount", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update topup amount: %v", err),
		}
	}

	// Retrieve the current balance from saldo
	currentSaldo, err := s.saldoRepository.ReadByUserID(request.UserID)
	if err != nil {
		s.logger.Error("Failed to retrieve current saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve current saldo: %v", err),
		}
	}

	if currentSaldo == nil {
		s.logger.Error("No saldo found for user", zap.Int("userID", request.UserID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to update saldo balance: %v", err),
		}
	}

	// Retrieve and return the updated topup
	updatedTopup, err := s.topupRepository.Read(request.TopupID)
	if err != nil || updatedTopup == nil {
		s.logger.Error("Failed to find updated topup by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Updated topup not found",
		}
	}

	return &response.ApiResponse[*record.TopupRecord]{
		Status:  "success",
		Message: "Topup successfully updated",
		Data:    updatedTopup,
	}, nil
}

func (s *topupService) Delete(topupID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.topupRepository.Delete(topupID)
	if err != nil {
		s.logger.Error("failed to delete topup", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete topup record",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Topup record deleted successfully",
		Data:    "Topup record with ID " + strconv.Itoa(topupID) + " has been deleted",
	}, nil
}
