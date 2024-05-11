package domain

import "time"

type UserAccount struct {
	UserAccountId uint      `json:"user_account_id"`
	UserId        uint      `json:"user_id"`
	AccountId     uint      `json:"account_id" binding:"required"`
	AccountName   string    `json:"account_name" binding:"required"`
	AccountCode   string    `json:"account_code"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	HistoryTransaction []Transaction `json:"history_transaction"`
}
