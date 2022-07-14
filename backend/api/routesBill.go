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
	c.JSON(http.StatusOK, service.GetAllBills())
}

// @Schemes
// @Summary create new bill
// @Description creates a new bill and returns it
// @Tags bills
// @Produce json
// @Param id path int true "Table ID"
// @Success 201 {object} service.Bill
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse
// @Router /bills/{id} [post]
// @Security Cookie
func (a *Api) createBill(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	table, err := service.DoesTableExist(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.CannotFind.String()})
		return
	}
	bill := service.CreateBill(&table)
	c.JSON(http.StatusCreated, bill)
}
