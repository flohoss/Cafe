package api

import (
	"cafe/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Schemes
// @Summary get all active tables
// @Description gets all active tables as array
// @Tags tables
// @Produce json
// @Success 200 {array} service.Table
// @Failure 401 "Unauthorized"
// @Router /tables [get]
// @Security Cookie
func (a *Api) getTables(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAllTables())
}

// @Schemes
// @Summary create new table
// @Description creates a new table and returns it
// @Tags tables
// @Accept json
// @Produce json
// @Success 201 {object} service.Table "Table has been created"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot create table"
// @Router /tables [post]
// @Security Cookie
func (a *Api) createTable(c *gin.Context) {
	table, err := service.CreateNewTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create table"})
	} else {
		c.JSON(http.StatusCreated, table)
	}
}

// @Schemes
// @Summary delete the latest table
// @Description deletes the latest table from the database
// @Tags tables
// @Produce json
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot delete table"
// @Router /tables [delete]
// @Security Cookie
func (a *Api) deleteTable(c *gin.Context) {
	err := service.DeleteLatestTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete table"})
	} else {
		c.Status(http.StatusOK)
	}
}
