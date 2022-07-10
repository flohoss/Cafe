package service

import (
	"cafe/config"
	"gorm.io/plugin/soft_delete"
)

type Table struct {
	ID         uint                  `gorm:"primaryKey" json:"id" validate:"optional"`
	OrderCount uint64                `json:"order_count" validate:"required"`
	Total      float64               `json:"total" validate:"required"`
	UpdatedAt  int64                 `json:"updated_at" validate:"optional"`
	IsDeleted  soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_deleted" swaggerignore:"true"`
}

func DoesTableExist(id string) (bool, Table) {
	var t Table
	result := config.C.Database.ORM.Limit(1).Find(&t, id)
	if result.RowsAffected == 0 {
		return false, t
	}
	return true, t
}
