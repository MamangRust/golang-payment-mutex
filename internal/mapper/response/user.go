package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type userResponseMapper struct {
}

func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{}
}

func (s *userResponseMapper) ToUserResponse(user record.UserRecord) *response.UserResponse {
	return &response.UserResponse{
		ID:          user.UserID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		NocTransfer: user.NocTransfer,
	}
}

func (s *userResponseMapper) ToUsersResponse(users []*record.UserRecord) []*response.UserResponse {
	var responses []*response.UserResponse

	for _, user := range users {
		responses = append(responses, s.ToUserResponse(*user))
	}

	return responses
}
