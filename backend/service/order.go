package service

import (
	"cafe/config"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	Order struct {
		ID          uint64    `gorm:"primaryKey" json:"id" validate:"optional"`
		TableID     uint64    `json:"table_id" validate:"required"`
		OrderItemID uint64    `json:"order_item_id" validate:"required"`
		OrderItem   OrderItem `json:"order_item" validate:"required"`
		UpdatedAt   int64     `json:"updated_at" validate:"optional"`
		IsServed    bool      `json:"is_served" default:"false" validate:"required"`
		Total       float32   `json:"total" validate:"required"`
		OrderCount  uint64    `json:"order_count" validate:"required"`
	}

	OrderItem struct {
		ID          uint64   `gorm:"primaryKey" json:"id" validate:"optional"`
		ItemType    ItemType `json:"item_type" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Price       float32  `json:"price" validate:"required"`
	}

	GetOrderOptions struct {
		TableId uint64   `json:"table_id"`
		Grouped bool     `json:"grouped"`
		Filter  []string `json:"filter"`
	}
)

func updateTableUpdatedAt(tx *gorm.DB, o *Order) {
	var table Table
	tx.Where("id = ?", o.TableID).First(&table)
	table.UpdatedAt = time.Now().Unix()
	tx.Save(&table)
}

func (o *Order) AfterCreate(tx *gorm.DB) (err error) {
	updateTableUpdatedAt(tx, o)
	return
}

func (o *Order) AfterDelete(tx *gorm.DB) (err error) {
	updateTableUpdatedAt(tx, o)
	return
}

func DoesOrderItemExist(id string) (OrderItem, error) {
	var orderItem OrderItem
	result := config.C.Database.ORM.Limit(1).Find(&orderItem, id)
	if result.RowsAffected == 0 {
		return orderItem, fmt.Errorf(config.CannotFind.String())
	}
	return orderItem, nil
}

func DoesOrderExist(id string) (Order, error) {
	var order Order
	result := config.C.Database.ORM.Limit(1).Find(&order, id)
	if result.RowsAffected == 0 {
		return order, fmt.Errorf(config.CannotFind.String())
	}
	return order, nil
}

func GetAllActiveOrders() []Order {
	var orders []Order
	config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").Where("is_served = ?", 0).Order("updated_at").Find(&orders)
	return orders
}

func GetAllOrdersForTable(options GetOrderOptions) []Order {
	var orders []Order
	if options.Grouped {
		if options.Filter[0] == "" {
			config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").Select("table_id, order_item_id, sum(price) as total, count(order_item_id) as order_count").Group("order_item_id").Where("table_id = ?", options.TableId).Order("item_type, description").Find(&orders)
		} else {
			config.C.Database.ORM.Model(&Order{}).Find(&orders, options.Filter).Joins("OrderItem").Select("table_id, order_item_id, sum(price) as total, count(order_item_id) as order_count").Group("order_item_id").Where("table_id = ?", options.TableId).Order("item_type, description").Find(&orders)
		}
	} else {
		config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").Where("table_id = ?", options.TableId).Order("item_type, description").Find(&orders)
	}
	return orders
}

func CreateOrder(order *Order) error {
	err := config.C.Database.ORM.Create(order).Error
	if err != nil {
		return fmt.Errorf(config.CannotCreate.String())
	}
	config.C.Database.ORM.Model(&Order{}).Joins("OrderItem").First(order)
	LiveCh <- WebSocketMsg{
		Type:    Create,
		Payload: *order,
	}
	return nil
}

func UpdateOrder(old *Order, new *Order) error {
	err := config.C.Database.ORM.First(old).Updates(new).Error
	if err != nil {
		return fmt.Errorf(config.CannotUpdate.String())
	}
	return nil
}

func DeleteOrder(tableId string, orderItemId string) error {
	var order Order
	err := config.C.Database.ORM.Where("table_id = ? AND order_item_id = ?", tableId, orderItemId).Last(&order).Error
	if err != nil {
		return fmt.Errorf(config.CannotFind.String())
	}
	err = config.C.Database.ORM.Delete(&order).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	LiveCh <- WebSocketMsg{
		Type:    Delete,
		Payload: order,
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

func UpdateOrderItem(old *OrderItem, new *OrderItem) error {
	err := config.C.Database.ORM.First(old).Updates(new).Error
	if err != nil {
		return fmt.Errorf(config.CannotUpdate.String())
	}
	return nil
}

func DeleteOrderItem(oderItem *OrderItem) error {
	err := config.C.Database.ORM.Delete(oderItem).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	return nil
}
