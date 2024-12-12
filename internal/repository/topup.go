package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"strings"
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

func (ds *topupRepository) ReadAll(page int, pageSize int, search string) ([]*record.TopupRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredTopups := make([]models.Topup, 0)

	for _, topup := range ds.topups {
		if search == "" ||
			strings.Contains(strings.ToLower(topup.CardNumber), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(topup.TopupNo), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(topup.TopupMethod), strings.ToLower(search)) {
			filteredTopups = append(filteredTopups, topup)
		}
	}

	totalRecords := len(filteredTopups)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedTopups := filteredTopups[start:end]

	return ds.mapping.ToTopupRecords(paginatedTopups), totalRecords, nil
}

func (ds *topupRepository) CountByDate(date string) (int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	count := 0
	for _, topup := range ds.topups {
		if topup.TopupTime.Format("2006-01-02") == date {
			count++
		}
	}

	return count, nil
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

func (ds *topupRepository) Create(request requests.CreateTopupRequest) (*record.TopupRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	topup := models.Topup{
		TopupID:     ds.nextID,
		CardNumber:  request.CardNumber,
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

	topup.CardNumber = request.CardNumber
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
