package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"sync"
)

type userRepository struct {
	mu      sync.RWMutex
	users   map[int]models.User
	nextID  int
	mapping recordmapper.UserRecordMapping
}

func NewUserRepository(mapping recordmapper.UserRecordMapping) *userRepository {
	return &userRepository{
		users:   make(map[int]models.User),
		nextID:  1,
		mapping: mapping,
	}
}

func (ds *userRepository) ReadAll() ([]*record.UserRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	users := make([]models.User, 0, len(ds.users))
	for _, user := range ds.users {
		users = append(users, user)
	}

	return ds.mapping.ToUsersRecord(users), nil
}

func (ds *userRepository) Read(userID int) (*record.UserRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, user := range ds.users {
		if user.UserID == userID {
			return ds.mapping.ToUserRecord(user), nil
		}
	}

	return nil, fmt.Errorf("user with ID %d not found", userID)
}

func (ds *userRepository) ReadByEmail(email string) (*record.UserRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, user := range ds.users {
		if user.Email == email {
			return ds.mapping.ToUserRecord(user), nil
		}
	}

	return nil, fmt.Errorf("user with email %s not found", email)
}

func (ds *userRepository) Create(request requests.CreateUserRequest) (*record.UserRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, existingUser := range ds.users {
		if existingUser.Email == request.Email {
			return nil, fmt.Errorf("user for email %s already exists", request.Email)
		}
	}

	user := models.User{
		UserID:    ds.nextID,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
	}

	user.UserID = ds.nextID
	ds.users[user.UserID] = user
	ds.nextID++

	return ds.mapping.ToUserRecord(user), nil
}

func (ds *userRepository) Update(request requests.UpdateUserRequest) (*record.UserRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	user, ok := ds.users[request.UserID]

	if !ok {
		return nil, fmt.Errorf("user with id %d not found", request.UserID)
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Password = request.Password

	ds.users[request.UserID] = user

	return ds.mapping.ToUserRecord(user), fmt.Errorf("user with ID %d not found", request.UserID)
}

func (ds *userRepository) Delete(userID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.users[userID]; ok {
		delete(ds.users, userID)
		return nil
	}

	return nil
}
