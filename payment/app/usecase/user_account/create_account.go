package useraccount

import (
	"context"
	"net/http"
	"paymentservice/app/domain"
	"paymentservice/app/models"
	customerror "paymentservice/config/custom_error"
	"strconv"
	"time"
)

func (usecase *userAccount) CreateUserAccount(ctx context.Context, email string, userAccount domain.UserAccount) (resp domain.UserAccount, err error) {

	// verify user
	user, err := usecase.repo.GetUser(ctx, email)
	if err != nil {
		return domain.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	accountCode := "00" + strconv.Itoa(int(time.Now().Unix()))

	// verify account
	accountResult, err := usecase.repo.FindTypeAccount(ctx, int(userAccount.AccountId))
	if err != nil {
		return domain.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	result, err := usecase.repo.CreateUserAccount(ctx, models.UserAccount{
		UserId:      user.ID,
		AccountId:   accountResult.ID,
		AccountName: userAccount.AccountName,
		AccountCode: accountCode,
	})

	if err != nil {
		return domain.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	resp = domain.UserAccount{
		UserAccountId: result.ID,
		UserId:        result.UserId,
		AccountId:     result.AccountId,
		AccountName:   result.AccountName,
		AccountCode:   result.AccountCode,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}

	return resp, nil
}
