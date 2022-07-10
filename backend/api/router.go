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
		tableGroup := api.Group("/tables")
		{
			tableGroup.Use(a.Auth.CookieAuthRequired())
			tableGroup.GET("", a.getTables)
			tableGroup.GET("/:id", a.getTable)
			tableGroup.POST("", a.createTable)
			tableGroup.PUT("", a.updateTable)
			tableGroup.DELETE("/:id", a.deleteTable)
		}
		orderGroup := api.Group("/orders")
		{
			orderGroup.Use(a.Auth.CookieAuthRequired())
			orderGroup.GET("", a.getOrders)
			orderGroup.GET("/:id", a.getOrder)
			orderGroup.POST("", a.createOrder)
			orderGroup.PUT("", a.updateOrder)
			orderGroup.DELETE("/:id", a.deleteOrder)
		}
	}

	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	logrus.WithField("amount", len(a.Router.Routes())).Debug("Routes initialized")
}
