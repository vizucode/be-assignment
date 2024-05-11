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
	GetUser(ctx context.Context, email string) (resp models.Users, err error)
	GetAllTypeAccount(ctx context.Context) (resp []models.Account, err error)
	FindTypeAccount(ctx context.Context, typeAccountId int) (resp models.Account, err error)
	GetUserAccount(ctx context.Context, userId int) (resp []models.UserAccount, err error)
	CreateUserAccount(ctx context.Context, userAccount models.UserAccount) (resp models.UserAccount, err error)
	GetAllTransaction(ctx context.Context, email int) (resp []models.Transaction, err error)
}
