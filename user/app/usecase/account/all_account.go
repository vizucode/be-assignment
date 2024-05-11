package account

import (
	"context"
	"net/http"
	"user_service/app/domain"
	customerror "user_service/config/custom_error"
)

func (usecase *account) GetAllTypeAccount(ctx context.Context) (resp []domain.Account, err error) {

	result, err := usecase.repo.GetAllTypeAccount(ctx)
	if err != nil {
		return []domain.Account{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	for _, v := range result {
		resp = append(resp, domain.Account{
			Id:            uint(v.ID),
			LimitBallance: v.LimitBallance,
			AccountType:   v.AccountType,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return resp, nil
}
