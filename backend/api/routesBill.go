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
// @Param year query int true "year"
// @Param month query int true "month (1-12)"
// @Param day query int true "day (1-31)"
// @Success 200 {array} service.Bill
// @Failure 401 "Unauthorized"
// @Router /bills [get]
// @Security Cookie
func (a *Api) getBills(c *gin.Context) {
	year, presentYear := c.GetQuery("year")
	month, presentMonth := c.GetQuery("month")
	day, presentDay := c.GetQuery("day")
	if !presentYear || !presentMonth || !presentDay {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	bills, err := service.GetAllBills(year, month, day)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
	} else {
		c.JSON(http.StatusOK, bills)
	}
}

// @Schemes
// @Summary create new bill
// @Description creates a new bill and returns it
// @Tags bills
// @Produce json
// @Param table query int true "Table ID"
// @Success 201 {object} service.Bill
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse
// @Router /bills [post]
// @Security Cookie
func (a *Api) createBill(c *gin.Context) {
	id, present := c.GetQuery("table")
	if !present {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	orders := service.GetAllOrdersForTable(id)
	bill := service.CreateBill(orders)
	c.JSON(http.StatusCreated, bill)
}
