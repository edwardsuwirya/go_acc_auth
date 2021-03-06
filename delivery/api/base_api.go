package api

import (
	"enigmacamp.com/goaccauth/delivery/commonresp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) Success(c *gin.Context, message string, data interface{}) {
	commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", message, data))
}

func (b *BaseApi) FailedUnauthorized(c *gin.Context, serviceName, errCode string, message string) {
	e := commonresp.NewErrorMessage(http.StatusUnauthorized, serviceName, errCode, message)
	c.Error(fmt.Errorf("%s", e.ToJson()))
}

func (b *BaseApi) FailedRequest(c *gin.Context, serviceName, errCode string, message string) {
	e := commonresp.NewErrorMessage(http.StatusBadRequest, serviceName, errCode, message)
	c.Error(fmt.Errorf("%s", e.ToJson()))
}
func (b *BaseApi) Failed(c *gin.Context, serviceName, errCode string, message string) {
	e := commonresp.NewErrorMessage(http.StatusInternalServerError, serviceName, errCode, message)
	c.Error(fmt.Errorf("%s", e.ToJson()))
}
