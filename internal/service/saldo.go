package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"strconv"

	"go.uber.org/zap"
)

type saldoService struct {
	cardRepository  repository.CardRepository
	saldoRepository repository.SaldoRepository
	logger          logger.Logger
	mapper          responseMapper.SaldoResponseMapper
}

func NewSaldoService(
	cardRepository repository.CardRepository,
	saldoRepository repository.SaldoRepository,
	logger logger.Logger,
	mapper responseMapper.SaldoResponseMapper,
) *saldoService {
	return &saldoService{
		cardRepository:  cardRepository,
		saldoRepository: saldoRepository,
		logger:          logger,
		mapper:          mapper,
	}
}

func (s *saldoService) FindAll() (*response.ApiResponse[[]*response.SaldoResponse], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.ReadAll()
	if err != nil {
		s.logger.Error("failed find all saldo", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch all saldo records",
		}
	}

	so := s.mapper.ToSaldoResponses(saldo)

	return &response.ApiResponse[[]*response.SaldoResponse]{
		Status:  "success",
		Message: "Successfully fetched all saldo records",
		Data:    so,
	}, nil
}

func (s *saldoService) FindById(saldoID int) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.Read(saldoID)
	if err != nil {
		s.logger.Error("failed find saldo by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo record not found",
		}
	}

	so := s.mapper.ToSaldoResponse(*saldo)

	return &response.ApiResponse[*response.SaldoResponse]{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    so,
	}, nil
}

func (s *saldoService) Create(requests requests.CreateSaldoRequest) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse) {
	_, err := s.cardRepository.ReadByCardNumber(requests.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
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

	so := s.mapper.ToSaldoResponse(*saldo)

	return &response.ApiResponse[*response.SaldoResponse]{
		Status:  "success",
		Message: "Saldo record created successfully",
		Data:    so,
	}, nil
}

func (s *saldoService) Update(requests requests.UpdateSaldoRequest) (*response.ApiResponse[*response.SaldoResponse], *response.ErrorResponse) {
	_, err := s.cardRepository.ReadByCardNumber(requests.CardNumber)
	if err != nil {
		s.logger.Error("failed to find card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
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

	so := s.mapper.ToSaldoResponse(*saldo)

	return &response.ApiResponse[*response.SaldoResponse]{
		Status:  "success",
		Message: "Saldo record updated successfully",
		Data:    so,
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
