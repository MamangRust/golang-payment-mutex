package service

import (
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/internal/repository"
	"payment-mutex/pkg/auth"
	"payment-mutex/pkg/hash"
	"payment-mutex/pkg/logger"

	"go.uber.org/zap"
)

type authService struct {
	hash       hash.Hashing
	repository repository.UserRepository
	token      auth.TokenManager
	logger     logger.Logger
}

func NewAuthService(hash hash.Hashing, repository repository.UserRepository, token auth.TokenManager, logger logger.Logger) *authService {
	return &authService{
		hash:       hash,
		repository: repository,
		token:      token,
		logger:     logger,
	}
}

func (s *authService) RegisterUser(request *requests.RegisterRequest) (*models.User, error) {
	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("Error hashing password: ", zap.Error(err))
		return nil, err
	}

	res, err := s.repository.Create(requests.CreateUserRequest{
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Email:           request.Email,
		Password:        hashing,
		ConfirmPassword: request.ConfirmPassword,
	})

	if err != nil {
		s.logger.Error("Error creating user: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *authService) Login(request *requests.AuthRequest) (string, error) {
	res, err := s.repository.ReadByEmail(request.Email)
	if err != nil {
		s.logger.Error("failed login: ", zap.Error(err))
		return "", err
	}

	err = s.hash.ComparePassword(res.Password, request.Password)

	if err != nil {
		s.logger.Error("Error comparing password: ", zap.Error(err))
	}

	token, err := s.createJwt(int(res.UserID))

	if err != nil {
		s.logger.Error("failed create jwt token: ", zap.Error(err))
		return "", err
	}

	return token, nil

}

func (s *authService) createJwt(id int) (string, error) {
	token, err := s.token.NewJwtToken(id)
	if err != nil {
		s.logger.Error("failed create jwt token: ", zap.Error(err))
		return "", err
	}
	return token, nil
}
