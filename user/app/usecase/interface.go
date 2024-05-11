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
