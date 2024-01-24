package repository

import (
	"fmt"
	"payment-mutex/internal/models"

	"sync"
)

type transferRepository struct {
	mu        sync.RWMutex
	transfers map[int]models.Transfer
	nextID    int
}

func NewTransferRepository() *transferRepository {
	return &transferRepository{
		transfers: make(map[int]models.Transfer),
		nextID:    1,
	}
}

func (ds *transferRepository) ReadAll() (*[]models.Transfer, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfers := make([]models.Transfer, 0, len(ds.transfers))

	for _, transfer := range ds.transfers {
		transfers = append(transfers, transfer)
	}

	if len(transfers) == 0 {
		return nil, fmt.Errorf("no transfer found")
	}

	return &transfers, nil
}

func (ds *transferRepository) Read(transferID int) (*models.Transfer, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfer, ok := ds.transfers[transferID]

	if !ok {
		return nil, fmt.Errorf("transfer with ID %d not found", transferID)
	}

	return &transfer, nil
}

func (ds *transferRepository) ReadByUserID(userID int) (*models.Transfer, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	transfer, ok := ds.transfers[userID]

	if !ok {
		return nil, fmt.Errorf("transfer with ID %d not found", userID)
	}

	return &transfer, nil
}

func (ds *transferRepository) Create(transfer models.Transfer) (*models.Transfer, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.transfers[transfer.TransferID]; exists {
		return nil, fmt.Errorf("transfer with ID %d already exists", transfer.TransferID)
	}

	transfer.TransferID = ds.nextID
	ds.nextID++

	return &transfer, nil
}

func (ds *transferRepository) Update(newTransfer models.Transfer) (*models.Transfer, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.transfers[newTransfer.TransferID]; ok {
		ds.transfers[newTransfer.TransferID] = newTransfer
		return &newTransfer, nil
	}

	return nil, fmt.Errorf("transfer with ID %d not found", newTransfer.TransferID)
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
