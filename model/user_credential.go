package model

type UserCredential struct {
	userName     string
	userPassword string
}

func (p *UserCredential) GetUserName() string {
	return p.userName
}
func (p *UserCredential) GetUserPassword() string {
	return p.userPassword
}

func NewUserCredential(userName, userPassword string) UserCredential {
	return UserCredential{
		userName:     userName,
		userPassword: userPassword,
	}
}
