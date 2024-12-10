package response

type OverviewData struct {
	TotalBalance   int            `json:"total_balance"`
	ActiveCards    int            `json:"active_cards"`
	TotalTransaksi int            `json:"total_transaksi"`
	TotalTopup     int            `json:"total_topup"`
	TopupAmount    int            `json:"topup_amount"`
	TotalWithdraw  int            `json:"total_withdraw"`
	WithdrawAmount int            `json:"withdraw_amount"`
	TotalTransfer  int            `json:"total_transfer"`
	ActivityTrends map[string]int `json:"activity_trends"`
}
