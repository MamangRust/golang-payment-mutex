package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type cardRecordMapper struct {
}

func NewCardRecordMapper() *cardRecordMapper {
	return &cardRecordMapper{}
}

func (s *cardRecordMapper) ToCardRecord(card models.Card) *record.CardRecord {
	return &record.CardRecord{
		CardID:       card.CardID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
	}
}

func (s *cardRecordMapper) ToCardsRecord(cards []models.Card) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecord(card))
	}
	return records
}
