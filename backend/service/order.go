package service

import "cafe/config"

type (
	Order struct {
		ID       uint `gorm:"primaryKey" json:"id"`
		TableID  uint
		ItemID   uint
		Item     OrderItem `json:"item"`
		IsServed bool      `json:"is_served"`
	}

	OrderItem struct {
		ID          uint     `gorm:"primaryKey" json:"id"`
		ItemType    ItemType `json:"item_type"`
		Description string   `json:"description"`
		Price       float64  `json:"price"`
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
