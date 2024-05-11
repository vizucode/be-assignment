package supabase

import (
	"context"
	"user_service/app/models"
)

func (supa *supabase) RefreshToken(ctx context.Context, accessToken string, refreshToken string) (resp models.Auth, err error) {
	result, err := supa.supaClient.Auth.RefreshUser(ctx, accessToken, refreshToken)
	if err != nil {
		return models.Auth{}, err
	}

	return models.Auth{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}, nil
}
