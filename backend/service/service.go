package service

import (
	"cafe/config"
	"github.com/sirupsen/logrus"
)

type (
	ItemType uint

	NotifierType uint

	WebSocketMsg struct {
		Type    NotifierType `json:"type"`
		Payload []Order      `json:"payload"`
	}
)

const (
	Create NotifierType = iota
	Delete
	DeleteAll
)

const (
	Food ItemType = iota
	Drink
)

var LiveCh chan WebSocketMsg

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

func Initialize() {
	migrateHelper(Table{}, "table")
	migrateHelper(Order{}, "order")
	migrateHelper(OrderItem{}, "orderItem")
	migrateHelper(Bill{}, "bill")
	migrateHelper(BillItem{}, "billItem")
	LiveCh = make(chan WebSocketMsg)
}
