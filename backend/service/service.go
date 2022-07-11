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

func ParseItemType(itemType string) ItemType {
	switch itemType {
	case "0":
		return Food
	case "1":
		return Drink
	default:
		return Food
	}
}

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
