package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type withdrawResponseMapper struct {
}

func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{}
}

func (s *withdrawResponseMapper) ToWithdrawResponse(withdraw record.WithdrawRecord) *response.WithdrawResponse {
	return &response.WithdrawResponse{
		ID:             withdraw.WithdrawID,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
	}
}

func (s *withdrawResponseMapper) ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse {
	var withdrawResponses []*response.WithdrawResponse
	for _, withdraw := range withdraws {
		withdrawResponses = append(withdrawResponses, s.ToWithdrawResponse(*withdraw))
	}
	return withdrawResponses
}
