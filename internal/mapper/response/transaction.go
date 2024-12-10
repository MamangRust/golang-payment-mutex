package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type transactionResponseMapper struct {
}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (s *transactionResponseMapper) ToTransactionResponse(transaction record.TransactionRecord) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              transaction.TransactionID,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
	}
}

func (s *transactionResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse {
	responses := make([]*response.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponse(*transaction))
	}
	return responses
}
