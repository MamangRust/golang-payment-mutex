package repository

import recordmapper "payment-mutex/internal/mapper/record"

type Repositories struct {
	User        UserRepository
	Saldo       SaldoRepository
	Topup       TopupRepository
	Transfer    TransferRepository
	Withdraw    WithdrawRepository
	Card        CardRepository
	Transaction TransactionRepository
	Merchant    MerchantRepository
}

type Deps struct {
	MapperRecord recordmapper.RecordMapper
}

func NewRepositorys(deps Deps) *Repositories {
	return &Repositories{
		User:        NewUserRepository(deps.MapperRecord.UserRecordMapper),
		Saldo:       NewSaldoRepository(deps.MapperRecord.SaldoRecordMapper),
		Topup:       NewTopupRepository(deps.MapperRecord.TopupRecordMapper),
		Transfer:    NewTransferRepository(deps.MapperRecord.TransferRecordMapper),
		Withdraw:    NewWithdrawRepository(deps.MapperRecord.WithdrawRecordMapper),
		Card:        NewCardRepository(deps.MapperRecord.CardRecordMapper),
		Transaction: NewTransactionRepository(deps.MapperRecord.TransactionRecordMapper),
		Merchant:    NewMerchantRepository(deps.MapperRecord.MerchantRecordMapper),
	}
}
