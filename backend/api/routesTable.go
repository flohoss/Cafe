package api

import (
	"cafe/config"
	"cafe/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Schemes
// @Summary get all tables
// @Description gets all tables as array
// @Tags tables
// @Produce json
// @Success 200 {array} service.Table
// @Failure 401 "Unauthorized"
// @Router /tables [get]
// @Security Cookie
func (a *Api) getTables(c *gin.Context) {
	var o []service.Table
	config.C.Database.ORM.Find(&o)
	c.JSON(http.StatusOK, o)
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
	exists, o := service.DoesExist(id, service.Table{})
	if !exists {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, o)
	}
}

// @Schemes
// @Summary create new table
// @Description creates a new table and returns it
// @Tags tables
// @Accept json
// @Produce json
// @Param table body service.Table true "the table as an object"
// @Success 201 {object} service.Table "Table has been created"
// @Failure 400 {object} errorResponse "Missing table body"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot create table"
// @Router /tables [post]
// @Security Cookie
func (a *Api) createTable(c *gin.Context) {
	var o service.Table
	err := c.ShouldBindJSON(&o)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing table body"})
		return
	}
	result := config.C.Database.ORM.Create(&o)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create table"})
	} else {
		c.JSON(http.StatusCreated, o)
	}
}

// @Schemes
// @Summary update a table
// @Description updates a table with provided information
// @Tags tables
// @Accept json
// @Produce json
// @Param table body service.Table true "the table with new values"
// @Success 200	{object} service.Table "Table has been updated"
// @Failure 400 {object} errorResponse "Missing table body"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot update table"
// @Router /tables [put]
// @Security Cookie
func (a *Api) updateTable(c *gin.Context) {
	var updated service.Table
	err := c.ShouldBindJSON(&updated)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing table body"})
		return
	}
	id := updated.ID
	exists, old := service.DoesExist(strconv.Itoa(int(id)), service.Table{})
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	result := config.C.Database.ORM.First(&old).Updates(&updated)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot update table"})
	} else {
		c.JSON(http.StatusOK, updated)
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
	exists, o := service.DoesExist(id, service.Table{})
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	result := config.C.Database.ORM.Delete(&o)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete table"})
		return
	}
	c.Status(http.StatusOK)
}
