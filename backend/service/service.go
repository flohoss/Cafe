package service

import (
	"cafe/config"
	"github.com/sirupsen/logrus"
	"time"
)

type ItemType uint

const (
	Food ItemType = iota
	Drink
)

type (
	Table struct {
		ID         uint      `gorm:"primaryKey" json:"id"`
		OrderCount uint      `json:"order_count"`
		Total      uint      `json:"total"`
		IsActive   bool      `json:"is_active"`
		UpdatedAt  time.Time `json:"updated_at"`
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
		Price       float32  `json:"price"`
	}

	Bill struct {
		ID        uint      `gorm:"primaryKey" json:"id"`
		Total     float32   `json:"total"`
		CreatedAt time.Time `json:"created_at"`
	}

	BillItem struct {
		ID          uint     `gorm:"primaryKey" json:"id"`
		ItemType    ItemType `json:"item_type"`
		Description string   `json:"description"`
		Price       float32  `json:"price"`
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
