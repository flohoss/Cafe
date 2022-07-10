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
// @Summary get a table
// @Description gets a single table by id
// @Tags tables
// @Produce json
// @Param id path int true "Table ID"
// @Success 200 {object} service.Table
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Router /tables/{id} [get]
// @Security Cookie
func (a *Api) getTable(c *gin.Context) {
	id := c.Param("id")
	table, err := service.DoesTableExist(id)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
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
// @Summary update a table
// @Description updates a table with provided information
// @Tags tables
// @Accept json
// @Produce json
// @Param id path int true "Table ID"
// @Param count query int true "new order count"
// @Param total query float64 true "new total amount"
// @Success 200	{object} service.Table "Table has been updated"
// @Failure 400 {object} errorResponse "Missing query parameter"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot update table"
// @Router /tables/{id} [put]
// @Security Cookie
func (a *Api) updateTable(c *gin.Context) {
	id := c.Param("id")
	table, err := service.DoesTableExist(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = service.UpdateTable(table, getQueryAsFloat64(c, "total"), getQueryAsUint64(c, "count"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot update table"})
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// @Schemes
// @Summary delete a table
// @Description deletes a table from the database
// @Tags tables
// @Produce json
// @Param id path int true "Table ID"
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot delete table"
// @Router /tables/{id} [delete]
// @Security Cookie
func (a *Api) deleteTable(c *gin.Context) {
	id := c.Param("id")
	table, err := service.DoesTableExist(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = service.DeleteTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete table"})
	} else {
		c.Status(http.StatusOK)
	}
}
