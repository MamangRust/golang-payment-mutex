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

type transferRepository struct {
	mu        sync.RWMutex
	transfers map[int]models.Transfer
	nextID    int
	mapping   recordmapper.TransferRecordMapping
}

func NewTransferRepository(mapping recordmapper.TransferRecordMapping) *transferRepository {
	return &transferRepository{
		transfers: make(map[int]models.Transfer),
		nextID:    1,
		mapping:   mapping,
	}
}

func (ds *transferRepository) ReadAll(page int, pageSize int, search string) ([]*record.TransferRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredTransfers := make([]models.Transfer, 0)

	for _, transfer := range ds.transfers {
		if search == "" ||
			strings.Contains(strings.ToLower(transfer.TransferFrom), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(transfer.TransferTo), strings.ToLower(search)) {
			filteredTransfers = append(filteredTransfers, transfer)
		}
	}

	totalRecords := len(filteredTransfers)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedTransfers := filteredTransfers[start:end]

	return ds.mapping.ToTransfersRecord(paginatedTransfers), totalRecords, nil
}

func (ds *transferRepository) CountByDate(date string) (int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	count := 0
	for _, transfer := range ds.transfers {
		if transfer.TransferTime.Format("2006-01-02") == date {
			count++
		}
	}

	return count, nil
}

func (ds *transferRepository) CountAll() (int, error) {
	return len(ds.transfers), nil
}

func (ds *transferRepository) Read(transferID int) (*record.TransferRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfer, ok := ds.transfers[transferID]

	if !ok {
		return nil, fmt.Errorf("transfer with ID %d not found", transferID)
	}

	return ds.mapping.ToTransferRecord(transfer), nil
}

func (ds *transferRepository) Create(request requests.CreateTransferRequest) (*record.TransferRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transfer := models.Transfer{
		TransferID:     ds.nextID,
		TransferFrom:   request.TransferFrom,
		TransferTo:     request.TransferTo,
		TransferAmount: request.TransferAmount,
		TransferTime:   time.Now(),
	}

	ds.transfers[transfer.TransferID] = transfer

	ds.nextID++

	return ds.mapping.ToTransferRecord(transfer), nil
}

func (ds *transferRepository) Update(request requests.UpdateTransferRequest) (*record.TransferRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transfer, exists := ds.transfers[request.TransferID]

	if !exists {
		return nil, fmt.Errorf("transfer with ID %d not found", request.TransferID)
	}

	transfer.TransferFrom = request.TransferFrom
	transfer.TransferTo = request.TransferTo
	transfer.TransferAmount = request.TransferAmount
	transfer.TransferTime = time.Now()
	ds.transfers[request.TransferID] = transfer

	return ds.mapping.ToTransferRecord(transfer), nil
}

func (ds *transferRepository) UpdateAmount(request requests.UpdateTransferAmountRequest) (*record.TransferRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transfer, exists := ds.transfers[request.TransferID]

	if !exists {
		return nil, fmt.Errorf("transfer with id %d not found", request.TransferID)
	}

	transfer.TransferAmount = request.TransferAmount
	transfer.TransferTime = time.Now()

	ds.transfers[request.TransferID] = transfer

	return ds.mapping.ToTransferRecord(transfer), nil

}

func (ds *transferRepository) Delete(transferID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.transfers[transferID]; ok {
		delete(ds.transfers, transferID)
		return nil
	}

	return fmt.Errorf("transfer with ID %d not found", transferID)
}
