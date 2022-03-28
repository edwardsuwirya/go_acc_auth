package usecase

import (
	"enigmacamp.com/goaccauth/model"
	"enigmacamp.com/goaccauth/repository"
	"enigmacamp.com/goaccauth/utils"
)

type AuthenticationUseCase interface {
	Login(userName string, userPassword string) error
}

type authenticationUseCase struct {
	repo repository.UserCredentialRepo
}

func (a *authenticationUseCase) Login(userName string, userPassword string) error {
	userAuth := model.NewUserCredential(userName, userPassword)
	if err := a.repo.GetByUserNameAndPassword(userAuth); err != nil {
		return utils.UnauthorizedError()
	}
	return nil
}

func NewAuthenticationUseCase(repo repository.UserCredentialRepo) AuthenticationUseCase {
	return &authenticationUseCase{
		repo: repo,
	}
}
