package service

import (
	"cafe/config"
	"fmt"
	"time"
)

type (
	Bill struct {
		ID        uint64  `gorm:"primaryKey" json:"id" validate:"optional"`
		TableID   uint64  `json:"table_id" validate:"required"`
		Total     float32 `json:"total" validate:"required"`
		CreatedAt int64   `json:"created_at" validate:"optional"`
	}

	BillItem struct {
		ID          uint64  `gorm:"primaryKey" json:"id" validate:"optional"`
		BillID      uint64  `json:"bill_id" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Total       float64 `json:"price" validate:"required"`
		Amount      float64 `json:"amount" validate:"required"`
	}
)

func GetAllBills(year string, month string, day string) ([]Bill, error) {
	var bills []Bill
	const layout = "2006-1-02"
	date := fmt.Sprintf("%s-%s-%s", year, month, day)
	today, err := time.Parse(layout, date)
	if err != nil {
		return bills, fmt.Errorf(config.CannotFind.String())
	}
	beginningOfDay := today.Unix()
	endOfDay := today.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second).Unix()
	config.C.Database.ORM.Where("created_at BETWEEN ? AND ?", beginningOfDay, endOfDay).Find(&bills)
	return bills, nil
}

func CreateBill(orders []Order) Bill {
	var bill Bill
	var total float32 = 0
	for _, order := range orders {
		total += order.Total
	}
	if len(orders) > 0 {
		bill.TableID = orders[0].TableID
	}
	bill.Total = total
	config.C.Database.ORM.Create(&bill)
	return bill
}
