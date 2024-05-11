package auth

import "context"

func (a *auth) ResetPassword(ctx context.Context, email string) (err error) {
	return a.repoAuth.ResetPassword(ctx, email)
}
