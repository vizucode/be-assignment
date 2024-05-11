package auth

import "context"

func (uc *auth) SignOut(ctx context.Context, accessToken string) (err error) {
	return uc.repoAuth.SignOut(ctx, accessToken)
}
