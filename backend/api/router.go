package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *Api) setupRouter() {
	api := a.Router.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", a.login)
			authGroup.POST("/logout", a.logout)
		}
	}

	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	logrus.WithField("amount", len(a.Router.Routes())).Debug("Routes initialized")
}
