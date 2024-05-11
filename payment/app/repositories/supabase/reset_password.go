package supabase

import (
	"context"
	"net/http"
	customerror "paymentservice/config/custom_error"
)

func (supa *supabase) ResetPassword(ctx context.Context, email string) (err error) {
	if err = supa.supaClient.Auth.ResetPasswordForEmail(ctx, email); err != nil {
		return customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}
