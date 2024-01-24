package repository

import (
	"fmt"
	"payment-mutex/internal/models"

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

func (ds *topupRepository) Create(topup models.Topup) (*models.Topup, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.topups[topup.TopupID]; exists {
		return nil, fmt.Errorf("topup with ID %d already exists", topup.TopupID)
	}

	topup.TopupID = ds.nextID
	ds.topups[topup.TopupID] = topup

	ds.nextID++

	return &topup, nil
}

func (ds *topupRepository) Update(newTopup models.Topup) (*models.Topup, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.topups[newTopup.TopupID]; ok {
		ds.topups[newTopup.TopupID] = newTopup
		return &newTopup, nil
	}

	return nil, fmt.Errorf("topup with ID %d not found", newTopup.TopupID)
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
