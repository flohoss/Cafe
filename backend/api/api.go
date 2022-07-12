package api

import (
	"cafe/config"
	"cafe/hub"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unjx-de/go-api-auth"
)

type Api struct {
	Router *gin.Engine
	Hub    hub.Hub
	Auth   *auth.Auth
}

type errorResponse struct {
	Error string `json:"error" validate:"required"`
}

func (a *Api) Run() {
	a.Router = gin.New()
	a.setMiddlewares()
	a.handleStaticFiles()
	a.setupSwagger()
	a.setupRouter()
	logrus.WithField("port", config.C.Port).Info("Server running")
	err := a.Router.Run(fmt.Sprintf(":%d", config.C.Port))
	if err != nil {
		logrus.WithField("error", err).Fatal("Cannot start server")
	}
}
