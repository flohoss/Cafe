package api

import (
	"cafe/config"
	"cafe/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Schemes
// @Summary get all bills
// @Description gets all bills as array
// @Tags bills
// @Produce json
// @Success 200 {array} service.Bill
// @Failure 401 "Unauthorized"
// @Router /bills [get]
// @Security Cookie
func (a *Api) getBills(c *gin.Context) {
	var bills []service.Bill
	config.C.Database.ORM.Find(&bills)
	c.JSON(http.StatusOK, bills)
}

// @Schemes
// @Summary create new bill
// @Description creates a new bill and returns it
// @Tags bills
// @Produce json
// @Param id path int true "Table ID"
// @Success 201 {object} service.Bill "Bill has been created"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot create bill"
// @Router /bills/{id} [post]
// @Security Cookie
func (a *Api) createBill(c *gin.Context) {
	id := c.Param("id")
	exists, table := service.DoesTableExist(id)
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	var bill service.Bill
	bill.TableID = table.ID
	bill.Total = table.Total
	config.C.Database.ORM.Create(&bill)
	c.JSON(http.StatusCreated, bill)
}
