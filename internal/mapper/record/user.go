package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type userRecordMapper struct {
}

func NewUserRecordMapper() *userRecordMapper {
	return &userRecordMapper{}
}

func (u *userRecordMapper) ToUserRecord(user models.User) *record.UserRecord {
	var password *string

	if user.Password != "" {
		password = &user.Password
	}

	return &record.UserRecord{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  password,
	}
}

func (u *userRecordMapper) ToUsersRecord(users []models.User) []*record.UserRecord {

	var responses []*record.UserRecord

	for _, response := range users {
		responses = append(responses, u.ToUserRecord(response))
	}
	return responses
}
