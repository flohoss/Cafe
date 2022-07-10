package service

import "cafe/config"

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

func DoesOrderExist(id string) (bool, Order) {
	var o Order
	result := config.C.Database.ORM.Limit(1).Find(&o, id)
	if result.RowsAffected == 0 {
		return false, o
	}
	return true, o
}

func DoesOrderItemExist(id string) (bool, OrderItem) {
	var o OrderItem
	result := config.C.Database.ORM.Limit(1).Find(&o, id)
	if result.RowsAffected == 0 {
		return false, o
	}
	return true, o
}

func GetAllOrders() []Order {
	var orders []Order
	config.C.Database.ORM.Find(&orders)
	return orders
}
