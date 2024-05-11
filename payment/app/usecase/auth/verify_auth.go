package auth

import (
	"context"
	"paymentservice/app/domain"
	"paymentservice/app/models"
)

func (uc *auth) VerifyAuth(ctx context.Context, tokenHash string, verifyType string) (resp domain.Auth, err error) {

	result, err := uc.repoAuth.ConfirmationEmail(ctx, tokenHash, verifyType)
	if err != nil {
		return domain.Auth{}, err
	}

	data := domain.Auth{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiredToken: int(result.ExpiredToken),
	}

	_, err = uc.repoDB.UpdateUser(ctx, []string{"status"}, models.Users{
		Email: result.Email,
	})

	if err != nil {
		return domain.Auth{}, err
	}

	return data, nil
}
