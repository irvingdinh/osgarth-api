package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/irvingdinh/osgarth-api/src/component/config"
)

//go:generate mockery --name=PingHandler --case=snake
type PingHandler interface {
	Ping(c *gin.Context)
}

func NewPingHandler() PingHandler {
	return &pingHandlerImpl{}
}

type pingHandlerImpl struct {
	//
}

func (i *pingHandlerImpl) Ping(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, config.GetAppConfig().RedirectUrl)
}
