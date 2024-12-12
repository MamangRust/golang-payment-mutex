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

type saldoRepository struct {
	mu      sync.RWMutex
	saldos  map[int]models.Saldo
	nextID  int
	mapping recordmapper.SaldoRecordMapping
}

func NewSaldoRepository(mapping recordmapper.SaldoRecordMapping) *saldoRepository {
	return &saldoRepository{
		saldos:  make(map[int]models.Saldo),
		nextID:  1,
		mapping: mapping,
	}
}

func (ds *saldoRepository) ReadAll(page int, pageSize int, search string) ([]*record.SaldoRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredSaldos := make([]models.Saldo, 0)

	for _, saldo := range ds.saldos {
		if search == "" ||
			strings.Contains(strings.ToLower(saldo.CardNumber), strings.ToLower(search)) {
			filteredSaldos = append(filteredSaldos, saldo)
		}
	}

	totalRecords := len(filteredSaldos)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedSaldos := filteredSaldos[start:end]

	return ds.mapping.ToSaldosRecord(paginatedSaldos), totalRecords, nil
}

func (ds *saldoRepository) Read(saldoID int) (*record.SaldoRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	saldo, ok := ds.saldos[saldoID]
	if !ok {
		return nil, fmt.Errorf("saldo with ID %d not found", saldoID)
	}
	return ds.mapping.ToSaldoRecord(saldo), nil
}

func (ds *saldoRepository) ReadByCardNumber(cardNumber string) (*record.SaldoRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	for _, saldo := range ds.saldos {
		if saldo.CardNumber == cardNumber {
			return ds.mapping.ToSaldoRecord(saldo), nil
		}
	}
	return nil, fmt.Errorf("saldo for user ID %s not found", cardNumber)
}

func (ds *saldoRepository) Create(request requests.CreateSaldoRequest) (*record.SaldoRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, existingSaldo := range ds.saldos {
		if existingSaldo.CardNumber == request.CardNumber {
			return nil, fmt.Errorf("saldo for user ID %s already exists", request.CardNumber)
		}
	}

	saldo := models.Saldo{
		SaldoID:      ds.nextID,
		CardNumber:   request.CardNumber,
		TotalBalance: request.TotalBalance,
	}

	ds.saldos[saldo.SaldoID] = saldo
	ds.nextID++

	return ds.mapping.ToSaldoRecord(saldo), nil
}

func (ds *saldoRepository) Update(request requests.UpdateSaldoRequest) (*record.SaldoRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	saldo, ok := ds.saldos[request.SaldoID]
	if !ok {
		return nil, fmt.Errorf("saldo with ID %d not found", request.SaldoID)
	}

	saldo.CardNumber = request.CardNumber
	saldo.TotalBalance = request.TotalBalance

	ds.saldos[request.SaldoID] = saldo

	return ds.mapping.ToSaldoRecord(saldo), nil
}

func (ds *saldoRepository) UpdateBalance(request requests.UpdateSaldoBalance) (*record.SaldoRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for id, saldo := range ds.saldos {
		if saldo.CardNumber == request.CardNumber {
			updatedSaldo := saldo
			updatedSaldo.TotalBalance = request.TotalBalance

			ds.saldos[id] = updatedSaldo

			return ds.mapping.ToSaldoRecord(saldo), nil
		}
	}

	return nil, fmt.Errorf("saldo for user ID %s not found", request.CardNumber)
}

func (ds *saldoRepository) UpdateSaldoWithdraw(request requests.UpdateSaldoWithdraw) (*record.SaldoRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for id, saldo := range ds.saldos {
		if saldo.CardNumber == request.CardNumber {
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

			return ds.mapping.ToSaldoRecord(saldo), nil
		}
	}

	return nil, fmt.Errorf("saldo for user ID %s not found", request.CardNumber)
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
