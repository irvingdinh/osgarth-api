package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/irvingdinh/osgarth-api/src/component/config"
	"github.com/irvingdinh/osgarth-api/src/component/logger"
	"github.com/irvingdinh/osgarth-api/src/http/handler"
	"github.com/irvingdinh/osgarth-api/src/http/middleware"
)

//go:generate mockery --name=Server --case=snake
type Server interface {
	Start() error
}

func New(handler handler.Handler) Server {
	server := &serverImpl{
		handler: handler,
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	handler handler.Handler

	router *gin.Engine
}

func (i *serverImpl) Start() error {
	log := logger.CToL(context.Background(), "server.Start")
	log.Infof("Listening and serving HTTP on :%d", config.GetHTTPConfig().Port)

	return i.router.Run(fmt.Sprintf(":%d", config.GetHTTPConfig().Port))
}

func (i *serverImpl) withRouter() {
	if config.GetAppConfig().Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(middleware.Logger())

	router.GET("/", i.handler.PingHandler().Ping)
	router.GET("/:slug", i.handler.ItemHandler().GetItem)

	i.router = router
}
