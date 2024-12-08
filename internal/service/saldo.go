package service

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"strconv"

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

func (s *saldoService) FindAll() (*response.ApiResponse[[]*record.SaldoRecord], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed find all saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch all saldo records",
		}
	}

	return &response.ApiResponse[[]*record.SaldoRecord]{
		Status:  "success",
		Message: "Successfully fetched all saldo records",
		Data:    saldo,
	}, nil
}

func (s *saldoService) FindById(saldoID int) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.Read(saldoID)
	if err != nil {
		s.logger.Error("failed find saldo by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo record not found",
		}
	}

	return &response.ApiResponse[*record.SaldoRecord]{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    saldo,
	}, nil
}

func (s *saldoService) FindByUserID(userID int) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.ReadByUserID(userID)
	if err != nil {
		s.logger.Error("Failed to find saldo record by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve saldo for the specified user.",
		}
	}

	if saldo == nil {
		s.logger.Error("No saldo record found for user ID", zap.Int("userID", userID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No saldo record found for the specified user ID.",
		}
	}

	return &response.ApiResponse[*record.SaldoRecord]{
		Status:  "success",
		Message: "Successfully retrieved saldo for the user.",
		Data:    saldo,
	}, nil
}

func (s *saldoService) FindByUsersID(userID int) (*response.ApiResponse[[]*record.SaldoRecord], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.ReadByUsersID(userID)
	if err != nil {
		s.logger.Error("failed find saldo by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No saldo records found for the given user",
		}
	}

	return &response.ApiResponse[[]*record.SaldoRecord]{
		Status:  "success",
		Message: "Successfully fetched saldo records for user",
		Data:    saldo,
	}, nil
}

func (s *saldoService) Create(requests requests.CreateSaldoRequest) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(requests.UserID)
	if err != nil {
		s.logger.Error("failed to find user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	saldo, err := s.saldoRepository.Create(requests)
	if err != nil {
		s.logger.Error("failed to create saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create saldo record",
		}
	}

	return &response.ApiResponse[*record.SaldoRecord]{
		Status:  "success",
		Message: "Saldo record created successfully",
		Data:    saldo,
	}, nil
}

func (s *saldoService) Update(requests requests.UpdateSaldoRequest) (*response.ApiResponse[*record.SaldoRecord], *response.ErrorResponse) {
	_, err := s.userRepository.Read(requests.UserID)
	if err != nil {
		s.logger.Error("failed to find user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	saldo, err := s.saldoRepository.Update(requests)
	if err != nil {
		s.logger.Error("failed to update saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo record",
		}
	}

	return &response.ApiResponse[*record.SaldoRecord]{
		Status:  "success",
		Message: "Saldo record updated successfully",
		Data:    saldo,
	}, nil
}

func (s *saldoService) Delete(saldoID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.saldoRepository.Delete(saldoID)
	if err != nil {
		s.logger.Error("failed to delete saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete saldo record",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Saldo record deleted successfully",
		Data:    "Saldo record with ID " + strconv.Itoa(saldoID) + " has been deleted",
	}, nil
}
