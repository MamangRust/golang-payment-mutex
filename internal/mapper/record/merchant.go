package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type merchantRecordMapper struct {
}

func NewMerchantRecordMapper() *merchantRecordMapper {
	return &merchantRecordMapper{}
}

func (m *merchantRecordMapper) ToMerchantRecord(merchant models.Merchant) *record.MerchantRecord {
	return &record.MerchantRecord{
		MerchantID: merchant.MerchantID,
		Name:       merchant.Name,
		ApiKey:     merchant.ApiKey,
		UserID:     merchant.UserID,
		Status:     merchant.Status,
	}
}

func (m *merchantRecordMapper) ToMerchantsRecord(merchants []models.Merchant) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantRecord(merchant))
	}
	return records
}
