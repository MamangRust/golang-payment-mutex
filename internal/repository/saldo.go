package repository

import (
	"fmt"
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

	saldo, ok := ds.saldos[userID]

	if !ok {
		return nil, fmt.Errorf("saldo with ID %d not found", userID)
	}

	return &saldo, nil
}

func (ds *saldoRepository) Create(saldo models.Saldo) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.saldos[saldo.SaldoID]; exists {
		return nil, fmt.Errorf("saldo with ID %d already exists", saldo.SaldoID)
	}

	saldo.SaldoID = ds.nextID
	ds.saldos[saldo.SaldoID] = saldo

	ds.nextID++

	return &saldo, nil
}

func (ds *saldoRepository) Update(newSaldo models.Saldo) (*models.Saldo, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.saldos[newSaldo.SaldoID]; ok {
		ds.saldos[newSaldo.SaldoID] = newSaldo
		return &newSaldo, nil
	}

	return nil, fmt.Errorf("saldo with ID %d not found", newSaldo.SaldoID)
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
