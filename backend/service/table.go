package service

import (
	"cafe/config"
	"fmt"
	"gorm.io/plugin/soft_delete"
)

type Table struct {
	ID         uint64                `gorm:"primaryKey" json:"id" validate:"optional"`
	OrderCount uint64                `json:"order_count" validate:"required"`
	Total      float64               `json:"total" validate:"required"`
	UpdatedAt  int64                 `json:"updated_at" validate:"optional"`
	IsDeleted  soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_deleted" swaggerignore:"true"`
}

func DoesTableExist(id string) (Table, error) {
	var table Table
	result := config.C.Database.ORM.Limit(1).Find(&table, id)
	if result.RowsAffected == 0 {
		return table, fmt.Errorf(config.CannotFind.String())
	}
	return table, nil
}

func GetAllTables() []Table {
	var tables []Table
	config.C.Database.ORM.Model(
		&Table{},
	).Joins(
		"left join orders on tables.id = orders.table_id",
	).Joins(
		"left join order_items on orders.order_item_id = order_items.id",
	).Select(
		"tables.id, tables.updated_at, sum(order_items.price) as total, count(orders.id) as order_count",
	).Group(
		"tables.id",
	).Order("tables.id").Find(&tables)
	return tables
}

func CreateNewTable() (Table, error) {
	var table Table
	var err error
	result := config.C.Database.ORM.Unscoped().Where("is_deleted = ?", 1).Limit(1).Find(&table)
	if result.RowsAffected == 0 {
		err = config.C.Database.ORM.Create(&table).Error
	} else {
		table.IsDeleted = 0
		err = config.C.Database.ORM.Save(&table).Error
	}
	if err != nil {
		return table, fmt.Errorf(config.CannotCreate.String())
	}
	return table, nil
}

func DeleteLatestTable() error {
	var table Table
	err := config.C.Database.ORM.Last(&table).Error
	if err != nil {
		return fmt.Errorf(config.CannotFind.String())
	}
	err = config.C.Database.ORM.Delete(&table).Error
	if err != nil {
		return fmt.Errorf(config.CannotDelete.String())
	}
	return nil
}
