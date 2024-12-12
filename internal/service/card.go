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

type cardService struct {
	cardRepository  repository.CardRepository
	userRepository  repository.UserRepository
	saldoRepostiroy repository.SaldoRepository
	logger          logger.Logger
	mapper          responseMapper.CardResponseMapper
}

func NewCardService(
	cardRepository repository.CardRepository,
	userRepository repository.UserRepository,
	saldoRepostiroy repository.SaldoRepository,
	logger logger.Logger,
	mapper responseMapper.CardResponseMapper,

) *cardService {
	return &cardService{
		cardRepository:  cardRepository,
		userRepository:  userRepository,
		saldoRepostiroy: saldoRepostiroy,
		logger:          logger,
		mapper:          mapper,
	}
}

func (s *cardService) FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.CardResponse], *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cards, totalRecords, err := s.cardRepository.ReadAll(page, pageSize, search)

	if err != nil {
		s.logger.Error("failed find all card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch all card records",
		}
	}

	if len(cards) == 0 {
		s.logger.Error("no cards found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No cards found",
		}
	}
	so := s.mapper.ToCardsResponse(cards)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return &response.APIResponsePagination[[]*response.CardResponse]{
		Status:  "success",
		Message: "Successfully fetched all card records",
		Data:    so,
		Meta: response.PaginationMeta{
			CurrentPage:  page,
			PageSize:     pageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
	}, nil
}

func (s *cardService) FindById(cardID int) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse) {
	card, err := s.cardRepository.Read(cardID)
	if err != nil {
		s.logger.Error("failed find card by id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}
	so := s.mapper.ToCardResponse(*card)

	return &response.ApiResponse[*response.CardResponse]{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardService) FindByUserID(userID int) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse) {
	card, err := s.cardRepository.ReadByUserID(userID)
	if err != nil {
		s.logger.Error("failed find card by user id", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	if card == nil {
		s.logger.Error("failed find card by user id", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}
	so := s.mapper.ToCardResponse(*card)

	return &response.ApiResponse[*response.CardResponse]{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardService) FindByCardNumber(cardNumber string) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse) {
	card, err := s.cardRepository.ReadByCardNumber(cardNumber)

	if err != nil {
		s.logger.Error("failed find card by card number", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}

	so := s.mapper.ToCardResponse(*card)

	return &response.ApiResponse[*response.CardResponse]{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil

}

func (s *cardService) FindByUsersID(userID int) (*response.ApiResponse[[]*response.CardResponse], *response.ErrorResponse) {
	cards, err := s.cardRepository.ReadByUsersID(userID)

	if err != nil {
		s.logger.Error("failed find by card id", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Card not found",
		}
	}
	so := s.mapper.ToCardsResponse(cards)

	return &response.ApiResponse[[]*response.CardResponse]{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    so,
	}, nil
}

func (s *cardService) Create(request requests.CreateCardRequest) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse) {
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("failed to find user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	card, err := s.cardRepository.Create(request)

	if err != nil {
		s.logger.Error("failed to create card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card record",
		}
	}

	so := s.mapper.ToCardResponse(*card)

	return &response.ApiResponse[*response.CardResponse]{
		Status:  "success",
		Message: "Card record created successfully",
		Data:    so,
	}, nil
}

func (s *cardService) Update(request requests.UpdateCardRequest) (*response.ApiResponse[*response.CardResponse], *response.ErrorResponse) {
	_, err := s.userRepository.Read(request.UserID)
	if err != nil {
		s.logger.Error("failed to find user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	card, err := s.cardRepository.Update(request)

	if err != nil {
		s.logger.Error("failed to update card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card record",
		}
	}
	so := s.mapper.ToCardResponse(*card)

	return &response.ApiResponse[*response.CardResponse]{
		Status:  "success",
		Message: "Card record updated successfully",
		Data:    so,
	}, nil
}

func (s *cardService) Delete(cardID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := s.cardRepository.Delete(cardID)

	if err != nil {
		s.logger.Error("failed to delete card", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card record",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "Card record deleted successfully",
		Data:    "Card record with ID " + strconv.Itoa(cardID) + " has been deleted",
	}, nil
}
