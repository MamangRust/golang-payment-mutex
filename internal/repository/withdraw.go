package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"strings"
	"sync"
)

type withdrawRepository struct {
	mu       sync.RWMutex
	withdraw map[int]models.Withdraw
	nextID   int
	mapping  recordmapper.WithdrawRecordMapping
}

func NewWithdrawRepository(mapping recordmapper.WithdrawRecordMapping) *withdrawRepository {
	return &withdrawRepository{
		withdraw: make(map[int]models.Withdraw),
		nextID:   1,
		mapping:  mapping,
	}
}

func (ds *withdrawRepository) ReadAll(page int, pageSize int, search string) ([]*record.WithdrawRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredWithdraws := make([]models.Withdraw, 0)

	for _, withdraw := range ds.withdraw {
		if search == "" ||
			strings.Contains(strings.ToLower(withdraw.CardNumber), strings.ToLower(search)) {
			filteredWithdraws = append(filteredWithdraws, withdraw)
		}
	}

	totalRecords := len(filteredWithdraws)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedWithdraws := filteredWithdraws[start:end]

	return ds.mapping.ToWithdrawsRecord(paginatedWithdraws), totalRecords, nil
}

func (ds *withdrawRepository) CountByDate(date string) (int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	count := 0
	for _, withdraw := range ds.withdraw {
		if withdraw.WithdrawTime.Format("2006-01-02") == date {
			count++
		}
	}

	return count, nil
}

func (ds *withdrawRepository) Read(withdrawID int) (*record.WithdrawRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	withdraw, ok := ds.withdraw[withdrawID]

	if !ok {
		return nil, fmt.Errorf("withdraw with ID %d not found", withdrawID)
	}

	return ds.mapping.ToWithdrawRecord(withdraw), nil
}

func (ds *withdrawRepository) Create(request requests.CreateWithdrawRequest) (*record.WithdrawRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	withdraw := models.Withdraw{
		WithdrawID:     ds.nextID,
		CardNumber:     request.CardNumber,
		WithdrawAmount: request.WithdrawAmount,
		WithdrawTime:   request.WithdrawTime,
	}

	withdraw.WithdrawID = ds.nextID
	ds.withdraw[withdraw.WithdrawID] = withdraw

	ds.nextID++

	return ds.mapping.ToWithdrawRecord(withdraw), nil
}

func (ds *withdrawRepository) Update(request requests.UpdateWithdrawRequest) (*record.WithdrawRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	withdraw, exists := ds.withdraw[request.WithdrawID]

	if !exists {
		return nil, fmt.Errorf("withdraw with id %d not found", request.WithdrawID)
	}

	withdraw.CardNumber = request.CardNumber
	withdraw.WithdrawAmount = request.WithdrawAmount
	withdraw.WithdrawTime = request.WithdrawTime

	ds.withdraw[request.WithdrawID] = withdraw

	return ds.mapping.ToWithdrawRecord(withdraw), nil
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
