package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"time"

	"sync"
)

type topupRepository struct {
	mu      sync.RWMutex
	topups  map[int]models.Topup
	nextID  int
	mapping recordmapper.TopupRecordMapping
}

func NewTopupRepository(mapping recordmapper.TopupRecordMapping) *topupRepository {
	return &topupRepository{
		topups:  make(map[int]models.Topup),
		nextID:  1,
		mapping: mapping,
	}
}

func (ds *topupRepository) ReadAll() ([]*record.TopupRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	topups := make([]models.Topup, 0, len(ds.topups))

	for _, topup := range ds.topups {
		topups = append(topups, topup)
	}

	if len(topups) == 0 {
		return nil, fmt.Errorf("no topup found")
	}

	return ds.mapping.ToTopupRecords(topups), nil

}
func (ds *topupRepository) Read(topupID int) (*record.TopupRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	topup, ok := ds.topups[topupID]

	if !ok {
		return nil, fmt.Errorf("topup with ID %d not found", topupID)
	}

	return ds.mapping.ToTopupRecord(topup), nil
}

func (ds *topupRepository) ReadByUserID(userID int) (*record.TopupRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	topup, ok := ds.topups[userID]

	if !ok {
		return nil, fmt.Errorf("topup with ID %d not found", userID)
	}

	return ds.mapping.ToTopupRecord(topup), nil
}

func (ds *topupRepository) ReadByUsersID(userID int) ([]*record.TopupRecord, error) {
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

	return ds.mapping.ToTopupRecords(topups), nil
}

func (ds *topupRepository) Create(request requests.CreateTopupRequest) (*record.TopupRecord, error) {
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

	return ds.mapping.ToTopupRecord(topup), nil
}

func (ds *topupRepository) Update(request requests.UpdateTopupRequest) (*record.TopupRecord, error) {
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

	return ds.mapping.ToTopupRecord(topup), nil
}

func (ds *topupRepository) UpdateAmount(request requests.UpdateTopupAmount) (*record.TopupRecord, error) {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	topup, exists := ds.topups[request.TopupID]

	if !exists {
		return nil, fmt.Errorf("topup with ID %d not found", request.TopupID)
	}

	topup.TopupAmount = request.TopupAmount

	ds.topups[request.TopupID] = topup

	return ds.mapping.ToTopupRecord(topup), nil
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
