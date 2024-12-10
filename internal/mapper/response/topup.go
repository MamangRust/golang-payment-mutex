package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type topupResponseMapper struct {
}

func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{}
}

func (s *topupResponseMapper) ToTopupResponse(topup record.TopupRecord) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          topup.TopupID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
	}
}

func (s *topupResponseMapper) ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponse(*response))
	}

	return responses
}
