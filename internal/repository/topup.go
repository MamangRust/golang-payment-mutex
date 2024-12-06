package repository

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"time"

	"sync"
)

type topupRepository struct {
	mu     sync.RWMutex
	topups map[int]models.Topup
	nextID int
}

func NewTopupRepository() *topupRepository {
	return &topupRepository{
		topups: make(map[int]models.Topup),
		nextID: 1,
	}
}

func (ds *topupRepository) ReadAll() (*[]models.Topup, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	topups := make([]models.Topup, 0, len(ds.topups))

	for _, topup := range ds.topups {
		topups = append(topups, topup)
	}

	if len(topups) == 0 {
		return nil, fmt.Errorf("no topup found")
	}

	return &topups, nil

}
func (ds *topupRepository) Read(topupID int) (*models.Topup, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	topup, ok := ds.topups[topupID]

	if !ok {
		return nil, fmt.Errorf("topup with ID %d not found", topupID)
	}

	return &topup, nil
}

func (ds *topupRepository) ReadByUserID(userID int) (*models.Topup, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	topup, ok := ds.topups[userID]

	if !ok {
		return nil, fmt.Errorf("topup with ID %d not found", userID)
	}

	return &topup, nil
}

func (ds *topupRepository) ReadByUsersID(userID int) (*[]models.Topup, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	topups := []models.Topup{}

	for _, topup := range ds.topups {
		if topup.UserID == userID {
			topups = append(topups, topup)
		}
	}

	if len(topups) == 0 {
		return nil, fmt.Errorf("no topups not found for user ID %d", userID)
	}

	return &topups, nil
}

func (ds *topupRepository) Create(request requests.CreateTopupRequest) (*models.Topup, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	topup := models.Topup{
		TopupID:     ds.nextID,
		UserID:      request.UserID,
		TopupNo:     request.TopupNo,
		TopupAmount: request.TopupAmount,
		TopupMethod: request.TopupMethod,
		TopupTime:   time.Now(),
	}

	ds.topups[topup.TopupID] = topup

	ds.nextID++

	return &topup, nil
}

func (ds *topupRepository) Update(request requests.UpdateTopupRequest) (*models.Topup, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	topup, exists := ds.topups[request.TopupID]

	if !exists {
		return nil, fmt.Errorf("topup with ID %d not found", request.TopupID)

	}

	topup.UserID = request.UserID
	topup.TopupAmount = request.TopupAmount
	topup.TopupMethod = request.TopupMethod
	topup.TopupTime = time.Now()
	ds.topups[request.TopupID] = topup

	return &topup, nil
}

func (ds *topupRepository) UpdateAmount(request requests.UpdateTopupAmount) (*models.Topup, error) {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	topup, exists := ds.topups[request.TopupID]

	if !exists {
		return nil, fmt.Errorf("topup with ID %d not found", request.TopupID)
	}

	topup.TopupAmount = request.TopupAmount

	ds.topups[request.TopupID] = topup

	return &topup, nil
}

func (ds *topupRepository) Delete(topupID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.topups[topupID]; ok {
		delete(ds.topups, topupID)
		return nil
	}

	return fmt.Errorf("topup with ID %d not found", topupID)
}
