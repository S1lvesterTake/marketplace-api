package domain

type Transaction struct {
	BaseModel
	CustomerID        uint64              `gorm:"column:customer_id;not null;index" json:"customer_id"`
	TransactionNumber uint64              `gorm:"column:transaction_number;not null;index" json:"transaction_number"`
	CodeID            string              `gorm:"column:code_id;not null;index" json:"code_id"`
	TotalAmount       int64               `gorm:"column:total_amount;not null" json:"total_amount"`
	TransactionDetail []TransactionDetail `json:"transaction_details"`
}
