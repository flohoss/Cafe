package service

import (
	"cafe/config"
	"fmt"
)

type (
	Order struct {
		ID          uint64    `gorm:"primaryKey" json:"id" validate:"optional"`
		TableID     uint64    `json:"table_id" validate:"required"`
		OrderItemID uint64    `json:"order_item_id" validate:"required"`
		OrderItem   OrderItem `json:"order_item" validate:"required"`
		IsServed    bool      `json:"is_served" default:"false" validate:"required"`
	}

	OrderItem struct {
		ID          uint64   `gorm:"primaryKey" json:"id" validate:"optional"`
		ItemType    ItemType `json:"item_type" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Price       float64  `json:"price" validate:"required"`
	}
)

func DoesOrderExist(id string) (Order, error) {
	var order Order
	result := config.C.Database.ORM.Limit(1).Find(&order, id)
	if result.RowsAffected == 0 {
		return order, fmt.Errorf("table not found")
	}
	return order, nil
}

func DoesOrderItemExist(id string) (OrderItem, error) {
	var orderItem OrderItem
	result := config.C.Database.ORM.Limit(1).Find(&orderItem, id)
	if result.RowsAffected == 0 {
		return orderItem, fmt.Errorf("table not found")
	}
	return orderItem, nil
}

func GetAllOrders(table string, itemType string) ([]Order, error) {
	var orders []Order
	err := config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").Where("table_id = ? AND item_type = ?", table, itemType).Find(&orders).Error
	return orders, err
}

func CreateOrder(order *Order) error {
	result := config.C.Database.ORM.Create(order)
	return result.Error
}

func DeleteOrder(order *Order) error {
	result := config.C.Database.ORM.Delete(order)
	return result.Error
}

func GetOrderItemsForType(itemType string) []OrderItem {
	var orderItems []OrderItem
	config.C.Database.ORM.Order("description").Where("item_type = ?", ParseItemType(itemType)).Find(&orderItems)
	return orderItems
}

func CreateOrderItem(oderItem *OrderItem) error {
	result := config.C.Database.ORM.Create(oderItem)
	return result.Error
}

func UpdateOrderItem(old OrderItem, new OrderItem) error {
	result := config.C.Database.ORM.First(&old).Updates(&new)
	return result.Error
}

func DeleteOrderItem(oderItem *OrderItem) error {
	result := config.C.Database.ORM.Delete(oderItem)
	return result.Error
}
