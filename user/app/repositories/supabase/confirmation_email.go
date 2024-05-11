package supabase

import (
	"context"
	"net/http"
	"user_service/app/models"
	customerror "user_service/config/custom_error"

	SUPABASE "github.com/nedpals/supabase-go"
)

func (supa *supabase) ConfirmationEmail(ctx context.Context, tokenHash string, verifyType string) (resp models.Auth, err error) {

	verifiedUser, err := supa.supaClient.Auth.ExchangeCode(ctx, SUPABASE.ExchangeCodeOpts{
		AuthCode:     tokenHash,
		CodeVerifier: verifyType,
	})

	if err != nil {
		return models.Auth{}, customerror.NewCustomError(err.Error(), http.StatusUnauthorized)
	}

	resp = models.Auth{
		AccessToken:  verifiedUser.AccessToken,
		RefreshToken: verifiedUser.RefreshToken,
		Email:        verifiedUser.User.Email,
		ExpiredToken: int(verifiedUser.ExpiresIn),
	}

	return resp, nil
}
