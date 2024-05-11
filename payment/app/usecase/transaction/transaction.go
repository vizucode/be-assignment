package transaction

import "paymentservice/app/repositories"

type transactionst struct {
	repo repositories.IRepoDatabase
}

func NewTransaction(
	repo repositories.IRepoDatabase,
) *transactionst {
	return &transactionst{
		repo: repo,
	}
}
