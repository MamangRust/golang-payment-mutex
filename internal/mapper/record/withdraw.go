package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type withdrawRecordMapper struct {
}

func NewWithdrawRecordMapper() *withdrawRecordMapper {
	return &withdrawRecordMapper{}
}

func (s *withdrawRecordMapper) ToWithdrawRecord(withdraw models.Withdraw) *record.WithdrawRecord {
	return &record.WithdrawRecord{
		WithdrawID:     withdraw.WithdrawID,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
	}
}

func (s *withdrawRecordMapper) ToWithdrawsRecord(withdraws []models.Withdraw) []*record.WithdrawRecord {
	var responses []*record.WithdrawRecord

	for _, response := range withdraws {
		responses = append(responses, s.ToWithdrawRecord(response))
	}

	return responses
}
