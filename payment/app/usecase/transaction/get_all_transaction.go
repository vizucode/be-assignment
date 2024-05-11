package transaction

import (
	"context"
	"net/http"
	"paymentservice/app/domain"
	customerror "paymentservice/config/custom_error"
)

func (uc *transactionst) GetAllTransaction(ctx context.Context, email string) (resp []domain.Transaction, err error) {

	user, err := uc.repo.GetUser(ctx, email)
	if err != nil {
		return []domain.Transaction{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	transactionResult, err := uc.repo.GetAllTransaction(ctx, int(user.ID))
	if err != nil {
		return []domain.Transaction{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	for _, transaction := range transactionResult {
		resp = append(resp, domain.Transaction{
			Id:                 transaction.ID,
			UserId:             transaction.UserId,
			AccountCode:        transaction.AccountCodeOwner,
			TargetAccountCode:  transaction.TargetAccountCode,
			TransactionType:    transaction.TransactionType,
			IncomeSpending:     transaction.IncomSpending,
			CurrencyType:       transaction.CurrencyType,
			Amount:             transaction.Amount,
			Status:             transaction.Status,
			TargetAccountOwner: transaction.AccountTarget.User.FirstName,
			AccountNameOwner:   transaction.AccountOwner.User.FirstName,
			CreatedAt:          transaction.CreatedAt,
			UpdatedAt:          transaction.UpdatedAt,
		})
	}

	return resp, nil
}
