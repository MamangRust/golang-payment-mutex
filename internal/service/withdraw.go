package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type withdrawService struct {
	userRepository     repository.UserRepository
	saldoRepository    repository.SaldoRepository
	withdrawRepository repository.WithdrawRepository
	logger             logger.Logger
	mapper             responseMapper.WithdrawResponseMapper
}

func NewWithdrawService(
	userRepository repository.UserRepository,
	withdrawRepository repository.WithdrawRepository, saldoRepository repository.SaldoRepository, logger logger.Logger, mapper responseMapper.WithdrawResponseMapper) *withdrawService {
	return &withdrawService{
		userRepository:     userRepository,
		saldoRepository:    saldoRepository,
		withdrawRepository: withdrawRepository,
		logger:             logger,
		mapper:             mapper,
	}
}

func (s *withdrawService) FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.WithdrawResponse], *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	withdraws, totalRecords, err := s.withdrawRepository.ReadAll(page, pageSize, search)

	if err != nil {
		s.logger.Error("failed to fetch withdraws", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraws",
		}
	}

	if len(withdraws) == 0 {
		s.logger.Error("no withdraws found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No withdraws found",
		}
	}

	withdrawResponse := s.mapper.ToWithdrawsResponse(withdraws)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return &response.APIResponsePagination[[]*response.WithdrawResponse]{
		Status:  "success",
		Message: "Successfully retrieved withdraw records.",
		Data:    withdrawResponse,
		Meta: response.PaginationMeta{
			CurrentPage:  page,
			PageSize:     pageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
	}, nil
}

func (s *withdrawService) FindById(withdrawID int) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse) {
	withdraw, err := s.withdrawRepository.Read(withdrawID)
	if err != nil {
		s.logger.Error("failed to find withdraw by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch withdraw record by ID.",
		}
	}
	so := s.mapper.ToWithdrawResponse(*withdraw)

	return &response.ApiResponse[*response.WithdrawResponse]{
		Status:  "success",
		Message: "Successfully retrieved withdraw record by ID.",
		Data:    so,
	}, nil
}

func (s *withdrawService) Create(request requests.CreateWithdrawRequest) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse) {
	saldo, err := s.saldoRepository.ReadByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to find saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo == nil {
		s.logger.Error("Saldo not found for user", zap.String("cardNumber", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Saldo not found for the specified user ID.",
		}
	}

	// Periksa saldo mencukupi
	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.String("cardNumber", request.CardNumber), zap.Int("requested", request.WithdrawAmount))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal.",
		}
	}

	// Update saldo setelah penarikan
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateData := requests.UpdateSaldoWithdraw{
		CardNumber:     request.CardNumber,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateData)
	if err != nil {
		s.logger.Error("Failed to update saldo after withdrawal", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo after withdrawal.",
		}
	}

	// Buat catatan withdraw
	withdrawRecord, err := s.withdrawRepository.Create(request)
	if err != nil {
		s.logger.Error("Failed to create withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create withdraw record.",
		}
	}

	so := s.mapper.ToWithdrawResponse(*withdrawRecord)

	return &response.ApiResponse[*response.WithdrawResponse]{
		Status:  "success",
		Message: "Withdrawal created successfully.",
		Data:    so,
	}, nil
}

func (s *withdrawService) Update(request requests.UpdateWithdrawRequest) (*response.ApiResponse[*response.WithdrawResponse], *response.ErrorResponse) {
	_, err := s.withdrawRepository.Read(request.WithdrawID)
	if err != nil {
		s.logger.Error("Failed to find withdraw record by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Withdraw record not found.",
		}
	}

	// Ambil saldo pengguna
	saldo, err := s.saldoRepository.ReadByCardNumber(request.CardNumber)
	if err != nil {
		s.logger.Error("Failed to fetch saldo by user ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch saldo for the user.",
		}
	}

	if saldo.TotalBalance < request.WithdrawAmount {
		s.logger.Error("Insufficient balance for user", zap.String("cardNumber", request.CardNumber))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Insufficient balance for withdrawal update.",
		}
	}

	// Update saldo baru
	newTotalBalance := saldo.TotalBalance - request.WithdrawAmount
	updateSaldoData := requests.UpdateSaldoWithdraw{
		CardNumber:     saldo.CardNumber,
		TotalBalance:   newTotalBalance,
		WithdrawAmount: &request.WithdrawAmount,
		WithdrawTime:   &request.WithdrawTime,
	}

	_, err = s.saldoRepository.UpdateSaldoWithdraw(updateSaldoData)
	if err != nil {
		s.logger.Error("Failed to update saldo balance", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update saldo balance.",
		}
	}

	updatedWithdraw, err := s.withdrawRepository.Update(request)
	if err != nil {
		rollbackData := requests.UpdateSaldoBalance{
			CardNumber:   saldo.CardNumber,
			TotalBalance: saldo.TotalBalance,
		}
		_, rollbackErr := s.saldoRepository.UpdateBalance(rollbackData)
		if rollbackErr != nil {
			s.logger.Error("Failed to rollback saldo after withdraw update failure", zap.Error(rollbackErr))
		}
		s.logger.Error("Failed to update withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update withdraw record.",
		}
	}

	so := s.mapper.ToWithdrawResponse(*updatedWithdraw)

	return &response.ApiResponse[*response.WithdrawResponse]{
		Status:  "success",
		Message: "Withdraw record updated successfully.",
		Data:    so,
	}, nil
}

func (s *withdrawService) Delete(withdrawID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.withdrawRepository.Delete(withdrawID)
	if err != nil {
		s.logger.Error("Failed to delete withdraw record", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete withdraw record.",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Withdraw record deleted successfully.",
		Data:    "Record deleted",
	}, nil
}
