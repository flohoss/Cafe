package api

import (
	"cafe/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unjx-de/go-api-auth"
)

type Api struct {
	Router *gin.Engine
	Auth   *auth.Auth
}

type errorResponse struct {
	Error string `json:"error" validate:"required"`
}

type ErrorResponses uint

const (
	MissingInformation ErrorResponses = iota
	CannotCreate
	CannotDelete
	CannotFind
	StillInUse
)

func (e ErrorResponses) String() string {
	switch e {
	case MissingInformation:
		return "fehlende Informationen"
	case CannotCreate:
		return "kann nicht gespeichert werden"
	case CannotDelete:
		return "kann nicht gelöscht werden"
	case CannotFind:
		return "kann nicht gefunden werden"
	case StillInUse:
		return "der Artikel wird noch verwendet"
	default:
		return "unbekannt"
	}
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
