package auth

import "paymentservice/app/repositories"

type auth struct {
	repoAuth repositories.IRepoAuth
	repoDB   repositories.IRepoDatabase
}

func NewAuth(
	repoAuth repositories.IRepoAuth,
	repoDB repositories.IRepoDatabase,
) *auth {
	return &auth{
		repoAuth: repoAuth,
		repoDB:   repoDB,
	}
}
