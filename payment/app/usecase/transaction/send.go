package transaction

import (
	"context"
	"paymentservice/app/domain"
	"paymentservice/app/models"
	"time"
)

func (uc *transactionst) Send(ctx context.Context, send domain.Send) (resp domain.Transaction, err error) {
	// check owner
	userM, err := uc.repo.GetUser(ctx, send.OwnerEmail)
	if err != nil {
		return domain.Transaction{}, err
	}

	// verify owner account is valid
	// check owner account
	OwnerAccount, err := uc.repo.CheckOwnerAccount(ctx, int(userM.ID), send.FromAccountCode)
	if err != nil {
		return domain.Transaction{}, err
	}

	// check target account
	targetAccount, err := uc.repo.FindUserAccount(ctx, send.TargetAccountCode)
	if err != nil {
		return domain.Transaction{}, err
	}

	time.Sleep(5 * time.Second)
	_, err = uc.repo.CreateTransaction(ctx, models.Transaction{
		Amount:            send.AmountTransfer,
		CurrencyType:      send.CurrencyType,
		TargetAccountCode: targetAccount.AccountCode,
		AccountCodeOwner:  OwnerAccount.AccountCode,
		UserId:            int(userM.ID),
		IncomSpending:     "spending",
		TransactionType:   "send",
		Status:            true,
	})
	if err != nil {
		return domain.Transaction{}, err
	}

	_, err = uc.repo.CreateTransaction(ctx, models.Transaction{
		Amount:            send.AmountTransfer,
		CurrencyType:      send.CurrencyType,
		TargetAccountCode: OwnerAccount.AccountCode,
		AccountCodeOwner:  targetAccount.AccountCode,
		UserId:            int(userM.ID),
		IncomSpending:     "income",
		TransactionType:   "send",
		Status:            true,
	})
	if err != nil {
		return domain.Transaction{}, err
	}

	resp = domain.Transaction{
		Amount:             send.AmountTransfer,
		CurrencyType:       send.CurrencyType,
		TargetAccountCode:  targetAccount.AccountCode,
		AccountCode:        OwnerAccount.AccountCode,
		TargetAccountOwner: targetAccount.User.FirstName,
		AccountNameOwner:   OwnerAccount.User.FirstName,
		UserId:             int(userM.ID),
		Status:             true,
		TransactionType:    "send",
		IncomeSpending:     "spending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return resp, nil
}
