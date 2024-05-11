package domain

import "time"

type Auth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=10"`

	ExpiredToken int    `json:"expired_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	EmailSendedAt time.Time `json:"email_sended_at"`
}
