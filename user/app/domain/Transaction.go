package domain

import "time"

type Transaction struct {
	Id                uint    `json:"id"`
	UserId            int     `json:"user_id"`
	AccountCode       string  `json:"account_code"`
	TargetAccountCode string  `json:"target_account_code"`
	TransactionType   string  `json:"transaction_type"`
	IncomeSpending    string  `json:"income_spending"`
	CurrencyType      string  `json:"currency_type"`
	Amount            float64 `json:"amount"`
	Status            bool    `json:"status"`

	TargetAccountOwner string `json:"target_account_owner"`
	AccountNameOwner   string `json:"account_owner"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
