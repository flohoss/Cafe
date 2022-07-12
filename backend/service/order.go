package service

import (
	"cafe/config"
	"fmt"
)

type (
	Order struct {
		ID          uint64    `gorm:"primaryKey" json:"id" validate:"optional"`
		TableID     uint64    `json:"table_id" validate:"required"`
		Table       Table     `json:"table" validate:"required"`
		OrderItemID uint64    `json:"order_item_id" validate:"required"`
		OrderItem   OrderItem `json:"order_item" validate:"required"`
		IsServed    bool      `json:"is_served" default:"false" validate:"required"`
		Total       uint64    `json:"total" validate:"optional"`
	}

	OrderItem struct {
		ID          uint64   `gorm:"primaryKey" json:"id" validate:"optional"`
		ItemType    ItemType `json:"item_type" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Price       float64  `json:"price" validate:"required"`
	}
)

func DoesOrderItemExist(id string) (OrderItem, error) {
	var orderItem OrderItem
	result := config.C.Database.ORM.Limit(1).Find(&orderItem, id)
	if result.RowsAffected == 0 {
		return orderItem, fmt.Errorf(config.CannotFind.String())
	}
	return orderItem, nil
}

func GetAllOrders(table string, itemType string) []Order {
	var orders []Order
	config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").Joins("Table").Select("table_id", "order_item_id", "count(order_item_id) as total").Group("order_item_id").Where("table_id = ? AND item_type = ?", table, itemType).Order("description").Find(&orders)
	return orders
}

func CreateOrder(order *Order) error {
	err := config.C.Database.ORM.Create(order).Error
	if err != nil {
		return fmt.Errorf(config.CannotCreate.String())
	}
	config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").First(order)
	return nil
}

func DeleteOrder(tableId string, orderItemId string) error {
	var order Order
	err := config.C.Database.ORM.Where("table_id = ? AND order_item_id = ?", tableId, orderItemId).First(&order).Error
	if err != nil {
		return fmt.Errorf(config.CannotFind.String())
	}
	err = config.C.Database.ORM.Delete(&order).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	return nil
}

func GetOrderItemsForType(itemType string) []OrderItem {
	var orderItems []OrderItem
	config.C.Database.ORM.Order("description").Where("item_type = ?", ParseItemType(itemType)).Find(&orderItems)
	return orderItems
}

func CreateOrderItem(oderItem *OrderItem) error {
	err := config.C.Database.ORM.Create(oderItem).Error
	if err != nil {
		return fmt.Errorf(config.CannotCreate.String())
	}
	return nil
}

func UpdateOrderItem(old OrderItem, new OrderItem) error {
	var order Order
	result := config.C.Database.ORM.Where("order_item_id = ?", old.ID).Limit(1).Find(&order)
	if result.RowsAffected != 0 {
		return fmt.Errorf(config.StillInUse.String())
	}
	config.C.Database.ORM.First(&old).Updates(&new)
	return nil
}

func DeleteOrderItem(oderItem *OrderItem) error {
	err := config.C.Database.ORM.Delete(oderItem).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	return nil
}
