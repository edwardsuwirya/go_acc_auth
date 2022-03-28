package repository

import (
	"enigmacamp.com/goaccauth/logger"
	"enigmacamp.com/goaccauth/model"
	"enigmacamp.com/goaccauth/utils"
	"errors"
	"github.com/jmoiron/sqlx"
)

type userCredentialRepoImpl struct {
	userCredDb *sqlx.DB
}

func (u *userCredentialRepoImpl) GetByUserNameAndPassword(user model.UserCredential) error {
	var isUserExist int
	err := u.userCredDb.Get(&isUserExist, "selectsssss count(id) from user_credentials where user_name=$1 and user_password=$2 and is_blocked=$3", user.GetUserName(), user.GetUserPassword(), false)
	if err != nil {
		logger.Log.Error().Err(err).Str("service", "db-usercred-select").Msg("Query user credentials failed")
		return errors.New(err.Error())
	}
	if isUserExist == 0 {
		return utils.DataNotFoundError()
	}
	return nil
}

func NewUserCredentialRepo(db *sqlx.DB) UserCredentialRepo {
	return &userCredentialRepoImpl{
		db,
	}
}
