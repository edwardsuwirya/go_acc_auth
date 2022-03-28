package repository

import "enigmacamp.com/goaccauth/model"

type UserCredentialRepo interface {
	GetByUserNameAndPassword(user model.UserCredential) error
}
