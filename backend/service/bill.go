package service

import "cafe/config"

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

func GetAllBills() []Bill {
	var bills []Bill
	config.C.Database.ORM.Find(&bills)
	return bills
}

func CreateBill(table *Table) Bill {
	var bill Bill
	bill.TableID = table.ID
	bill.Total = table.Total
	config.C.Database.ORM.Create(&bill)
	return bill
}
