package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type CardResponseMapper interface {
	ToCardResponse(card record.CardRecord) *response.CardResponse
	ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse
}

type UserResponseMapper interface {
	ToUserResponse(user record.UserRecord) *response.UserResponse
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse
}

type SaldoResponseMapper interface {
	ToSaldoResponse(saldo record.SaldoRecord) *response.SaldoResponse
	ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse
}

type TopupResponseMapper interface {
	ToTopupResponse(topup record.TopupRecord) *response.TopupResponse
	ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse
}

type TransactionResponseMapper interface {
	ToTransactionResponse(transaction record.TransactionRecord) *response.TransactionResponse
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse
}

type TransferResponseMapper interface {
	ToTransferResponse(transfer record.TransferRecord) *response.TransferResponse
	ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse
}

type WithdrawResponseMapper interface {
	ToWithdrawResponse(withdraw record.WithdrawRecord) *response.WithdrawResponse
	ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse
}

type MerchantResponseMapper interface {
	ToMerchantResponse(merchant record.MerchantRecord) *response.MerchantResponse
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse
}
