package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type cardResponseMapper struct {
}

func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{}
}

func (s *cardResponseMapper) ToCardResponse(card record.CardRecord) *response.CardResponse {
	return &response.CardResponse{
		ID:           card.CardID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
	}
}

func (s *cardResponseMapper) ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse {
	var cardResponses []*response.CardResponse

	for _, card := range cards {
		cardResponses = append(cardResponses, s.ToCardResponse(*card))
	}

	return cardResponses
}
