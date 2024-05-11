package account

import "paymentservice/app/repositories"

type account struct {
	repo repositories.IRepoDatabase
}

func NewAccount(
	repo repositories.IRepoDatabase,
) *account {
	return &account{
		repo: repo,
	}
}
