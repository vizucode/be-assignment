package auth

import (
	"context"
	"paymentservice/app/domain"
	"paymentservice/app/models"
)

func (uc *auth) SignUp(ctx context.Context, creds domain.Users) (resp domain.Auth, err error) {

	_, err = uc.repoDB.CreateUser(ctx, models.Users{
		Email:     creds.Email,
		FirstName: creds.FirstName,
		LastName:  creds.LastName,
		Status:    false,
	})

	if err != nil {
		return domain.Auth{}, err
	}

	user, err := uc.repoAuth.SignUp(ctx, models.Auth{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		return domain.Auth{}, err
	}

	return domain.Auth{
		EmailSendedAt: user.EmailSendedAt,
	}, nil
}
