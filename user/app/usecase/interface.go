package usecase

import (
	"context"
	"user_service/app/domain"
)

type IAuth interface {
	SignUp(ctx context.Context, creds domain.Users) (resp domain.Auth, err error)
	SignIn(ctx context.Context, creds domain.Auth) (resp domain.Auth, err error)
	SignOut(ctx context.Context, accessToken string) (err error)
	ResetPassword(ctx context.Context, email string) (err error)
	VerifyAuth(ctx context.Context, tokenHash string, verifyType string) (resp domain.Auth, err error)
}

type IUserAccount interface {
	GetUserAccount(ctx context.Context, email string) (resp []domain.UserAccount, err error)
	CreateUserAccount(ctx context.Context, email string, userAccount domain.UserAccount) (resp domain.UserAccount, err error)
}

type IAccount interface {
	GetAllTypeAccount(ctx context.Context) (resp []domain.Account, err error)
}

type ITransaction interface {
	GetAllTransaction(ctx context.Context, email string) (resp []domain.Transaction, err error)
}
