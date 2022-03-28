package manager

import (
	"enigmacamp.com/goaccauth/repository"
)

type RepoManager interface {
	UserCredentialRepo() repository.UserCredentialRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) UserCredentialRepo() repository.UserCredentialRepo {
	return repository.NewUserCredentialRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
