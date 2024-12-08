package recordmapper

type RecordMapper struct {
	UserRecordMapper     UserRecordMapping
	SaldoRecordMapper    SaldoRecordMapping
	TopupRecordMapper    TopupRecordMapping
	TransferRecordMapper TransferRecordMapping
	WithdrawRecordMapper WithdrawRecordMapping
}

func NewRecordMapper() *RecordMapper {
	return &RecordMapper{
		UserRecordMapper:     NewUserRecordMapper(),
		SaldoRecordMapper:    NewSaldoRecordMapper(),
		TopupRecordMapper:    NewTopupRecordMapper(),
		TransferRecordMapper: NewTransferRecordMapper(),
		WithdrawRecordMapper: NewWithdrawRecordMapper(),
	}
}
