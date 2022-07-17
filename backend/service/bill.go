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
		ID          uint64   `gorm:"primaryKey" json:"id" validate:"optional"`
		BillID      uint64   `json:"bill_id" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Total       float32  `json:"total" validate:"required"`
		Price       float32  `json:"price" validate:"required"`
		Amount      uint64   `json:"amount" validate:"required"`
		ItemType    ItemType `json:"item_type" validate:"required"`
	}
)

func DoesBillExist(id string) (Bill, error) {
	var bill Bill
	result := config.C.Database.ORM.Limit(1).Find(&bill, id)
	if result.RowsAffected == 0 {
		return bill, fmt.Errorf(config.CannotFind.String())
	}
	return bill, nil
}

func GetAllBillItems(billId uint64) ([]BillItem, error) {
	var billItems []BillItem
	result := config.C.Database.ORM.Where("bill_id = ?", billId).Find(&billItems)
	if result.RowsAffected == 0 {
		return billItems, fmt.Errorf(config.CannotFind.String())
	}
	return billItems, nil
}

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
	config.C.Database.ORM.Where("created_at BETWEEN ? AND ?", beginningOfDay, endOfDay).Order("table_id, created_at").Find(&bills)
	return bills, nil
}

func CreateBill(options GetOrderOptions) (Bill, error) {
	orders := GetAllOrdersForTable(options)
	var bill Bill
	var total float32 = 0
	for _, order := range orders {
		total += order.Total
	}
	bill.TableID = options.TableId
	bill.Total = total
	err := config.C.Database.ORM.Create(&bill).Error
	if err != nil {
		return bill, fmt.Errorf(config.CannotCreate.String())
	}
	for _, order := range orders {
		billItem := BillItem{
			BillID:      bill.ID,
			Description: order.OrderItem.Description,
			Total:       order.Total,
			Price:       order.OrderItem.Price,
			Amount:      order.OrderCount,
			ItemType:    order.OrderItem.ItemType,
		}
		config.C.Database.ORM.Create(&billItem)
	}
	ordersToDelete := GetAllOrdersForTable(GetOrderOptions{TableId: options.TableId, Grouped: false, Filter: options.Filter})
	config.C.Database.ORM.Delete(&ordersToDelete)
	return bill, nil
}
