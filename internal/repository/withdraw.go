package repository

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"sync"
)

type withdrawRepository struct {
	mu       sync.RWMutex
	withdraw map[int]models.Withdraw
	nextID   int
}

func NewWithdrawRepository() *withdrawRepository {
	return &withdrawRepository{
		withdraw: make(map[int]models.Withdraw),
		nextID:   1,
	}
}

func (ds *withdrawRepository) ReadAll() (*[]models.Withdraw, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	withdraws := make([]models.Withdraw, 0, len(ds.withdraw))

	for _, withdraw := range ds.withdraw {
		withdraws = append(withdraws, withdraw)
	}
	if len(withdraws) == 0 {
		return nil, fmt.Errorf("no withdraw found")
	}

	return &withdraws, nil

}

func (ds *withdrawRepository) Read(withdrawID int) (*models.Withdraw, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	withdraw, ok := ds.withdraw[withdrawID]

	if !ok {
		return nil, fmt.Errorf("withdraw with ID %d not found", withdrawID)
	}

	return &withdraw, nil
}

func (ds *withdrawRepository) ReadByUserID(userID int) (*models.Withdraw, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	withdraw, ok := ds.withdraw[userID]

	if !ok {
		return nil, fmt.Errorf("withdraw with ID %d not found", userID)
	}

	return &withdraw, nil
}

func (ds *withdrawRepository) ReadByUsersID(userID int) (*[]models.Withdraw, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	withdraws := []models.Withdraw{}

	for _, withdraw := range ds.withdraw {
		if withdraw.UserID == userID {
			withdraws = append(withdraws, withdraw)
		}
	}

	if len(withdraws) == 0 {
		return nil, fmt.Errorf("no withdraws not found for user ID %d", userID)
	}

	return &withdraws, nil
}

func (ds *withdrawRepository) Create(request requests.CreateWithdrawRequest) (*models.Withdraw, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	withdraw := models.Withdraw{
		WithdrawID:     ds.nextID,
		UserID:         request.UserID,
		WithdrawAmount: request.WithdrawAmount,
		WithdrawTime:   request.WithdrawTime,
	}

	withdraw.WithdrawID = ds.nextID
	ds.withdraw[withdraw.WithdrawID] = withdraw

	ds.nextID++

	return &withdraw, nil
}

func (ds *withdrawRepository) Update(request requests.UpdateWithdrawRequest) (*models.Withdraw, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	withdraw, exists := ds.withdraw[request.WithdrawID]

	if !exists {
		return nil, fmt.Errorf("withdraw with id %d not found", request.WithdrawID)
	}

	withdraw.UserID = request.UserID
	withdraw.WithdrawAmount = request.WithdrawAmount
	withdraw.WithdrawTime = request.WithdrawTime

	ds.withdraw[request.WithdrawID] = withdraw

	return &withdraw, nil
}

func (ds *withdrawRepository) Delete(withdrawID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.withdraw[withdrawID]; ok {
		delete(ds.withdraw, withdrawID)
		return nil
	}

	return fmt.Errorf("withdraw with ID %d not found", withdrawID)
}
