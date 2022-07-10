package service

import (
	"cafe/config"
	"github.com/sirupsen/logrus"
	"gorm.io/plugin/soft_delete"
)

type ItemType uint

const (
	Food ItemType = iota
	Drink
)

type (
	Table struct {
		ID         uint                  `gorm:"primaryKey" json:"id" validate:"optional"`
		OrderCount uint64                `json:"order_count" validate:"required"`
		Total      float64               `json:"total" validate:"required"`
		UpdatedAt  int64                 `json:"updated_at" validate:"optional"`
		IsDeleted  soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_deleted" swaggerignore:"true"`
	}

	Order struct {
		ID       uint `gorm:"primaryKey" json:"id"`
		TableID  uint
		Table    Table `json:"table"`
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

	Bill struct {
		ID        uint    `gorm:"primaryKey" json:"id"`
		Total     float64 `json:"total"`
		CreatedAt int64   `json:"created_at"`
	}

	BillItem struct {
		ID          uint     `gorm:"primaryKey" json:"id"`
		ItemType    ItemType `json:"item_type"`
		Description string   `json:"description"`
		Price       float64  `json:"price"`
		BillID      uint
		Bill        Bill `json:"bill"`
	}
)

func migrateHelper(i interface{}, name string) {
	err := config.C.Database.ORM.AutoMigrate(i)
	if err != nil {
		logrus.WithField("error", err).Fatalf("Failed to migrate %s", name)
	}
}

func MigrateToDb() {
	migrateHelper(Table{}, "table")
	migrateHelper(Order{}, "order")
	migrateHelper(OrderItem{}, "orderItem")
	migrateHelper(Bill{}, "bill")
	migrateHelper(BillItem{}, "billItem")
}

func DoesOrderExist(id string) (bool, Order) {
	var o Order
	result := config.C.Database.ORM.Limit(1).Find(&o, id)
	if result.RowsAffected == 0 {
		return false, o
	}
	return true, o
}

func DoesTableExist(id string) (bool, Table) {
	var t Table
	result := config.C.Database.ORM.Limit(1).Find(&t, id)
	if result.RowsAffected == 0 {
		return false, t
	}
	return true, t
}
