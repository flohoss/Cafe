package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// @Schemes
// @Summary login client
// @Description logs in client with a password
// @Tags authorization
// @Produce json
// @Param password header string true "Password to authorize"
// @Success 200 "OK"
// @Failure 400 {object} errorResponse "Missing credentials"
// @Failure 401 "Unauthorized"
// @Router /auth/login [post]
func (a *Api) login(c *gin.Context) {
	password := strings.Trim(c.GetHeader("password"), " ")
	if password == "" {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing credentials"})
	} else if !a.Auth.PasswordEquals(password) {
		c.Status(http.StatusUnauthorized)
		logrus.WithField("client", c.ClientIP()).Warning("Authentication failed")
	} else {
		a.Auth.SetSessionCookie(c)
		c.Status(http.StatusOK)
	}
}

// @Schemes
// @Summary logout client
// @Description logs the client out
// @Tags authorization
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Router /auth/logout [post]
// @Security Cookie
func (a *Api) logout(c *gin.Context) {
	if a.Auth.CookieAuthIsValid(c) {
		a.Auth.DeleteSessionCookie(c)
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}
