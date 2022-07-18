package service

import (
	"cafe/config"
	"fmt"
	"strconv"
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
	yearI, yearErr := strconv.Atoi(year)
	if yearErr != nil {
		return bills, fmt.Errorf("jahr " + config.CannotParse.String())
	}
	monthI, monthErr := strconv.Atoi(month)
	if monthErr != nil {
		return bills, fmt.Errorf("monat " + config.CannotParse.String())
	}
	dayI, dayErr := strconv.Atoi(day)
	if dayErr != nil {
		return bills, fmt.Errorf("tag " + config.CannotParse.String())
	}
	loc, locErr := time.LoadLocation("Local")
	if locErr != nil {
		return bills, fmt.Errorf("zeitzone " + config.CannotParse.String())
	}
	today := time.Date(yearI, time.Month(monthI), dayI, 0, 0, 0, 0, loc)
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
	err = config.C.Database.ORM.Delete(&ordersToDelete).Error
	if err != nil {
		return bill, fmt.Errorf(config.CannotDelete.String())
	}
	LiveCh <- WebSocketMsg{
		Type:    DeleteAll,
		Payload: ordersToDelete,
	}
	return bill, nil
}

func DeleteBill(bill *Bill) error {
	err := config.C.Database.ORM.Delete(bill).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	billItemsToDelete, _ := GetAllBillItems(bill.ID)
	config.C.Database.ORM.Delete(&billItemsToDelete)
	return nil
}
