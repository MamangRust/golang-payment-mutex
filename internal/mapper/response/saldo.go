package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type saldoResponseMapper struct {
}

func NewSaldoResponseMapper() *saldoResponseMapper {
	return &saldoResponseMapper{}
}

func (s *saldoResponseMapper) ToSaldoResponse(saldo record.SaldoRecord) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             saldo.SaldoID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
	}
}

func (s *saldoResponseMapper) ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse {
	var responses []*response.SaldoResponse

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponse(*response))
	}

	return responses
}
