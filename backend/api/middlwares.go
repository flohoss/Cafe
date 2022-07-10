package api

import (
	"cafe/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func myLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if logrus.GetLevel() != logrus.TraceLevel {
			return
		}
		reqUri := c.Request.RequestURI
		if strings.Contains(reqUri, "/storage") {
			return
		}
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		logrus.WithFields(logrus.Fields{
			"status":  http.StatusText(c.Writer.Status()),
			"latency": latencyTime,
			"client":  c.ClientIP(),
			"method":  c.Request.Method,
		}).Trace(reqUri)
	}
}

func (a *Api) setMiddlewares() {
	allowedHeaders := []string{"content-type", "password"}
	a.Router.Use(myLogger())
	a.Router.Use(gin.Recovery())
	a.Router.Use(cors.New(cors.Config{
		AllowOrigins:     config.C.AllowedHosts,
		AllowCredentials: true,
		AllowHeaders:     allowedHeaders,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	}))
	_ = a.Router.SetTrustedProxies(config.C.AllowedHosts)
	a.Router.MaxMultipartMemory = 8 << 20 // 8 MiB
	logrus.WithFields(logrus.Fields{
		"allowedOrigins": config.C.AllowedHosts,
	}).Debug("Middlewares set")
}
