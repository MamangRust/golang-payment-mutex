package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type transferRecordMapper struct {
}

func NewTransferRecordMapper() *transferRecordMapper {
	return &transferRecordMapper{}
}

func (t *transferRecordMapper) ToTransferRecord(transfer models.Transfer) *record.TransferRecord {
	return &record.TransferRecord{
		TransferID:     transfer.TransferID,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
	}
}

func (t *transferRecordMapper) ToTransfersRecord(transfers []models.Transfer) []*record.TransferRecord {
	var responses []*record.TransferRecord

	for _, response := range transfers {
		responses = append(responses, t.ToTransferRecord(response))
	}

	return responses
}
