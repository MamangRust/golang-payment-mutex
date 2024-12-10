package responseMapper

type ResponseMapper struct {
	CardResponseMapper        CardResponseMapper
	SaldoResponseMapper       SaldoResponseMapper
	TransactionResponseMapper TransactionResponseMapper
	TransferResponseMapper    TransferResponseMapper
	TopupResponseMapper       TopupResponseMapper
	WithdrawResponseMapper    WithdrawResponseMapper
	UserResponseMapper        UserResponseMapper
	MerchantResponseMapper    MerchantResponseMapper
}

func NewResponseMapper() *ResponseMapper {
	return &ResponseMapper{
		CardResponseMapper:        NewCardResponseMapper(),
		SaldoResponseMapper:       NewSaldoResponseMapper(),
		TransactionResponseMapper: NewTransactionResponseMapper(),
		TransferResponseMapper:    NewTransferResponseMapper(),
		TopupResponseMapper:       NewTopupResponseMapper(),
		WithdrawResponseMapper:    NewWithdrawResponseMapper(),
		UserResponseMapper:        NewUserResponseMapper(),
		MerchantResponseMapper:    NewMerchantResponseMapper(),
	}
}
