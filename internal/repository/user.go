package repository

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"payment-mutex/pkg/randomvcc"
	"sync"
)

type userRepository struct {
	mu     sync.RWMutex
	users  map[int]models.User
	nextID int
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (ds *userRepository) ReadAll() (*[]models.User, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	users := make([]models.User, 0, len(ds.users))
	for _, user := range ds.users {
		users = append(users, user)
	}

	return &users, nil
}

func (ds *userRepository) Read(userID int) (*models.User, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, user := range ds.users {
		if user.UserID == userID {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user with ID %d not found", userID)
}

func (ds *userRepository) ReadByEmail(email string) (*models.User, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, user := range ds.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user with email %s not found", email)
}

func (ds *userRepository) Create(request requests.CreateUserRequest) (*models.User, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, existingUser := range ds.users {
		if existingUser.Email == request.Email {
			return nil, fmt.Errorf("user for email %s already exists", request.Email)
		}
	}

	random, err := randomvcc.RandomVCC()

	if err != nil {
		return nil, fmt.Errorf("random vcc error: %d", err)
	}

	user := models.User{
		UserID:      ds.nextID,
		Email:       request.Email,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Password:    request.Password,
		NocTransfer: int(random),
	}

	user.UserID = ds.nextID
	ds.users[user.UserID] = user
	ds.nextID++

	return &user, nil
}

func (ds *userRepository) Update(request requests.UpdateUserRequest) (*models.User, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	user, ok := ds.users[request.UserID]

	if !ok {
		return nil, fmt.Errorf("user with id %d not found", request.UserID)
	}

	random, err := randomvcc.RandomVCC()

	if err != nil {
		return nil, fmt.Errorf("random vcc error: %d", err)
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Password = request.Password
	user.NocTransfer = int(random)

	ds.users[request.UserID] = user

	return nil, fmt.Errorf("user with ID %d not found", request.UserID)
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
