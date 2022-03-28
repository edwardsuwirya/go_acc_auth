package manager

import "enigmacamp.com/goaccauth/usecase"

type UseCaseManager interface {
	UserLoginUseCase() usecase.AuthenticationUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserLoginUseCase() usecase.AuthenticationUseCase {
	return usecase.NewAuthenticationUseCase(u.repo.UserCredentialRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
