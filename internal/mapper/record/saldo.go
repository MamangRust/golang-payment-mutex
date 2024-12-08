package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type saldoRecordMapper struct {
}

func NewSaldoRecordMapper() *saldoRecordMapper {
	return &saldoRecordMapper{}
}

func (s *saldoRecordMapper) ToSaldoRecord(saldo models.Saldo) *record.SaldoRecord {
	return &record.SaldoRecord{
		SaldoID:        saldo.SaldoID,
		UserID:         saldo.UserID,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
	}
}

func (s *saldoRecordMapper) ToSaldosRecord(saldos []models.Saldo) []*record.SaldoRecord {
	var responses []*record.SaldoRecord

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoRecord(response))
	}

	return responses
}
