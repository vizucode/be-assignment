package useraccount

import (
	"context"
	"net/http"
	"user_service/app/domain"
	customerror "user_service/config/custom_error"
)

func (usecase *userAccount) GetUserAccount(ctx context.Context, email string) (resp []domain.UserAccount, err error) {

	result, err := usecase.repo.GetUser(ctx, email)
	if err != nil {
		return []domain.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	userAccountResult, err := usecase.repo.GetUserAccount(ctx, int(result.ID))
	if err != nil {
		return []domain.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	for _, v := range userAccountResult {
		resp = append(resp, domain.UserAccount{
			UserAccountId: v.ID,
			UserId:        v.UserId,
			AccountId:     v.AccountId,
			AccountName:   v.AccountName,
			AccountCode:   v.AccountCode,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return resp, nil
}
