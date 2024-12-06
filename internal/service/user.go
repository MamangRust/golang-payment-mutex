package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
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

func (ds *userService) FindAll() (*[]models.User, error) {
	user, err := ds.userRepository.ReadAll()

	if err != nil {
		ds.logger.Error("failed find all user: ", zap.Error(err))

		return nil, err
	}

	return user, nil

}

func (ds *userService) FindByID(id int) (*models.User, error) {
	user, err := ds.userRepository.Read(id)

	if err != nil {
		ds.logger.Error("failed find user by id: ", zap.Error(err))

		return nil, err
	}

	return user, nil
}

func (ds *userService) Create(request requests.CreateUserRequest) (*models.User, error) {
	_, err := ds.userRepository.ReadByEmail(request.Email)

	if err != nil {
		ds.logger.Error("failed find user by email: ", zap.Error(err))

		return nil, err
	}

	res, err := ds.userRepository.Create(request)

	if err != nil {
		ds.logger.Error("failed create user: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (ds *userService) Update(request requests.UpdateUserRequest) (*models.User, error) {
	_, err := ds.userRepository.ReadByEmail(request.Email)

	if err != nil {
		ds.logger.Error("failed find user by email: ", zap.Error(err))

		return nil, err
	}

	res, err := ds.userRepository.Update(request)

	if err != nil {
		ds.logger.Error("failed update user: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (ds *userService) Delete(userID int) error {
	err := ds.userRepository.Delete(userID)

	if err != nil {
		ds.logger.Error("failed delete user: ", zap.Error(err))
	}

	return nil
}
