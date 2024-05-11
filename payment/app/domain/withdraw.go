package domain

type Withdraw struct {
	FromAccountCode string  `json:"from_account_code" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	CurrencyType    string  `json:"currency_type" binding:"required"`
	OwnerEmail      string  `json:"omitempty"`
}
