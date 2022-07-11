package service

import (
	"cafe/api"
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
		return table, fmt.Errorf(api.CannotFind.String())
	}
	return table, nil
}

func GetAllTables() []Table {
	var tables []Table
	config.C.Database.ORM.Find(&tables)
	return tables
}

func CreateNewTable() (Table, error) {
	var table Table
	result := config.C.Database.ORM.Unscoped().Where("is_deleted = ?", 1).Limit(1).Find(&table)
	if result.RowsAffected == 0 {
		result = config.C.Database.ORM.Create(&table)
	} else {
		table.IsDeleted = 0
		config.C.Database.ORM.Save(&table)
	}
	return table, result.Error
}

func DeleteLatestTable() error {
	var table Table
	config.C.Database.ORM.Last(&table)
	table.Total = 0
	table.OrderCount = 0
	config.C.Database.ORM.Save(&table)
	result := config.C.Database.ORM.Delete(&table)
	return result.Error
}
