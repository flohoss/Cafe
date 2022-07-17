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
			tableGroup.POST("", a.createTable)
			tableGroup.DELETE("", a.deleteTable)
		}
		orderGroup := api.Group("/orders")
		{
			orderGroup.Use(a.Auth.CookieAuthRequired())
			orderGroup.GET("", a.getOrders)
			orderGroup.POST("", a.createOrder)
			orderGroup.DELETE("", a.deleteOrder)
			orderGroup.PUT("", a.updateOrder)
			orderGroup.GET("/ws", a.serveWs)
			orderItemGroup := orderGroup.Group("/items")
			{
				orderItemGroup.GET("", a.getOrderItems)
				orderItemGroup.POST("", a.createOrderItem)
				orderItemGroup.PUT("", a.updateOrderItem)
				orderItemGroup.DELETE("/:id", a.deleteOrderItem)
			}
		}
		billGroup := api.Group("/bills")
		{
			billGroup.Use(a.Auth.CookieAuthRequired())
			billGroup.GET("", a.getBills)
			billGroup.POST("", a.createBill)
			billGroup.DELETE("/:id", a.deleteBill)
			billItemGroup := billGroup.Group("/items")
			{
				billItemGroup.GET("", a.getBillItems)
			}
		}
	}

	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	logrus.WithField("amount", len(a.Router.Routes())).Debug("Routes initialized")
}
