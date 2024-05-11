package useraccount

import "paymentservice/app/repositories"

type userAccount struct {
	repo repositories.IRepoDatabase
}

func NewUserAccount(
	repo repositories.IRepoDatabase,
) *userAccount {
	return &userAccount{
		repo: repo,
	}
}
