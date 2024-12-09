package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type transactionRecordMapper struct {
}

func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{}
}

func (s *transactionRecordMapper) ToTransactionRecord(transfer models.Transaction) *record.TransactionRecord {
	return &record.TransactionRecord{
		TransactionID:   transfer.TransactionID,
		CardNumber:      transfer.CardNumber,
		Amount:          transfer.Amount,
		PaymentMethod:   transfer.PaymentMethod,
		TransactionTime: transfer.TransactionTime,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecord(transfers []models.Transaction) []*record.TransactionRecord {
	var transferRecords []*record.TransactionRecord
	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransactionRecord(transfer))
	}
	return transferRecords
}
