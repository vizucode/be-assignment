package transaction

import (
	"context"
	"paymentservice/app/domain"
	"paymentservice/app/models"
	"time"
)

func (uc *transactionst) Withdraw(ctx context.Context, withdraw domain.Withdraw) (resp domain.Transaction, err error) {

	// check owner
	userM, err := uc.repo.GetUser(ctx, withdraw.OwnerEmail)
	if err != nil {
		return domain.Transaction{}, err
	}

	// verify owner account is valid
	OwnerAccount, err := uc.repo.CheckOwnerAccount(ctx, int(userM.ID), withdraw.FromAccountCode)
	if err != nil {
		return domain.Transaction{}, err
	}

	time.Sleep(5 * time.Second)

	_, err = uc.repo.CreateTransaction(ctx, models.Transaction{
		UserId:            int(userM.ID),
		TargetAccountCode: OwnerAccount.AccountCode,
		AccountCodeOwner:  OwnerAccount.AccountCode,
		Amount:            withdraw.Amount,
		CurrencyType:      withdraw.CurrencyType,
		IncomSpending:     "spending",
		TransactionType:   "withdraw",
		Status:            true,
	})
	if err != nil {
		return domain.Transaction{}, err
	}

	resp = domain.Transaction{
		Amount:            withdraw.Amount,
		CurrencyType:      withdraw.CurrencyType,
		TargetAccountCode: OwnerAccount.AccountCode,
		AccountCode:       OwnerAccount.AccountCode,
		AccountNameOwner:  OwnerAccount.User.FirstName,
		UserId:            int(userM.ID),
		Status:            true,
		TransactionType:   "withdraw",
		IncomeSpending:    "spending",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	return
}
