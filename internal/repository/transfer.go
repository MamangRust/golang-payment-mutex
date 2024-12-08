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

func (ds *transferRepository) ReadAll() ([]*record.TransferRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfers := make([]models.Transfer, 0, len(ds.transfers))

	for _, transfer := range ds.transfers {
		transfers = append(transfers, transfer)
	}

	if len(transfers) == 0 {
		return nil, fmt.Errorf("no transfer found")
	}

	return ds.mapping.ToTransfersRecord(transfers), nil
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

func (ds *transferRepository) ReadByUsersID(userID int) ([]*record.TransferRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	transfers := make([]models.Transfer, 0, len(ds.transfers))

	for _, transfer := range ds.transfers {
		if transfer.TransferFrom == userID || transfer.TransferTo == userID {
			transfers = append(transfers, transfer)
		}
	}

	if len(transfers) == 0 {
		return nil, fmt.Errorf("not transfer not found for user id %d", userID)
	}

	return ds.mapping.ToTransfersRecord(transfers), nil
}

func (ds *transferRepository) ReadByUserID(userID int) (*record.TransferRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfer, ok := ds.transfers[userID]

	if !ok {
		return nil, fmt.Errorf("transfer with ID %d not found", userID)
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
