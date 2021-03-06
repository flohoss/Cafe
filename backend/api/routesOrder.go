package api

import (
	"cafe/config"
	"cafe/hub"
	"cafe/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

// @Schemes
// @Summary get all orders
// @Description gets all orders as array
// @Tags orders
// @Produce json
// @Param table query int false "Table ID"
// @Param grouping query bool false "grouping"
// @Param filter query string false "filter"
// @Success 200 {array} service.Order
// @Failure 401 "Unauthorized"
// @Router /orders [get]
// @Security Cookie
func (a *Api) getOrders(c *gin.Context) {
	table, _ := strconv.ParseUint(c.Query("table"), 10, 64)
	grouping, _ := strconv.ParseBool(c.Query("grouping"))
	stringFiler, filterPresent := c.GetQuery("filter")
	var filter []string
	if filterPresent {
		filter = strings.Split(stringFiler, ",")
	}
	options := service.GetOrderOptions{TableId: table, Grouped: grouping, Filter: filter}
	var orders []service.Order
	if options.TableId == 0 {
		orders = service.GetAllActiveOrders()
	} else {
		orders = service.GetAllOrdersForTable(options)
	}
	c.JSON(http.StatusOK, orders)
}

// @Schemes
// @Summary create new order
// @Description creates a new order and returns it
// @Tags orders
// @Accept json
// @Produce json
// @Param item query int true "OrderItem ID"
// @Param table query int true "Table ID"
// @Success 201 {object} service.Order
// @Failure 400 {object} errorResponse
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse
// @Router /orders [post]
// @Security Cookie
func (a *Api) createOrder(c *gin.Context) {
	table, err1 := strconv.ParseUint(c.Query("table"), 10, 64)
	item, err2 := strconv.ParseUint(c.Query("item"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	order := service.Order{TableID: table, OrderItemID: item, IsServed: false}
	err := service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
	} else {
		c.JSON(http.StatusCreated, order)
	}
}

// @Schemes
// @Summary delete an order
// @Description deletes an order from the database
// @Tags orders
// @Produce json
// @Param item query int true "OrderItem ID"
// @Param table query int true "Table ID"
// @Success 200 "OK"
// @Failure 400 {object} errorResponse
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse
// @Router /orders [delete]
// @Security Cookie
func (a *Api) deleteOrder(c *gin.Context) {
	item := c.Query("item")
	table := c.Query("table")
	if table == "" || item == "" {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	err := service.DeleteOrder(table, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}

// @Schemes
// @Summary update an order
// @Description updates an order with provided information
// @Tags orders
// @Accept json
// @Produce json
// @Param order body service.Order true "updated Order"
// @Success 200	{object} service.Order
// @Failure 400 {object} errorResponse
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found" errorResponse
// @Failure 500 {object} errorResponse
// @Router /orders [put]
// @Security Cookie
func (a *Api) updateOrder(c *gin.Context) {
	var newOrder service.Order
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	oldOrder, err := service.DoesOrderExist(strconv.Itoa(int(newOrder.ID)))
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	}
	err = service.UpdateOrder(&oldOrder, &newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
	} else {
		c.JSON(http.StatusOK, newOrder)
	}
}

func (a *Api) serveWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.WithField("error", err).Warning("Cannot upgrade websocket")
		return
	}
	messageChan := make(hub.NotifierChan)
	a.Hub.NewClients <- messageChan
	defer func() {
		a.Hub.ClosingClients <- messageChan
		conn.Close()
	}()
	go readPump(conn)
	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				err := conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}
			err := conn.WriteJSON(msg)
			if err != nil {
				return
			}
		}
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
	if orderType == "" {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
	} else {
		c.JSON(http.StatusOK, service.GetOrderItemsForType(orderType))
	}
}

// @Schemes
// @Summary create new orderItem
// @Description creates a new orderItem and returns it
// @Tags orderItems
// @Accept json
// @Produce json
// @Param order body service.OrderItem true "OrderItem ID"
// @Success 201 {object} service.OrderItem
// @Failure 400 {object} errorResponse
// @Failure 401 "Unauthorized"
// @Failure 500 {object} errorResponse
// @Router /orders/items [post]
// @Security Cookie
func (a *Api) createOrderItem(c *gin.Context) {
	var orderItem service.OrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	err = service.CreateOrderItem(&orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
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
// @Param orderItem body service.OrderItem true "updated OrderItem"
// @Success 200	{object} service.OrderItem
// @Failure 400 {object} errorResponse
// @Failure 401 "Unauthorized"
// @Failure 404 "Not Found" errorResponse
// @Failure 500 {object} errorResponse
// @Router /orders/items [put]
// @Security Cookie
func (a *Api) updateOrderItem(c *gin.Context) {
	var newOrderItem service.OrderItem
	err := c.ShouldBindJSON(&newOrderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{config.MissingInformation.String()})
		return
	}
	oldOrderItem, err := service.DoesOrderItemExist(strconv.Itoa(int(newOrderItem.ID)))
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	}
	err = service.UpdateOrderItem(&oldOrderItem, &newOrderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
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
// @Failure 500 {object} errorResponse
// @Router /orders/items/{id} [delete]
// @Security Cookie
func (a *Api) deleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	orderItem, err := service.DoesOrderItemExist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse{err.Error()})
		return
	}
	err = service.DeleteOrderItem(&orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}
