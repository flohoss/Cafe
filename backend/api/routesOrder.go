package api

import (
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
// @Param table query int true "Table ID"
// @Success 200 {array} service.Order
// @Failure 401 "Unauthorized"
// @Router /orders [get]
// @Security Cookie
func (a *Api) getOrders(c *gin.Context) {
	table := c.Query("table")
	c.JSON(http.StatusOK, service.GetAllOrders(table))
}

// @Schemes
// @Summary create new order
// @Description creates a new order and returns it
// @Tags orders
// @Accept json
// @Produce json
// @Param item query int true "OrderItem ID"
// @Param table query int true "Table ID"
// @Success 201 {object} service.Order "Order has been created"
// @Failure 400 {object} errorResponse "Missing query parameter"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot create order"
// @Router /orders [post]
// @Security Cookie
func (a *Api) createOrder(c *gin.Context) {
	table, err1 := strconv.ParseUint(c.Query("table"), 10, 64)
	item, err2 := strconv.ParseUint(c.Query("item"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing query parameter"})
		return
	}
	order := service.Order{TableID: table, ItemID: item, IsServed: false}
	err := service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create order"})
	} else {
		c.JSON(http.StatusCreated, order)
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
	order, err := service.DoesOrderExist(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = service.DeleteOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete order"})
	} else {
		c.Status(http.StatusOK)
	}
}

// @Schemes
// @Summary get all orderItems
// @Description gets all orderItems as array
// @Tags orderItems
// @Produce json
// @Param type query int true "ItemType"
// @Success 200 {array} service.OrderItem
// @Failure 401 "Unauthorized"
// @Router /orders/items [get]
// @Security Cookie
func (a *Api) getOrderItems(c *gin.Context) {
	orderType := c.Query("type")
	c.JSON(http.StatusOK, service.GetOrderItemsForType(orderType))
}

// @Schemes
// @Summary create new orderItem
// @Description creates a new orderItem and returns it
// @Tags orderItems
// @Accept json
// @Produce json
// @Param order body service.OrderItem true "the orderItem as an object"
// @Success 201 {object} service.OrderItem "OrderItem has been created"
// @Failure 400 {object} errorResponse "Missing orderItem body"
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse "Cannot orderItem order"
// @Router /orders/items [post]
// @Security Cookie
func (a *Api) createOrderItem(c *gin.Context) {
	var orderItem service.OrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing orderItem body"})
		return
	}
	err = service.CreateOrderItem(&orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot create orderItem"})
	} else {
		c.JSON(http.StatusCreated, orderItem)
	}
}

// @Schemes
// @Summary update a orderItem
// @Description updates a orderItem with provided information
// @Tags orderItems
// @Accept json
// @Produce json
// @Param bookmark body service.OrderItem true "the orderItem with new values"
// @Success 200	{object} service.OrderItem "OrderItem has been updated"
// @Failure 400 {object} errorResponse "Missing orderItem body"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot update orderItem"
// @Router /orders/items [put]
// @Security Cookie
func (a *Api) updateOrderItem(c *gin.Context) {
	var newOrderItem service.OrderItem
	err := c.ShouldBindJSON(&newOrderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"Missing bookmark body"})
		return
	}
	oldOrderItem, err := service.DoesOrderItemExist(strconv.Itoa(int(newOrderItem.ID)))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = service.UpdateOrderItem(oldOrderItem, newOrderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot update bookmark"})
	} else {
		c.JSON(http.StatusOK, newOrderItem)
	}
}

// @Schemes
// @Summary delete an orderItem
// @Description deletes an orderItem from the database
// @Tags orderItems
// @Produce json
// @Param id path int true "OrderItem ID"
// @Success 200 "OK"
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found"
// @Failure 500 {object} errorResponse "Cannot delete orderItem"
// @Router /orders/items/{id} [delete]
// @Security Cookie
func (a *Api) deleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	orderItem, err := service.DoesOrderItemExist(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = service.DeleteOrderItem(&orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{"Cannot delete orderItem"})
	} else {
		c.Status(http.StatusOK)
	}
}
