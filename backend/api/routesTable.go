package api

import (
	"cafe/config"
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
	var tables []service.Table
	config.C.Database.ORM.Find(&tables)
	c.JSON(http.StatusOK, tables)
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
	exists, table := service.DoesTableExist(id)
	if !exists {
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
	var obj service.Table
	result := config.C.Database.ORM.Unscoped().Where("is_deleted = ?", 1).Limit(1).Find(&obj)
	if result.RowsAffected == 0 {
		result = config.C.Database.ORM.Create(&obj)
	} else {
		obj.IsDeleted = 0
		config.C.Database.ORM.Save(&obj)
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create table"})
	} else {
		c.JSON(http.StatusCreated, obj)
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
	exists, table := service.DoesTableExist(id)
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	table.Total = getQueryAsFloat64(c, "total")
	table.OrderCount = getQueryAsUint64(c, "count")
	result := config.C.Database.ORM.Save(&table)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot update table"})
	} else {
		c.JSON(http.StatusOK, table)
	}
}

// @Schemes
// @Summary set a table inactive
// @Description sets a table inactive in the database
// @Tags tables
// @Produce json
// @Param id path int true "Table ID"
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot set table inactive"
// @Router /tables/{id} [delete]
// @Security Cookie
func (a *Api) deleteTable(c *gin.Context) {
	id := c.Param("id")
	exists, old := service.DoesTableExist(id)
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	old.Total = 0
	old.OrderCount = 0
	result := config.C.Database.ORM.Save(&old)
	config.C.Database.ORM.Delete(&old)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot set table inactive"})
		return
	}
	c.Status(http.StatusOK)
}
