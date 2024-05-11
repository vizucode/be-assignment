package repositories

import (
	"context"
	"user_service/app/models"
)

type IRepoAuth interface {
	SignUp(ctx context.Context, creds models.Auth) (resp models.Auth, err error)
	SignIn(ctx context.Context, creds models.Auth) (resp models.Auth, err error)
	SignOut(ctx context.Context, accessToken string) (err error)
	ResetPassword(ctx context.Context, email string) (err error)
	ConfirmationEmail(ctx context.Context, tokenHash string, verifyType string) (resp models.Auth, err error)
	RefreshToken(ctx context.Context, accessToken string, refreshToken string) (resp models.Auth, err error)
}

type IRepoDatabase interface {
	CreateUser(ctx context.Context, user models.Users) (resp models.Users, err error)
	UpdateUser(ctx context.Context, selectedField []string, user models.Users) (resp models.Users, err error)
}
