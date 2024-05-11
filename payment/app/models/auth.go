package models

import "time"

type Auth struct {
	Email    string
	Password string

	ExpiredToken int
	AccessToken  string
	RefreshToken string

	EmailSendedAt time.Time
}
