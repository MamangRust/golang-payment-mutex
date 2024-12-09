package recordmapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/models"
)

type UserRecordMapping interface {
	ToUserRecord(user models.User) *record.UserRecord
	ToUsersRecord(users []models.User) []*record.UserRecord
}

type SaldoRecordMapping interface {
	ToSaldoRecord(saldo models.Saldo) *record.SaldoRecord
	ToSaldosRecord(saldos []models.Saldo) []*record.SaldoRecord
}

type TopupRecordMapping interface {
	ToTopupRecord(topup models.Topup) *record.TopupRecord
	ToTopupRecords(topups []models.Topup) []*record.TopupRecord
}

type TransferRecordMapping interface {
	ToTransferRecord(transfer models.Transfer) *record.TransferRecord
	ToTransfersRecord(transfers []models.Transfer) []*record.TransferRecord
}

type WithdrawRecordMapping interface {
	ToWithdrawRecord(withdraw models.Withdraw) *record.WithdrawRecord
	ToWithdrawsRecord(withdraws []models.Withdraw) []*record.WithdrawRecord
}

type CardRecordMapping interface {
	ToCardRecord(card models.Card) *record.CardRecord
	ToCardsRecord(cards []models.Card) []*record.CardRecord
}

type TransactionRecordMapping interface {
	ToTransactionRecord(transaction models.Transaction) *record.TransactionRecord
	ToTransactionsRecord(transactions []models.Transaction) []*record.TransactionRecord
}
