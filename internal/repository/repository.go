package repository

type Repositories struct {
	User     UserRepository
	Saldo    SaldoRepository
	Topup    TopupRepository
	Transfer TransferRepository
	Withdraw WithdrawRepository
}

func NewRepositorys() *Repositories {
	return &Repositories{
		User:     NewUserRepository(),
		Saldo:    NewSaldoRepository(),
		Topup:    NewTopupRepository(),
		Transfer: NewTransferRepository(),
		Withdraw: NewWithdrawRepository(),
	}
}
