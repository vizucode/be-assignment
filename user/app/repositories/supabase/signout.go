package supabase

import (
	"context"
	"net/http"
	customerror "user_service/config/custom_error"
)

func (supa *supabase) SignOut(ctx context.Context, accessToken string) (err error) {
	if err = supa.supaClient.Auth.SignOut(ctx, accessToken); err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}
