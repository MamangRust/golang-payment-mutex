package recordmapper

type RecordMapper struct {
	UserRecordMapper     UserRecordMapping
	SaldoRecordMapper    SaldoRecordMapping
	TopupRecordMapper    TopupRecordMapping
	TransferRecordMapper TransferRecordMapping
	WithdrawRecordMapper WithdrawRecordMapping
	CardRecordMapper     CardRecordMapping
	TransactionMapper    TransactionRecordMapping
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{
		UserRecordMapper:     NewUserRecordMapper(),
		SaldoRecordMapper:    NewSaldoRecordMapper(),
		TopupRecordMapper:    NewTopupRecordMapper(),
		TransferRecordMapper: NewTransferRecordMapper(),
		WithdrawRecordMapper: NewWithdrawRecordMapper(),
		CardRecordMapper:     NewCardRecordMapper(),
		TransactionMapper:    NewTransactionRecordMapper(),
	}
}
