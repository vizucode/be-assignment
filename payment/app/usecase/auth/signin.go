package auth

import (
	"context"
	"paymentservice/app/domain"
	"paymentservice/app/models"
)

func (uc *auth) SignIn(ctx context.Context, creds domain.Auth) (resp domain.Auth, err error) {

	result, err := uc.repoAuth.SignIn(ctx, models.Auth{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		return domain.Auth{}, err
	}

	_, err = uc.repoDB.UpdateUser(ctx, []string{"status"}, models.Users{
		Email:  creds.Email,
		Status: true,
	})

	if err != nil {
		return domain.Auth{}, err
	}

	return domain.Auth{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiredToken: int(result.ExpiredToken),
	}, nil
}
