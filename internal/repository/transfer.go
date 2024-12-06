package repository

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"time"

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

func (ds *transferRepository) ReadByUsersID(userID int) (*[]models.Transfer, error) {
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

	return &transfers, nil
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

func (ds *transferRepository) Create(request requests.CreateTransferRequest) (*models.Transfer, error) {
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

	return &transfer, nil
}

func (ds *transferRepository) Update(request requests.UpdateTransferRequest) (*models.Transfer, error) {
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

	return &transfer, nil
}

func (ds *transferRepository) UpdateAmount(request requests.UpdateTransferAmountRequest) (*models.Transfer, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transfer, exists := ds.transfers[request.TransferID]

	if !exists {
		return nil, fmt.Errorf("transfer with id %d not found", request.TransferID)
	}

	transfer.TransferAmount = request.TransferAmount
	transfer.TransferTime = time.Now()

	ds.transfers[request.TransferID] = transfer

	return &transfer, nil

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
