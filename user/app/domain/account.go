package domain

import "time"

type Account struct {
	Id            uint      `json:"id"`
	AccountType   string    `json:"account_type"`
	LimitBallance float64   `json:"limit_ballance"`
	Status        bool      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
