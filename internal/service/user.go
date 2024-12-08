package service

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type userService struct {
	userRepository repository.UserRepository
	logger         logger.Logger
}

func NewUserService(
	userRepository repository.UserRepository,
	logger logger.Logger,
) *userService {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (ds *userService) FindAll() (*response.ApiResponse[[]*record.UserRecord], *response.ErrorResponse) {
	users, err := ds.userRepository.ReadAll()
	if err != nil {
		ds.logger.Error("failed to find all users", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve users",
		}
	}

	return &response.ApiResponse[[]*record.UserRecord]{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	}, nil
}

func (ds *userService) FindByID(id int) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse) {
	user, err := ds.userRepository.Read(id)
	if err != nil {
		ds.logger.Error("failed to find user by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		}
	}

	return &response.ApiResponse[record.UserRecord]{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    *user,
	}, nil
}

func (ds *userService) Create(request requests.CreateUserRequest) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse) {
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

	res.Password = nil

	return &response.ApiResponse[record.UserRecord]{
		Status:  "success",
		Message: "User created successfully",
		Data:    *res,
	}, nil
}

func (ds *userService) Update(request requests.UpdateUserRequest) (*response.ApiResponse[record.UserRecord], *response.ErrorResponse) {
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

	res.Password = nil

	return &response.ApiResponse[record.UserRecord]{
		Status:  "success",
		Message: "User updated successfully",
		Data:    *res,
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
