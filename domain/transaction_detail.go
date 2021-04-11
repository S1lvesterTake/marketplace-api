package domain

type TransactionDetail struct {
	BaseModel
	TransactionID uint64 `gorm:"column:transaction_id;not null" json:"transaction_id"`
	ProductID     uint64 `gorm:"column:product_id;not null;index" json:"product_id"`
	Quantity      string `gorm:"column:quantity;not null" json:"quantity"`
	DetailAmount  int64  `gorm:"column:detail_amount;not null" json:"detail_amount"`
}
