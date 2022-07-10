package service

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
