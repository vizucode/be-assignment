package supabase

import (
	"context"
	"net/http"
	"paymentservice/app/models"
	customerror "paymentservice/config/custom_error"

	SUPABASE "github.com/nedpals/supabase-go"
)

func (supa *supabase) SignIn(ctx context.Context, creds models.Auth) (resp models.Auth, err error) {

	user, err := supa.supaClient.Auth.SignIn(ctx, SUPABASE.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		return models.Auth{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return models.Auth{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}, nil
}
