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

type transactionRepository struct {
	mu           sync.RWMutex
	transactions map[int]models.Transaction
	nextID       int
	mapping      recordmapper.TransactionRecordMapping
}

func NewTransactionRepository(mapping recordmapper.TransactionRecordMapping) *transactionRepository {
	return &transactionRepository{
		transactions: make(map[int]models.Transaction),
		nextID:       1,
		mapping:      mapping,
	}
}

func (ds *transactionRepository) ReadAll(page int, pageSize int, search string) ([]*record.TransactionRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredTransactions := make([]models.Transaction, 0)

	for _, transaction := range ds.transactions {
		if search == "" ||
			strings.Contains(strings.ToLower(transaction.CardNumber), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(transaction.PaymentMethod), strings.ToLower(search)) {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	totalRecords := len(filteredTransactions)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedTransactions := filteredTransactions[start:end]

	return ds.mapping.ToTransactionsRecord(paginatedTransactions), totalRecords, nil
}

func (ds *transactionRepository) CountByDate(date string) (int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	count := 0
	for _, transaction := range ds.transactions {
		if transaction.TransactionTime.Format("2006-01-02") == date {
			count++
		}
	}

	return count, nil
}

func (ds *transactionRepository) CountAll() (int, error) {
	return len(ds.transactions), nil
}

func (ds *transactionRepository) Read(transactionID int) (*record.TransactionRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	transaction, ok := ds.transactions[transactionID]

	if !ok {
		return nil, fmt.Errorf("transaction with ID %d not found", transactionID)
	}

	return ds.mapping.ToTransactionRecord(transaction), nil
}

func (ds *transactionRepository) Create(request requests.CreateTransactionRequest) (*record.TransactionRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transaction := models.Transaction{
		TransactionID:   ds.nextID,
		CardNumber:      request.CardNumber,
		Amount:          request.Amount,
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      *request.MerchantID,
		TransactionTime: request.TransactionTime,
	}

	ds.transactions[transaction.TransactionID] = transaction
	ds.nextID++

	return ds.mapping.ToTransactionRecord(transaction), nil
}

func (ds *transactionRepository) Update(request requests.UpdateTransactionRequest) (*record.TransactionRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	transaction, ok := ds.transactions[request.TransactionID]
	if !ok {
		return nil, fmt.Errorf("transaction with ID %d not found", request.TransactionID)
	}

	transaction.CardNumber = request.CardNumber
	transaction.Amount = request.Amount
	transaction.MerchantID = *request.MerchantID
	transaction.PaymentMethod = request.PaymentMethod
	transaction.TransactionTime = request.TransactionTime

	ds.transactions[transaction.TransactionID] = transaction

	return ds.mapping.ToTransactionRecord(transaction), nil
}

func (ds *transactionRepository) Delete(transactionID int) error {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	if _, ok := ds.transactions[transactionID]; ok {
		delete(ds.transactions, transactionID)
		return nil
	}

	return fmt.Errorf("transaction with ID %d not found", transactionID)
}
