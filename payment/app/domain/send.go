package domain

type Send struct {
	TargetAccountCode string  `json:"target_account_code" binding:"required"`
	FromAccountCode   string  `json:"from_account_code" binding:"required"`
	AmountTransfer    float64 `json:"amount_transfer" binding:"required"`
	CurrencyType      string  `json:"currency_type" binding:"required"`
	OwnerEmail        string  `json:"omitempty"`
}
