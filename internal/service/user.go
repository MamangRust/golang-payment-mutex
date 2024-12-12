package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	responseMapper "payment-mutex/internal/mapper/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type userService struct {
	userRepository repository.UserRepository
	logger         logger.Logger
	mapper         responseMapper.UserResponseMapper
}

func NewUserService(
	userRepository repository.UserRepository,
	logger logger.Logger,
	mapper responseMapper.UserResponseMapper,
) *userService {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
		mapper:         mapper,
	}
}

func (ds *userService) FindAll(page int, pageSize int, search string) (*response.APIResponsePagination[[]*response.UserResponse], *response.ErrorResponse) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := ds.userRepository.ReadAll(page, pageSize, search)
	if err != nil {
		ds.logger.Error("failed to fetch users", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch users",
		}
	}

	if len(users) == 0 {
		ds.logger.Error("no users found")
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "No users found",
		}
	}

	userResponses := ds.mapper.ToUsersResponse(users)

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return &response.APIResponsePagination[[]*response.UserResponse]{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    userResponses,
		Meta: response.PaginationMeta{
			CurrentPage:  page,
			PageSize:     pageSize,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
	}, nil
}

func (ds *userService) FindByID(id int) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse) {
	user, err := ds.userRepository.Read(id)
	if err != nil {
		ds.logger.Error("failed to find user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	so := ds.mapper.ToUserResponse(*user)

	return &response.ApiResponse[*response.UserResponse]{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    so,
	}, nil
}

func (ds *userService) Create(request requests.CreateUserRequest) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse) {
	existingUser, err := ds.userRepository.ReadByEmail(request.Email)
	if existingUser != nil {
		ds.logger.Error("user already exists with the given email", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User with the given email already exists",
		}
	}

	res, err := ds.userRepository.Create(request)
	if err != nil {
		ds.logger.Error("failed to create user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create user",
		}
	}

	so := ds.mapper.ToUserResponse(*res)

	return &response.ApiResponse[*response.UserResponse]{
		Status:  "success",
		Message: "User created successfully",
		Data:    so,
	}, nil
}

func (ds *userService) Update(request requests.UpdateUserRequest) (*response.ApiResponse[*response.UserResponse], *response.ErrorResponse) {
	existingUser, err := ds.userRepository.ReadByEmail(request.Email)
	if existingUser == nil {
		ds.logger.Error("user not found with the given email", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User with the given email not found",
		}
	}

	res, err := ds.userRepository.Update(request)
	if err != nil {
		ds.logger.Error("failed to update user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update user",
		}
	}

	so := ds.mapper.ToUserResponse(*res)

	return &response.ApiResponse[*response.UserResponse]{
		Status:  "success",
		Message: "User updated successfully",
		Data:    so,
	}, nil
}

func (ds *userService) Delete(userID int) (*response.ApiResponse[string], *response.ErrorResponse) {
	err := ds.userRepository.Delete(userID)
	if err != nil {
		ds.logger.Error("failed to delete user", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete user",
		}
	}

	return &response.ApiResponse[string]{
		Status:  "success",
		Message: "User deleted successfully",
		Data:    "",
	}, nil
}
