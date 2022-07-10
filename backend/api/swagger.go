package api

import (
	"cafe/config"
	"cafe/docs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"net/url"
)

func (a *Api) setupSwagger() {
	if config.C.Swagger {
		docs.SwaggerInfo.Title = "Cafe"
		docs.SwaggerInfo.Description = "This is the backend of a cafe\n"
		docs.SwaggerInfo.Description += "When logged in successfully, the api will send a JWT Cookie for authorization"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.BasePath = "/api"
		parsed, _ := url.Parse(config.C.AllowedHosts[0])
		docs.SwaggerInfo.Host = parsed.Host

		// @contact.url https://github.com/flohoss/Cafe

		// @license.name MIT License
		// @license.url https://github.com/flohoss/Cafe/blob/main/LICENSE

		a.Router.GET("/swagger", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		logrus.WithField("url", config.C.AllowedHosts[0]+"/swagger").Info("Swagger running")
	}
}
