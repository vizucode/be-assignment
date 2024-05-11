package transaction

import "user_service/app/repositories"

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
