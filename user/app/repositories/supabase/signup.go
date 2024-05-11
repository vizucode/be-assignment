package supabase

import (
	"context"
	"net/http"
	"user_service/app/models"
	customerror "user_service/config/custom_error"

	SUPABASE "github.com/nedpals/supabase-go"
)

func (supa *supabase) SignUp(ctx context.Context, creds models.Auth) (resp models.Auth, err error) {

	user, err := supa.supaClient.Auth.SignUp(ctx, SUPABASE.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		return models.Auth{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	resp = models.Auth{
		EmailSendedAt: user.ConfirmationSentAt,
	}

	return
}
