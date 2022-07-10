package service

import (
	"cafe/config"
	"github.com/sirupsen/logrus"
)

type ItemType uint

const (
	Food ItemType = iota
	Drink
)

type (
	Bill struct {
		ID        uint    `gorm:"primaryKey" json:"id"`
		TableID   uint    `json:"table_id"`
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
