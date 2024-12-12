package service

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"
	"strconv"

	"go.uber.org/zap"
)

type merchantService struct {
	merchantRepository repository.MerchantRepository
	logger             logger.Logger
	mapper             responseMapper.MerchantResponseMapper
}

func NewMerchantService(
	merchantRepository repository.MerchantRepository,
	logger logger.Logger,
	mapper responseMapper.MerchantResponseMapper,
) *merchantService {
	return &merchantService{
		merchantRepository: merchantRepository,
		logger:             logger,
		mapper:             mapper,
	}
}

func (s *merchantService) FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.MerchantResponse], *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.ReadAll(page, pageSize, search)

	if err != nil {
		s.logger.Error("failed to fetch merchants", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch merchants",
		}
	}

	if len(merchants) == 0 {
		s.logger.Error("no merchants found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No merchants found",
		}
	}

	so := s.mapper.ToMerchantsResponse(merchants)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return &response.APIResponsePagination[[]*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchants retrieved successfully",
		Data:    so,
		Meta: response.PaginationMeta{
			CurrentPage:  page,
			PageSize:     pageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
	}, nil

}

func (s *merchantService) FindByID(id int) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse) {
	merchant, err := s.merchantRepository.Read(id)
	if err != nil {
		s.logger.Error("failed to find merchant by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	so := s.mapper.ToMerchantResponse(*merchant)

	return &response.ApiResponse[*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchant retrieved successfully",
		Data:    so,
	}, nil
}

func (s *merchantService) FindByApiKey(apiKey string) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse) {
	merchant, err := s.merchantRepository.ReadByApiKey(apiKey)

	if err != nil {
		s.logger.Error("failed to find merchant by api key", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	so := s.mapper.ToMerchantResponse(*merchant)

	return &response.ApiResponse[*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchant retrieved successfully",
		Data:    so,
	}, nil
}

func (s *merchantService) FindByName(name string) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse) {
	merchant, err := s.merchantRepository.ReadByName(name)
	if err != nil {
		s.logger.Error("failed to find merchant by name", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Merchant not found",
		}
	}

	so := s.mapper.ToMerchantResponse(*merchant)

	return &response.ApiResponse[*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchant retrieved successfully",
		Data:    so,
	}, nil
}

func (s *merchantService) Create(request requests.CreateMerchantRequest) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse) {
	merchant, err := s.merchantRepository.Create(request)

	if err != nil {
		fmt.Println(err.Error())

		s.logger.Error("failed to create merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create merchant",
		}

	}

	so := s.mapper.ToMerchantResponse(*merchant)

	return &response.ApiResponse[*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchant created successfully",
		Data:    so,
	}, nil
}

func (s *merchantService) Update(request requests.UpdateMerchantRequest) (*response.ApiResponse[*response.MerchantResponse], *response.ErrorResponse) {
	merchant, err := s.merchantRepository.Update(request)

	if err != nil {
		s.logger.Error("failed to update merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant",
		}

	}

	so := s.mapper.ToMerchantResponse(*merchant)

	return &response.ApiResponse[*response.MerchantResponse]{
		Status:  "success",
		Message: "Merchant updated successfully",
		Data:    so,
	}, nil
}

func (s *merchantService) Delete(id int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.merchantRepository.Delete(id)

	if err != nil {
		s.logger.Error("failed to delete merchant", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete merchant",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Merchant deleted successfully",
		Data:    "Merchant with ID " + strconv.Itoa(id) + " has been deleted",
	}, nil
}
