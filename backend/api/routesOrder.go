package api

import (
	"cafe/config"
	"cafe/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Schemes
// @Summary get all orders
// @Description gets all orders as array
// @Tags orders
// @Produce json
// @Success 200 {array} service.Order
// @Failure 401 "Unauthorized"
// @Router /orders [get]
// @Security Cookie
func (a *Api) getOrders(c *gin.Context) {
	var o []service.Order
	config.C.Database.ORM.Find(&o)
	c.JSON(http.StatusOK, o)
}

// @Schemes
// @Summary get an order
// @Description gets a single order by id
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} service.Order
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Router /orders/{id} [get]
// @Security Cookie
func (a *Api) getOrder(c *gin.Context) {
	id := c.Param("id")
	exists, o := service.DoesOrderExist(id)
	if !exists {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, o)
	}
}

// @Schemes
// @Summary create new order
// @Description creates a new order and returns it
// @Tags orders
// @Accept json
// @Produce json
// @Param order body service.Order true "the order as an object"
// @Success 201 {object} service.Order "Order has been created"
// @Failure 400 {object} errorResponse "Missing order body"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot create order"
// @Router /orders [post]
// @Security Cookie
func (a *Api) createOrder(c *gin.Context) {
	var o service.Order
	err := c.ShouldBindJSON(&o)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing order body"})
		return
	}
	result := config.C.Database.ORM.Create(&o)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create order"})
	} else {
		c.JSON(http.StatusCreated, o)
	}
}

// @Schemes
// @Summary update an order
// @Description updates an order with provided information
// @Tags orders
// @Accept json
// @Produce json
// @Param order body service.Order true "the order with new values"
// @Success 200	{object} service.Order "Order has been updated"
// @Failure 400 {object} errorResponse "Missing order body"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot update order"
// @Router /orders [put]
// @Security Cookie
func (a *Api) updateOrder(c *gin.Context) {
	var updated service.Order
	err := c.ShouldBindJSON(&updated)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing order body"})
		return
	}
	id := updated.ID
	exists, old := service.DoesOrderExist(strconv.Itoa(int(id)))
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	result := config.C.Database.ORM.First(&old).Updates(&updated)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot update order"})
	} else {
		c.JSON(http.StatusOK, updated)
	}
}

// @Schemes
// @Summary delete an order
// @Description deletes an order from the database
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot delete order"
// @Router /orders/{id} [delete]
// @Security Cookie
func (a *Api) deleteOrder(c *gin.Context) {
	id := c.Param("id")
	exists, o := service.DoesOrderExist(id)
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}
	result := config.C.Database.ORM.Delete(&o)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete order"})
		return
	}
	c.Status(http.StatusOK)
}
