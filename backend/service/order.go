package service

import (
	"cafe/config"
	"fmt"
)

type (
	Order struct {
		ID       uint `gorm:"primaryKey" json:"id" validate:"optional"`
		TableID  uint `json:"table_id" validate:"required"`
		ItemID   uint `json:"item_id" validate:"required"`
		IsServed bool `json:"is_served" default:"false" validate:"required"`
	}

	OrderItem struct {
		ID          uint     `gorm:"primaryKey" json:"id" validate:"optional"`
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

func GetAllOrders() []Order {
	var orders []Order
	config.C.Database.ORM.Find(&orders)
	return orders
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
