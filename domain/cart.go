package domain

type Cart struct {
	BaseModel
	CustomerID       uint64       `gorm:"column:customer_id;not null;index" json:"customer_id"`
	CodeID           string       `gorm:"column:code_id;not null;index" json:"code_id"`
	Quantity         int          `gorm:"column:quantity" json:"quantity"`
	CartTotalAmmount int64        `gorm:"column:cart_total_amount" json:"cart_total_amount"`
	CartDetail       []CartDetail `gorm:"foreignKey" json:"cart_details"`
}
