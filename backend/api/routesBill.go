package api

import (
	"cafe/config"
	"cafe/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
// @Param filter query string false "filter"
// @Success 201 {object} service.Bill
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse
// @Router /bills [post]
// @Security Cookie
func (a *Api) createBill(c *gin.Context) {
	table, tableErr := strconv.ParseUint(c.Query("table"), 10, 64)
	if tableErr != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	stringFiler, filterPresent := c.GetQuery("filter")
	var filter []string
	if filterPresent {
		filter = strings.Split(stringFiler, ",")
	}
	bill := service.CreateBill(service.GetOrderOptions{TableId: table, Grouped: true, Filter: filter})
	c.JSON(http.StatusCreated, bill)
}
