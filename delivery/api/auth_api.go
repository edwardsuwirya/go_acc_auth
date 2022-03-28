package api

import (
	"enigmacamp.com/goaccauth/delivery/appreq"
	"enigmacamp.com/goaccauth/usecase"
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	BaseApi
	authtenticationUseCase usecase.AuthenticationUseCase
}

func (p *AuthApi) userAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authReq appreq.AuthRequest
		err := p.ParseRequestBody(c, &authReq)
		if err != nil {
			p.FailedRequest(c, "api-userAuthenticate", "02", "Can not parse body")
			return
		}
		err = p.authtenticationUseCase.Login(authReq.UserName, authReq.Password)
		if err != nil {
			p.FailedUnauthorized(c, "api-userAuthenticate", "05", "Unauthorized")
			return
		}
		p.Success(c, "User", "OK!!")
	}
}
func NewAuthApi(authRoute *gin.RouterGroup, authenticationUseCase usecase.AuthenticationUseCase) {
	api := AuthApi{
		authtenticationUseCase: authenticationUseCase,
	}
	authRoute.POST("", api.userAuthenticate())
}
