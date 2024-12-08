package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type topupRecordMapper struct {
}

func NewTopupRecordMapper() *topupRecordMapper {
	return &topupRecordMapper{}
}

func (t *topupRecordMapper) ToTopupRecord(topup models.Topup) *record.TopupRecord {
	return &record.TopupRecord{
		TopupID:     topup.TopupID,
		UserID:      topup.UserID,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
	}
}

func (t *topupRecordMapper) ToTopupRecords(topups []models.Topup) []*record.TopupRecord {
	var responses []*record.TopupRecord

	for _, response := range topups {
		responses = append(responses, t.ToTopupRecord(response))
	}

	return responses
}
