package account

import "user_service/app/repositories"

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
