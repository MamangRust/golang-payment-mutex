package repository

import (
	"fmt"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/models"
	"sync"
)

type saldoRepository struct {
	mu     sync.RWMutex
	saldos map[int]models.Saldo
	nextID int
}

func NewSaldoRepository() *saldoRepository {
	return &saldoRepository{
		saldos: make(map[int]models.Saldo),
		nextID: 1,
	}
}

func (ds *saldoRepository) ReadAll() (*[]models.Saldo, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	saldos := make([]models.Saldo, 0, len(ds.saldos))
	for _, saldo := range ds.saldos {
		saldos = append(saldos, saldo)
	}
	if len(saldos) == 0 {
		return nil, fmt.Errorf("no saldo found")
	}
	return &saldos, nil
}

func (ds *saldoRepository) Read(saldoID int) (*models.Saldo, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	saldo, ok := ds.saldos[saldoID]
	if !ok {
		return nil, fmt.Errorf("saldo with ID %d not found", saldoID)
	}
	return &saldo, nil
}

func (ds *saldoRepository) ReadByUserID(userID int) (*models.Saldo, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	for _, saldo := range ds.saldos {
		if saldo.UserID == userID {
			return &saldo, nil
		}
	}
	return nil, fmt.Errorf("saldo for user ID %d not found", userID)
}

func (ds *saldoRepository) ReadByUsersID(userID int) (*[]models.Saldo, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	var userSaldos []models.Saldo
	for _, saldo := range ds.saldos {
		if saldo.UserID == userID {
			userSaldos = append(userSaldos, saldo)
		}
	}
	if len(userSaldos) == 0 {
		return nil, fmt.Errorf("no saldo found for user ID %d", userID)
	}
	return &userSaldos, nil
}

func (ds *saldoRepository) Create(request requests.CreateSaldoRequest) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, existingSaldo := range ds.saldos {
		if existingSaldo.UserID == request.UserID {
			return nil, fmt.Errorf("saldo for user ID %d already exists", request.UserID)
		}
	}

	saldo := models.Saldo{
		SaldoID:      ds.nextID,
		UserID:       request.UserID,
		TotalBalance: request.TotalBalance,
	}

	ds.saldos[saldo.SaldoID] = saldo
	ds.nextID++
	return &saldo, nil
}

func (ds *saldoRepository) Update(request requests.UpdateSaldoRequest) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	saldo, ok := ds.saldos[request.SaldoID]
	if !ok {
		return nil, fmt.Errorf("saldo with ID %d not found", request.SaldoID)
	}

	saldo.UserID = request.UserID
	saldo.TotalBalance = request.TotalBalance
	saldo.WithdrawAmount = request.WithdrawAmount
	saldo.WithdrawTime = request.WithdrawTime

	ds.saldos[request.SaldoID] = saldo
	return &saldo, nil
}

func (ds *saldoRepository) UpdateBalance(request requests.UpdateSaldoBalance) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for id, saldo := range ds.saldos {
		if saldo.UserID == request.UserID {
			updatedSaldo := saldo
			updatedSaldo.TotalBalance = request.TotalBalance

			ds.saldos[id] = updatedSaldo

			return &updatedSaldo, nil
		}
	}

	return nil, fmt.Errorf("saldo for user ID %d not found", request.UserID)
}

func (ds *saldoRepository) UpdateSaldoWithdraw(request requests.UpdateSaldoWithdraw) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for id, saldo := range ds.saldos {
		if saldo.UserID == request.UserID {
			updatedSaldo := saldo
			if request.WithdrawAmount != nil {
				if *request.WithdrawAmount > updatedSaldo.TotalBalance {
					return nil, fmt.Errorf("insufficient balance")
				}
				updatedSaldo.WithdrawAmount = *request.WithdrawAmount
				updatedSaldo.TotalBalance -= *request.WithdrawAmount
			}

			if request.WithdrawTime != nil {
				updatedSaldo.WithdrawTime = *request.WithdrawTime
			}

			ds.saldos[id] = updatedSaldo

			return &updatedSaldo, nil
		}
	}

	return nil, fmt.Errorf("saldo for user ID %d not found", request.UserID)
}

func (ds *saldoRepository) Delete(saldoID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	if _, ok := ds.saldos[saldoID]; ok {
		delete(ds.saldos, saldoID)
		return nil
	}
	return fmt.Errorf("saldo with ID %d not found", saldoID)
}
